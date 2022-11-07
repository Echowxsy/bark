package bark

import (
	"fmt"
	"testing"
)

func TestBark_SendMessage(t *testing.T) {
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
	fmt.Println(a, err)
}
