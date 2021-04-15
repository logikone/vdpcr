package poller

import (
	"testing"
)

func Test_isUpOrDown(t *testing.T) {
	type args struct {
		code int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "200",
			args: args{
				code: 200,
			},
			want: 0,
		},
		{
			name: "500",
			args: args{
				code: 500,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUpOrDown(tt.args.code); got != tt.want {
				t.Errorf("isUpOrDown() = %v, want %v", got, tt.want)
			}
		})
	}
}
