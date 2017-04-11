package qutils

import (
	"strings"
	"fmt"
)


// IsInput checks if a list of inputs (e.g. [one, two]) matches the suffix of item (two-one)
func IsInput(list []string, item string) bool {
	prefixItem := fmt.Sprintf("->%s", item)
	for _, i := range list {
		if item == i {
			return true
		}
		if strings.HasSuffix(prefixItem, i) {
			return true
		}
	}
	return false
}
