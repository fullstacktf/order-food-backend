package enums

type OrderStatus string

const (
	Ordered   OrderStatus = "ordered"
	Preparing OrderStatus = "preparing"
	Ready     OrderStatus = "ready"
	Complete  OrderStatus = "complete"
)
