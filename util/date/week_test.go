package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLastMonday(t *testing.T) {

	monday := time.Date(2003, 2, 24, 0, 0, 0, 0, time.Local)

	type args struct {
		time time.Time
	}

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "should return last monday",
			args: args{time.Date(2003, 2, 24, 0, 0, 0, 0, time.Local)},
			want: monday,
		},
		{
			name: "should return last monday",
			args: args{time.Date(2003, 2, 25, 0, 0, 0, 0, time.Local)},
			want: monday,
		},
		{
			name: "should return last monday",
			args: args{time.Date(2003, 2, 26, 0, 0, 0, 0, time.Local)},
			want: monday,
		}, {
			name: "should return last monday",
			args: args{time.Date(2003, 2, 27, 0, 0, 0, 0, time.Local)},
			want: monday,
		},
		{
			name: "should return last monday",
			args: args{time.Date(2003, 2, 28, 0, 0, 0, 0, time.Local)},
			want: monday,
		},
		{
			name: "should return last monday",
			args: args{time.Date(2003, 3, 1, 0, 0, 0, 0, time.Local)},
			want: monday,
		},
		{
			name: "should return last monday",
			args: args{time.Date(2003, 3, 2, 0, 0, 0, 0, time.Local)},
			want: monday,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := LastMonday(test.args.time)
			assert.Equal(t, test.want, m)
		})
	}
}
