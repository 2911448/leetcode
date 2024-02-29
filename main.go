package main

type aa struct {
	age  int
	name string
}

func main() {
	mp1 := make(map[aa]int)
	a1 := aa{age: 1, name: "abc"}
	mp1[a1] = 100
	if v, ok := mp1[a1]; ok {
		println(v)
	}

	mp2 := make(map[*aa]int)
	a2 := &aa{age: 1, name: "abc"}
	mp2[a2] = 200
	if v, ok := mp2[a2]; ok {
		println(v)
	}
}
