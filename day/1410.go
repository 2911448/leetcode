package day

func EntityParser(text string) string {
	var (
		res    = ""
		start  = byte('&')
		end    = byte(';')
		strMap = map[string]string{
			"&quot;":  "\\\"",
			"&apos;":  "'",
			"&amp;":   "&",
			"&gt;":    ">",
			"&lt;":    "<",
			"&frasl;": "/",
		}
	)
	chars := []byte(text)
	left, right := -1, 0
	fund := false
	for i, char := range chars {
		if char == start {
			left = i
		}
		if left >= 0 && char == end {
			if v, ok := strMap[text[left:i+1]]; ok {
				if fund {
					res += text[right+1:left] + v
				} else {
					res += text[0:left] + v
				}
				fund = true
				right = i
			} else {
				left = -1
			}
		}
	}
	if !fund {
		return text
	}
	return res + text[right+1:]
}
