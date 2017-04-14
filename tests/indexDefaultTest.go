// home/index.html page interface testing

package tests

import (
	"github.com/revel/revel/testing"
	"net/url"
)

type IndexDefaultTest struct {
	testing.TestSuite
}

func (t *IndexDefaultTest) Before() {
	println("Set up")
}

// ------ 未登录时主页检查 ------
// 检查 views/home/header.html footer.html
func (t *IndexDefaultTest) TestIndexPageNonLogin() {

	t.Get("/")
	// 返回码是 200
	t.AssertOk()
	// 检查 title ，返回的模板中包含<title>YingNote</title>
	title_text := "\\<title\\>YingNote"
	t.AssertContainsRegex(title_text)

	english_link := "a data-lang.*en-us.*English"
	t.AssertContainsRegex(english_link)

	chinese_link := "a data-lang.*zh-cn.*简体中文"
	t.AssertContainsRegex(chinese_link)

	home_link := "a href.*index.*主页"
	t.AssertContainsRegex(home_link)

	desktopApp_link := "a.*yingNoteApp.*客户端"
	t.AssertContainsRegex(desktopApp_link)

	websiteURL_link := "a href.*yingnote.cn"
	t.AssertContainsRegex(websiteURL_link)

	registry_button := "a.*btn.*register.*注册"
	t.AssertContainsRegex(registry_button)

	login_button := "a.*btn.*login.*登录"
	t.AssertContainsRegex(login_button)

	demo_button := "a.*btn.*demo.*体验"
	t.AssertContainsRegex(demo_button)
}

// ------ 登录时主页元素检查 ------
// 检查 views/home/header.html footer.html
func (t *IndexDefaultTest) TestIndexPageLoggedIn() {

	data := url.Values{}
	data.Add("email", "demo@leanote.com")
	data.Add("pwd", "demo@leanote.com")

	t.PostForm("/doLogin", data)
	// 返回码是 200
	t.AssertOk()

	t.Get("/")
	// 返回码是 200
	t.AssertOk()
	// 检查 title ，返回的模板中包含<title>YingNote</title>
	title_text := "\\<title\\>YingNote"
	t.AssertContainsRegex(title_text)

	english_link := "a data-lang.*en-us.*English"
	t.AssertContainsRegex(english_link)

	chinese_link := "a data-lang.*zh-cn.*简体中文"
	t.AssertContainsRegex(chinese_link)

	home_link := "a href.*index.*主页"
	t.AssertContainsRegex(home_link)

	desktopApp_link := "a.*yingNoteApp.*客户端"
	t.AssertContainsRegex(desktopApp_link)

	websiteURL_link := "a href.*yingnote.cn"
	t.AssertContainsRegex(websiteURL_link)

	registry_button := "a.*btn.*register.*注册"
	t.AssertContainsRegex(registry_button)

	login_button := "登录"
	t.AssertNotContains(login_button)

	demo_button := "a.*btn.*demo.*体验"
	t.AssertContainsRegex(demo_button)
}

func (t *IndexDefaultTest) After() {
	println("Tear down")
}
