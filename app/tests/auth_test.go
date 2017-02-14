package tests

import (
	"github.com/xiaozi0lei/YingNote/app/db"
	"github.com/xiaozi0lei/YingNote/app/service"
	"testing"
	//	. "github.com/leanote/leanote/app/lea"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func init() {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
	service.InitService()
}

// 测试登录
func TestAuth(t *testing.T) {
	_, err := service.AuthS.Login("admin", "abc123")
	if err != nil {
		t.Error("Admin User Auth Error")
	}
}
