package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
    for i := len(m) - 1; i >= 0; i-- {
        h = m[i](h)
    }
    return h
}
