// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package sections

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/onsonr/sonr/app/nebula/global/ui"
	models "github.com/onsonr/sonr/app/nebula/models"
)

// ╭───────────────────────────────────────────────────────────╮
// │                         Data Model                        │
// ╰───────────────────────────────────────────────────────────╯

// hero is the (1st) home page hero section
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
		{Value: "476", Label: "Tradeable Crypto Assets", Denom: "+"},
		{Value: "1.44", Label: "Verified Identities", Denom: "K"},
		{Value: "1.5", Label: "Encrypted Messages Sent", Denom: "M+"},
		{Value: "50", Label: "Decentralized Global Nodes", Denom: "+"},
	},
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Render Section View                      │
// ╰───────────────────────────────────────────────────────────╯

// Hero Section
func Hero() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Hero --><section class=\"relative before:absolute before:inset-0 before:h-80 before:pointer-events-none before:bg-gradient-to-b before:from-zinc-100 before:-z-10\"><div class=\"pt-32 pb-12 md:pt-40 md:pb-20\"><!-- Section content --><div class=\"px-4 sm:px-6\"><div class=\"max-w-3xl mx-auto\"><div class=\"text-center pb-12 md:pb-16\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = heroTitle(hero.TitleFirst, hero.TitleEmphasis, hero.TitleSecond).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-lg text-zinc-500 mb-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(hero.Subtitle)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 49, Col: 22}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"max-w-xs mx-auto sm:max-w-none sm:inline-flex sm:justify-center space-y-4 sm:space-y-0 sm:space-x-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ui.PrimaryButton("/register", "Get Started").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ui.SecondaryButton("/about", "Learn More").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = heroImage(hero).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = stats(hero.Stats).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func heroTitle(first string, emphasis string, second string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"font-inter-tight text-4xl md:text-5xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-zinc-500 via-zinc-900 to-zinc-900 pb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(first)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 66, Col: 9}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <em class=\"italic relative inline-flex justify-center items-center text-zinc-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(emphasis)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 68, Col: 13}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <svg class=\"absolute fill-zinc-300 w-[calc(100%+1rem)] -z-10\" xmlns=\"http://www.w3.org/2000/svg\" width=\"223\" height=\"62\" viewBox=\"0 0 223 62\" aria-hidden=\"true\" preserveAspectRatio=\"none\"><path d=\"M45.654 53.62c17.666 3.154 35.622 4.512 53.558 4.837 17.94.288 35.91-.468 53.702-2.54 8.89-1.062 17.742-2.442 26.455-4.352 8.684-1.945 17.338-4.3 25.303-7.905 3.94-1.81 7.79-3.962 10.634-6.777 1.38-1.41 2.424-2.994 2.758-4.561.358-1.563-.078-3.143-1.046-4.677-.986-1.524-2.43-2.96-4.114-4.175a37.926 37.926 0 0 0-5.422-3.32c-3.84-1.977-7.958-3.563-12.156-4.933-8.42-2.707-17.148-4.653-25.95-6.145-8.802-1.52-17.702-2.56-26.622-3.333-17.852-1.49-35.826-1.776-53.739-.978-8.953.433-17.898 1.125-26.79 2.22-8.887 1.095-17.738 2.541-26.428 4.616-4.342 1.037-8.648 2.226-12.853 3.676-4.197 1.455-8.314 3.16-12.104 5.363-1.862 1.13-3.706 2.333-5.218 3.829-1.52 1.47-2.79 3.193-3.285 5.113-.528 1.912-.127 3.965.951 5.743 1.07 1.785 2.632 3.335 4.348 4.68 2.135 1.652 3.2 2.672 2.986 3.083-.18.362-1.674.114-4.08-1.638-1.863-1.387-3.63-3.014-4.95-5.09C.94 35.316.424 34.148.171 32.89c-.275-1.253-.198-2.579.069-3.822.588-2.515 2.098-4.582 3.76-6.276 1.673-1.724 3.612-3.053 5.57-4.303 3.96-2.426 8.177-4.278 12.457-5.868 4.287-1.584 8.654-2.89 13.054-4.036 8.801-2.292 17.74-3.925 26.716-5.19C70.777 2.131 79.805 1.286 88.846.723c18.087-1.065 36.236-.974 54.325.397 9.041.717 18.07 1.714 27.042 3.225 8.972 1.485 17.895 3.444 26.649 6.253 4.37 1.426 8.697 3.083 12.878 5.243a42.11 42.11 0 0 1 6.094 3.762c1.954 1.44 3.823 3.2 5.283 5.485a12.515 12.515 0 0 1 1.63 3.88c.164.706.184 1.463.253 2.193-.063.73-.094 1.485-.247 2.195-.652 2.886-2.325 5.141-4.09 6.934-3.635 3.533-7.853 5.751-12.083 7.688-8.519 3.778-17.394 6.09-26.296 7.998-8.917 1.86-17.913 3.152-26.928 4.104-18.039 1.851-36.17 2.295-54.239 1.622-18.062-.713-36.112-2.535-53.824-6.23-5.941-1.31-5.217-2.91.361-1.852\"></path></svg></em> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(second)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 73, Col: 10}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func heroImage(hero *models.Hero) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Image --><div class=\"max-w-5xl motion-preset-slide-up motion-opacity-in-0 motion-scale-in-50 mx-auto px-4 sm:px-6 flex justify-center pb-12 md:pb-20 relative before:absolute before:-top-12 before:w-96 before:h-96 before:bg-zinc-900 before:opacity-[.15] before:rounded-full before:blur-3xl before:-z-10 from-zinc-100 to-white\"><img class=\"rounded-lg \" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(hero.Image.Src)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 82, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(hero.Image.Width)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 83, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(hero.Image.Height)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 84, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"Hero\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func stats(stats []*models.Stat) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var11 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var11 == nil {
			templ_7745c5c3_Var11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Stats --><div class=\"max-w-4xl mx-auto px-4 sm:px-6 justify-center items-center\"><div class=\"max-w-sm mx-auto grid gap-12 sm:grid-cols-2 md:grid-cols-4 md:-mx-5 md:gap-0 items-end md:max-w-none\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range stats {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative text-center md:px-5 after:hidden md:after:block after:absolute after:right-0 after:top-1/2 after:-translate-y-1/2 after:w-px after:h-8 after:border-l after:border-zinc-300 after:border-dashed last:after:hidden\"><h4 class=\"font-inter-tight text-2xl md:text-3xl font-bold tabular-nums mb-2\"><span x-data=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(counterXData(item.Value))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 96, Col: 122}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-text=\"counterValue\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(item.Value)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 96, Col: 159}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(item.Denom)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 96, Col: 180}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h4><p class=\"text-sm text-zinc-500\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(item.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/sections/hero.templ`, Line: 97, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = counterAnimation().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func counterXData(value string) string {
	return fmt.Sprintf("counter(%s)", value)
}

func counterAnimation() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_counterAnimation_cac3`,
		Function: `function __templ_counterAnimation_cac3(){document.addEventListener('alpine:init', () => {
                      Alpine.data('counter', (target = 0, duration = 3000) => ({
                          startTimestamp: null,
                          step: null,
                          rawValue: 0,
                          counterValue: 0,
                          target: target,
                          precision: (target % 1 === 0) ? 0 : (target.toString().split('.')[1] || []).length,
                          animationRequestId: null,
                          animationCompleted: false,
                          observer: null,
                          init() {
                              // Intersection observer to watch visibility
                              this.observer = new IntersectionObserver(entries => {
                                  entries.forEach(entry => {
                                      // Check if element is in view
                                      if (entry.isIntersecting && !this.animationCompleted) {
                                          this.startAnimation()
                                      }
                                  })
                              })
                              this.observer.observe(this.$el)
                          },
                          startAnimation() {
                              this.step = (timestamp) => {
                                  if (!this.startTimestamp) this.startTimestamp = timestamp
                                  const progress = Math.min((timestamp - this.startTimestamp) / duration, 1)
                                  const easedProgress = this.easeOut(progress)
                                  this.rawValue = parseFloat((easedProgress * this.target).toFixed(this.precision))
                                  this.counterValue = this.rawValue.toFixed(this.precision)
                                  if (progress < 1) {
                                      this.animationRequestId = window.requestAnimationFrame(this.step)
                                  } else {
                                      this.animationCompleted = true
                                  }
                              }
                              this.animationRequestId = window.requestAnimationFrame(this.step);
                          },
                          easeOut(t) {
                              return 1 - Math.pow(1 - t, 5)
                          },
                          destroy() {
                              // Detach the handler, avoiding memory and side-effect leakage
                              this.animationRequestId && window.cancelAnimationFrame(this.step)
                              this.observer && this.observer.disconnect()
                          },
                      }))
                  })
}`,
		Call:       templ.SafeScript(`__templ_counterAnimation_cac3`),
		CallInline: templ.SafeScriptInline(`__templ_counterAnimation_cac3`),
	}
}

var _ = templruntime.GeneratedTemplate
