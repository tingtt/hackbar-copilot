package middleware

import (
	"fmt"
	"net/http"
)

func Wrap(h http.Handler) wrapper {
	return wrapper{h, make(map[contextKey]bool)}
}

type wrapper struct {
	http.Handler
	usedContextKeys map[contextKey]bool
}

type contextKey string

func (m *wrapper) Use(middleware func(http.Handler) http.Handler, usedContextKeys []contextKey) {
	m.Handler = middleware(m.Handler)
	for _, key := range usedContextKeys {
		if used, alreadyUsed := m.usedContextKeys[key]; alreadyUsed || used {
			panic(fmt.Sprintf("Duplicated context key: \"%s\"", key))
		}
		m.usedContextKeys[key] = true
	}
}
