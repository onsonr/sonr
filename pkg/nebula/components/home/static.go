package home

import "github.com/onsonr/sonr/pkg/nebula/models"

var hero = &models.Hero{
	TitleFirst:      "Simplified",
	TitleEmphasis:   "self-custody",
	TitleSecond:     "for everyone",
	Subtitle:        "Sonr is a modern re-imagination of online user identity, empowering users to take ownership of their digital footprint and unlocking a new era of self-sovereignty.",
	PrimaryButton:   &models.Button{Text: "Get Started", Href: "/register"},
	SecondaryButton: &models.Button{Text: "Learn More", Href: "/about"},
	Image: &models.Image{
		Src:    "https://cdn.sonr.id/img/hero-clipped.svg",
		Width:  "500",
		Height: "500",
	},
	Stats: []*models.Stat{
		{Value: "476", Label: "Assets packed with power beyond your imagination.", Denom: "K"},
		{Value: "1.44", Label: "Assets packed with power beyond your imagination.", Denom: "K"},
		{Value: "1.5", Label: "Assets packed with power beyond your imagination.", Denom: "M+"},
		{Value: "750", Label: "Assets packed with power beyond your imagination.", Denom: "K"},
	},
}

var highlights = &models.Highlights{
	Heading:  "The Internet Rebuilt for You",
	Subtitle: "Sonr is a comprehensive system for Identity Management which proteects users across their digital personas while providing Developers a cost-effective solution for decentralized authentication.",
	Features: []*models.Feature{
		{
			Title: "Infinite-Factor Authentication",
			Desc:  "Sonr is designed to work across all platforms and devices, building a encrypted and anonymous identity layer for each user on the internet.",
			Icon:  nil,
		},
		{
			Title: "Self-Custody BTC & ETH",
			Desc:  "Sonr leverages advanced cryptography to permit facilitating Wallet Operations directly on-chain, without the need for a centralized server.",
			Icon:  nil,
		},
		{
			Title: "mAiNsTrEaM Ready",
			Desc:  "Sonr follows the latest specifications from W3C, DIF, and ICF to essentially have an Interchain-Connected, Smart Account System - seamlessly authenticated with PassKeys.",
			Icon:  nil,
		},
		{
			Title: "DAO Governed",
			Desc:  "Sonr is a proudly American Project which operates under the new Wyoming DUNA Legal Framework, ensuring the protection of your digital rights.",
			Icon:  nil,
		},
	},
}

var mission = &models.Mission{
	Eyebrow:  "L1 Blockchain",
	Heading:  "The Protocol for Decentralized Identity & Authentication",
	Subtitle: "We're creating the Global Standard for Decentralized Identity. Authenticate users with PassKeys, Issue Crypto Wallets, Build Payment flows, Send Encrypted Messages - all on a single platform.",
	Experience: &models.Feature{
		Title: "Experience",
		Desc:  "Sonr is a comprehensive system for Identity Management which proteects users across their digital personas while providing Developers a cost-effective solution for decentralized authentication.",
		Icon:  nil,
	},
	Compliance: &models.Feature{
		Title: "Compliance",
		Desc:  "Sonr is designed to work across all platforms and devices, building a encrypted and anonymous identity layer for each user on the internet.",
		Icon:  nil,
	},
	Interoperability: &models.Feature{
		Title: "Interoperability",
		Desc:  "Sonr follows the latest specifications from W3C, DIF, and ICF to essentially have an Interchain-Connected, Smart Account System - seamlessly authenticated with PassKeys.",
		Icon:  nil,
	},
	Standards: []*models.Feature{
		{
			Title: "Standards",
			Desc:  "Sonr is a proudly American Project which operates under the new Wyoming DUNA Legal Framework, ensuring the protection of your digital rights.",
			Icon:  nil,
		},
	},
}
