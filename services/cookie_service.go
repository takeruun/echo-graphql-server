package services

import (
	"app/helpers"
	"context"
	"net/http"
)

type CookieService interface {
	GetCookieValue(ctx context.Context, key string) (string, error)
	SetCookie(ctx context.Context, key string, value string) error
	DeleteCookie(ctx context.Context, key string) error
}

type cookieService struct {
}

func NewSessionService() CookieService {
	return &cookieService{}
}

func (service *cookieService) GetCookieValue(ctx context.Context, key string) (string, error) {
	httpContext := ctx.Value(helpers.HTTPKey("http")).(helpers.HTTP)

	cookie, err := httpContext.R.Cookie(key)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (service *cookieService) SetCookie(ctx context.Context, key string, value string) error {
	httpContext := ctx.Value(helpers.HTTPKey("http")).(helpers.HTTP)

	http.SetCookie(*httpContext.W, &http.Cookie{
		HttpOnly: true,
		MaxAge:   60 * 60 * 24,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Name:     key,
		Value:    value,
	})

	return nil
}

func (service *cookieService) DeleteCookie(ctx context.Context, key string) error {
	httpContext := ctx.Value(helpers.HTTPKey("http")).(helpers.HTTP)

	http.SetCookie(*httpContext.W, &http.Cookie{
		HttpOnly: true,
		MaxAge:   -1,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Name:     key,
	})

	return nil
}
