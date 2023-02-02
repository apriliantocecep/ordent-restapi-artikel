package config

import helper "github.com/apriliantocecep/ordent-restapi-artikel/app/Helper"

func NewJwtWrapper(config *Config) helper.JwtWrapper {
	return helper.JwtWrapper{
		SecretKey:       config.JWT_SECRET_KEY,
		Issuer:          config.JWT_ISSUER,
		ExpirationHours: config.JWT_EXPIRATION_HOURS,
	}
}
