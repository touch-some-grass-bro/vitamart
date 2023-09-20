package handlers

import (
	"io"
	"net/http"
	"strconv"

	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
	"github.com/touch-some-grass-bro/vitamart/models"
	"github.com/touch-some-grass-bro/vitamart/utils"
)

func AddItemHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{} = make(map[string]interface{})

		jwtTokenCookie, err := r.Cookie("token")

		if err != nil {
			resp["error"] = "No token found"
			utils.JSON(w, http.StatusUnauthorized, resp)
			return
		}

		email, errMsg, status := utils.GetUserEmailFromJWT(jwtTokenCookie.Value)
		if errMsg != "" {
			resp["error"] = errMsg
			utils.JSON(w, status, resp)
			return
		}

		// Parse the mutlipart form data
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}

		name := r.FormValue("name")
		description := r.FormValue("description")

		// Get the image as a file
		imageFile, _, err := r.FormFile("image")

		image, err := io.ReadAll(imageFile)
		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}

		priceStr := r.FormValue("price")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}

		// Get the form data
		data := models.AddItemsData{
			Name:        name,
			Description: description,
			Image:       image,
			Price:       price,
		}

		item, err := queries.CreateOrUpdateItem(r.Context(), db.CreateOrUpdateItemParams{
			Name:        data.Name,
			Description: data.Description,
			ImageBinary: data.Image,
			Price:       int32(data.Price),
			SellerEmail: email,
		})

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		resp["item"] = item
		resp["message"] = "Item added successfully"

		utils.JSON(w, http.StatusOK, resp)
	}
}

func GetItemsHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{} = make(map[string]interface{})
		jwtTokenCookie, err := r.Cookie("token")

		if err != nil {
			resp["error"] = "No token found"
			utils.JSON(w, http.StatusUnauthorized, resp)
			return
		}

		_, errMsg, status := utils.GetUserEmailFromJWT(jwtTokenCookie.Value)
		if errMsg != "" {
			resp["error"] = errMsg
			utils.JSON(w, status, resp)
			return
		}

		var offset int

		offsetStr := r.URL.Query().Get("offset")
		if offsetStr == "" {
			offsetStr = "0"
		} else {

			offset, err = strconv.Atoi(offsetStr)
			if err != nil {
				resp["error"] = err.Error()
				utils.JSON(w, http.StatusBadRequest, resp)
				return
			}
		}

		items, err := queries.GetItems(r.Context(), db.GetItemsParams{
			Limit:  10,
			Offset: int32(offset),
		})

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		utils.JSON(w, http.StatusOK, items)
	}
}
