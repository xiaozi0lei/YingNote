package tests

import (
	"github.com/xiaozi0lei/YingNote/app/db"
	"testing"
	//	. "github.com/leanote/leanote/app/lea"
	//	"github.com/leanote/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}
