package bark

// 通知铃声
type Sound string

const (
	Alarm              Sound = "alarm"
	Anticipate         Sound = "anticipate"
	Bell               Sound = "bell"
	Birdsong           Sound = "birdsong"
	Bloom              Sound = "bloom"
	Calypso            Sound = "calypso"
	Chime              Sound = "chime"
	Choo               Sound = "choo"
	Descent            Sound = "descent"
	Electronic         Sound = "electronic"
	Fanfare            Sound = "fanfare"
	Glass              Sound = "glass"
	Gotosleep          Sound = "gotosleep"
	Healthnotification Sound = "healthnotification"
	Lorn               Sound = "horn"
	Ladder             Sound = "ladder"
	Mailsent           Sound = "mailsent"
	Minuet             Sound = "minuet"
	Multiwayinvitation Sound = "multiwayinvitation"
	Newmail            Sound = "newmail"
	Newsflash          Sound = "newsflash"
	Noir               Sound = "noir"
	Paymentsuccess     Sound = "paymentsuccess"
	Shake              Sound = "shake"
	Sherwoodforest     Sound = "sherwoodforest"
	Silence            Sound = "silence"
	Spell              Sound = "spell"
	Suspense           Sound = "suspense"
	Telegraph          Sound = "telegraph"
	Tiptoes            Sound = "tiptoes"
	Typewriters        Sound = "typewriters"
	Update             Sound = "update"
)

type Switch int

const (
	ON  Switch = 1 // 开启通知消息保存
	OFF Switch = 0 // 关闭通知消息保存
)

type Level string

const (
	Active        Level = "active"        // 默认值，系统会立即亮屏显示通知
	TimeSensitive Level = "timeSensitive" // 时效性通知，可在专注状态下显示通知
	Passive       Level = "passive"       // 仅将通知添加到通知列表，不会亮屏提醒
)

// Bark 数据类型
type Bark struct {
	ServerUrl string
	Title     string `json:"title"` // 通知标题 (必需要有）
	DeviceKey string `json:"device_key"`
	Category  string `json:"category"`
	Body      string `json:"body"` // 通知内容

	Level     Level  `json:"level,omitempty"`             // 时效性通知	'active', 'timeSensitive', or 'passive'
	Badge     int    `json:"badge,omitempty"`             //设置角标 optional
	AutoCopy  Switch `json:"automaticallyCopy,omitempty"` // 自动复制 optional
	Copy      string `json:"copy,omitempty"`              //只复制此值 optional
	Sound     Sound  `json:"sound,omitempty"`             // 推送铃声 optional
	Icon      string `json:"icon,omitempty"`              // 通知图标 optional
	Group     string `json:"group,omitempty"`             // 接受消息分组的组名
	IsArchive Switch `json:"isArchive,omitempty"`         // 是否保存通知消息 optional
	Jump2Url  string `json:"url,omitempty"`               // 通知 跳转Url optional

}

type BarkRes struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}
