package decodeways

import (
	"fmt"
	"strconv"
)

// NumDecodings will return the amount of ways s can be decoded.
func NumDecodings(s string) (ret int) {
	if len(s) == 0 {
		return 0
	}

	numWays := make([][]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		inDigit := s[i] - '0'

		if len(numWays) == 0 {
			if inDigit == 0 {
				return 0
			}
			numWays = append(numWays, []byte{s[i] - '0'})
		} else {
			tmpNumWays := make([][]byte, 0, len(numWays)+1)
			for _, bytArr := range numWays {
				comb := fmt.Sprintf("%d%d", bytArr[len(bytArr)-1], inDigit)
				fmt.Printf("comb %s\n", comb)
				combInt, _ := strconv.Atoi(comb)
				fmt.Printf("inDigit1 %d\n", inDigit)
				fmt.Printf("combInt %d\n", combInt)

				if combInt <= 26 {
					bytArrCpy := make([]byte, len(bytArr))
					copy(bytArrCpy, bytArr[:len(bytArr)-1])
					bytArrCpy = append(bytArrCpy, byte(combInt))
					tmpNumWays = append(tmpNumWays, bytArrCpy)
				} else {
					if inDigit == 0 {
						continue
					}
				}

				if inDigit != 0 {
					tmpNumWays = append(tmpNumWays, append(bytArr, inDigit))
				}
			}

			if len(tmpNumWays) == 0 {
				return 0
			}

			numWays = tmpNumWays
		}

	}

	return len(numWays)
}
