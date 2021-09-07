package appctl

import (
	"blog/conf"
	"github.com/labstack/echo/v4"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/zxysilent/logs"
	"sync"
)

var one sync.Once
var account *officialaccount.OfficialAccount
// ViewWxGet 公众号服务地址
func ViewWxGet(ctx echo.Context) error {
	one.Do(func() {
		wc := NewWechat()
		memory := cache.NewMemory()
		cfg := &config.Config{
			AppID:     conf.App.Wechat.GzhAppid,
			AppSecret: conf.App.Wechat.GzhSecret,
			Token:     conf.App.Wechat.GzhToken,
			EncodingAESKey: conf.App.Wechat.GzhAESKey,
			Cache: memory,
		}
		account = wc.GetOfficialAccount(cfg)
	})

	// 传入request和responseWriter
	server := account.GetServer(ctx.Request(), ctx.Response())
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		switch msg.Content{
		case "0":
			return &message.Reply{MsgData: "http://blog.go2live.cn/archives", MsgType: message.MsgTypeLink}
		case "1":
			return &message.Reply{MsgData: "http://blog.go2live.cn/tags", MsgType: message.MsgTypeLink}
		case "2":
			return &message.Reply{MsgData: "http://blog.go2live.cn/", MsgType: message.MsgTypeLink}
		default:
			text = message.NewText("0: 获取归档\r\n1: 获取所有tag\r\n2: 获取最新文章\r\n")
		}
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		logs.Debug(err)
		return err
	}
	//发送回复的消息
	server.Send()
	return nil
}
