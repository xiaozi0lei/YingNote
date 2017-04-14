// home Demo interface testing

package tests

import (
	"github.com/revel/revel/testing"
)

type IndexDemoTest struct {
	testing.TestSuite
}

func (t *IndexDemoTest) Before() {
	println("Set up")
}

// ------ 登录页检查 ------
// 检查 views/home/login.html footer.html
func (t *IndexDemoTest) TestIndexDemoLink() {

	t.Get("/demo")
	// 返回码是 200
	t.AssertOk()
	// 检查 logo text
	logo_text := "a.*logo.*YingNote"
	t.AssertContainsRegex(logo_text)

	leftSwitcher_i := "<i.*隐藏左侧"
	t.AssertContainsRegex(leftSwitcher_i)

	searchNote_input := "input.*搜索笔记"
	t.AssertContainsRegex(searchNote_input)

	createNormalNote_span := "<span.*新建普通笔记"
	t.AssertContainsRegex(createNormalNote_span)

	createMarkdownNote_span := "<span.*新建Markdown笔记"
	t.AssertContainsRegex(createMarkdownNote_span)

	registry_link := "<a.*register.*立即注册"
	t.AssertContainsRegex(registry_link)

	writing_link := "<a.*writing.*写作模式"
	t.AssertContainsRegex(writing_link)

	myBlog_link := "<a.*blog.*\n.*我的博客"
	t.AssertContainsRegex(myBlog_link)

	profile_span := "span.*username.*\n.*demo"
	t.AssertContainsRegex(profile_span)

	notebook_span := "<span.*\n.*笔记本"
	t.AssertContainsRegex(notebook_span)

	addNoteBook_i := "<i.*添加笔记本"
	t.AssertContainsRegex(addNoteBook_i)

	tag_span := "<span.*\n.*标签"
	t.AssertContainsRegex(tag_span)

	share_span := "<span.*\n.*分享"
	t.AssertContainsRegex(share_span)
}

func (t *IndexDemoTest) After() {
	println("Tear down")
}
