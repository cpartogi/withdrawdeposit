package response

// SwaggerGetBalance for

type SwaggerGetBalance struct {
	Base
	Data struct {
		Balance  int    `json:"balance,omitempty"`
		SellerId string `json:"seller_id,omitempty"`
	} `json:"data"`
}
