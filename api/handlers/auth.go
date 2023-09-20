package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
	"github.com/touch-some-grass-bro/vitamart/models"
	"github.com/touch-some-grass-bro/vitamart/utils"
	"golang.org/x/oauth2"
)

func GetAuthURLHandler(oauthConf *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{} = make(map[string]interface{})
		base_url := oauthConf.AuthCodeURL(models.Config.API.OAuthState)
		URL, err := url.Parse(base_url)
		if err != nil {
			log.Error().Msg(err.Error())
		}

		// Add the parameters of login_hint and hd to the URL
		parameters := URL.Query()
		parameters.Add("login_hint", "martin.garrix1996@vitstudent.ac.in")
		parameters.Add("hd", "vitstudent.ac.in")
		URL.RawQuery = parameters.Encode()
		base_url = URL.String()

		resp["url"] = base_url
		utils.JSON(w, http.StatusOK, resp)
	}
}

func CallbackHandler(queries *db.Queries, oauthConf *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{} = make(map[string]interface{})
		state := r.FormValue("state")
		if state != models.Config.API.OAuthState {
			resp["error"] = "Invalid state"
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}
		code := r.FormValue("code")
		googleTokens, err := oauthConf.Exchange(r.Context(),
			code,
			oauth2.SetAuthURLParam("client_id", oauthConf.ClientID),
			oauth2.SetAuthURLParam("client_secret", oauthConf.ClientSecret),
		)
		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		googleUser, err := utils.GetGoogleProfile(googleTokens.AccessToken, googleTokens.TokenType)

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		user, err := queries.CreateOrUpdateUser(r.Context(), db.CreateOrUpdateUserParams{
			Email:             googleUser.Email,
			Name:              googleUser.Name,
			JoinYear:          utils.GetUserJoinYear(googleUser.Email),
			ProfilePictureUrl: googleUser.Picture,
		})

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		now := time.Now().UTC()

		token, err := queries.CreateOrUpdateGoogleTokens(r.Context(), db.CreateOrUpdateGoogleTokensParams{
			UserEmail:    user.Email,
			CreatedAt:    now,
			RefreshToken: googleTokens.RefreshToken,
			AccessToken:  googleTokens.AccessToken,
			ExpiresAt:    googleTokens.Expiry.UTC(),
			TokenType:    googleTokens.TokenType,
		})

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		err = utils.SetJWTOnCookie(user.Email, token.ExpiresAt, now, w)

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		if user.Hostel.Valid {
			http.Redirect(w, r, models.Config.API.FrontendUrl+"/buy", http.StatusFound)
		} else {
			http.Redirect(w, r, models.Config.API.FrontendUrl+"/gender", http.StatusFound)
		}
	}
}

func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var resp map[string]interface{} = make(map[string]interface{})

		err := utils.SetJWTOnCookie("69420", time.Now().Add(time.Duration(5)), time.Now(), w)
		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		http.Redirect(w, r, models.Config.API.FrontendUrl + "/login", http.StatusFound)
	}
}

func IsAuthenticatedHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{} = make(map[string]interface{})
		jwtToken, err := r.Cookie("token")
		var respStatus int
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				resp["error"] = "No token provided"
				resp["is_authenticated"] = false
				respStatus = http.StatusUnauthorized
			} else {
				resp["error"] = err.Error()
				respStatus = http.StatusInternalServerError
			}
			utils.JSON(w, respStatus, resp)
			return
		}

		userEmail, errMsg, status := utils.GetUserEmailFromJWT(jwtToken.Value)

		if errMsg != "" {
			resp["error"] = errMsg
			utils.JSON(w, status, resp)
			return
		}

		_, err = utils.GetOrUpdateGoogleToken(userEmail, queries, r.Context(), w)
		if err != nil {
			resp["error"] = err.Error()
			resp["is_authenticated"] = false
			utils.JSON(w, http.StatusUnauthorized, resp)
			return
		}

		user, err := queries.GetUser(r.Context(), userEmail)
		if err != nil {
			resp["error"] = err.Error()
			resp["is_authenticated"] = false
			utils.JSON(w, http.StatusInternalServerError, resp)
			return
		}

		resp["is_authenticated"] = true
		resp["user"] = user
		utils.JSON(w, http.StatusOK, resp)
	}
}
