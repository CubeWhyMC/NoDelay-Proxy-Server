package service

import (
	"log"
	"net"
	"github.com/layou233/ZBProxy/common"
	"github.com/layou233/ZBProxy/config"
	"github.com/layou233/ZBProxy/outbound"
	"github.com/layou233/ZBProxy/outbound/socks"
	"github.com/layou233/ZBProxy/service/access"
	"github.com/layou233/ZBProxy/service/transfer"

	"github.com/fatih/color"
)

var Listeners []net.Listener

func StartNewService(s *config.ConfigProxyService) {
	// Check Settings
	var (
		isTLSHandleNeeded = s.TLSSniffing.RejectNonTLS ||
			s.TLSSniffing.RejectIfNonMatch ||
			len(s.TLSSniffing.SNIAllowListTags) != 0
		isMinecraftHandleNeeded = s.Minecraft.EnableHostnameRewrite ||
			s.Minecraft.EnableAnyDest ||
			s.Minecraft.MotdDescription != "" && s.Minecraft.MotdDescription != config.DefaultMotd ||
			s.Minecraft.MotdFavicon != ""
	)
	if isTLSHandleNeeded && isMinecraftHandleNeeded {
		log.Panic(color.HiRedString("Service %s: The current version can't handle TLS and Minecraft at the same time.", s.Name))
	}
	flowType := getFlowType(s.Flow)
	if flowType == -1 {
		log.Panic(color.HiRedString("Service %s: Unknown flow type '%s'.", s.Name, s.Flow))
	}
	if s.Minecraft.EnableHostnameRewrite && s.Minecraft.RewrittenHostname == "" {
		s.Minecraft.RewrittenHostname = s.TargetAddress
	}
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   nil, // listens on all available IP addresses of the local system
		Port: int(s.Listen),
	})
	if err != nil {
		log.Panic(color.HiRedString("Service %s: Can't start listening on port %v: %v", s.Name, s.Listen, err.Error()))
	}
	Listeners = append(Listeners, listen) // add to Listeners

	switch s.IPAccess.Mode {
	case access.DefaultMode:
	case access.AllowMode, access.BlockMode, access.DownMode, access.JokeMode:
		if s.IPAccess.ListTags == nil {
			log.Panic(color.HiRedString("Service %s: ListTags can't be null when access control enabled.", s.Name))
		}
		for _, tag := range s.IPAccess.ListTags {
			if _, err = access.GetTargetList(tag); err != nil {
				log.Panic(color.HiRedString("Service %s: %s", s.Name, err.Error()))
			}
		}
	default:
		log.Panicf("Unknown access control mode: %s", s.IPAccess.Mode)
	}

	if isMinecraftHandleNeeded {
		switch s.Minecraft.NameAccess.Mode {
		case access.DefaultMode:
		case access.AllowMode, access.BlockMode, access.DownMode, access.JokeMode:
			if s.Minecraft.NameAccess.ListTags == nil {
				log.Panic(color.HiRedString("Service %s: ListTags can't be null when access control enabled.", s.Name))
			}
			for _, tag := range s.Minecraft.NameAccess.ListTags {
				if _, err = access.GetTargetList(tag); err != nil {
					log.Panic(color.HiRedString("Service %s: %s", s.Name, err.Error()))
				}
			}
		default:
			log.Panicf("Unknown access control mode: %s", s.Minecraft.NameAccess.Mode)
		}
	}

	out := outbound.NewSystemOutbound(
		outbound.NewDialerControlFromOptions(s.SocketOptions))
	switch s.Outbound.Type {
	case "socks", "socks5", "socks4a", "socks4":
		out = &socks.Client{
			Dialer:  out,
			Version: s.Outbound.Type,
			Network: s.Outbound.Network,
			Address: s.Outbound.Address,
		}
	}

	options := &transfer.Options{
		Out:                     out,
		IsTLSHandleNeeded:       isTLSHandleNeeded,
		IsMinecraftHandleNeeded: isMinecraftHandleNeeded,
		FlowType:                flowType,
	}
	for {
		conn, err := listen.AcceptTCP()
		if err == nil {
			if s.IPAccess.Mode != access.DefaultMode {
				// https://stackoverflow.com/questions/29687102/how-do-i-get-a-network-clients-ip-converted-to-a-string-in-golang
				ip := conn.RemoteAddr().(*net.TCPAddr).IP.String()
				hit := false
				for _, list := range s.IPAccess.ListTags {
					if hit = common.Must(access.GetTargetList(list)).Has(ip); hit {
						break
					}
				}
				switch s.IPAccess.Mode {
				case access.AllowMode:
					if !hit {
						forciblyCloseTCP(conn)
						continue
					}
				case access.BlockMode:
					if hit {
						forciblyCloseTCP(conn)
						continue
					}
				case access.DownMode:
					if hit {
						forciblyCloseTCP(conn)
						continue
					}
				case access.JokeMode:
					if hit {
						forciblyCloseTCP(conn)
						continue
					}
				}
			}
			go newConnReceiver(s, conn, options)
		}
	}
}

func getFlowType(flow string) int {
	switch flow {
	case "origin":
		return transfer.FLOW_ORIGIN
	case "linux-zerocopy":
		return transfer.FLOW_LINUX_ZEROCOPY
	case "zerocopy":
		return transfer.FLOW_ZEROCOPY
	case "multiple":
		return transfer.FLOW_MULTIPLE
	case "auto", "":
		return transfer.FLOW_AUTO
	default:
		return -1
	}
}

func forciblyCloseTCP(conn *net.TCPConn) {
	//nolint:errcheck
	conn.SetLinger(0) // let Close send RST to forcibly close the connection
	conn.Close()      // forcibly close
}