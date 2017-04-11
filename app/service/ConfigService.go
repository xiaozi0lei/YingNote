package service

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/xiaozi0lei/YingNote/app/db"
	"github.com/xiaozi0lei/YingNote/app/info"
	. "github.com/xiaozi0lei/YingNote/app/lea"
	"gopkg.in/mgo.v2/bson"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 配置服务
// 只是全局的, 用户的配置没有
type ConfigService struct {
	adminUserId   string
	siteUrl       string
	adminUsername string
	// 全局的
	GlobalAllConfigs    map[string]interface{}
	GlobalStringConfigs map[string]string
	GlobalArrayConfigs  map[string][]string
	GlobalMapConfigs    map[string]map[string]string
	GlobalArrMapConfigs map[string][]map[string]string
}

// appStart 时 从数据库获取全局配置
func (c *ConfigService) InitGlobalConfigs() bool {
	c.GlobalAllConfigs = map[string]interface{}{}
	c.GlobalStringConfigs = map[string]string{}
	c.GlobalArrayConfigs = map[string][]string{}
	c.GlobalMapConfigs = map[string]map[string]string{}
	c.GlobalArrMapConfigs = map[string][]map[string]string{}

	c.adminUsername, _ = revel.Config.String("adminUsername")
	if c.adminUsername == "" {
		c.adminUsername = "admin"
	}
	c.siteUrl, _ = revel.Config.String("site.url")

	userInfo := userService.GetUserInfoByAny(c.adminUsername)
	if userInfo.UserId == "" {
		return false
	}
	c.adminUserId = userInfo.UserId.Hex()

	configs := []info.Config{}
	// db.ListByQ(db.Configs, bson.M{"UserId": userInfo.UserId}, &configs)
	db.ListByQ(db.Configs, bson.M{}, &configs)

	for _, config := range configs {
		if config.IsArr {
			c.GlobalArrayConfigs[config.Key] = config.ValueArr
			c.GlobalAllConfigs[config.Key] = config.ValueArr
		} else if config.IsMap {
			c.GlobalMapConfigs[config.Key] = config.ValueMap
			c.GlobalAllConfigs[config.Key] = config.ValueMap
		} else if config.IsArrMap {
			c.GlobalArrMapConfigs[config.Key] = config.ValueArrMap
			c.GlobalAllConfigs[config.Key] = config.ValueArrMap
		} else {
			c.GlobalStringConfigs[config.Key] = config.ValueStr
			c.GlobalAllConfigs[config.Key] = config.ValueStr
		}
	}

	// site URL
	if s, ok := c.GlobalStringConfigs["siteUrl"]; !ok || s != "" {
		c.GlobalStringConfigs["siteUrl"] = c.siteUrl
	}

	return true
}

func (c *ConfigService) GetSiteUrl() string {
	s := c.GetGlobalStringConfig("siteUrl")
	if s != "" {
		return s
	}

	return c.siteUrl
}
func (c *ConfigService) GetAdminUsername() string {
	return c.adminUsername
}
func (c *ConfigService) GetAdminUserId() string {
	return c.adminUserId
}

