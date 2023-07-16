package masking

var DefaultSensitiveKeys = []string{"password"}

var sensitiveKeys = DefaultSensitiveKeys

func SensitiveKeys() []string {
	return sensitiveKeys
}
