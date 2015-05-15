package spartan

import "testing"

func TestSanity(t *testing.T) {
	app := New()
	app.Run()
}