// 通用方法
func (c *ConfigService) updateGlobalConfig(userId, key string, value interface{}, isArr, isMap, isArrMap bool) bool {
	// 判断是否存在
	if _, ok := c.GlobalAllConfigs[key]; !ok {
		// 需要添加
		config := info.Config{ConfigId: bson.NewObjectId(),
			UserId:      bson.ObjectIdHex(userId), // 没用
			Key:         key,
			IsArr:       isArr,
			IsMap:       isMap,
			IsArrMap:    isArrMap,
			UpdatedTime: time.Now(),
		}
		if isArr {
			v, _ := value.([]string)
			config.ValueArr = v
			c.GlobalArrayConfigs[key] = v
		} else if isMap {
			v, _ := value.(map[string]string)
			config.ValueMap = v
			c.GlobalMapConfigs[key] = v
		} else if isArrMap {
			v, _ := value.([]map[string]string)
			config.ValueArrMap = v
			c.GlobalArrMapConfigs[key] = v
		} else {
			v, _ := value.(string)
			config.ValueStr = v
			c.GlobalStringConfigs[key] = v
		}
		return db.Insert(db.Configs, config)
	} else {
		i := bson.M{"UpdatedTime": time.Now()}
		c.GlobalAllConfigs[key] = value
		if isArr {
			v, _ := value.([]string)
			i["ValueArr"] = v
			c.GlobalArrayConfigs[key] = v
		} else if isMap {
			v, _ := value.(map[string]string)
			i["ValueMap"] = v
			c.GlobalMapConfigs[key] = v
		} else if isArrMap {
			v, _ := value.([]map[string]string)
			i["ValueArrMap"] = v
			c.GlobalArrMapConfigs[key] = v
		} else {
			v, _ := value.(string)
			i["ValueStr"] = v
			c.GlobalStringConfigs[key] = v
		}
		// return db.UpdateByQMap(db.Configs, bson.M{"UserId": bson.ObjectIdHex(userId), "Key": key}, i)
		return db.UpdateByQMap(db.Configs, bson.M{"Key": key}, i)
	}
}

// 更新用户配置
func (c *ConfigService) UpdateGlobalStringConfig(userId, key string, value string) bool {
	return c.updateGlobalConfig(userId, key, value, false, false, false)
}
func (c *ConfigService) UpdateGlobalArrayConfig(userId, key string, value []string) bool {
	return c.updateGlobalConfig(userId, key, value, true, false, false)
}
func (c *ConfigService) UpdateGlobalMapConfig(userId, key string, value map[string]string) bool {
	return c.updateGlobalConfig(userId, key, value, false, true, false)
}
func (c *ConfigService) UpdateGlobalArrMapConfig(userId, key string, value []map[string]string) bool {
	return c.updateGlobalConfig(userId, key, value, false, false, true)
}

// 获取全局配置, 博客平台使用
func (c *ConfigService) GetGlobalStringConfig(key string) string {
	return c.GlobalStringConfigs[key]
}
func (c *ConfigService) GetGlobalArrayConfig(key string) []string {
	arr := c.GlobalArrayConfigs[key]
	if arr == nil {
		return []string{}
	}
	return arr
}
func (c *ConfigService) GetGlobalMapConfig(key string) map[string]string {
	m := c.GlobalMapConfigs[key]
	if m == nil {
		return map[string]string{}
	}
	return m
}
func (c *ConfigService) GetGlobalArrMapConfig(key string) []map[string]string {
	m := c.GlobalArrMapConfigs[key]
	if m == nil {
		return []map[string]string{}
	}
	return m
}

func (c *ConfigService) IsOpenRegister() bool {
	return c.GetGlobalStringConfig("openRegister") != ""
}

