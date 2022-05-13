package dispatcher

import (
	"testing"
)

func TestDispatcher(t *testing.T) {

	params := []string{"a", "b", "c"}

	d := NewDispatcher(3, 10240, func(v interface{}) {

		if  !(v.(string) == "a" || v.(string) == "b" || v.(string) == "c") {
			t.Error("wrong params", v.(string))
		}

	})
	d.Start()

	for _, v := range params {
		d.Add(v)
	}
	d.Wait()

}