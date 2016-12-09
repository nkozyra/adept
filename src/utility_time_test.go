package adept

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	s := time.Now().Unix()
	n := Now()
	if s != n {
		t.Error("Expected ", s, ", got", n)
	}
}
