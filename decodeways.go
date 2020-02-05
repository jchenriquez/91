package decodeways

import "strconv"

func decoder (s string, memo map[string] bool) int {
  
  

  if count, saw := memo[s]; saw {
    return count
  }

  var count int

  for i := 0; i < len(s); i++ {
    count += decoder(s[:i+1], memo)
    count += decoder(s[i+1:], memo)
  }

  memo[s] = count
  return count
}

// NumDecodings will return the amount of ways s can be decoded.
func NumDecodings(s string) (ret int) {
  if len(s) == 0 {
    return 0
  }
	return decoder(s, make(map[string] bool))
}
