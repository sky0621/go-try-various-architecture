package main

import (
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}

// 「商品」を定義
type item struct {
	ID        string    `json:"id"`         // 商品ID
	Name      string    `json:"name"`       // 商品名
	Price     int       `json:"price"`      // 金額
	CreatedAt time.Time `json:"created_at"` // 商品登録日
}