//-------
// 修改共享笔记的配置
func (c *ConfigService) UpdateShareNoteConfig(registerSharedUserId string,
	registerSharedNotebookPerms, registerSharedNotePerms []int,
	registerSharedNotebookIds, registerSharedNoteIds, registerCopyNoteIds []string) (ok bool, msg string) {

	defer func() {
		if err := recover(); err != nil {
			ok = false
			msg = fmt.Sprint(err)
		}
	}()

	// 用户是否存在?
	if registerSharedUserId == "" {
		ok = true
		msg = "share userId is blank, So it share nothing to register"
		c.UpdateGlobalStringConfig(c.adminUserId, "registerSharedUserId", "")
		return
	} else {
		user := userService.GetUserInfo(registerSharedUserId)
		if user.UserId == "" {
			ok = false
			msg = "no such user: " + registerSharedUserId
			return
		} else {
			c.UpdateGlobalStringConfig(c.adminUserId, "registerSharedUserId", registerSharedUserId)
		}
	}

	notebooks := []map[string]string{}
	// 共享笔记本
	if len(registerSharedNotebookIds) > 0 {
		for i := 0; i < len(registerSharedNotebookIds); i++ {
			// 判断笔记本是否存在
			notebookId := registerSharedNotebookIds[i]
			if notebookId == "" {
				continue
			}
			notebook := notebookService.GetNotebook(notebookId, registerSharedUserId)
			if notebook.NotebookId == "" {
				ok = false
				msg = "The user has no such notebook: " + notebookId
				return
			} else {
				perm := "0"
				if registerSharedNotebookPerms[i] == 1 {
					perm = "1"
				}
				notebooks = append(notebooks, map[string]string{"notebookId": notebookId, "perm": perm})
			}
		}
	}
	c.UpdateGlobalArrMapConfig(c.adminUserId, "registerSharedNotebooks", notebooks)

	notes := []map[string]string{}
	// 共享笔记
	if len(registerSharedNoteIds) > 0 {
		for i := 0; i < len(registerSharedNoteIds); i++ {
			// 判断笔记本是否存在
			noteId := registerSharedNoteIds[i]
			if noteId == "" {
				continue
			}
			note := noteService.GetNote(noteId, registerSharedUserId)
			if note.NoteId == "" {
				ok = false
				msg = "The user has no such note: " + noteId
				return
			} else {
				perm := "0"
				if registerSharedNotePerms[i] == 1 {
					perm = "1"
				}
				notes = append(notes, map[string]string{"noteId": noteId, "perm": perm})
			}
		}
	}
	c.UpdateGlobalArrMapConfig(c.adminUserId, "registerSharedNotes", notes)

	// 复制
	noteIds := []string{}
	if len(registerCopyNoteIds) > 0 {
		for i := 0; i < len(registerCopyNoteIds); i++ {
			// 判断笔记本是否存在
			noteId := registerCopyNoteIds[i]
			if noteId == "" {
				continue
			}
			note := noteService.GetNote(noteId, registerSharedUserId)
			if note.NoteId == "" {
				ok = false
				msg = "The user has no such note: " + noteId
				return
			} else {
				noteIds = append(noteIds, noteId)
			}
		}
	}
	c.UpdateGlobalArrayConfig(c.adminUserId, "registerCopyNoteIds", noteIds)

	ok = true
	return
}

// 添加备份
func (c *ConfigService) AddBackup(path, remark string) bool {
	backups := c.GetGlobalArrMapConfig("backups") // [{}, {}]
	n := time.Now().Unix()
	nstr := fmt.Sprintf("%v", n)
	backups = append(backups, map[string]string{"createdTime": nstr, "path": path, "remark": remark})
	return c.UpdateGlobalArrMapConfig(c.adminUserId, "backups", backups)
}

func (c *ConfigService) getBackupDirname() string {
	n := time.Now()
	y, m, d := n.Date()
	return strconv.Itoa(y) + "_" + m.String() + "_" + strconv.Itoa(d) + "_" + fmt.Sprintf("%v", n.Unix())
}
func (c *ConfigService) Backup(remark string) (ok bool, msg string) {
	binPath := configService.GetGlobalStringConfig("mongodumpPath")
	config := revel.Config
	dbname, _ := config.String("db.dbname")
	host, _ := revel.Config.String("db.host")
	port, _ := revel.Config.String("db.port")
	username, _ := revel.Config.String("db.username")
	password, _ := revel.Config.String("db.password")
	// mongodump -h localhost -d leanote -o /root/mongodb_backup/leanote-9-22/ -u leanote -p nKFAkxKnWkEQy8Vv2LlM
	binPath = binPath + " -h " + host + " -d " + dbname + " --port " + port
	if username != "" {
		binPath += " -u " + username + " -p " + password
	}
	// 保存的路径
	dir := revel.BasePath + "/mongodb_backup/" + c.getBackupDirname()
	binPath += " -o " + dir
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		ok = false
		msg = fmt.Sprintf("%v", err)
		return
	}

	cmd := exec.Command("/bin/sh", "-c", binPath)
	Log(binPath)
	b, err := cmd.Output()
	if err != nil {
		msg = fmt.Sprintf("%v", err)
		ok = false
		Log("error:......")
		Log(string(b))
		return
	}
	ok = configService.AddBackup(dir, remark)
	return ok, msg
}

