package test

import (
	"net/http"

	"github.com/keratin/authn-server/app/data"
	"github.com/keratin/authn-server/app/models"
	"github.com/keratin/authn-server/app/tokens/sessions"
	"github.com/keratin/authn-server/conf"
	jose "gopkg.in/square/go-jose.v2"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

func CreateSession(tokenStore data.RefreshTokenStore, cfg *conf.Config, accountID int) *http.Cookie {
	sessionToken, err := sessions.New(tokenStore, cfg, accountID, cfg.ApplicationDomains[0].String())
	if err != nil {
		panic(err)
	}

	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.HS256, Key: cfg.SessionSigningKey},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		panic(err)
	}
	sessionString, err := jwt.Signed(signer).Claims(sessionToken).CompactSerialize()
	if err != nil {
		panic(err)
	}

	return &http.Cookie{
		Name:  cfg.SessionCookieName,
		Value: sessionString,
	}
}

func RevokeSession(store data.RefreshTokenStore, cfg *conf.Config, session *http.Cookie) {
	claims, err := sessions.Parse(session.Value, cfg)
	if err != nil {
		panic(err)
	}
	err = store.Revoke(models.RefreshToken(claims.Subject))
	if err != nil {
		panic(err)
	}
}
