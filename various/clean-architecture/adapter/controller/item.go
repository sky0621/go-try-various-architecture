package controller

import (
	"clean-architecture/usecase/inputport"
	"github.com/labstack/echo"
)

type Item struct {
	input inputport.ItemInputPort
}

func NewUser(input inputport.ItemInputPort) *Item {
	return &Item{input: input}
}

func (i *Item) GetItemByID(c echo.Context) {
	itemID := c.Param("item_id")
	i.input.GetItemByID(itemID)
}