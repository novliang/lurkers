package response

type Response struct {
	Code    int         `json:"Code"`
	Data    interface{} `json:"Data"`
	Message string      `json:"Message"`
}

