package jwt_service

// GetToken generate token by given object
func (jwtService JWT) GetToken(object any) (string, error) {
	return generate(object, jwtService.key)
}

// Parse parse jwt string token & store it to given object
func (jwtService JWT) Parse(token string, object any) error {
	return parse(jwtService.key, token, object)
}

// GetMap parse jwt string token & returns map
func (jwtService JWT) GetMap(token string) (map[string]any, error) {
	return parseMap(jwtService.key, token)
}
