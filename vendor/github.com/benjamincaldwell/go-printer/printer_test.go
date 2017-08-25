package printer

import "testing"

func TestSuccess(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Success(tt.args.text, tt.args.a...)
		})
	}
}

func TestFail(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fail(tt.args.text, tt.args.a...)
		})
	}
}

func TestError(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.text, tt.args.a...)
		})
	}
}

func TestInfo(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.text, tt.args.a...)
		})
	}
}

func TestWarning(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warning(tt.args.text, tt.args.a...)
		})
	}
}

func TestSuccessBar(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessBar(tt.args.text, tt.args.a...)
		})
	}
}

func TestErrorBar(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorBar(tt.args.text, tt.args.a...)
		})
	}
}

func TestInfoBar(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoBar(tt.args.text, tt.args.a...)
		})
	}
}

func TestWarningBar(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningBar(tt.args.text, tt.args.a...)
		})
	}
}

func TestSuccessBarIcon(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessBarIcon(tt.args.text, tt.args.a...)
		})
	}
}

func TestErrorBarIcon(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorBarIcon(tt.args.text, tt.args.a...)
		})
	}
}

func TestInfoBarIcon(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoBarIcon(tt.args.text, tt.args.a...)
		})
	}
}

func TestWarningBarIcon(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningBarIcon(tt.args.text, tt.args.a...)
		})
	}
}

func TestSuccessLine(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessLine()
		})
	}
}

func TestErrorLine(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorLine()
		})
	}
}

func TestInfoLine(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoLine()
		})
	}
}

func TestWarningLine(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningLine()
		})
	}
}

func TestSuccessLineText(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessLineText(tt.args.text, tt.args.a...)
		})
	}
}

func TestErrorLineText(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorLineText(tt.args.text, tt.args.a...)
		})
	}
}

func TestInfoLineText(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoLineText(tt.args.text, tt.args.a...)
		})
	}
}

func TestWarningLineText(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningLineText(tt.args.text, tt.args.a...)
		})
	}
}

func TestSuccessLineTop(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessLineTop()
		})
	}
}

func TestErrorLineTop(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorLineTop()
		})
	}
}

func TestInfoLineTop(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoLineTop()
		})
	}
}

func TestWarningLineTop(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningLineTop()
		})
	}
}

func TestSuccessLineTextTop(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessLineTextTop(tt.args.text, tt.args.a...)
		})
	}
}

func TestErrorLineTextTop(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorLineTextTop(tt.args.text, tt.args.a...)
		})
	}
}

func TestInfoLineTextTop(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoLineTextTop(tt.args.text, tt.args.a...)
		})
	}
}

func TestWarningLineTextTop(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningLineTextTop(tt.args.text, tt.args.a...)
		})
	}
}

func TestSuccessLineBottom(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessLineBottom()
		})
	}
}

func TestErrorLineBottom(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorLineBottom()
		})
	}
}

func TestInfoLineBottom(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoLineBottom()
		})
	}
}

func TestWarningLineBottom(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningLineBottom()
		})
	}
}

func TestVerboseSuccess(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseSuccess(tt.args.text, tt.args.a...)
		})
	}
}

func TestVerboseFail(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseFail(tt.args.text, tt.args.a...)
		})
	}
}

func TestVerboseError(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseError(tt.args.text, tt.args.a...)
		})
	}
}

func TestVerboseInfo(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseInfo(tt.args.text, tt.args.a...)
		})
	}
}

func TestVerboseWarning(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseWarning(tt.args.text, tt.args.a...)
		})
	}
}

func TestVerboseSuccessBar(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseSuccessBar(tt.args.text, tt.args.a...)
		})
	}
}

func TestVerboseErrorBar(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseErrorBar(tt.args.text, tt.args.a...)
		})
	}
}

func TestVerboseInfoBar(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseInfoBar(tt.args.text, tt.args.a...)
		})
	}
}

func TestVerboseWarningBar(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"sup %s", a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseWarningBar(tt.args.text, tt.args.a...)
		})
	}
}

func TestPrintColored(t *testing.T) {
	var a []interface{}
	a = append(a, "dawg")

	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing!",
			args: args{"{{green:sup}} {{bold:{{red:%s}}}}", a},
		},
		{
			name: "all colors!",
			args: args{"{{green:sup}} {{bold:{{red:bold red!}}}} {{cyan: cyan!!!}} {{blue:blue}}", nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintColored(tt.args.text, tt.args.a...)
		})
	}
}

func Test_getSize(t *testing.T) {
	type args struct {
		fd int
	}
	tests := []struct {
		name       string
		args       args
		wantWidth  int
		wantHeight int
		wantErr    bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWidth, gotHeight, err := getSize(tt.args.fd)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("getSize() gotWidth = %v, want %v", gotWidth, tt.wantWidth)
			}
			if gotHeight != tt.wantHeight {
				t.Errorf("getSize() gotHeight = %v, want %v", gotHeight, tt.wantHeight)
			}
		})
	}
}
