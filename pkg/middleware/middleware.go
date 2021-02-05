package middleware

import (
	"net/http"
)

type HandlerFunc func(next http.Handler) http.Handler
type Middleware struct {
	middlewareChain []HandlerFunc
}

func NewMiddleware() *Middleware {
	return &Middleware{middlewareChain: make([]HandlerFunc, 0)}
}
func (m *Middleware) Use(h ...HandlerFunc) *Middleware {
	m.middlewareChain = append(m.middlewareChain, h...)
	return m
}

func (m *Middleware) Add(h http.Handler) http.Handler {
	var mergedHandler = h
	for i := len(m.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = m.middlewareChain[i](mergedHandler)
	}
	return mergedHandler
}
