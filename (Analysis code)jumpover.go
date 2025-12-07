package piscine

func JumpOver(str string) string {
	length := len(str)
	if length < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < length; i += 3 {
		result += string(str[i])
	}

	return result + "\n"
}
