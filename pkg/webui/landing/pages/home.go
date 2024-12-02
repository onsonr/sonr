package pages

import "github.com/onsonr/sonr/pkg/webui/landing/models"

var header = &models.NavHeader{
	Logo: &models.Image{
		Src:    "https://cdn.sonr.id/logo-ivory.svg",
		Width:  "20",
		Height: "20",
	},
	Primary: &models.NavItem{
		Text: "Get Access",
		Href: "/register",
	},
	Items: []*models.NavItem{
		{
			Text: "Docs",
			Href: "#",
		},
		{
			Text: "Blog",
			Href: "#",
		},
		{
			Text: "Changelog",
			Href: "#",
		},
		{
			Text: "About",
			Href: "#",
		},
	},
}

// hero is the (1st) home page hero section
var hero = &models.Hero{
	TitleFirst:      "Sonr puts",
	TitleEmphasis:   "digital identity",
	TitleSecond:     "back in your hands",
	Subtitle:        "We're creating the Global Standard for Decentralized Identity. Authenticate users with PassKeys, Issue Crypto Wallets, Build Payment flows, Send Encrypted Messages - all on a single platform.",
	PrimaryButton:   &models.Button{Text: "Get Started", Href: "/register"},
	SecondaryButton: &models.Button{Text: "Learn More", Href: "/about"},
	Image: &models.Image{
		Src:    "https://cdn.sonr.id/img/hero-clipped.svg",
		Width:  "500",
		Height: "500",
	},
	Stats: []*models.Stat{
		{Value: "476", Label: "Tradeable Crypto Assets", Denom: "+"},
		{Value: "1.44", Label: "Verified Identities", Denom: "K"},
		{Value: "1.5", Label: "Encrypted Messages Sent", Denom: "M+"},
		{Value: "50", Label: "Decentralized Global Nodes", Denom: "+"},
	},
}

// highlights is the (2nd) home page highlights section
var highlights = &models.Highlights{
	Heading:  "The Internet Rebuilt for You",
	Subtitle: "Sonr is a comprehensive system for Identity Management which proteects users across their digital personas while providing Developers a cost-effective solution for decentralized authentication.",
	Features: []*models.Feature{
		{
			Title: "Crypto Wallet",
			Desc:  "Sonr is designed to work across all platforms and devices, building a encrypted and anonymous identity layer for each user on the internet.",
			Icon:  nil,
			Image: &models.Image{
				Src:    "",
				Width:  "44",
				Height: "44",
			},
		},
		{
			Title: "PassKey Authenticator",
			Desc:  "Sonr leverages advanced cryptography to permit facilitating Wallet Operations directly on-chain, without the need for a centralized server.",
			Icon:  nil,
			Image: &models.Image{
				Src:    "",
				Width:  "44",
				Height: "44",
			},
		},
		{
			Title: "Anonymous Identity",
			Desc:  "Sonr follows the latest specifications from W3C, DIF, and ICF to essentially have an Interchain-Connected, Smart Account System - seamlessly authenticated with PassKeys.",
			Icon:  nil,
			Image: &models.Image{
				Src:    "",
				Width:  "44",
				Height: "44",
			},
		},
		{
			Title: "Tokenized Authorization",
			Desc:  "Sonr anonymously associates your online identities with a Quantum-Resistant Vault which only you can access.",
			Icon:  nil,
			Image: &models.Image{
				Src:    "",
				Width:  "44",
				Height: "44",
			},
		},
	},
}

// mission is the (3rd) home page mission section
var mission = &models.Mission{
	Eyebrow:  "L1 Blockchain",
	Heading:  "The Protocol for Decentralized Identity & Authentication",
	Subtitle: "We're creating the Global Standard for Decentralized Identity. Authenticate users with PassKeys, Issue Crypto Wallets, Build Payment flows, Send Encrypted Messages - all on a single platform.",
	Experience: &models.Feature{
		Title: "UX First Approach",
		Desc:  "Sonr is a comprehensive system for Identity Management which proteects users across their digital personas while providing Developers a cost-effective solution for decentralized authentication.",
		Icon:  nil,
	},
	Compliance: &models.Feature{
		Title: "Universal Interoperability",
		Desc:  "Sonr is designed to work across all platforms and devices, building a encrypted and anonymous identity layer for each user on the internet.",
		Icon:  nil,
	},
	Interoperability: &models.Feature{
		Title: "Made in the USA",
		Desc:  "Sonr follows the latest specifications from W3C, DIF, and ICF to essentially have an Interchain-Connected, Smart Account System - seamlessly authenticated with PassKeys.",
		Icon:  nil,
	},
}

