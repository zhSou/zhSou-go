package highlight

import (
	"fmt"
	"strings"
)

func HighLight(source string, replaceWords []string, head string, tail string) string {
	s := source
	for _, word := range replaceWords {
		s = strings.ReplaceAll(s, word, fmt.Sprintf("%s%s%s", head, word, tail))
	}
	return s
}
