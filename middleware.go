package absol

import "net/http"

type Middleware func(http.Handler) http.Handler

func compose(a Middleware, b Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		return a(b(handler))
	}
}
