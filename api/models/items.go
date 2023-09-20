package models

type AddItemsData struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Image []byte `json:"image"`
  Price float64 `json:"price"`
}
