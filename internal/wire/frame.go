package wire

import (
	"bytes"

	"github.com/shravan9912/mpquic_actor_critic_v1/internal/protocol"
)

// A Frame in QUIC
type Frame interface {
	Write(b *bytes.Buffer, version protocol.VersionNumber) error
	MinLength(version protocol.VersionNumber) (protocol.ByteCount, error)
}
