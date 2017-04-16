package tako

import "testing"

func Test_octopod_at(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		pod  octopod
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pod.at(tt.args.i); got != tt.want {
				t.Errorf("octopod.at() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_octopod_turnOn(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		pod  octopod
		args args
		want octopod
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pod.turnOn(tt.args.i); got != tt.want {
				t.Errorf("octopod.turnOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_octopod_turnOff(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		pod  octopod
		args args
		want octopod
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pod.turnOff(tt.args.i); got != tt.want {
				t.Errorf("octopod.turnOff() = %v, want %v", got, tt.want)
			}
		})
	}
}
