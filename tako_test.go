package tako

import (
	"reflect"
	"testing"
)

func TestNewDefault(t *testing.T) {
	tests := []struct {
		name string
		want *tako
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		size uint32
		keys int
	}
	tests := []struct {
		name string
		args args
		want *tako
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.size, tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tako_Reset(t *testing.T) {
	type fields struct {
		podMap []octopod
		size   uint32
		keys   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &tako{
				podMap: tt.fields.podMap,
				size:   tt.fields.size,
				keys:   tt.fields.keys,
			}
			tk.Reset()
		})
	}
}

func Test_tako_MayExist(t *testing.T) {
	type fields struct {
		podMap []octopod
		size   uint32
		keys   int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := tako{
				podMap: tt.fields.podMap,
				size:   tt.fields.size,
				keys:   tt.fields.keys,
			}
			if got := tk.MayExist(tt.args.data); got != tt.want {
				t.Errorf("tako.MayExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tako_Add(t *testing.T) {
	type fields struct {
		podMap []octopod
		size   uint32
		keys   int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &tako{
				podMap: tt.fields.podMap,
				size:   tt.fields.size,
				keys:   tt.fields.keys,
			}
			tk.Add(tt.args.data)
		})
	}
}

func Test_tako_calcPosition(t *testing.T) {
	type fields struct {
		podMap []octopod
		size   uint32
		keys   int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantMapIdx uint32
		wantPodIdx uint8
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := tako{
				podMap: tt.fields.podMap,
				size:   tt.fields.size,
				keys:   tt.fields.keys,
			}
			gotMapIdx, gotPodIdx := tk.calcPosition(tt.args.data)
			if gotMapIdx != tt.wantMapIdx {
				t.Errorf("tako.calcPosition() gotMapIdx = %v, want %v", gotMapIdx, tt.wantMapIdx)
			}
			if gotPodIdx != tt.wantPodIdx {
				t.Errorf("tako.calcPosition() gotPodIdx = %v, want %v", gotPodIdx, tt.wantPodIdx)
			}
		})
	}
}
