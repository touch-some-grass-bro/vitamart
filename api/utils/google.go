package utils

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
	"github.com/touch-some-grass-bro/vitamart/models"
)

func GetGoogleProfile(accessToken, tokenType string) (*models.GoogleAuthResponse, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, err
	}

	var googleUser models.GoogleAuthResponse

	req.Header.Set("Authorization", tokenType+" "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, err
	}

	return &googleUser, nil
}

// func RefreshSpotifyToken(refresh_token string) (*db.CreateOrUpdateGoogleTokensParams, error) {
//   // TODO: Kya karu iska
// }

// func GetOrUpdateSpotifyToken(spotifyId string, queries *db.Queries, ctx context.Context, w http.ResponseWriter) (*db.SpotifyToken, error) {
// }

func RefreshGoogleToken(refresh_token string) (*db.CreateOrUpdateGoogleTokensParams, error) {
	reqBody := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refresh_token},
		"client_id":     {models.Config.Google.ClientID},
		"client_secret": {models.Config.Google.ClientSecret},
	}

	req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/token", strings.NewReader(reqBody.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  googleResp, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }

  defer googleResp.Body.Close()

  // Convert to string and print
  body, err := io.ReadAll(googleResp.Body)
  fmt.Println(string(body))

  return nil, nil

}

func GetOrUpdateGoogleToken(email string, queries *db.Queries, ctx context.Context, w http.ResponseWriter) (*db.GoogleToken, error) {
	token, err := queries.GetGoogleToken(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("No such user exists.")
	}
	if err != nil {
		return nil, err
	}

	if time.Now().UTC().After(token.ExpiresAt) {
		if token.RefreshToken == "" {
			log.Warn().Msg("Refresh token is empty.")
			return nil, errors.New("Refresh token is empty.")
		}

		updateParams, err := RefreshGoogleToken(token.RefreshToken)
		if err != nil {
			return nil, err
		}
		updateParams.UserEmail = email

		token, err = queries.CreateOrUpdateGoogleTokens(
			ctx,
			*updateParams,
		)

		if err != nil {
			return nil, err
		}
	}

	return &token, nil
}
