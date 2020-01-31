package gcd

func Gcd(a, b uint) uint {
	if a == 0 {
		return 0
	}
	for {
		c := a % b
		if c == 0 {
			return b
		}
		b, a = c, b
	}
}

func GcdSlice(s []uint) uint {
	switch len(s) {
	case 0:
		return 0
	case 1:
		return s[0]
	case 2:
		return Gcd(s[0], s[1])
	}
	car := s[0]
	cdr := s[1:]
	for _, b := range cdr {
		car = Gcd(car, b)
	}
	return car
}
