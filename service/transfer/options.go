package transfer

import (
	"sync/atomic"

	"github.com/CubeWhyMC/NoDelay-Proxy-Server/outbound"
)

type Options struct {
	Out                     outbound.Outbound
	IsTLSHandleNeeded       bool
	IsMinecraftHandleNeeded bool
	FlowType                int
	OnlineCount             atomic.Int32
}
