package sysagent

type FormattedLine struct {
	Font      string
	FontSize  [2]uint8
	Align     string
	Emphasize uint8
	Smooth    uint8
	Underline uint8
	FormFeed  int
	Text      string
}

func Font(font string) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Font = font
	}
}
func FontSize(size [2]uint8) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.FontSize = size
	}
}
func Align(align string) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Align = align
	}
}
func Emphasize(emphasize uint8) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Emphasize = emphasize
	}
}
func Smooth(smooth uint8) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Smooth = smooth
	}
}
func Underline(underline uint8) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Underline = underline
	}
}

func FormFeed(formFeed int) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.FormFeed = formFeed
	}
}

func Line(text string, opts ...func(*FormattedLine)) *FormattedLine {
	options := &FormattedLine{
		Font:      "A",            // Default font
		FontSize:  [2]uint8{1, 1}, // Default font size
		Align:     "left",         // Default alignment
		Emphasize: 0,              // Default emphasize
		Smooth:    0,              // Default smooth
		Underline: 0,              // Default underline
		FormFeed:  1,              // Default form feed
		Text:      text,           // Default text
	}

	for _, opt := range opts {
		opt(options)
	}
	return options
}
