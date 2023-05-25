package evenodd

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/escalopa/itis-tables/core"
)

func TestEvenOddDateTimeGetWeek(t *testing.T) {
	type fields struct {
		startDate time.Time
	}
	type args struct {
		ctx context.Context
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   core.WeekType
	}{
		{
			name: "success odd date 1",
			fields: fields{
				startDate: time.Date(time.Now().Year(), 2, 9, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				ctx: context.Background(),
				now: time.Date(time.Now().Year(), 2, 26, 0, 0, 0, 0, time.UTC),
			},
			want: core.WeekOdd,
		},
		{
			name: "success odd date 2",
			fields: fields{
				startDate: time.Date(time.Now().Year(), 2, 9, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				ctx: context.Background(),
				now: time.Date(time.Now().Year(), 3, 10, 0, 0, 0, 0, time.UTC),
			},
			want: core.WeekOdd,
		},
		{
			name: "success even date 0",
			fields: fields{
				startDate: time.Date(time.Now().Year(), 2, 9, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				ctx: context.Background(),
				now: time.Date(time.Now().Year(), 2, 16, 0, 0, 0, 0, time.UTC),
			},
			want: core.WeekEven,
		},
		{
			name: "success even date 1",
			fields: fields{
				startDate: time.Date(time.Now().Year(), 2, 9, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				ctx: context.Background(),
				now: time.Date(time.Now().Year(), 3, 3, 0, 0, 0, 0, time.UTC),
			},
			want: core.WeekEven,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eod := NewEvenOddDateTime(tt.fields.startDate)
			if got := eod.GetWeek(tt.args.ctx, tt.args.now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvenOddDateTime.GetWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
