package dto

type Response[T any] struct {
	Status  string `json:"status" swaggertype:"primitive,string"`
	Message string `json:"message" swaggertype:"primitive,string"`
	Data    T      `json:"data" swaggertype:"primitive,object"`
}
