package wire

import (
	"bytes"

	"github.com/shravan9912/mpquic_ml_vb/internal/protocol"
	"github.com/shravan9912/mpquic_ml_vb/internal/utils"
)

// ComposeVersionNegotiation composes a Version Negotiation Packet
func ComposeVersionNegotiation(connectionID protocol.ConnectionID, versions []protocol.VersionNumber) []byte {
	fullReply := &bytes.Buffer{}
	responsePublicHeader := PublicHeader{
		ConnectionID: connectionID,
		PacketNumber: 1,
		VersionFlag:  true,
	}
	err := responsePublicHeader.Write(fullReply, protocol.VersionWhatever, protocol.PerspectiveServer)
	if err != nil {
		utils.Errorf("error composing version negotiation packet: %s", err.Error())
	}
	for _, v := range versions {
		utils.LittleEndian.WriteUint32(fullReply, protocol.VersionNumberToTag(v))
	}
	return fullReply.Bytes()
}
