package rotationalcipher


func RotationalCipher(plain string, shiftKey int) string {
  ret := []rune{}
  for _, letter := range plain {
    base := 0
    switch {
      case letter >= 'A' && letter <= 'Z':
        base = int('A')
      case letter >= 'a' && letter <= 'z':
        base = int('a')
    }
    if base != 0 {
      letter = rune(((int(letter) - base + shiftKey) % 26) + base)
    }
    ret = append(ret, letter)
  }
  return string(ret)
}
