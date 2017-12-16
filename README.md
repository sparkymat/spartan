# spartan
A windowing toolkit based on termbox. Most of the paradigms used here are inspired from the Android UI system.

### Currently Implemented
* LinearLayout
* TextView

### Example

``` go
	app := spartan.New()

	layout := spartan.LinearLayout{
		Direction: direction.Vertical,
		View: spartan.View{
			Width:  size.MatchParent,
			Height: size.MatchParent,
		},
	}

	helloBox := spartan.TextView{
		Text: "Hello, World!",
		View: spartan.View{
			Width:           20,
			Height:          1,
			ForegroundColor: termbox.ColorWhite,
			BackgroundColor: termbox.ColorRed,
		},
	}

	triumphBox := spartan.TextView{
		Text: "This was a triumph",
		View: spartan.View{
			Width:           size.MatchParent,
			Height:          3,
			ForegroundColor: termbox.ColorBlack,
			BackgroundColor: termbox.ColorMagenta,
		},
	}

	noteBox := spartan.TextView{
		Text: "I am making a note here",
		View: spartan.View{
			Width:           6,
			Height:          size.MatchParent,
			LayoutGravity:   gravity.Center,
			ForegroundColor: termbox.ColorRed,
			BackgroundColor: termbox.ColorBlue,
		},
	}

	successBox := spartan.TextView{
		Text: "Huge success",
		View: spartan.View{
			Width:           15,
			Height:          size.MatchParent,
			RightMargin:     10,
			LayoutGravity:   gravity.Right,
			ForegroundColor: termbox.ColorGreen,
			BackgroundColor: termbox.ColorYellow,
		},
	}

	layout.AddChild(&helloBox)
	layout.AddChild(&triumphBox)
	layout.AddChild(&noteBox)
	layout.AddChild(&successBox)

	app.SetContent(&layout)

	eventChannel := make(chan termbox.Event)

	go EventHandler(eventChannel)

	app.Run(eventChannel)
```

produces the following

![Linear layout example](/screenshots/screenshot1.png?raw=true)
