package lxivFilter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		size uint64
		k    int
	}
	tests := []struct {
		name string
		args args
		want *lxivFilter
	}{
		{"", args{1, 1}, &lxivFilter{make([]cell, 1), 1, 1}},
		{"", args{2, 1}, &lxivFilter{make([]cell, 2), 2, 1}},
		{"", args{4, 1}, &lxivFilter{make([]cell, 4), 4, 1}},
		{"", args{(1 << 10), 1}, &lxivFilter{make([]cell, (1 << 10)), (1 << 10), 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.size, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lxivFilter_Reset(t *testing.T) {
	type fields struct {
		cells []cell
		size  uint64
		k     int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"", fields{make([]cell, 1), 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lf := &lxivFilter{
				cells: tt.fields.cells,
				size:  tt.fields.size,
				k:     tt.fields.k,
			}
			lf.Reset()
		})
	}
}

func Test_lxivFilter_Size(t *testing.T) {
	type fields struct {
		cells []cell
		size  uint64
		k     int
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{"", fields{make([]cell, 1), 1, 1}, 1},
		{"", fields{make([]cell, 2), 2, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lf := lxivFilter{
				cells: tt.fields.cells,
				size:  tt.fields.size,
				k:     tt.fields.k,
			}
			if got := lf.Size(); got != tt.want {
				t.Errorf("lxivFilter.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lxivFilter_K(t *testing.T) {
	type fields struct {
		cells []cell
		size  uint64
		k     int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"", fields{make([]cell, 1), 1, 3}, 3},
		{"", fields{make([]cell, 1), 1, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lf := lxivFilter{
				cells: tt.fields.cells,
				size:  tt.fields.size,
				k:     tt.fields.k,
			}
			if got := lf.K(); got != tt.want {
				t.Errorf("lxivFilter.K() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lxivFilter_Add_MayExist(t *testing.T) {
	lf := New(1<<21, 4)
	tests := []struct {
		input  []byte
		before bool
		after  bool
	}{
		{[]byte("rewer"), false, true},
		// {[]byte("1231"), false, true},
	}

	for _, test := range tests {
		if before := lf.MayExist(test.input); before != test.before {
			t.Errorf("before add, lf.MayExist(%q) = %v", test.input, before)
		}
		fmt.Println()
		lf.Add(test.input)
		fmt.Println()
		if after := lf.MayExist(test.input); after != test.after {
			t.Errorf("after add, lf.MayExist(%q) = %v", test.input, after)
		}
	}
}
