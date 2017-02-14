package tests

import (
	"github.com/revel/revel"
	"github.com/xiaozi0lei/YingNote/app/db"
	"github.com/xiaozi0lei/YingNote/app/service"
	"testing"
	//	. "github.com/leanote/leanote/app/lea"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func init() {
	revel.Init("dev", "github.com/xiaozi0lei/YingNote", "/Users/life/Documents/Go/package_base/src")
	db.Init("mongodb://localhost:27017/leanote", "leanote")
	service.InitService()
	service.ConfigS.InitGlobalConfigs()
}

// 测试登录
func TestSendMail(t *testing.T) {
	ok, err := service.EmailS.SendEmail("life@leanote.com", "你好", "你好吗")
	t.Log(ok)
	t.Log(err)
}
