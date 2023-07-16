package pointer

func String(v string) *string {
	return &v
}

func Pointer[T any](v T) *T {
	return &v
}
