package models

type Button struct {
	Text string
	Href string
}

type Image struct {
	Src    string
	Width  string
	Height string
}

type Stat struct {
	Value string
	Label string
}

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
