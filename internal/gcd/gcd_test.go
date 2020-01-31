package gcd

import (
	"testing"
)

func TestGcd(t *testing.T) {
	type args struct {
		a uint
		b uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			args: args{0, 0},
			want: 0,
		},
		{
			args: args{1, 100},
			want: 1,
		},
		{
			args: args{3, 9},
			want: 3,
		},
		{
			args: args{3, 100},
			want: 1,
		},
		{
			args: args{123, 321},
			want: 3,
		},
		{
			args: args{456, 654},
			want: 6,
		},
		{
			args: args{2222, 1122},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGcdSlice(t *testing.T) {
	type args struct {
		s []uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			args: args{[]uint{}},
			want: 0,
		},
		{
			args: args{[]uint{1}},
			want: 1,
		},
		{
			args: args{[]uint{2, 4}},
			want: 2,
		},
		{
			args: args{[]uint{8, 40, 20}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GcdSlice(tt.args.s); got != tt.want {
				t.Errorf("GcdSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
