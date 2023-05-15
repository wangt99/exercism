package romannumerals

import (
	"errors"
	"sort"
)

var romanNum = map[int]string{
  1 : "I",
  4 : "IV",
  5 : "V",
  9 : "IX",
  10 : "X",
  40 : "XL",
  50 : "L",
  90 : "XC",
  100 : "C",
  400 : "CD",
  500 : "D",
  900 : "CM",
  1000 : "M",
}

func ToRomanNumeral(input int) (string, error) {
  if input <= 0 || input >= 4000 {
    return "", errors.New("invalid")
  }
  keys := []int{}
  for k := range romanNum {
    keys = append(keys, k);
  }
  ret := ""
  sort.Sort(sort.Reverse(sort.IntSlice(keys)))
  for _, k := range keys {
    for input >= k {
      ret += romanNum[k]
      input -= k
    }
  }
  return ret, nil
}
