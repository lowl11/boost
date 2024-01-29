package middlewares

import (
	"encoding/base64"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/pkg/system/types"
	"net/http"
	"strings"
)

func BasicAuth(username, password string) types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		authorizationToken := ctx.Authorization()

		parsedToken, err := base64.StdEncoding.DecodeString(authorizationToken)
		if err != nil {
			return ctx.Error(errorBasicAuthParseHeader(err))
		}

		parsedUsername, parsedPassword, found := strings.Cut(types.BytesToString(parsedToken), ":")
		if !found {
			return ctx.Error(errorBasicAuthParseToken())
		}

		if parsedUsername != username || parsedPassword != password {
			return ctx.Error(errorBasicAuthCredentials())
		}

		if err = ctx.Next(); err != nil {
			return ctx.Error(err)
		}

		return nil
	}
}

func errorBasicAuthParseHeader(err error) error {
	return errors.New(types.StringFromError(err, "Parse header error")).
		SetType("ErrorBasicAuthParseHeader").
		SetHttpCode(http.StatusUnauthorized)
}

func errorBasicAuthParseToken() error {
	return errors.New("Parse token error").
		SetType("ErrorBasicAuthParseToken").
		SetHttpCode(http.StatusUnauthorized)
}

func errorBasicAuthCredentials() error {
	return errors.New("Wrong credentials").
		SetType("ErrorBasicAuthCredentials").
		SetHttpCode(http.StatusUnauthorized)
}
