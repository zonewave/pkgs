package config

import (
	"errors"
	"testing"
)

func Test_checkObject(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "ok struct",
			args: args{
				struct {
					id string
				}{id: "test"},
			},
			wantErr: InvalidConfigTypeError("should be a pointer to a struct"),
		},
		{
			name: "ok ptr",
			args: args{
				&struct {
					id string
				}{id: "test"},
			},
			wantErr: nil,
		},
		{
			name: "no",
			args: args{
				32,
			},
			wantErr: InvalidConfigTypeError("should be a pointer to a struct"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkObject(tt.args.obj); !errors.Is(err, tt.wantErr) {
				t.Errorf("checkObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
