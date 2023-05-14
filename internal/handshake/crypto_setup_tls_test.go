package handshake

import (
	"fmt"

	"github.com/shravan9912/mpquic_ml_vb/internal/crypto"
	"github.com/shravan9912/mpquic_ml_vb/internal/protocol"
	"github.com/shravan9912/mpquic_ml_vb/internal/testdata"
	"github.com/bifurcation/mint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeMintController struct {
	result mint.Alert
}

var _ crypto.MintController = &fakeMintController{}

func (h *fakeMintController) Handshake() mint.Alert {
	return h.result
}
func (h *fakeMintController) GetCipherSuite() mint.CipherSuiteParams { panic("not implemented") }
func (h *fakeMintController) ComputeExporter(label string, context []byte, keyLength int) ([]byte, error) {
	panic("not implemented")
}

func mockKeyDerivation(crypto.MintController, protocol.Perspective) (crypto.AEAD, error) {
	return &mockAEAD{encLevel: protocol.EncryptionForwardSecure}, nil
}

var _ = Describe("TLS Crypto Setup", func() {
	var (
		cs          *cryptoSetupTLS
		aeadChanged chan protocol.EncryptionLevel
	)

	BeforeEach(func() {
		aeadChanged = make(chan protocol.EncryptionLevel, 2)
		csInt, err := NewCryptoSetupTLS(
			"",
			protocol.PerspectiveServer,
			protocol.VersionTLS,
			testdata.GetTLSConfig(),
			nil,
			aeadChanged,
		)
		Expect(err).ToNot(HaveOccurred())
		cs = csInt.(*cryptoSetupTLS)
	})

	It("errors when the handshake fails", func() {
		alert := mint.AlertBadRecordMAC
		cs.conn = &fakeMintController{result: alert}
		err := cs.HandleCryptoStream()
		Expect(err).To(MatchError(fmt.Errorf("TLS handshake error: %s (Alert %d)", alert.String(), alert)))
	})

	It("derives keys", func() {
		cs.conn = &fakeMintController{result: mint.AlertNoAlert}
		cs.keyDerivation = mockKeyDerivation
		err := cs.HandleCryptoStream()
		Expect(err).ToNot(HaveOccurred())
		Expect(aeadChanged).To(Receive(Equal(protocol.EncryptionForwardSecure)))
		Expect(aeadChanged).To(BeClosed())
	})

	Context("escalating crypto", func() {
		var barbarFNVSigned []byte // a "barbar", FNV signed

		doHandshake := func() {
			cs.conn = &fakeMintController{result: mint.AlertNoAlert}
			cs.keyDerivation = mockKeyDerivation
			err := cs.HandleCryptoStream()
			Expect(err).ToNot(HaveOccurred())
		}

		BeforeEach(func() {
			nullAEAD := crypto.NewNullAEAD(protocol.PerspectiveServer, protocol.VersionTLS)
			barbarFNVSigned = nullAEAD.Seal(nil, []byte("barbar"), 0, nil)
		})

		Context("null encryption", func() {
			It("is used initially", func() {
				enc, sealer := cs.GetSealer()
				Expect(enc).To(Equal(protocol.EncryptionUnencrypted))
				d := sealer.Seal(nil, []byte("barbar"), 0, []byte{})
				Expect(d).To(Equal(barbarFNVSigned))
			})

			It("is accepted initially", func() {
				d, enc, err := cs.Open(nil, barbarFNVSigned, 0, []byte{})
				Expect(err).ToNot(HaveOccurred())
				Expect(d).To(Equal([]byte("barbar")))
				Expect(enc).To(Equal(protocol.EncryptionUnencrypted))
			})

			It("is used for crypto stream", func() {
				enc, sealer := cs.GetSealerForCryptoStream()
				Expect(enc).To(Equal(protocol.EncryptionUnencrypted))
				d := sealer.Seal(nil, []byte("barbar"), 0, []byte{})
				Expect(d).To(Equal(barbarFNVSigned))
			})

			It("errors if the has the wrong hash", func() {
				barbarFNVSigned[0]++
				_, enc, err := cs.Open(nil, barbarFNVSigned, 0, []byte{})
				Expect(err).To(MatchError("NullAEAD: failed to authenticate received data"))
				Expect(enc).To(Equal(protocol.EncryptionUnspecified))
			})

			It("is not accepted after the handshake completes", func() {
				doHandshake()
				_, enc, err := cs.Open(nil, barbarFNVSigned, 0, []byte{})
				Expect(err).To(MatchError("authentication failed"))
				Expect(enc).To(Equal(protocol.EncryptionUnspecified))
			})
		})

		Context("forward-secure encryption", func() {
			It("is used for sealing after the handshake completes", func() {
				doHandshake()
				enc, sealer := cs.GetSealer()
				Expect(enc).To(Equal(protocol.EncryptionForwardSecure))
				d := sealer.Seal(nil, []byte("barbar"), 0, nil)
				Expect(d).To(Equal([]byte("barbar forward sec")))
			})

			It("is used for opening after the handshake completes", func() {
				doHandshake()
				d, enc, err := cs.Open(nil, []byte("forward secure encrypted"), 0, nil)
				Expect(err).ToNot(HaveOccurred())
				Expect(enc).To(Equal(protocol.EncryptionForwardSecure))
				Expect(d).To(Equal([]byte("decrypted")))
			})
		})

		Context("forcing encryption levels", func() {
			It("forces null encryption", func() {
				doHandshake()
				sealer, err := cs.GetSealerWithEncryptionLevel(protocol.EncryptionUnencrypted)
				Expect(err).ToNot(HaveOccurred())
				d := sealer.Seal(nil, []byte("barbar"), 0, []byte{})
				Expect(d).To(Equal(barbarFNVSigned))
			})

			It("forces forward-secure encryption", func() {
				doHandshake()
				sealer, err := cs.GetSealerWithEncryptionLevel(protocol.EncryptionForwardSecure)
				Expect(err).ToNot(HaveOccurred())
				d := sealer.Seal(nil, []byte("barbar"), 0, []byte{})
				Expect(d).To(Equal([]byte("barbar forward sec")))
			})

			It("errors if the forward-secure AEAD is not available", func() {
				sealer, err := cs.GetSealerWithEncryptionLevel(protocol.EncryptionForwardSecure)
				Expect(err).To(MatchError("CryptoSetup: no sealer with encryption level forward-secure"))
				Expect(sealer).To(BeNil())
			})

			It("never returns a secure AEAD (they don't exist with TLS)", func() {
				doHandshake()
				sealer, err := cs.GetSealerWithEncryptionLevel(protocol.EncryptionSecure)
				Expect(err).To(MatchError("CryptoSetup: no sealer with encryption level encrypted (not forward-secure)"))
				Expect(sealer).To(BeNil())
			})

			It("errors if no encryption level is specified", func() {
				seal, err := cs.GetSealerWithEncryptionLevel(protocol.EncryptionUnspecified)
				Expect(err).To(MatchError("CryptoSetup: no sealer with encryption level unknown"))
				Expect(seal).To(BeNil())
			})
		})
	})
})
