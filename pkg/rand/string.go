package rand

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/utils/constants"
)

func String(n int) (string, error) {
	sb := strings.Builder{}
	sb.Grow(n)

	for range n {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(constants.LettersAndNumbers))))
		if err != nil {
			return "", err
		}

		sb.WriteByte(constants.LettersAndNumbers[num.Int64()])
	}

	return sb.String(), nil
}
