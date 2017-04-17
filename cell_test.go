package lxivFilter

import "testing"

func Test_cell_at(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		c    cell
		args args
		want bool
	}{
		{"", 0, args{0}, false},
		{"", 0, args{1}, false},
		{"", 0, args{63}, false},
		{"", 1 << 0, args{0}, true},
		{"", 1 << 1, args{1}, true},
		{"", 1 << 12, args{12}, true},
		{"", 1 << 63, args{63}, true},
		{"", 1 << 63, args{63}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.at(tt.args.i); got != tt.want {
				t.Errorf("cell.at() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cell_turnOn(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		c    cell
		args args
		want cell
	}{
		{"", cell(0), args{0}, cell(1)},
		{"", cell(0), args{1}, cell(1 << 1)},
		{"", cell(0), args{2}, cell(1 << 2)},
		{"", cell(0), args{63}, cell(1 << 63)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.turnOn(tt.args.i); got != tt.want {
				t.Errorf("cell.turnOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cell_turnOff(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		c    cell
		args args
		want cell
	}{
		{"", cell(1), args{0}, cell(0)},
		{"", cell(2), args{1}, cell(0)},
		{"", cell(4), args{2}, cell(0)},
		{"", cell(1 << 63), args{63}, cell(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.turnOff(tt.args.i); got != tt.want {
				t.Errorf("cell.turnOff() = %v, want %v", got, tt.want)
			}
		})
	}
}
