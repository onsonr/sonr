package marketing

type Button struct {
	Text string
	Href string
}

type Image struct {
	Src    string
	Width  string
	Height string
}

// ╭──────────────────────────────────────────────────────────╮
// │                  Generic Models                          │
// ╰──────────────────────────────────────────────────────────╯

type Feature struct {
	Title string
	Desc  string
	Icon  *string
	Image *Image
}

type Stat struct {
	Value string
	Denom string
	Label string
}

type Technology struct {
	Title string
	Desc  string
	Icon  *string
	Image *Image
}

type Testimonial struct {
	FullName string
	Username string
	Avatar   string
	Quote    string
}

// ╭───────────────────────────────────────────────────────────╮
// │                  HomePage Models                          │
// ╰───────────────────────────────────────────────────────────╯

type Hero struct {
	TitleFirst      string
	TitleEmphasis   string
	TitleSecond     string
	Subtitle        string
	PrimaryButton   *Button
	SecondaryButton *Button
	Image           *Image
	Stats           []*Stat
}

type Highlights struct {
	Heading  string
	Subtitle string
	Features []*Feature
}

type Mission struct {
	Eyebrow          string
	Heading          string
	Subtitle         string
	Experience       *Feature
	Compliance       *Feature
	Interoperability *Feature
	Standards        []*Feature // Display 6 Standards applied by the Sonr Network
}

type Architecture struct {
	Heading    string
	Subtitle   string
	Primary    *Technology
	Secondary  *Technology
	Tertiary   *Technology
	Quaternary *Technology
	Quinary    *Technology
}

type Lowlights struct {
	Heading string
	Quotes  []*Testimonial
}

type CallToAction struct {
	Logo      *Image
	Heading   string
	Subtitle  string
	Primary   *Button
	Secondary *Button
	Partners  []*Image
}
