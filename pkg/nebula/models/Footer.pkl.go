// Code generated from Pkl module `models`. DO NOT EDIT.
package models

type Footer struct {
	Logo *Image `pkl:"logo"`

	MediumLink *SocialLink `pkl:"mediumLink"`

	TwitterLink *SocialLink `pkl:"twitterLink"`

	DiscordLink *SocialLink `pkl:"discordLink"`

	GithubLink *SocialLink `pkl:"githubLink"`

	CompanyLinks []*Link `pkl:"companyLinks"`

	ResourcesLinks []*Link `pkl:"resourcesLinks"`
}
