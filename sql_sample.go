package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:<pass>@tcp(<host>:<port>)/<dbname>")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT name FROM members_Tsumita WHERE grade_id = 1")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	// 名前が入った配列を"name"をkeyに呼び出せる辞書
	var namemap map[string][]string
	namemap = make(map[string][]string)

	// 名前が入るリスト
	var namelist []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err.Error())
		}
		namelist = append(namelist, name)
	}

	namemap[columns[0]] = namelist
	fmt.Println(namemap)

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println("Finish!")

}
