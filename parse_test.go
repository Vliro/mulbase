package humus

import (
	"fmt"
	"testing"
	"time"
)

var testErrorValue = []byte(`{
    "q0": [
      {
        "Error.errorType": "test",
        "Error.time": "2019-12-10T13:47:47.967597476+01:00",
        "Error.message": "testagain"
      }
    ]
  }`)

func TestDeserialize(t *testing.T) {
	var res dbError
	err := handleResponse(testErrorValue, []interface{}{&res}, []string{"q0"})
	if err != nil {
		t.Fail()
		return
	}
	if res.Message != "testagain" || res.ErrorType != "test" {
		t.Fail()
	}
}

func TestSerialize(t *testing.T) {
	var res = dbError{
		Message:   "testagain",
		ErrorType: "test",
		Time:      time.Now(),
	}
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func BenchmarkDeserialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res dbError
		err := handleResponse(testErrorValue, []interface{}{&res}, []string{"q0"})
		if err != nil {
			return
		}
	}
	b.ReportAllocs()
}
