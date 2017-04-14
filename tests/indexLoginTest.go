// home/login.html page interface testing

package tests

import (
	"github.com/revel/revel/testing"
)

type IndexLoginTest struct {
	testing.TestSuite
}

func (t *IndexLoginTest) Before() {
	println("Set up")
}

// ------ 登录页检查 ------
// 检查 views/home/login.html footer.html
func (t *IndexLoginTest) TestIndexLoginPage() {

	t.Get("/login")
	// 返回码是 200
	t.AssertOk()
	// 检查 title ，返回的模板中包含<title>YingNote</title>
	title_text := "<title>登录"
	t.AssertContainsRegex(title_text)

	username_label := "label.*用户名或Email"
	t.AssertContainsRegex(username_label)

	password_label := "label.*密码"
	t.AssertContainsRegex(password_label)

	findPassword_link := "a href.*忘记密码"
	t.AssertContainsRegex(findPassword_link)

	login_button := "button.*btn.*登录"
	t.AssertContainsRegex(login_button)

	registry_link := "a.*btn.*注册"
	t.AssertContainsRegex(registry_link)

	demo_button := "a.*btn.*体验"
	t.AssertContainsRegex(demo_button)
}

func (t *IndexLoginTest) After() {
	println("Tear down")
}
