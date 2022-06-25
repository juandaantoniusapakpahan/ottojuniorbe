package middlewares

import "net/http"

type CustomMux struct {
	http.ServeMux //essentially a request router that compares incoming requests against a list of predefined URL paths and executes the associated handler for the path whenever a match is found.
	middlewares   []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTPP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}
