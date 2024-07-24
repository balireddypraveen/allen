package constants

const (
	ApiPath    = "/allen-digital/api/"
	V1         = "v1"
	Deals      = "/deals"
	TestHealth = "/health"

	CreateDeal  = Deals
	CloseDeal   = Deals + "/:dealId"
	UpdateDeal  = Deals
	CreateOrder = "/orders"
)
