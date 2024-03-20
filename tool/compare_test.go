package tool

import "testing"

func TestCompare(t *testing.T) {
	mp := map[string]string{}

	sl := []string{}
	for _, s := range sl {
		if _, ok := mp[s]; !ok {
			println(s)
		}
	}
}
