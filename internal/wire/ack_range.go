package wire

import "github.com/shravan9912/mpquic_actor_critic_v1/internal/protocol"

// AckRange is an ACK range
type AckRange struct {
	First protocol.PacketNumber
	Last  protocol.PacketNumber
}
