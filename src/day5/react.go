package day5

import "unicode"

func React(line string) string {
	lastLen := len(line)
	for {
		for i := 0; i < len(line)-1; i++ {
			a := rune(line[i])
			b := rune(line[i+1])

			if a != b && unicode.ToLower(a) == unicode.ToLower(b) {
				line = line[:i] + line[i+2:]
			}
		}

		newLen := len(line)

		if newLen == lastLen {
			// done, since no more mods were made
			return line
		} else {
			lastLen = newLen
		}
	}
}
