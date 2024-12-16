package models

import "encoding/json"

func NewWebManifest() ([]byte, error) {
	return json.Marshal(baseWebManifest)
}

var baseWebManifest = WebManifest{
	Name:      "Sonr Vault",
	ShortName: "Sonr.ID",
	StartURL:  "/index.html",
	Display:   "standalone",
	DisplayOverride: []string{
		"fullscreen",
		"minimal-ui",
	},
	Icons: []IconDefinition{
		{
			Src:   "/icons/icon-192x192.png",
			Sizes: "192x192",
			Type:  "image/png",
		},
	},
	ServiceWorker: ServiceWorker{
		Scope:    "/",
		Src:      "/sw.js",
		UseCache: true,
	},
	ProtocolHandlers: []ProtocolHandler{
		{
			Scheme: "did.sonr",
			URL:    "/resolve/sonr/%s",
		},
		{
			Scheme: "did.eth",
			URL:    "/resolve/eth/%s",
		},
		{
			Scheme: "did.btc",
			URL:    "/resolve/btc/%s",
		},
		{
			Scheme: "did.usdc",
			URL:    "/resolve/usdc/%s",
		},
		{
			Scheme: "did.ipfs",
			URL:    "/resolve/ipfs/%s",
		},
	},
}

type WebManifest struct {
	// Required fields
	Name      string `json:"name"`       // Full name of the application
	ShortName string `json:"short_name"` // Short version of the name

	// Display and appearance
	Description     string   `json:"description,omitempty"` // Purpose and features of the application
	Display         string   `json:"display,omitempty"`     // Preferred display mode: fullscreen, standalone, minimal-ui, browser
	DisplayOverride []string `json:"display_override,omitempty"`
	ThemeColor      string   `json:"theme_color,omitempty"`      // Default theme color for the application
	BackgroundColor string   `json:"background_color,omitempty"` // Background color during launch
	Orientation     string   `json:"orientation,omitempty"`      // Default orientation: any, natural, landscape, portrait

	// URLs and scope
	StartURL      string        `json:"start_url"`       // Starting URL when launching
	Scope         string        `json:"scope,omitempty"` // Navigation scope of the web application
	ServiceWorker ServiceWorker `json:"service_worker,omitempty"`

	// Icons
	Icons []IconDefinition `json:"icons,omitempty"`

	// Optional features
	RelatedApplications       []RelatedApplication `json:"related_applications,omitempty"`
	PreferRelatedApplications bool                 `json:"prefer_related_applications,omitempty"`
	Shortcuts                 []Shortcut           `json:"shortcuts,omitempty"`

	// Experimental features (uncomment if needed)
	FileHandlers     []FileHandler     `json:"file_handlers,omitempty"`
	ProtocolHandlers []ProtocolHandler `json:"protocol_handlers,omitempty"`
}

type FileHandler struct {
	Action string              `json:"action"`
	Accept map[string][]string `json:"accept"`
}

type LaunchHandler struct {
	Action string `json:"action"`
}

type IconDefinition struct {
	Src     string `json:"src"`
	Sizes   string `json:"sizes"`
	Type    string `json:"type,omitempty"`
	Purpose string `json:"purpose,omitempty"`
}

type ProtocolHandler struct {
	Scheme string `json:"scheme"`
	URL    string `json:"url"`
}

type RelatedApplication struct {
	Platform string `json:"platform"`
	URL      string `json:"url,omitempty"`
	ID       string `json:"id,omitempty"`
}

type Shortcut struct {
	Name        string           `json:"name"`
	ShortName   string           `json:"short_name,omitempty"`
	Description string           `json:"description,omitempty"`
	URL         string           `json:"url"`
	Icons       []IconDefinition `json:"icons,omitempty"`
}

type ServiceWorker struct {
	Scope    string `json:"scope"`
	Src      string `json:"src"`
	UseCache bool   `json:"use_cache"`
}
