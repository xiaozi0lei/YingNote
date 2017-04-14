// home/login.html page interface testing

package tests

import (
	"github.com/revel/revel/testing"
)

type IndexRegistryTest struct {
	testing.TestSuite
}

func (t *IndexRegistryTest) Before() {
	println("Set up")
}

// ------ 登录页检查 ------
// 检查 views/home/login.html footer.html
func (t *IndexRegistryTest) TestIndexRegistryPage() {

	t.Get("/register")
	// 返回码是 200
	t.AssertOk()
	// 检查 title
	title_text := "\\<title\\>注册"
	t.AssertContainsRegex(title_text)

	email_label := "label.*Email"
	t.AssertContainsRegex(email_label)

	password_label := "label.*密码"
	t.AssertContainsRegex(password_label)

	passwordAgain_label := "label.*确认密码"
	t.AssertContainsRegex(passwordAgain_label)

	registry_button := "button.*btn.*注册"
	t.AssertContainsRegex(registry_button)

	login_link := "a.*btn.*登录"
	t.AssertContainsRegex(login_link)

	home_link := "a.*index.*主页"
	t.AssertContainsRegex(home_link)
}

func (t *IndexRegistryTest) After() {
	println("Tear down")
}
