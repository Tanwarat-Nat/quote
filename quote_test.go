package quote

import "testing"

func TestSay(t *testing.T) {
	want := "Hello"
	if got := Say(); got != want {
		t.Errorf("Say() = %q, want : %q", got, want)
	}
}