// 还原
func (c *ConfigService) Restore(createdTime string) (ok bool, msg string) {
	backups := c.GetGlobalArrMapConfig("backups") // [{}, {}]
	var i int
	var backup map[string]string
	for i, backup = range backups {
		if backup["createdTime"] == createdTime {
			break
		}
	}
	if i == len(backups) {
		return false, "Backup Not Found"
	}

	// 先备份当前
	ok, msg = c.Backup("Auto backup when restore from " + backup["createdTime"])
	if !ok {
		return
	}

	// mongorestore -h localhost -d leanote --directoryperdb /home/user1/gopackage/src/github.com/leanote/leanote/mongodb_backup/leanote_install_data/
	binPath := configService.GetGlobalStringConfig("mongorestorePath")
	config := revel.Config
	dbname, _ := config.String("db.dbname")
	host, _ := revel.Config.String("db.host")
	port, _ := revel.Config.String("db.port")
	username, _ := revel.Config.String("db.username")
	password, _ := revel.Config.String("db.password")
	// mongorestore -h localhost -d leanote -o /root/mongodb_backup/leanote-9-22/ -u leanote -p nKFAkxKnWkEQy8Vv2LlM
	binPath = binPath + " --drop -h " + host + " -d " + dbname + " --port " + port
	if username != "" {
		binPath += " -u " + username + " -p " + password
	}

	path := backup["path"] + "/" + dbname
	// 判断路径是否存在
	if !IsDirExists(path) {
		return false, path + " Is Not Exists"
	}

	binPath += " " + path

	cmd := exec.Command("/bin/sh", "-c", binPath)
	Log(binPath)
	b, err := cmd.Output()
	if err != nil {
		msg = fmt.Sprintf("%v", err)
		ok = false
		Log("error:......")
		Log(string(b))
		return
	}

	return true, ""
}
func (c *ConfigService) DeleteBackup(createdTime string) (bool, string) {
	backups := c.GetGlobalArrMapConfig("backups") // [{}, {}]
	var i int
	var backup map[string]string
	for i, backup = range backups {
		if backup["createdTime"] == createdTime {
			break
		}
	}
	if i == len(backups) {
		return false, "Backup Not Found"
	}

	// 删除文件夹之
	err := os.RemoveAll(backups[i]["path"])
	if err != nil {
		return false, fmt.Sprintf("%v", err)
	}

	// 删除之
	backups = append(backups[0:i], backups[i+1:]...)

	ok := c.UpdateGlobalArrMapConfig(c.adminUserId, "backups", backups)
	return ok, ""
}

func (c *ConfigService) UpdateBackupRemark(createdTime, remark string) (bool, string) {
	backups := c.GetGlobalArrMapConfig("backups") // [{}, {}]
	var i int
	var backup map[string]string
	for i, backup = range backups {
		if backup["createdTime"] == createdTime {
			break
		}
	}
	if i == len(backups) {
		return false, "Backup Not Found"
	}
	backup["remark"] = remark

	ok := c.UpdateGlobalArrMapConfig(c.adminUserId, "backups", backups)
	return ok, ""
}

// 得到备份
func (c *ConfigService) GetBackup(createdTime string) (map[string]string, bool) {
	backups := c.GetGlobalArrMapConfig("backups") // [{}, {}]
	var i int
	var backup map[string]string
	for i, backup = range backups {
		if backup["createdTime"] == createdTime {
			break
		}
	}
	if i == len(backups) {
		return map[string]string{}, false
	}
	return backup, true
}

//--------------
// sub domain
var defaultDomain string
var schema = "http://"
var port string

