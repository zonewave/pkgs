package byteutil

import (
	"reflect"
	"testing"
)

func TestUint64BigEndianToByte(t *testing.T) {
	type args struct {
		src uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{
				src: 1024,
			},
			want: []byte{0, 0, 0, 0, 0, 0, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64BigEndianToByte(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64BigEndianToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
