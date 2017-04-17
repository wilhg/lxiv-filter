package lxivFilter

type cell uint64

// i should less than 64
// coz it's the inner func, no need to check, but be caurful
func (c cell) at(i uint8) bool {
	return c&(1<<i) != 0
}

func (c cell) turnOn(i uint8) cell {
	return c | 1<<i
}

func (c cell) turnOff(i uint8) cell {
	return c ^ 1<<i
}
