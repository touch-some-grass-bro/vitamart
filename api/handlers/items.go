package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
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

		fmt.Println(email)

		// TODO: Add item to database
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

		email, errMsg, status := utils.GetUserEmailFromJWT(jwtTokenCookie.Value)
		if errMsg != "" {
			resp["error"] = errMsg
			utils.JSON(w, status, resp)
			return
		}

		fmt.Println(email)

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
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
