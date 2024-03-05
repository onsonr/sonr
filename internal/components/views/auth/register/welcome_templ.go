// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package register

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/sonrhq/sonr/internal/components/views/auth"
)

func welcomeAccountScreen() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"grow\"><section class=\"relative\"><!-- Illustration --><div class=\"md:block absolute left-1/2 -translate-x-1/2 -mt-36 blur-2xl opacity-70 pointer-events-none -z-10\" aria-hidden=\"true\"><img src=\"https://cdn.sonr.build/images/auth-illustration.svg\" class=\"max-w-none\" width=\"1440\" height=\"450\" alt=\"Page Illustration\"></div><div class=\"relative max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"pt-32 pb-12 md:pt-40 md:pb-20\"><!-- Page header --><div class=\"max-w-3xl mx-auto text-center pb-12\"><!-- Logo --><div class=\"mb-5\"><a class=\"inline-flex\" href=\"/\"><div class=\"relative flex items-center justify-center w-16 h-16 border border-transparent rounded-2xl shadow-2xl [background:linear-gradient(theme(colors.slate.900),_theme(colors.slate.900))_padding-box,_conic-gradient(theme(colors.slate.400),_theme(colors.slate.700)_25%,_theme(colors.slate.700)_75%,_theme(colors.slate.400)_100%)_border-box] before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-2xl\"><img class=\"relative\" src=\"https://cdn.sonr.build/images/logo.svg\" width=\"42\" height=\"42\" alt=\"Sonr\"></div></a></div><!-- Page title --><h1 class=\"h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60\">Welcome to Sonr</h1></div><!-- Form --><div class=\"max-w-sm mx-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = emailIdentifierForm().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-center mt-4\"><div class=\"text-sm text-slate-400\">Already have an account? <a class=\"font-medium text-purple-500 hover:text-purple-400 transition duration-150 ease-in-out\" href=\"/login\">Sign in</a></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = auth.RecoveryFooter().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></section></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}