package decodeways

import (
	"fmt"
	"strconv"
)

func toKey(arr []byte) (key string) {

	for _, byt := range arr {
		key = fmt.Sprintf("%s%c", key, byt+'a')
	}

	return
}

func decoder(s []byte, memo map[string]int) int {

	fmt.Printf("s %v\n", s)
	key := toKey(s)

	if _, saw := memo[key]; saw {
		return 0
	}

	var count int

	for i := 0; i < len(s)-1; i++ {
		inDigit := s[i] - '0'

		if inDigit == 0 && i == 0 {
			return 0
		}

		nDigit := s[i+1] - '0'

		fmted := fmt.Sprintf("%d%d", inDigit, nDigit)
		toDigit, _ := strconv.Atoi(fmted)

		switch {
		case toDigit <= 26:
			nArray := make([]byte, 0, len(s)-1)
			nArray = append(nArray, s[:i]...)
			nArray = append(nArray, byte(toDigit))
			count += decoder(append(nArray, s[i+2:]...), memo) + 1
		case toDigit > 26 && inDigit == 0:
			return 0
		}
	}

	memo[key] = count

	return count

}

// NumDecodings will return the amount of ways s can be decoded.
func NumDecodings(s string) (ret int) {
	if len(s) == 0 {
		return 0
	}

	count := 1

	sbyt := []byte(s)

	for _, byt := range sbyt {
		if byt-'0' == 0 {
			count--
		}
	}

	return count + decoder(sbyt, make(map[string]int, len(s)))
}
