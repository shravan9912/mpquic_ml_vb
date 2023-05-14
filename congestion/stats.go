package congestion

import "github.com/shravan9912/mpquic_ml_vb/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}
