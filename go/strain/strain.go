package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var ret Ints = nil
	for _, num := range i {
		if filter(num) {
			ret = append(ret, num)
		}
	}
	return ret
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var ret Ints = nil
	for _, num := range i {
		if !filter(num) {
			ret = append(ret, num)
		}
	}
	return ret
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	ret := Lists{}
	for _, v := range l {
		if filter(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s Strings) Keep(filter func(string) bool) Strings {
	ret := Strings{}
	for _, v := range s {
		if filter(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
