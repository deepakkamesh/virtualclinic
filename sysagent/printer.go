package sysagent

// FormattedLine represents a line of text with various formatting options.
type FormattedLine struct {
	Font      string   // Font type
	FontSize  [2]uint8 // Font size as width and height multipliers
	Align     string   // Text alignment (e.g., "left", "center", "right")
	Emphasize uint8    // Emphasize text (0 for no, 1 for yes)
	Smooth    uint8    // Smooth text (0 for no, 1 for yes)
	Underline uint8    // Underline text (0 for no, 1 for yes)
	FormFeed  int      // Number of form feeds after the line
	Text      string   // The actual text content
}

// Font sets the font type for the FormattedLine.
func Font(font string) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Font = font
	}
}

// FontSize sets the font size for the FormattedLine.
func FontSize(size [2]uint8) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.FontSize = size
	}
}

// Align sets the text alignment for the FormattedLine.
func Align(align string) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Align = align
	}
}

// Emphasize sets the emphasize option for the FormattedLine.
func Emphasize(emphasize uint8) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Emphasize = emphasize
	}
}

// Smooth sets the smooth option for the FormattedLine.
func Smooth(smooth uint8) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Smooth = smooth
	}
}

// Underline sets the underline option for the FormattedLine.
func Underline(underline uint8) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.Underline = underline
	}
}

// FormFeed sets the number of form feeds after the FormattedLine.
func FormFeed(formFeed int) func(*FormattedLine) {
	return func(o *FormattedLine) {
		o.FormFeed = formFeed
	}
}

// Line creates a new FormattedLine with the specified text and formatting options.
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
