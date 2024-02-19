## NoDelay

📝 **本项目简介**  

该项目是一个基于ZBProxy魔改的Minecraft服务器代理程序。NoDelay是一个功能强大的代理工具，用于优化和管理Minecraft游戏的网络连接。

⚙️ **添加的配置文件结构**

该项目添加了私有配置，包括白名单API地址、自定义表头、联系名称和链接。

```json
{
    "PrivateConfig": {
        "ListAPI": "http://bind.jsip.fun/isWhitelist.php",
        "Header": "HLN-Boost",
        "ContactName": "官方QQ售后群",
        "ContactLink": "666259678"
    }
}
```

🔨 **配置说明**

- 如果你使用的是`master`版本，请保证你的API能通过Get形式传入playerName参数，例如：`https://example.com/isWhitelist.php?playerName=`，ListAPI中不要带有`?playerName=`,并且当playerName正确或查询到的情况下，返回playerName。
- 如果你使用的是`OtherVerifications`版本，请保证你的API返回的是同原版ZBProxy的List形式，例如：
  ```json
  {
      "Lists":{
          "Example":["player1","player2"]
      }
  }
  ```

❗️ **注意事项**

- 请根据实际需求修改配置文件中的各项设置。
- 详细的配置说明和使用方法，请参考项目文档或相关资源。

✨ **感谢使用**

感谢您选择使用NoDelay代理程序，希望它能帮助您优化和管理Minecraft游戏的网络连接。如果您有任何建议或意见，欢迎随时反馈给我们。祝您游戏愉快！

🌐 **参考链接**

- ZBProxy 原作者: [Layou233](https://github.com/Layou233)
- NoDelay 作者: [MKyiwuQwQ](https://github.com/Mengke15)
- NoDelay GitHub 项目页：[https://github.com/CubewhyMC/NoDelay-Proxy-Server](https://github.com/CubewhyMC/NoDelay-Proxy-Server)