func init() {
	revel.OnAppStart(func() {
		/*
			不用配置的, 因为最终通过命令可以改, 而且有的使用nginx代理
			port  = strconv.Itoa(revel.HttpPort)
			if port != "80" {
				port = ":" + port
			} else {
				port = "";
			}
		*/

		siteUrl, _ := revel.Config.String("site.url") // 已包含:9000, http, 去掉成 yingnote.cn
		if strings.HasPrefix(siteUrl, "http://") {
			defaultDomain = siteUrl[len("http://"):]
		} else if strings.HasPrefix(siteUrl, "https://") {
			defaultDomain = siteUrl[len("https://"):]
			schema = "https://"
		}

		// port localhost:9000
		ports := strings.Split(defaultDomain, ":")
		if len(ports) == 2 {
			port = ports[1]
		}
		if port == "80" {
			port = ""
		} else {
			port = ":" + port
		}
	})
}

func (c *ConfigService) GetSchema() string {
	return schema
}

// 默认
func (c *ConfigService) GetDefaultDomain() string {
	return defaultDomain
}

// note
func (c *ConfigService) GetNoteDomain() string {
	return "/note"
}
func (c *ConfigService) GetNoteUrl() string {
	return c.GetNoteDomain()
}

// blog
func (c *ConfigService) GetBlogDomain() string {
	return "/blog"
}
func (c *ConfigService) GetBlogUrl() string {
	return c.GetBlogDomain()
}

// lea
func (c *ConfigService) GetLeaDomain() string {
	return "/lea"
}
func (c *ConfigService) GetLeaUrl() string {
	return schema + c.GetLeaDomain()
}

func (c *ConfigService) GetUserUrl(domain string) string {
	return schema + domain + port
}
func (c *ConfigService) GetUserSubUrl(subDomain string) string {
	return schema + subDomain + "." + c.GetDefaultDomain()
}

// 是否允许自定义域名
func (c *ConfigService) AllowCustomDomain() bool {
	return configService.GetGlobalStringConfig("allowCustomDomain") != ""
}

// 是否是好的自定义域名
func (c *ConfigService) IsGoodCustomDomain(domain string) bool {
	blacks := c.GetGlobalArrayConfig("blackCustomDomains")
	for _, black := range blacks {
		if strings.Contains(domain, black) {
			return false
		}
	}
	return true
}
func (c *ConfigService) IsGoodSubDomain(domain string) bool {
	blacks := c.GetGlobalArrayConfig("blackSubDomains")
	LogJ(blacks)
	for _, black := range blacks {
		if domain == black {
			return false
		}
	}
	return true
}

// 上传大小
func (c *ConfigService) GetUploadSize(key string) float64 {
	f, _ := strconv.ParseFloat(c.GetGlobalStringConfig(key), 64)
	return f
}
func (c *ConfigService) GetInt64(key string) int64 {
	f, _ := strconv.ParseInt(c.GetGlobalStringConfig(key), 10, 64)
	return f
}
func (c *ConfigService) GetInt32(key string) int32 {
	f, _ := strconv.ParseInt(c.GetGlobalStringConfig(key), 10, 32)
	return int32(f)
}
func (c *ConfigService) GetUploadSizeLimit() map[string]float64 {
	return map[string]float64{
		"uploadImageSize":    c.GetUploadSize("uploadImageSize"),
		"uploadBlogLogoSize": c.GetUploadSize("uploadBlogLogoSize"),
		"uploadAttachSize":   c.GetUploadSize("uploadAttachSize"),
		"uploadAvatarSize":   c.GetUploadSize("uploadAvatarSize"),
	}
}

// 为用户得到全局的配置
// NoteController调用
func (c *ConfigService) GetGlobalConfigForUser() map[string]interface{} {
	uploadSizeConfigs := c.GetUploadSizeLimit()
	config := map[string]interface{}{}
	for k, v := range uploadSizeConfigs {
		config[k] = v
	}
	return config
}

// 主页是否是管理员的博客页
func (c *ConfigService) HomePageIsAdminsBlog() bool {
	return c.GetGlobalStringConfig("homePage") == ""
}

func (c *ConfigService) GetVersion() string {
	return "2.3"
}
