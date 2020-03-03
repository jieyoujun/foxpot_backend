package utils

var (
	// SessionKey session key
	SessionKey string
	// Secret session secret
	Secret string
)

// LoadConfig ...
func LoadConfig() {
	SessionKey = "foxpot"
	Secret = "l2019i110k970i"
}
