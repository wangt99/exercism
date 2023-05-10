package reverse

func Reverse(input string) string {
	ret := []rune(input)
	for i, j := 0, len(ret)-1; i < len(ret)/2; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}
