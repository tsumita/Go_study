package main

import (
	action "./handlers"

	"github.com/labstack/echo"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// Routing

	e.GET("/api/show/address_count", action.Domain_count)

	e.GET("/api/show/grade_all", action.Grade_mem)

	e.GET("/api/show/team_member_count", action.Team_mem_count)

	e.GET("/api/data/member_add", action.Add_mem)

	e.GET("/api/data/member_delete", action.Delete_mem)

	e.GET("/api/data/member_update", action.Update_mem)

	// Start server
	e.Logger.Fatal(e.Start(":1232"))
	// http://localhost:49155/
}
