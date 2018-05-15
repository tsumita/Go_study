package main

import (
	"fmt"
	"strings"
	"time"

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

type Project struct {
	Id         int         `json:id`
	Member_id  int         `json:member_id`
	Project_id int         `json:project_id`
	Updatedat  null.String `json:updatedAt`
	Createdat  null.String `json:createdAt`
}

func (Member) TableName() string {
	return "members_Tsumita"
}

func (Project) TableName() string {
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
	//cnn.SingularTable(true)

	return cnn
}

// ドメインをカウントする
func domain_count() {
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

	fmt.Println(domainsMap)
}

// 指定したgradeのmemberのnameを取得する
func grade_mem(grade int) {
	db := dbconnect()
	memberEx := []Member{}

	nameMap := map[string][]string{}
	var nameList []string

	db.Where("grade_id = ?", grade).Find(&memberEx)

	for _, record := range memberEx {
		nameList = append(nameList, record.Name)
	}
	nameMap["name"] = nameList

	fmt.Println(nameMap)
}

// 指定したteamのmember数を取得する
func team_mem_count(pid int) {
	db := dbconnect()
	projectEx := []Project{}

	db.Where("project_id = ?", pid).Find(&projectEx)
	pcountMap := map[string]int{"projectmembercount": len(projectEx)}

	fmt.Println(pcountMap)
}

// 新しいmember情報を追加する
func add_mem(name string, grade int, mailaddress string) {
	db := dbconnect()

	db.Create(&Member{Name: name, Grade_id: grade, Mail_address: mailaddress})
}

// 指定したmemberの情報を削除
func delete_mem(name string) {
	db := dbconnect()
	memberEx := Member{}

	db.Where("name = ?", name).Find(&memberEx)
	db.Delete(&memberEx)
}

// 指定したmemberの情報を更新する
func update_mem(id int, name string, grade_id int, mailaddress string) {
	db := dbconnect()
	memberExBefore := Member{}

	db.Where("name = ?", name).Find(&memberExBefore)
	memberExAfter := memberExBefore

	if id != 0 {
		memberExAfter.Id = id
	}
	if grade_id != 0 {
		memberExAfter.Grade_id = grade_id
	}
	if mailaddress != "" {
		memberExAfter.Mail_address = mailaddress
	}

	now := time.Now()
	memberExAfter.Updatedat = null.NewString(now.String(), true)
	db.Model(&memberExBefore).Update(&memberExAfter)
}

func main() {
}
