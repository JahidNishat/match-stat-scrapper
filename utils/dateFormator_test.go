package utils

import (
	"testing"
)

func TestBleagueDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BleagueDate",
			args: args{"2006.03.09"},
			want: "03/09",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BleagueDate(tt.args.date); got != tt.want {
				t.Errorf("BleagueDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBnxtDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BnxtDate",
			args: args{"02-Jan-2006"},
			want: "01/02",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BnxtDate(tt.args.date); got != tt.want {
				t.Errorf("BnxtDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBritishBasketBallDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BBBDate",
			args: args{"Jan 02, 2006, 3:04 PM"},
			want: "01/02",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BritishBasketBallDate(tt.args.date); got != tt.want {
				t.Errorf("BritishBasketBallDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEspnDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "EspnDate",
			args: args{"Sun 12/10"},
			want: "12/10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EspnDate(tt.args.date); got != tt.want {
				t.Errorf("EspnDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEuroBasketDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BnxtDate",
			args: args{"12/11/2006"},
			want: "12/11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EuroBasketDate(tt.args.date); got != tt.want {
				t.Errorf("EuroBasketDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNblDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "NblDate",
			args: args{"Nov 18th"},
			want: "11/18",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NblDate(tt.args.date); got != tt.want {
				t.Errorf("NblDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
