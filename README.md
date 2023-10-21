## NoDelay

📝 **简介**  
该项目是一个基于ZBProxy魔改的Minecraft服务器代理程序。NoDelay是一个功能强大的代理工具，用于优化和管理Minecraft游戏的网络连接。

⚙️ **配置文件结构**

```json
{
    "Services": [
        {
            "Name": "HypixelDefault",
            "TargetAddress": "mc.hypixel.net",
            "TargetPort": 25565,
            "Listen": 25565,
            "Flow": "auto",
            "IPAccess": {
                "Mode": ""
            },
            "Minecraft": {
                "EnableHostnameRewrite": true,
                "EnableHostnameAccess": false,
                "OnlineCount": {
                    "Max": 10,
                    "Online": -1,
                    "EnableMaxLimit": true
                },
                "NameAccess": {
                    "Mode": ""
                },
                "AnyDestSettings": {},
                "PingMode": "",
                "MotdFavicon": "{LOGO}",
                "MotdDescription": "§aHypixel Network §c[1.8-1.20]\n§bDROPPER v1.0 §7- §6NEW ARCADE LOBBY"
            },
            "TLSSniffing": {
                "RejectNonTLS": false
            },
            "Outbound": {
                "Type": ""
            }
        }
    ],
    "PrivateConfig": {
        "ListAPI": "http://whitelist.hln-network.xyz/isWhitelist.php",
        "Header": "HLN-Boost",
        "ContactName": "官方QQ售后群",
        "ContactLink": "666259678"
    },
    "Lists": {}
}
```

🔨 **配置说明**

- `Services`: 描述了代理服务的详细设置，包括名称、目标地址和端口、监听端口、流量控制、IP访问控制等。
- `Minecraft`: 针对Minecraft游戏的特定设置，包括主机名重写、在线人数限制、名称访问模式、Ping模式、MOTD Favicon和描述等。
- `TLSSniffing`: TLS嗅探设置，用于判断非TLS连接是否被拒绝。
- `PrivateConfig`: 私有配置，包括白名单API地址、自定义请求头、联系名称和联系链接。

❗️ **注意事项**
- 请根据实际需求修改配置文件中的各项设置。
- 详细的配置说明和使用方法，请参考项目文档或相关资源。

✨ **感谢使用**

感谢您选择使用NoDelay Minecraft服务器代理程序配置文件。希望它能帮助您优化和管理Minecraft游戏的网络连接。如果您有任何建议或意见，欢迎随时反馈给我们。祝您游戏愉快！

🌐 **参考链接**

- NoDelay GitHub 项目页：[https://github.com/Mengke15/NoDelay](https://github.com/Mengke15/NoDelay)
