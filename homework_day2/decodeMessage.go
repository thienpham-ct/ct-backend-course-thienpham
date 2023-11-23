package homework_day2

func DecodeMessage(key string, message string) string {
	keymap := make(map[int]rune)
	for _, char := range key {
		_, isFound := keymap[int(char)]
		if isFound {
			continue
		} else if char != ' ' {
			keymap[int(char)] = rune(int('a') + len(keymap))
		}

	}

	ans := make([]rune, len(message))
	for i, char := range message {
		if char == ' ' {
			ans[i] = ' '
		} else {
			ans[i] = keymap[int(char)]
		}
	}

	return string(ans)
}
