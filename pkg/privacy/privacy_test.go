package privacy

import "testing"

func TestRedactWithOptions(t *testing.T) {
	input := "User test@example.com connected from 192.168.1.1"

	opts := Options{RedactEmails: true, RedactIPs: true, RedactCards: true}
	expected := "User [REDACTED EMAIL] connected from [REDACTED IP]"
	if got := RedactWithOptions(input, opts); got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}

	optsNoIP := Options{RedactEmails: true, RedactIPs: false, RedactCards: true}
	expectedNoIP := "User [REDACTED EMAIL] connected from 192.168.1.1"
	if got := RedactWithOptions(input, optsNoIP); got != expectedNoIP {
		t.Errorf("got %q, want %q", got, expectedNoIP)
	}
}