// architecture is the (4th) home page architecture section
var arch = &models.Architecture{
	Heading:  "Onchain Security with Offchain Privacy",
	Subtitle: "Whenever you are ready, just hit publish to turn your site sketches into an actual designs. No creating, no skills, no reshaping.",
	Primary: &models.Technology{
		Title: "Decentralized Identity",
		Desc:  "Sonr leverages the latest specifications from W3C, DIF, and ICF to essentially have an Interchain-Connected, Smart Account System - seamlessly authenticated with PassKeys.",
		Image: &models.Image{
			Src:    models.HardwareWallet.Src(),
			Width:  "721",
			Height: "280",
		},
	},
	Secondary: &models.Technology{
		Title: "IPFS Vaults",
		Desc:  "Completely distributed, encrypted, and decentralized storage for your data.",
		Image: &models.Image{
			Src:    models.DecentralizedNetwork.Src(),
			Width:  "342",
			Height: "280",
		},
	},
	Tertiary: &models.Technology{
		Title: "Service Records",
		Desc:  "On-chain validated services created by Developers for secure transmission of user data.",
		Image: &models.Image{
			Src:    models.DefiDashboard.Src(),
			Width:  "342",
			Height: "280",
		},
	},
	Quaternary: &models.Technology{
		Title: "Authentication & Authorization",
		Desc:  "Sonr leverages decentralized Macaroons and Multi-Party Computation to provide a secure and decentralized authentication and authorization system.",
		Image: &models.Image{
			Src:    models.PrivateKey.Src(),
			Width:  "342",
			Height: "280",
		},
	},
	Quinary: &models.Technology{
		Title: "First-Class Exchange",
		Desc:  "Sonr integrates with the IBC protocol allowing for seamless integration with popular exchanges such as OKX, Binance, and Osmosis.",
		Image: &models.Image{
			Src:    models.CrossChainBridge.Src(),
			Width:  "342",
			Height: "280",
		},
	},
}

var lowlights = &models.Lowlights{
	Heading: "The Fragmentation Problem in the Existing Web is seeping into Crypto",
	UpperQuotes: []*models.Testimonial{
		{
			FullName: "0xDesigner",
			Username: "@0xDesigner",
			Avatar: &models.Image{
				Src:    models.Avatar0xDesigner.Src(),
				Width:  "44",
				Height: "44",
			},
			Quote: "what if the wallet ui appeared next to the click instead of in a new browser window?",
		},
		{
			FullName: "Alex Recouso",
			Username: "@alexrecouso",
			Avatar: &models.Image{
				Src:    models.AvatarAlexRecouso.Src(),
				Width:  "44",
				Height: "44",
			},
			Quote: "2024 resembles 1984, but it doesn't have to be that way for you",
		},
		{
			FullName: "Chjango Unchained",
			Username: "@chjango",
			Avatar: &models.Image{
				Src:    models.AvatarChjango.Src(),
				Width:  "44",
				Height: "44",
			},
			Quote: "IBC is the inter-blockchain highway of @cosmos. While not very cypherpunk, charging a 1.5 basis pt fee would go a long way if priced in $ATOM.",
		},
		{
			FullName: "Gwart",
			Username: "@GwartyGwart",
			Avatar: &models.Image{
				Src:    models.AvatarGwart.Src(),
				Width:  "44",
				Height: "44",
			},
			Quote: "	Base is incredible. Most centralized l2. Least details about their plans to decentralize. Keeps OP cabal quiet by pretending to care about quadratic voting and giving 10% tithe. Pays Ethereum mainnet virtually nothing. Runs yuppie granola ad campaigns.",
		},
	},
	LowerQuotes: []*models.Testimonial{
		{
			FullName: "winnie",
			Username: "@winnielaux_",
			Avatar: &models.Image{
				Src:    models.AvatarWinnieLaux.Src(),
				Width:  "44",
				Height: "44",
			},
			Quote: "the ability to download apps directly from the web or from “crypto-only” app stores will be a massive unlock for web3",
		},
		{
			FullName: "Jelena",
			Username: "@jelena_noble",
			Avatar: &models.Image{
				Src:    models.AvatarJelenaNoble.Src(),
				Width:  "44",
				Height: "44",
			},
			Quote: "Excited for all the @cosmos nerds to be vindicated in the next bull run",
		},
		{
			FullName: "accountless",
			Username: "@alexanderchopan",
			Avatar: &models.Image{
				Src:    models.AvatarAccountless.Src(),
				Width:  "44",
				Height: "44",
			},
			Quote: "sounds like webThree. Single key pair Requires the same signer At risk of infinite approvals Public history of all transactions different account on each chain different addresses for each account",
		},
		{
			FullName: "Unusual Whales",
			Username: "@unusual_whales",
			Avatar: &models.Image{
				Src:    models.AvatarUnusualWhales.Src(),
				Width:  "44",
				Height: "44",
			},
			Quote: "BREAKING: Fidelity & Fidelity Investments has confirmed that over 77,000 customers had personal information compromised, including Social Security numbers and driver’s licenses.",
		},
	},
}

var cta = &models.CallToAction{
	Logo: &models.Image{
		Src:    "https://cdn.sonr.id/logo-zinc.svg",
		Width:  "60",
		Height: "60",
	},
	Heading:  "Take control of your Identity",
	Subtitle: "Sonr is a decentralized, permissionless, and censorship-resistant identity network.",
	Primary: &models.Button{
		Href: "request-demo.html",
		Text: "Register",
	},
	Secondary: &models.Button{
		Href: "#0",
		Text: "Learn More",
	},
}
