package quote

import "testing"

func TestSay(t *testing.T) {
	want := "Hello App"
	if got := Say() ; got != want {
		t.Errorf("Say() = %q , want : %q", got, want)
	}
}

func TestSpeak(t *testing.T) {
	want := "Hi App"
	if got := Speak() ; got != want {
		t.Errorf("Speak() = %q , want : %q", got, want)
	}
}