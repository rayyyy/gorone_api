package util

// UniqString 文字列の配列の重複を削除
func UniqString(array []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range array {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
