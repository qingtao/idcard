// Package idcard 校验证号码的有效性,但不校验真实性
package idcard

// 校验方法参考: GB 11643-1999

var (
	// 身份证预判断正则, 使用前置正则和直接判断字节长度性能差距过大,所以取消正则匹配
	// re = regexp.MustCompile(`^(\d{6})(\d{4})(0[1-9]|1[0-2])([0-2][1-9]|123[0]|31)(\d{2})(\d)([0-9X])$`)
	// 前17位数字对应的权重
	weights = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 最后一位的数组
	last = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

// sum 前17位加权的和
func sum(b []byte) (n int) {
	for i := 0; i < len(b)-1; i++ {
		n += weights[i] * int(b[i]-'0')
	}
	return n
}

// checkCode 计算校验码
func checkCode(s []byte) int {
	return sum(s) % 11
}

// validate 校验身份证号码
func validate(b []byte) bool {
	s := b[len(b)-1]
	if s == 'x' {
		s = 'X'
	}
	return s == last[checkCode(b)]
}

// Validate 校验身份证号码, s为要校验的身份证,gender为性别校验,单数为男,双数为女
func Validate(s string, gender ...int) bool {
	if len(s) != 18 {
		return false
	}
	if len(gender) != 0 && int(s[16]-'0')&0x1 != gender[0]&0x1 {
		return false
	}
	return validate([]byte(s))
}

// GetGender 从身份号码中提取性别, 如果是男性返回1,女性返回0, 如果身份证号码不符合规则, 则返回-1,
func GetGender(s string) (bool, int) {
	if len(s) != 18 {
		return false, -1
	}
	ok := validate([]byte(s))
	if !ok {
		return false, -1
	}
	return true, int(s[16]-'0') & 0x1
}
