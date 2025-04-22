package responces

type UserResponces struct {
	Status  int         `json:"status" bson:"status"`
	Success bool        `json:"success" bson:"success"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}
