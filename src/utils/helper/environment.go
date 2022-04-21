package helper

import "os"

func GetJWTToken() string {
	secret, ok := os.LookupEnv("JWT_SECRET_KEY")
	if !ok {
		panic("Could not find JWT_SECRET_KEY as environment variable.")
	}

	return secret
}
