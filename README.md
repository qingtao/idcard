# idcard

[![Go Reference](https://pkg.go.dev/badge/github.com/qingtao/idcard.svg)](https://pkg.go.dev/github.com/qingtao/idcard)

校验身份证号码

## Usage

go get github.com/qingtao/idcard@latest


```go
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
``
