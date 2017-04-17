package lxivFilter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestManualNew(t *testing.T) {
	type args struct {
		size uint64
		k    int
	}
	tests := []struct {
		name string
		args args
		want *lxivFilter
	}{
		{"", args{64, 2}, &lxivFilter{make([]cell, 1), 1, 2}},
		{"", args{1 << 16, 2}, &lxivFilter{make([]cell, 1<<10), 1 << 10, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ManualNew(tt.args.size, tt.args.k); !reflect.DeepEqual(got, tt.want) {
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

func Test_lxivFilter_Add_MayExist(t *testing.T) {
	const amount = 1000000
	done := make(chan struct{}, amount)
	bWrong := make(chan struct{}, amount/1000)
	aWrong := make(chan struct{}, amount/1000)
	// lf := NewDefault()
	lf := New(amount, 0.0001)
	fmt.Printf("k=%d\n", lf.K())
	fmt.Printf("m/n=%d\n", lf.Size()/amount)
	type tt struct {
		input  []byte
		before bool
		after  bool
	}
	tests := make([]*tt, amount)
	for i := range tests {
		tests[i] = &tt{genRandByteArray(32), false, true}
	}

	for _, test := range tests {
		go func(test *tt) {
			if before := lf.MayExist(test.input); before != test.before {
				bWrong <- struct{}{}
			}
			lf.Add(test.input)
			if after := lf.MayExist(test.input); after != test.after {
				aWrong <- struct{}{}
			}
			done <- struct{}{}
		}(test)
	}
	bErrors := 0
	aErrors := 0
	go func() {
		for {
			select {
			case <-bWrong:
				bErrors++
			case <-aWrong:
				aErrors++
			}
		}
	}()
	for i := 0; i < amount; i++ {
		<-done
	}
	fmt.Printf("before errors number: %d\n", bErrors)
	fmt.Printf("after errors number: %d\n", aErrors)
}
