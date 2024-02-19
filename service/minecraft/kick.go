package minecraft

import (
	"fmt"
	"time"

	"github.com/CubeWhyMC/NoDelay-Proxy-Server/common/mcprotocol"
	"github.com/CubeWhyMC/NoDelay-Proxy-Server/config"
)
func generateNewMessage(s *config.ConfigProxyService, name string) mcprotocol.Message {
	return mcprotocol.Message{
		Color: mcprotocol.White,
		Extra: []mcprotocol.Message{
			{Bold: true, Color: mcprotocol.Green, Text: "=========首次进入提示=========\n"},
			{Color: mcprotocol.Red, Text: "检测到您当前第一次进入本IP!\n"},
			{Color: mcprotocol.LightPurple, Text: "本IP暂不支持防安全警报。\n"},
			{Color: mcprotocol.Blue, Text: "请使用21+或已经历安全警报的账号进入本IP!\n"},
			{Color: mcprotocol.Gold, Text: "一旦被安全警报我们概不负责!\n"},
			{Color: mcprotocol.Green, Text: "如果已经使用21+或已经历安全警报的账号，请尝试重新进入。\n"},
			{Color: mcprotocol.White, Text: "还有其他问题，请开票获取支持!"},
		},
	}
}

func generateKickMessage(s *config.ConfigProxyService, name string) mcprotocol.Message {
	return mcprotocol.Message{
		Color: mcprotocol.White,
		Extra: []mcprotocol.Message{
			{Bold: true, Color: mcprotocol.Yellow, Text: fmt.Sprintf("%s", config.Config.PrivateConfig.Header)},
			{Text: " ‖ "},
			{Bold: true, Color: mcprotocol.Red, Text: "已拒绝服务\n"},

			{Text: "您无法加入当前服务器！\n"},
			{Text: "理由: "},
			{Color: mcprotocol.LightPurple, Text: "你的连接可能未经处理，或者你没有权限加入此服务器。\n"},
			{Text: "请联系管理员寻求帮助！\n\n"},

			{
				Color: mcprotocol.Gray,
				Text: fmt.Sprintf("时间戳: %d | 玩家名称: %s\n | 服务节点: %s\n",
					time.Now().UnixMilli(), name, s.Name),
			},
			{Text: fmt.Sprintf("%s", config.Config.PrivateConfig.ContactName)},
			{
				Color: mcprotocol.Blue, UnderLined: true,
				Text: fmt.Sprintf("%s", config.Config.PrivateConfig.ContactLink),
				// ClickEvent: chat.OpenURL("http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=eV_W6FV6hkjbeA35MNJ2lulA7M67JMig&authKey=E5hHr6NTSJ9u9z7eurOavBW9U6tE94P1EazZSGMGV71LCjsfvgMt0kXRaXyaDF4d&noverify=0&group_code=666259678"),
			},
		},
	}
}

func generatePlayerNumberLimitExceededMessage(s *config.ConfigProxyService, name string) mcprotocol.Message {
	return mcprotocol.Message{
		Color: mcprotocol.White,
		Extra: []mcprotocol.Message{
			{Bold: true, Color: mcprotocol.Yellow, Text: fmt.Sprintf("%s", config.Config.PrivateConfig.Header)},
			{Text: " ‖ "},
			{Bold: true, Color: mcprotocol.Red, Text: "已拒绝服务\n"},

			{Text: "你无法加入当前服务器！\n"},
			{Text: "Reason: "},
			{Color: mcprotocol.LightPurple, Text: "服务器当前人数已满载！\n"},
			{Text: "请联系管理员寻求帮助！\n\n"},

			{
				Color: mcprotocol.Gray,
				Text: fmt.Sprintf("时间戳: %d | 玩家名称: %s\n | 服务节点: %s\n",
					time.Now().UnixMilli(), name, s.Name),
			},
			{Text: fmt.Sprintf("%s", config.Config.PrivateConfig.ContactName)},
			{
				Color: mcprotocol.Blue, UnderLined: true,
				Text: fmt.Sprintf("%s", config.Config.PrivateConfig.ContactLink),
				// ClickEvent: chat.OpenURL("http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=eV_W6FV6hkjbeA35MNJ2lulA7M67JMig&authKey=E5hHr6NTSJ9u9z7eurOavBW9U6tE94P1EazZSGMGV71LCjsfvgMt0kXRaXyaDF4d&noverify=0&group_code=666259678"),
			},
		},
	}
}

func generateNewMessage(s *config.ConfigProxyService, name string) mcprotocol.Message {
	return mcprotocol.Message{
		Color: mcprotocol.White,
		Extra: []mcprotocol.Message{
			{Bold: true, Color: mcprotocol.Green, Text: "=========首次进入提示=========\n"},
			{Color: mcprotocol.Red, Text: "检测到您当前第一次进入本IP!\n"},
			{Color: mcprotocol.LightPurple, Text: "本IP暂不支持防安全警报。\n"},
			{Color: mcprotocol.Blue, Text: "请使用21+或已经历安全警报的账号进入本IP!\n"},
			{Color: mcprotocol.Gold, Text: "一旦被安全警报我们概不负责!\n"},
			{Color: mcprotocol.Green, Text: "如果已经使用21+或已经历安全警报的账号，请尝试重新进入。\n"},
			{Color: mcprotocol.White, Text: "还有其他问题，请开票获取支持!"},
		},
	}
}