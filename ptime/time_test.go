package ptime

import (
	"testing"
)

func TestCountLeapYears(t *testing.T) {
	type args struct {
		syear uint
		eyear uint
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "syear is zero",
			args: args{
				syear: 0,
				eyear: 2000,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "syear is bigger",
			args: args{
				syear: 2000,
				eyear: 1999,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test1",
			args: args{
				syear: 1999,
				eyear: 2003,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				syear: 1999,
				eyear: 2000,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				syear: 1999,
				eyear: 2001,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				syear: 2000,
				eyear: 2007,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				syear: 2000,
				eyear: 2008,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "test6",
			args: args{
				syear: 2000,
				eyear: 2009,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "test7",
			args: args{
				syear: 2001,
				eyear: 2003,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "test8",
			args: args{
				syear: 2001,
				eyear: 2004,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "test9",
			args: args{
				syear: 2001,
				eyear: 2005,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountLeapYears(tt.args.syear, tt.args.eyear)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountLeapYears() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountLeapYears() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yearDay(t *testing.T) {
	type args struct {
		md     string
		isleap bool
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"2023-01-01", args{"01-01", false}, 1},
		{"2023-02-15", args{"02-15", false}, 46},
		{"2023-12-31", args{"12-31", false}, 365},
		{"2024-01-01", args{"01-01", true}, 1},
		{"2024-02-15", args{"02-15", true}, 46},
		{"2024-12-31", args{"12-31", true}, 366},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := YearDay(tt.args.md, tt.args.isleap); got != tt.want {
				t.Errorf("yearDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcDiffDay(t *testing.T) {
	type args struct {
		d1 string
		d2 string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"same day", args{"2023-06-30", "2023-06-30"}, 0, false},
		{"less", args{"2023-06-30", "2023-07-31"}, 31, false},
		{"big", args{"2023-06-30", "2023-06-27"}, -3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcDiffDay(tt.args.d1, tt.args.d2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc DiffDayEx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc DiffDayEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
