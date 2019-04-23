package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// 登録した「商品」の保持用（現時点では永続化は無し）
	itemMap := sync.Map{} // ※管理するデータがタイプセーフではないが、永続化対応までの一時しのぎのため許容

	// 「商品」を登録
	e.POST("/item", func(c echo.Context) error {
		i := &item{}
		if err := c.Bind(i); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		if _, ok := itemMap.Load(i.ID); ok {
			return c.JSON(http.StatusBadRequest, nil)
		}
		itemMap.Store(i.ID, i)
		return c.JSON(http.StatusOK, nil)
	})

	// 「商品」一覧を返却
	e.GET("/item", func(c echo.Context) error {
		// FIXME:
		return c.JSON(http.StatusOK, nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// 「商品」を定義
type item struct {
	ID        string    `json:"id"`         // 商品ID
	Name      string    `json:"name"`       // 商品名
	Price     int       `json:"price"`      // 金額
	CreatedAt time.Time `json:"created_at"` // 商品登録日
}
