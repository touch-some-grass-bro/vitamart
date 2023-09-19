package models

import "image"

type AddItemsBody struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Image image.Image `json:"image"`
  Price float64 `json:"price"`
}
