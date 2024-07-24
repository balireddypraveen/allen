package constants

const (
	ApiPath    = "/allen-digital/api/"
	V1         = "v1"
	OnCallV1   = "oncall/v1"
	Allen    = "/allen"
	TestHealth = "/health"

	CreateOrder = Allen + "/orders"
	GetOrders   = Allen + "/orders/query"
	CancelOrder = Allen + "/orders/:orderId"
)
