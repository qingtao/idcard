package idcard

import (
	"fmt"
	"strconv"
	"testing"
)

func ExampleValidate() {
	var (
		s      = "11010519491231002X"
		gender = 0
	)
	fmt.Printf("Validate(%s): %t\n", s, Validate(s))
	fmt.Printf("Validate(%s, %d): %t\n", s, gender, Validate(s, gender))
	ok, got := GetGender(s)
	fmt.Printf("GetGender(%s): %t, %d", s, ok, got)

	// Output:
	// Validate(11010519491231002X): true
	// Validate(11010519491231002X, 0): true
	// GetGender(11010519491231002X): true, 0
}

func Test_Validate(t *testing.T) {
	type args struct {
		s      string
		gender []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s:      "11010519491231002X",
				gender: []int{0},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s:      "440524188001010014",
				gender: []int{1},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s:      "370683198901117667",
				gender: []int{1},
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				s:      "370683198901007667",
				gender: []int{1},
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				s: "34052419800101001X",
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				s: "abc52419800101001X",
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				s: "身份证号校验", // 不是数字或者Xx
			},
			want: false,
		},
		{
			name: "8",
			args: args{
				s: "34052419800101001",
			},
			want: false,
		},
		{
			// 这是号码能过规则, 但是是一个无效的
			name: "9",
			args: args{
				s: "372925198811050000",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.s, tt.args.gender...); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkCode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{b: []byte("34052419800101001X")},
			want: 2,
		},
		{
			name: "2",
			args: args{b: []byte("34052419800101001x")},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkCode(tt.args.b); got != tt.want {
				t.Errorf("checkCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{b: []byte("34052419800101001X")},
			want: 189,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.b); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{b: []byte("34052419800101001X")},
			want: true,
		},
		{
			name: "2",
			args: args{b: []byte("34052419800100001X")},
			want: false,
		},
		{
			name: "3",
			args: args{b: []byte("34052419800101001x")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.b); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Validate(b *testing.B) {
	type args struct {
		idcard string
		gender int
	}
	tests := []args{
		{"34052419800101001X", 1},
		{"370683198901117657", 1},
		{"370683198901117657", 0},
		{"身份证号校验", 0},
		{"37068319890111657", 0},
		{"abc6*3198901117657", 0},
	}
	for ii, tt := range tests {
		var ok bool
		b.Run(strconv.Itoa(ii), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ok = Validate(tt.idcard, tt.gender)
			}
		})
		b.Log(ok)
	}
}

func Benchmark_validate(b *testing.B) {
	tests := [][]byte{
		[]byte("34052419800101001X"),
		[]byte("370683198901117657"),
	}
	for ii, tt := range tests {
		var ok bool
		b.Run(strconv.Itoa(ii), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ok = validate(tt)
			}
		})
		b.Log(ok)
	}
}

func TestGetGender(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name:  "1",
			args:  args{s: "34052419800101001X"},
			want:  true,
			want1: 1,
		},
		{
			name:  "2",
			args:  args{s: "11010519491231002X"},
			want:  true,
			want1: 0,
		},
		{
			name:  "3",
			args:  args{s: "34052419800101001x"},
			want:  true,
			want1: 1,
		},
		{
			name:  "4",
			args:  args{s: "34052419800101021X"},
			want:  false,
			want1: -1,
		},
		{
			name:  "5",
			args:  args{s: "3405241980010102X"},
			want:  false,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetGender(tt.args.s)
			if got != tt.want {
				t.Errorf("GetGender() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetGender() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
