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
		{"", args{64, 2}, &lxivFilter{make([]cell, 1), 1, 2}},
		{"", args{1 << 16, 2}, &lxivFilter{make([]cell, 1<<10), 1 << 10, 2}},
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

func Test_lxivFilter_Add_MayExist(t *testing.T) {
	const amount = 1000
	done := make(chan struct{}, amount)
	bWrong := make(chan struct{}, amount/1000)
	aWrong := make(chan struct{}, amount/1000)
	lf := NewDefault()
	// lf := NewWithEstimate(amount, 0.0001)
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

func genRandByteArrayX(len int, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genRandByteArray(len)
	}
}

// func Benchmark_genRandByteArray16(b *testing.B)  { genRandByteArrayX(16, b) }
// func Benchmark_genRandByteArray32(b *testing.B)  { genRandByteArrayX(32, b) }
// func Benchmark_genRandByteArray64(b *testing.B)  { genRandByteArrayX(64, b) }
// func Benchmark_genRandByteArray128(b *testing.B) { genRandByteArrayX(128, b) }
// func Benchmark_genRandByteArray256(b *testing.B) { genRandByteArrayX(256, b) }
// func Benchmark_genRandByteArray512(b *testing.B) { genRandByteArrayX(512, b) }

func Benchmark_lxivFilter_Add(b *testing.B) {
	const amount = 1000000
	const times = 100000
	lf := NewDefault()
	// lf := NewWithEstimate(amount, 0.0001)
	ss := make([][]byte, times)
	b.N = times
	for i := 0; i < b.N; i++ {
		ss[i] = genRandByteArray(32)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lf.Add(ss[i])
	}
}

func Benchmark_lxivFilter_MayExist_allMiss(b *testing.B) {
	const amount = 1000000
	const times = 100000
	lf := NewDefault()
	// lf := NewWithEstimate(amount, 0.0001)
	ss := make([][]byte, times)
	b.N = times
	for i := 0; i < b.N; i++ {
		ss[i] = genRandByteArray(32)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lf.MayExist(ss[i])
	}
}

func Benchmark_lxivFilter_MayExist_allHit(b *testing.B) {
	const amount = 1000000
	const times = 100000
	lf := NewDefault()
	// lf := NewWithEstimate(amount, 0.0001)
	ss := make([][]byte, times)
	b.N = times
	for i := 0; i < b.N; i++ {
		ss[i] = genRandByteArray(32)
	}
	for i := 0; i < b.N; i++ {
		lf.Add(ss[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lf.MayExist(ss[i])
	}
}
