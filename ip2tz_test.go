package ip2tz

import (
	"testing"
)

func TestStrToInt(t *testing.T) {
	cases := []string{"0.0.0.0", "211.81.55.60", "8.8.8.8"}
	for _, c := range cases {
		t.Log(Find(c))
	}
}
