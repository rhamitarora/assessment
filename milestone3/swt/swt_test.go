package swt

import (
	"testing"
)

var want = "2022-06-07@385.19"

func TestCurrentTime(t *testing.T) {
	got := CurrentTime()
	
	if got != want {
			t.Errorf("Test fail! want: '%s', got: '%s'", want, got)
	}
}
