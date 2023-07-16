package masking

import "github.com/samber/lo"

type MaskedValue[V any] struct {
	Value       V
	MaskedValue string
}

type MaskedMap[V any] struct {
	m map[string]V
}

func (m MaskedMap[V]) UnmarshalJSON(b []byte) error {
	panic("not implemented") // TODO: Implement
}

func (m MaskedMap[V]) Get(key string) (V, bool) {
	v, ok := m.m[key]
	if !ok {
		return v, false
	}

	if lo.Contains(SensitiveKeys(), key) {
	}
	return v, true
}
