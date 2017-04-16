package lxivFilter

type cell uint64

func (c cell) at(i uint8) bool {
	return c&(1<<i) != 0
}

func (c cell) turnOn(i uint8) cell {
	return c | 1<<i
}

func (c cell) turnOff(i uint8) cell {
	return c ^ 1<<i
}
