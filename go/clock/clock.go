package clock

import "fmt"

// Define the Clock type here.
type Clock int 

const (
  dayMinutes = 1440
  hourMinutes = 60
)

func normalize(m int) Clock{
  if m < 0 {
    return Clock(dayMinutes + (m % dayMinutes))
  } else if m >= dayMinutes {
    return Clock(m % dayMinutes)
  }
  return Clock(m)
}

func New(h, m int) Clock {
  return normalize(h * hourMinutes + m)
}

func (c Clock) Add(m int) Clock {
  return normalize(int(c) + m)
}

func (c Clock) Subtract(m int) Clock {
  return normalize(int(c) - m)
}

func (c Clock) String() string {
  return fmt.Sprintf("%02d:%02d", int(c) / hourMinutes, int(c) % hourMinutes)
}
