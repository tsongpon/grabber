package grabber

import "testing"

func TestGrab(t *testing.T) {
	want := ""
	if got := Grab("https://xonly8.com/index.php?topic=253382.0", ""); got != want {
		t.Errorf("Grab() = %q, want %q", got, want)
	}
}
