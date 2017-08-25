package printer

import (
	"fmt"
	"testing"
)

func TestGreen(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "color test",
			args: args{"hello"},
			want: GreenColor + "hello" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Green(tt.args.text)
			if got != tt.want {
				t.Errorf("Green() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestGreenf(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "colorf test",
			args: args{"Hello %s", a},
			want: GreenColor + "Hello dawg" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greenf(tt.args.text, tt.args.a...)
			if got != tt.want {
				t.Errorf("Greenf() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestRed(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "color test",
			args: args{"hello"},
			want: RedColor + "hello" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Red(tt.args.text)
			if got != tt.want {
				t.Errorf("Red() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestRedf(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "colorf test",
			args: args{"Hello %s", a},
			want: RedColor + "Hello dawg" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Redf(tt.args.text, tt.args.a...)
			if got != tt.want {
				t.Errorf("Redf() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestBlue(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "color test",
			args: args{"hello"},
			want: BlueColor + "hello" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Blue(tt.args.text)
			if got != tt.want {
				t.Errorf("Blue() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestBluef(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "colorf test",
			args: args{"Hello %s", a},
			want: BlueColor + "Hello dawg" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Bluef(tt.args.text, tt.args.a...)
			if got != tt.want {
				t.Errorf("Bluef() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestYellow(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "color test",
			args: args{"hello"},
			want: YellowColor + "hello" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Yellow(tt.args.text)
			if got != tt.want {
				t.Errorf("Yellow() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestYellowf(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "colorf test",
			args: args{"Hello %s", a},
			want: YellowColor + "Hello dawg" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Yellowf(tt.args.text, tt.args.a...)
			if got != tt.want {
				t.Errorf("Yellowf() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestCyan(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "color test",
			args: args{"hello"},
			want: CyanColor + "hello" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cyan(tt.args.text)
			if got != tt.want {
				t.Errorf("Cyan() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestCyanf(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "colorf test",
			args: args{"Hello %s", a},
			want: CyanColor + "Hello dawg" + NoColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cyanf(tt.args.text, tt.args.a...)
			if got != tt.want {
				t.Errorf("Cyaf() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestBold(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "color test",
			args: args{"hello"},
			want: BoldString + "hello" + NoboldString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Bold(tt.args.text)
			if got != tt.want {
				t.Errorf("Bold() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}

func TestBoldf(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "colorf test",
			args: args{"Hello %s", a},
			want: BoldString + "Hello dawg" + NoboldString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Boldf(tt.args.text, tt.args.a...)
			if got != tt.want {
				t.Errorf("Boldf() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
