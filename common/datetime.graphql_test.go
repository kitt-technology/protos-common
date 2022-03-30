package common

import (
	"reflect"
	"testing"
)

func Test_DateTimeFromArgs(t *testing.T) {
	type args struct {
		args map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *DateTime
		wantErr bool
	}{
		{
			name: "Test_DateTimeFromArgs",
			args: args{
				args: map[string]interface{}{
					"dateTime": map[string]interface{}{
						"year":    2022,
						"month":   03,
						"day":     30,
						"hours":   17,
						"minutes": 34,
						"seconds": 00,
						"nanos":   00,
					},
				},
			},
			want: &DateTime{
				Year:    2022,
				Month:   3,
				Day:     30,
				Hours:   17,
				Minutes: 34,
				Seconds: 0,
				Nanos:   0,
			},
			wantErr: false,
		},
		{
			name: "Test_DateTimeFromArgs",
			args: args{
				args: map[string]interface{}{
					"ISOString": "2022-03-30T17:34:00Z",
				},
			},
			want: &DateTime{
				Year:     2021,
				Month:    3,
				Day:      30,
				Hours:    17,
				Minutes:  34,
				Seconds:  0,
				Nanos:    0,
				TimeZone: "UTC",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DateTimeFromArgs(tt.args.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTimeFromArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
