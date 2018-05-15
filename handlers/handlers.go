package handlers

import (
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	null "gopkg.in/guregu/null.v3"
)

type Member struct {
	Id           int         `json:id`
	Name         string      `json:name`
	Grade_id     int         `json:grade_id`
	Mail_address string      `json:mail_address`
	Updatedat    null.String `json:updatedAt`
	Createdat    null.String `json:createdAt`
}

type Memproject struct {
	Id         int         `json:id`
	Member_id  int         `json:member_id`
	Project_id int         `json:project_id`
	Updatedat  null.String `json:updatedAt`
	Createdat  null.String `json:createdAt`
}

type Grade struct {
	Id   int    `json:id`
	Name string `json:name`
}

type Project struct {
	Id   int    `json:id`
	Name string `json:name`
}

func (Member) TableName() string {
	return "members_Tsumita"
}

func (Memproject) TableName() string {
	return "member_projects_Tsumita"
}

// DB接続
func dbconnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "****"
	HOST := "****"
	DBNAME := "****"

	dbconf := USER + ":" + PASS + "@tcp(" + HOST + ":3306)/" + DBNAME
	cnn, err := gorm.Open(DBMS, dbconf)
	if err != nil {
		panic(err.Error())
	}

	return cnn
}

func grade_to_id(grade string) int {
	db := dbconnect()
	gradeEx := Grade{}
	db.Where("name = ?", grade).First(&gradeEx)
	grade_id := gradeEx.Id

	return grade_id
}

func project_to_id(project string) int {
	db := dbconnect()
	projectEx := Project{}
	db.Where("name = ?", project).First(&projectEx)
	project_id := projectEx.Id

	return project_id
}

// ドメインをカウントする
func Domain_count(c echo.Context) error {
	db := dbconnect()
	memberEx := []Member{}

	domainsMap := map[string]map[string]int{}
	domainMap := map[string]int{}

	db.Find(&memberEx)

	var mailconst []string

	for _, record := range memberEx {
		mailconst = strings.Split(record.Mail_address, "@")
		if _, ok := domainMap[mailconst[1]]; ok {
			domainMap[mailconst[1]]++
		} else {
			domainMap[mailconst[1]] = 1
		}
	}
	domainsMap["domain_type"] = domainMap

	return c.JSON(http.StatusOK, domainsMap)
}

// 指定したgradeのmemberのnameを取得する
func Grade_mem(c echo.Context) error {
	db := dbconnect()
	memberEx := []Member{}

	grade_id := grade_to_id(c.QueryParam("grade"))

	nameMap := map[string][]string{}
	var nameList []string

	db.Where("grade_id = ?", grade_id).Find(&memberEx)

	for _, record := range memberEx {
		nameList = append(nameList, record.Name)
	}
	nameMap["name"] = nameList

	return c.JSON(http.StatusOK, nameMap)
}

// 指定したteamのmember数を取得する
func Team_mem_count(c echo.Context) error {
	db := dbconnect()
	projectEx := []Memproject{}

	project_id := project_to_id(c.QueryParam("team"))

	db.Where("project_id = ?", project_id).Find(&projectEx)
	pcountMap := map[string]int{"projectmembercount": len(projectEx)}

	return c.JSON(http.StatusOK, pcountMap)
}

// 新しいmember情報を追加する
func Add_mem(c echo.Context) error {
	db := dbconnect()
	memberEx := Member{}

	name := c.QueryParam("name")
	mailaddress := c.QueryParam("mail_address")
	grade_id := grade_to_id(c.QueryParam("grade"))
	project_id := project_to_id(c.QueryParam("project"))

	now := time.Now()

	db.Create(&Member{Name: name, Grade_id: grade_id, Mail_address: mailaddress, Createdat: null.NewString(now.String(), true)})
	db.Where("name = ?", name).First(&memberEx)

	db.Create(&Memproject{Member_id: memberEx.Id, Project_id: project_id, Createdat: null.NewString(now.String(), true)})

	return c.String(http.StatusOK, "Complete ADD!!")
}

// 指定したmemberの情報を削除
func Delete_mem(c echo.Context) error {
	db := dbconnect()
	memberEx := Member{}
	projectEx := Memproject{}

	name := c.QueryParam("name")

	db.Where("name = ?", name).Find(&memberEx)
	db.Delete(&memberEx)

	db.Where("member_id = ?", memberEx.Id).Find(&projectEx)
	db.Delete(&projectEx)

	return c.String(http.StatusOK, "Complete DELETE!!")
}

// 指定したmemberの情報を更新する
func Update_mem(c echo.Context) error {
	db := dbconnect()
	memberEx := Member{}
	memberExBefore := Member{}
	projectExBefore := Memproject{}

	id, _ := strconv.Atoi(c.QueryParam("new_id"))
	name := c.QueryParam("name")
	grade_id := grade_to_id(c.QueryParam("new_grade"))
	mailaddress := c.QueryParam("new_mailaddress")
	project_id := project_to_id(c.QueryParam("project"))

	db.Where("name = ?", name).First(&memberEx)
	memberExBefore.Id = memberEx.Id
	memberExAfter := memberExBefore
	db.First(&memberExAfter)

	projectExBefore.Id = project_id
	projectExAfter := projectExBefore
	db.First(&projectExAfter)

	memberExAfter.Id = id
	memberExAfter.Grade_id = grade_id
	memberExAfter.Mail_address = mailaddress
	projectExAfter.Project_id = project_id
	/*
		if id != 0 {
			memberExAfter.Id = id
		}
		if grade_id != 0 {
			memberExAfter.Grade_id = grade_id
		}
		if mailaddress != "" {
			memberExAfter.Mail_address = mailaddress
		}
		if project_id != 0 {
			projectExAfter.Project_id = project_id
		}
	*/
	now := time.Now()
	memberExAfter.Updatedat = null.NewString(now.String(), true)
	db.Model(&memberExBefore).Update(&memberExAfter)

	projectExAfter.Updatedat = null.NewString(now.String(), true)
	db.Model(&projectExBefore).Update(&projectExAfter)

	return c.String(http.StatusOK, "Complete UPDATE!!")
}
