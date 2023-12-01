package chain

import (
	"net/http"
)

func ChainHandlers(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for index := len(middlewares) - 1; index >= 0; index-- {
			next = middlewares[index](next)
		}

		return next
	}
}
