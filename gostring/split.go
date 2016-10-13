package gostring

import (
	"unicode"
	"unicode/utf8"
)

// SplitTextToWords 将text中的一串字母或数字当做一个Text，其他字当做一个Text输出。
// 比如 `中国abc世界`, 分离结果是 [中, 国, abc, 世, 界]。
func SplitTextToWords(text []byte) [][]byte {
	output := make([][]byte, 0, len(text)/3)
	current := 0
	inAlphanumeric := true
	alphanumericStart := 0
	for current < len(text) {
		r, size := utf8.DecodeRune(text[current:])
		if size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r)) {
			// 当前是拉丁字母或数字（非中日韩文字）
			if !inAlphanumeric {
				// 上一次遍历时不是字母或数字
				alphanumericStart = current
				inAlphanumeric = true
			}
		} else {
			// 中日韩文字
			if inAlphanumeric {
				// 上一次遍历是数字或字母，这次不是了，所以要先将字母或数字串存起来
				inAlphanumeric = false
				if current != 0 {
					output = append(output, text[alphanumericStart:current])
				}
			}
			output = append(output, text[current:current+size]) // 当前指向的rune
		}
		current += size
	}

	// 处理最后一个字元是英文的情况
	if inAlphanumeric {
		if current != 0 {
			output = append(output, text[alphanumericStart:current])
		}
	}

	return output
}
