package handlers

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "kood/social-network/docs"
)

func SwaggerHandler() http.HandlerFunc {
	return httpSwagger.WrapHandler
}
