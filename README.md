# Golang勉強用

前にGoでAPI作った時のやつとか、Echoの使い方とか、GoのORMの使い方とか

- Framework: echo
- ORM: GORM

## echo

<参考文献>  
[echo](https://github.com/labstack/echo)  
[Echo documentation](https://echo.labstack.com/guide)  
[Go言語フレームワーク『Echo』の導入方法＆かんたんな使い方](http://vdeep.net/go-echo)  
[GoのechoってWebサーバーでサクッとRESTしちゃう](https://qiita.com/ezaki/items/62e806ae42828bb3567a)

---

インストール
```shell
go get github.com/labstack/echo
```

練習コード
```go
package main

import (
	"github.com/labstack/echo"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// Routing（通常Handler使う）
  	e.GET("/", c.String(http.StatusOK, "Hello, World!"))

	// Start server
	e.Logger.Fatal(e.Start(":1000"))
	// http://localhost:1000/
}
```
