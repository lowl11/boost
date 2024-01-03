package jwt_service

import "github.com/lowl11/boost/internal/helpers/jwt_helper"

// GetToken generate token by given object
func (jwtService JWT) GetToken(object any) (string, error) {
	return jwt_helper.Generate(object, jwtService.key)
}

// Parse parse jwt string token & store it to given object
func (jwtService JWT) Parse(token string, object any) error {
	return jwt_helper.Parse(jwtService.key, token, object)
}

// GetMap parse jwt string token & returns map
func (jwtService JWT) GetMap(token string) (map[string]any, error) {
	return jwt_helper.ParseMap(jwtService.key, token)
}
