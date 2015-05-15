package spartan

import "testing"

func TestSanity(t *testing.T) {
	app := New()

	helloBox := TextBox{text: "Hello, World!"}
	app.AddView(helloBox)

	app.Run()
}
