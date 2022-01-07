package jpush

type Notification struct {
	AiOpportunity bool                   `json:"ai_opportunity,omitempty"` // 如需采用“智能时机”策略下发通知，必须指定该字段。
	Alert         string                 `json:"alert,omitempty"`          // 通知的内容在各个平台上，都可能只有这一个最基本的属性 "alert"。
	Android       *AndroidNotification   `json:"android,omitempty"`        // Android通知
	Ios           *IosNotification       `json:"ios,omitempty"`            // iOS通知
	QuickApp      *QuickAppNotification  `json:"quick_app,omitempty"`      // 快应用通知
	WinPhone      *WinPhoneNotification  `json:"winphone,omitempty"`       // Windows Phone通知
	Voip          map[string]interface{} `json:"voip,omitempty"`           // iOS VOIP功能。
}

type AndroidNotification struct {
	Alert             interface{}            `json:"alert"`                        // 通知内容
	Title             interface{}            `json:"title,omitempty"`              // 通知标题
	BuilderID         int                    `json:"builder_id,omitempty"`         // 通知栏样式 ID
	ChannelId         string                 `json:"channel_id,omitempty"`         // android通知channel_id
	Priority          int                    `json:"priority,omitempty"`           // 通知栏展示优先级, 默认为 0，范围为 -2～2。
	Category          string                 `json:"category,omitempty"`           // 通知栏条目过滤或排序
	Style             int                    `json:"style,omitempty"`              // 通知栏样式类型, 默认为 0，还有 1，2，3 可选
	AlertType         int                    `json:"alert_type,omitempty"`         // 通知提醒方式, 可选范围为 -1～7
	BigText           string                 `json:"big_text,omitempty"`           // 大文本通知栏样式, 当 style = 1 时可用，内容会被通知栏以大文本的形式展示出来
	Inbox             interface{}            `json:"inbox,omitempty"`              // 文本条目通知栏样式, 当 style = 2 时可用， json 的每个 key 对应的 value 会被当作文本条目逐条展示
	BigPicPath        string                 `json:"big_pic_path,omitempty"`       // 大图片通知栏样式, 当 style = 3 时可用，可以是网络图片 url，或本地图片的 path
	Extras            map[string]interface{} `json:"extras,omitempty"`             // 扩展字段
	LargeIcon         string                 `json:"large_icon,omitempty"`         // 通知栏大图标
	SmallIconUri      string                 `json:"small_icon_uri,omitempty"`     // 通知栏小图标
	Intent            interface{}            `json:"intent,omitempty"`             // 指定跳转页面
	UriActivity       string                 `json:"uri_activity,omitempty"`       // 指定跳转页面, 该字段用于指定开发者想要打开的 activity，值为 activity 节点的 “android:name”属性值; 适配华为、小米、vivo厂商通道跳转；
	UriAction         string                 `json:"uri_action,omitempty"`         // 指定跳转页面, 该字段用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值; 适配 oppo、fcm跳转；
	BadgeAddNum       int                    `json:"badge_add_num,omitempty"`      // 角标数字，取值范围1-99
	BadgeClass        string                 `json:"badge_class,omitempty"`        // 桌面图标对应的应用入口Activity类， 配合badge_add_num使用，二者需要共存，缺少其一不可；
	Sound             string                 `json:"sound,omitempty"`              // 填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`    //定时展示开始时间（yyyy-MM-dd HH:mm:ss）
	ShowEndTime       string                 `json:"show_end_time,omitempty"`      //定时展示结束时间（yyyy-MM-dd HH:mm:ss）
	DisplayForeground string                 `json:"display_foreground,omitempty"` //APP在前台，通知是否展示, 值为 "1" 时，APP 在前台会弹出通知栏消息；值为 "0" 时，APP 在前台不会弹出通知栏消息。
}

type IosNotification struct {
	Alert             interface{}            `json:"alert"`                        // 通知内容
	Sound             interface{}            `json:"sound,omitempty"`              // 通知提示声音或警告通知
	Badge             interface{}            `json:"badge,omitempty"`              // 应用角标, 如果不填，表示不改变角标数字，否则把角标数字改为指定的数字；为 0 表示清除。
	ContentAvailable  bool                   `json:"content-available,omitempty"`  // 推送唤醒
	MutableContent    bool                   `json:"mutable-content,omitempty"`    // 通知扩展
	Category          string                 `json:"category,omitempty"`           // 通知类别, IOS 8 才支持。设置 APNs payload 中的 "category" 字段值
	Extras            map[string]interface{} `json:"extras,omitempty"`             // 扩展字段
	ThreadId          string                 `json:"thread-id,omitempty"`          // 通知分组, ios 的远程通知通过该属性来对通知进行分组，同一个 thread-id 的通知归为一组。
	InterruptionLevel string                 `json:"interruption-level,omitempty"` // 通知优先级和交付时间的中断级别, ios15 的通知级别，取值只能是active,critical,passive,timeSensitive中的一个。
}

type QuickAppNotification struct {
	Title  string                 `json:"title"`            // 通知标题, 必填字段，快应用推送通知的标题
	Alert  string                 `json:"alert"`            // 通知内容, 这里指定了，则会覆盖上级统一指定的 alert 信息。
	Page   string                 `json:"page"`             // 通知跳转页面, 必填字段，快应用通知跳转地址。
	Extras map[string]interface{} `json:"extras,omitempty"` // 扩展字段, 这里自定义 Key / value 信息，以供业务使用。
}

type WinPhoneNotification struct {
	Alert    string                 `json:"alert"`                // 通知内容, 必填字段，会填充到 toast 类型 text2 字段上。这里指定了，将会覆盖上级统一指定的 alert 信息；内容为空则不展示到通知栏。
	Title    string                 `json:"title,omitempty"`      // 通知标题, 会填充到 toast 类型 text1 字段上。
	OpenPage string                 `json:"_open_page,omitempty"` // 点击打开的页面名称, 点击打开的页面。会填充到推送信息的 param 字段上，表示由哪个 App 页面打开该通知。可不填，则由默认的首页打开。
	Extras   map[string]interface{} `json:"extras,omitempty"`     // 扩展字段, 这里自定义 Key / value 信息，以供业务使用。
}

// SetAlert 设置通知内容
func (n *Notification) SetAlert(alert string) {
	n.Alert = alert
}

// SetAiOpportunity 设置智能推送是否开启
func (n *Notification) SetAiOpportunity(use bool) {
	n.AiOpportunity = use
}

// SetAndroid 设置 Android 通知
func (n *Notification) SetAndroid(android *AndroidNotification) {
	n.Android = android
}

// SetIos 设置 iOS 通知
func (n *Notification) SetIos(ios *IosNotification) {
	n.Ios = ios
}

// SetQuickApp 设置 QuickApp 通知
func (n *Notification) SetQuickApp(quickApp *QuickAppNotification) {
	n.QuickApp = quickApp
}

// SetWinPhone 设置 WinPhone 通知
func (n *Notification) SetWinPhone(winPhone *WinPhoneNotification) {
	n.WinPhone = winPhone
}

// SetVoip 设置 Voip 通知
func (n *Notification) SetVoip(value map[string]interface{}) {
	n.Voip = value
}
