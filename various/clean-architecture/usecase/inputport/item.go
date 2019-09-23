package inputport

// adapter層のcontrollerから呼ばれる。
type ItemInputPort interface {
	GetItemByID(id string) error
}
