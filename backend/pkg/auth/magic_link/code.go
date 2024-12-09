package magic_link

import (
	"math/rand"
)

const charset = "QSTVXYZ012345678901234567890123456789"

func generateCode() string {
	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
