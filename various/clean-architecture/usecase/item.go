package usecase

import "clean-architecture/usecase/inputport"

type Item struct {
	// FIXME:
}

func NewItem() inputport.ItemInputPort {
	return &Item{}
}

func (i *Item) GetItemByID(id string) error {
	// FIXME:
	return nil
}
