package middleware

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func WithRequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

type sensitiveFilteredMap struct {
	sensitiveKeys []string
	json          map[string]any
}

func newSensitiveFilteredMap() *sensitiveFilteredMap {
	return &sensitiveFilteredMap{
		sensitiveKeys: []string{"password"},
		json:          map[string]any{},
	}
}

func (s *sensitiveFilteredMap) UnmarshalJSON(b []byte) (_ error) {
	var m map[string]any
	if err := json.Unmarshal(b, m); err != nil {
		return err
	}
	return nil
}

func (s sensitiveFilteredMap) masking(m map[string]any) map[string]any {
	lo.MapValues[string, any, any](m, func(value any, key string) any {
		if !s.isSensitive(key) {
			return value
		}
		if v, ok := value.(map[string]any); ok {
			return s.masking(v)
		}
		return s.maskingString()
	})

	return nil
}

func (s sensitiveFilteredMap) isSensitive(key string) bool {
	return lo.Contains(s.sensitiveKeys, key)
}

func (s sensitiveFilteredMap) maskingString() string {
	return "*****"
}
