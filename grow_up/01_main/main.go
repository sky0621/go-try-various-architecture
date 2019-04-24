package main

import (
	"net/http"
	"sync"

	"github.com/labstack/echo"
)

// 注意: プロダクション品質ではありません。
func main() {
	e := echo.New()

	// 登録した「商品」の保持用（現時点では永続化は無し）
	itemMap := sync.Map{} // ※管理するデータがタイプセーフではないが、永続化対応までの一時しのぎのため許容

	// 「商品」を登録
	e.POST("/item", func(c echo.Context) error {
		i := &item{} // バリデーションは後回し
		if err := c.Bind(i); err != nil {
			return sendResponse(c, http.StatusBadRequest)
		}
		if _, ok := itemMap.Load(i.ID); ok {
			return sendResponse(c, http.StatusBadRequest)
		}
		itemMap.Store(i.ID, i)
		return sendResponse(c, http.StatusOK)
	})

	// 「商品」一覧を返却
	e.GET("/item", func(c echo.Context) error {
		var res []*item
		itemMap.Range(func(k, v interface{}) bool {
			res = append(res, v.(*item))
			return true
		})
		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// 「商品」を定義
type item struct {
	ID        string `json:"id"`         // 商品ID
	Name      string `json:"name"`       // 商品名
	Price     int    `json:"price"`      // 金額
	CreatedAt int    `json:"created_at"` // 商品登録日(Unixタイムスタンプ)
}

func sendResponse(c echo.Context, code int) error {
	return c.JSON(code, struct {
		Message string `json:"message"`
	}{Message: http.StatusText(code)})
}
