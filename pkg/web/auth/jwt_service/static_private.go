package jwt_service

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lowl11/boost/data/errors"
)

func generate(object any, key []byte) (string, error) {
	objectInBytes, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	objectMap := make(map[string]any)
	if err = json.Unmarshal(objectInBytes, &objectMap); err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for claimsKey, claimsValue := range objectMap {
		claims[claimsKey] = claimsValue
	}

	stringToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

func parse(key []byte, token string, destination any) error {
	claimsMap, err := getMap(key, token)
	if err != nil {
		return err
	}

	claimsMapInBytes, err := json.Marshal(claimsMap)
	if err != nil {
		return err
	}

	return json.Unmarshal(claimsMapInBytes, &destination)
}

func parseMap(key []byte, token string) (map[string]any, error) {
	return getMap(key, token)
}

func getMap(key []byte, token string) (map[string]any, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("parse JWT error")
		}

		return key, nil
	})
	if err != nil {
		return nil, err
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}
