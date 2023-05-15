package isbn

func IsValidISBN(isbn string) bool {
  pos := 10
  sum := 0

  for i := 0; i < len(isbn); i++ {
    chr := isbn[i]
    if digit, ok := getDigit(chr); ok {
      sum += (digit * pos)
      pos--
    } else if pos == 1 && chr == 'X' {
      sum += 10
      pos--
    } else if chr != '-' {
      return false
    }
  }
  return pos == 0 && sum % 11 == 0
}

func getDigit(chr byte) (digit int, ok bool) {
  if chr < '0' || chr > '9' {
    return digit, ok
  }
  return int(chr - '0'), true
}

