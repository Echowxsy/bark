## Bark 的通知工具 🎉🎉🎉

<img src="https://wx3.sinaimg.cn/mw690/0060lm7Tly1g0nfnjjxbbj30sg0sg757.jpg" width=200px height=200px />

### 🍺用法：👇

#### SendMessage

```go
 bark := Bark{
  // 请求的URL (必需要有）
  ServerUrl: "https://api.day.app/",
  // Key (必需要有）
  DeviceKey: "xxxx",
  // 通知标题 (必需要有）
  Title: "title",
  // 通知内容
  Body: "body",
  // 推送铃声
  Sound: Alarm,
  // 是否保存通知消息
  IsArchive: ON,
  // 通知图标
  Icon: "https://s1.ax1x.com/2022/09/16/vzIC9K.png",
  // 接受消息分组的组名
  Group: "test",
  // 时效性通知
  Level:    Active,
  Jump2Url: "https://www.baidu.com",
  // 只复制 copy 参数到值
  Copy: "copyText",
  // 设置角标
  Badge: 1,
  // 自动复制 需要与 Copy 组合使用
  AutoCopy: OFF,
 }
 a, err := bark.SendMessage() // 配置好 bark 之后直接发送消息即可

 // 打印：
 // {"code":200,"message":"success","timestamp":1663581572} <nil>
```

### ⚠️注意：👇

创建 Bark 时，至少需要 ServerUrl、DeviceKey 和 Title 这三个字段，其他字段可以根据自己的需求进行添加

### Refer

[Bark 常见问题](https://day.app/2021/06/barkfaq/)

[API_V2](https://github.com/Finb/bark-server/blob/master/docs/API_V2.md)
