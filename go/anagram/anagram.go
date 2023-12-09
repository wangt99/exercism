package anagram

import (
	"sort"
	"strings"
)

func Anagram(subject, candidate string) bool {
  if strings.ToLower(subject) == strings.ToLower(candidate) {
    return false
  }
  s := []rune(strings.ToLower(subject))
  cmp := func (a, b int) bool { return s[a] < s[b] }
  sort.Slice(s, cmp)
  subject = string(s)

  s = []rune(strings.ToLower(candidate))
  sort.Slice(s, cmp)
  candidate = string(s)

  return subject == candidate
}

func Detect(subject string, candidates []string) []string {
  res := []string{}
  for _, cand := range candidates {
    if Anagram(subject, cand) {
      res = append(res, cand)
    }
  }
  return res
}
