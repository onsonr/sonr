// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package landing

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import models "github.com/onsonr/sonr/pkg/webapp/models"

// ╭───────────────────────────────────────────────────────────╮
// │                         Data Model                        │
// ╰───────────────────────────────────────────────────────────╯

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

// ╭───────────────────────────────────────────────────────────╮
// │                  Render Section View                      │
// ╰───────────────────────────────────────────────────────────╯

// Lowlights is the (4th) home page lowlights section
func Lowlights() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"bg-zinc-800\"><div class=\"py-12 md:py-20\"><div class=\"max-w-5xl mx-auto px-4 sm:px-6\"><div class=\"max-w-3xl mx-auto text-center pb-12 md:pb-20\"><h2 class=\"font-inter-tight text-3xl md:text-4xl font-bold text-zinc-200\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(lowlights.Heading)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/landing/lowlights.templ`, Line: 108, Col: 25}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2></div></div><div class=\"max-w-[94rem] mx-auto space-y-6\"><!-- Row #1 --><div x-data=\"{}\" x-init=\"$nextTick(() =&gt; {\n                                let ul = $refs.testimonials;\n                                ul.insertAdjacentHTML(&#39;afterend&#39;, ul.outerHTML);\n                                ul.nextSibling.setAttribute(&#39;aria-hidden&#39;, &#39;true&#39;);\n                            })\" class=\"w-full inline-flex flex-nowrap overflow-hidden [mask-image:_linear-gradient(to_right,transparent_0,_black_28%,_black_calc(100%-28%),transparent_100%)] group\"><div x-ref=\"testimonials\" class=\"flex items-start justify-center md:justify-start [&amp;&gt;div]:mx-3 animate-infinite-scroll group-hover:[animation-play-state:paused]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, quote := range lowlights.UpperQuotes {
			templ_7745c5c3_Err = quoteItem(quote).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><!-- Row #2 --><div x-data=\"{}\" x-init=\"$nextTick(() =&gt; {\n                                let ul = $refs.testimonials;\n                                ul.insertAdjacentHTML(&#39;afterend&#39;, ul.outerHTML);\n                                ul.nextSibling.setAttribute(&#39;aria-hidden&#39;, &#39;true&#39;);\n                            })\" class=\"w-full inline-flex flex-nowrap overflow-hidden [mask-image:_linear-gradient(to_right,transparent_0,_black_28%,_black_calc(100%-28%),transparent_100%)] group\"><div x-ref=\"testimonials\" class=\"flex items-start justify-center md:justify-start [&amp;&gt;div]:mx-3 animate-infinite-scroll-inverse group-hover:[animation-play-state:paused] [animation-delay:-7.5s]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, quote := range lowlights.LowerQuotes {
			templ_7745c5c3_Err = quoteItem(quote).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func quoteItem(quote *models.Testimonial) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"rounded h-full w-[22rem] border border-transparent [background:linear-gradient(#323237,#323237)_padding-box,linear-gradient(120deg,theme(colors.zinc.700),theme(colors.zinc.700/0),theme(colors.zinc.700))_border-box] p-5\"><div class=\"flex items-center mb-4\"><img class=\"shrink-0 rounded-full mr-3\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(quote.Avatar.Src)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/landing/lowlights.templ`, Line: 153, Col: 65}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(quote.Avatar.Width)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/landing/lowlights.templ`, Line: 153, Col: 94}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(quote.Avatar.Height)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/landing/lowlights.templ`, Line: 153, Col: 125}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"Testimonial 01\"><div><div class=\"font-inter-tight font-bold text-zinc-200\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(quote.FullName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/landing/lowlights.templ`, Line: 155, Col: 74}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div><a class=\"text-sm font-medium text-zinc-500 hover:text-zinc-300 transition\" href=\"#0\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(quote.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/landing/lowlights.templ`, Line: 157, Col: 107}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div></div></div><div class=\"text-zinc-500 before:content-[&#39;\\0022&#39;] after:content-[&#39;\\0022&#39;]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(quote.Quote)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/webapp/components/landing/lowlights.templ`, Line: 162, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
