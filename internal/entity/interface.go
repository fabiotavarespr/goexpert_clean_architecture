package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	FindOrders() ([]*Order, error)
}
