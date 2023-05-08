package luhn

func Valid(id string) bool {
	total := 0
	pos := 0
	for i := len(id) - 1; i > -1; i-- {
		if id[i] == ' ' {
			continue
		}
		if id[i] < '0' || id[i] > '9' {
			return false
		}
		x := int(id[i] - '0')
		if pos%2 == 0 {
			total += x
		} else {
			switch x {
			case 1, 2, 3, 4:
				total += 2 * x
			case 5, 6, 7, 8:
				total += (2 * x) - 9
			case 9:
				total += x
			}
		}
		pos++
	}
	return pos > 1 && total%10 == 0
}
