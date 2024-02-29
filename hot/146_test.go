package hot

import "testing"

func Test146(t *testing.T) {
	lru := Constructor(5)
	lru.Put(1, 1)
	lru.Put(2, 2)
	println(lru.Get(1))
}
