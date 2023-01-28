//go:generate mockgen -destination=mocks/mock_routeruser.go -package=mocks github.com/superlinkx/go-skeleton/middleware RouterUser

package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type RouterUser interface {
	Use(...func(http.Handler) http.Handler)
}

func Register(router RouterUser) {
	router.Use(middleware.Logger)
}
