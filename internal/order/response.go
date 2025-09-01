package order
type CreateOrderResponse struct {
	OrderID     int64   `json:"orderId"`
	TotalAmount float64 `json:"totalAmount"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
}