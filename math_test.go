package main

import (
	"testing"
)

func Test_maxmin(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name       string
		args       args
		wantMax    int
		wantMin    int
		wantMaxabs int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				[]int{-4, 3, 2, -1, 2, 6, -9, 7},
			},
			wantMax:    7,
			wantMin:    -9,
			wantMaxabs: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotMin, gotMaxabs := maxmin(tt.args.data...)
			if gotMax != tt.wantMax {
				t.Errorf("maxmin() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotMin != tt.wantMin {
				t.Errorf("maxmin() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMaxabs != tt.wantMaxabs {
				t.Errorf("maxmin() gotMaxabs = %v, want %v", gotMaxabs, tt.wantMaxabs)
			}
		})
	}
}

func Test_minabs(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				[]float64{1.2, -4.5, -0.4, 3.9, 4, -7, 9, -22.1},
			},
			want: 0.4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minabs(tt.args.data...); got != tt.want {
				t.Errorf("minabs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlope(t *testing.T) {
	type args struct {
		data    []int
		step    int
		nSmooth int
	}
	tests := []struct {
		name      string
		args      args
		wantXUp   int
		wantXDown int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				data:    []int{1, 2, 4, 5, 6, 9, 10, 3, 8, 8},
				step:    3,
				nSmooth: 1,
			},
			wantXUp:   8,
			wantXDown: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotXUp, gotXDown := maxSlope(tt.args.data, tt.args.step, tt.args.nSmooth)
			if gotXUp != tt.wantXUp {
				t.Errorf("maxSlope() gotXUp = %v, want %v", gotXUp, tt.wantXUp)
			}
			if gotXDown != tt.wantXDown {
				t.Errorf("maxSlope() gotXDown = %v, want %v", gotXDown, tt.wantXDown)
			}
		})
	}
}
