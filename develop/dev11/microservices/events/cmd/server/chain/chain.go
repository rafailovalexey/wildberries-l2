package chain

import (
	"net/http"
)

func ChainHandlers(handlers ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for index := len(handlers) - 1; index >= 0; index-- {
			next = handlers[index](next)
		}

		return next
	}
}
