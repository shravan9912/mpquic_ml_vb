package utils

import "github.com/shravan9912/mpquic_actor_critic_v1/internal/protocol"

// ByteInterval is an interval from one ByteCount to the other
// +gen linkedlist
type ByteInterval struct {
	Start protocol.ByteCount
	End   protocol.ByteCount
}
