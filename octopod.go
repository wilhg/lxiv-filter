package takoyaki

type octopod uint64

func (pod octopod) at(i uint8) bool {
	return pod&(1<<i) != 0
}

func (pod octopod) turnOn(i uint8) octopod {
	return pod | 1<<i
}

func (pod octopod) turnOff(i uint8) octopod {
	return pod ^ 1<<i
}
