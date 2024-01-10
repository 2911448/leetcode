package day

func MinLength(s string) int {
	var stack []byte
	for _, a := range s {
		stack = append(stack, byte(a))
		length := len(stack)
		if length >= 2 {
			tmp := string(stack[length-2:])
			if tmp == "AB" || tmp == "CD" {
				stack = stack[:length-2]
			}
		}
	}
	return len(stack)
}
