package decodeways

import (
  "testing"
  "os"
  "encoding/json"
  "bufio"
)

type Test struct {
  Input string `json:"input"`
  Output int  `json:"output"`
}

func TestDecodeWays (test *testing.T) {
  f, err := os.Open(testing.json)

  if err != nil {
    test.Error(err)
  }

  defer f.close()

  reader := bufio.NewReader(f)
  decoder := json.NewDecoder(reader)

  for {
    var tests map[string]Test
    err = decoder.Decode(&tests)

    if err == nil {
      for name, tst := range tests {
        test.Run(name, func (st *testing.T) {
          if NumDecodings(tst.Input) != tst.Output {
            st.Errorf("Use case %v failed\n", tst)
          }
        })
      }
    }
  }

}

