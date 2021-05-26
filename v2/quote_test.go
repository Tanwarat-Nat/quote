package quote

import "testing"

func TestSay(t *testing.T) {
	want := "Hello"
	if got := Say2(); got != want {
		t.Errorf("Say2() = %q, want : %q", got, want)
	}
}

func TestSpeak2(t *testing.T) {
	want := "Hi, mate."
	if got := Speak2(); got != want {
		t.Errorf("Speak2() = %q, want : %q", got, want)
	}
}

func TestTalk2(t *testing.T) {
	want := "Gopher!"
	if got := Talk2(); got != want {
		t.Errorf("Talk2() = %q, want : %q", got, want)
	}
}
