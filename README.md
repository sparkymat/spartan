# spartan
A windowing toolkit based on termbox. Most of the paradigms used here are inspired from the Android UI system.

### Currently Implemented
* LinearLayout
* TextView

### Example

``` go
	app := New()

	layout := LinearLayout{}
	layout.SetDirection(direction.Vertical)
	layout.SetWidth(size.MatchParent)
	layout.SetHeight(size.MatchParent)

	helloBox := TextView{text: "Hello, World!"}
	helloBox.SetWidth(20)
	helloBox.SetHeight(1)
	helloBox.SetColor(termbox.ColorWhite)
	helloBox.SetBackgroundColor(termbox.ColorRed)

	triumphBox := TextView{text: "This was a triumph"}
	triumphBox.SetWidth(size.MatchParent)
	triumphBox.SetHeight(3)
	triumphBox.SetColor(termbox.ColorBlack)
	triumphBox.SetBackgroundColor(termbox.ColorMagenta)

	noteBox := TextView{text: "I am making a note here"}
	noteBox.SetWidth(6)
	noteBox.SetHeight(size.MatchParent)
	noteBox.SetLayoutGravity(gravity.Center)
	noteBox.SetColor(termbox.ColorRed)
	noteBox.SetBackgroundColor(termbox.ColorBlue)

	successBox := TextView{text: "Huge success"}
	successBox.SetWidth(15)
	successBox.SetHeight(size.MatchParent)
	successBox.SetRightMargin(10)
	successBox.SetLayoutGravity(gravity.Right)
	successBox.SetColor(termbox.ColorGreen)
	successBox.SetBackgroundColor(termbox.ColorYellow)

	layout.AddView(&helloBox)
	layout.AddView(&triumphBox)
	layout.AddView(&noteBox)
	layout.AddView(&successBox)

	app.SetLayout(&layout)

	app.Run()
```

produces the following

![Linear layout example](/sparkymat/spartan/master/screenshots/screenshot1.png?raw=true)
