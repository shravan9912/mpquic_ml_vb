package congestion

import "github.com/shravan9912/mpquic_actor_critic_v1/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}
