// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package research

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/didao-org/sonr/internal/components/base"
	"github.com/didao-org/sonr/internal/views/landing"
	"github.com/labstack/echo/v4"
)

func Page(ctx echo.Context) templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = landing.Navbar().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <main class=\"grow\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = contentSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ctaSection().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = landing.Footer().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = base.PageLayout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func contentSection() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Content --><section class=\"relative\"><!-- Radial gradient --><div class=\"absolute flex items-center justify-center top-0 -translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-[800px] aspect-square\" aria-hidden=\"true\"><div class=\"absolute inset-0 translate-z-0 bg-purple-500 rounded-full blur-[120px] opacity-30\"></div><div class=\"absolute w-64 h-64 translate-z-0 bg-purple-400 rounded-full blur-[80px] opacity-70\"></div></div><!-- Particles animation --><div class=\"absolute inset-0 h-96 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"20\"></canvas></div><!-- Illustration --><div class=\"md:block absolute left-1/2 -translate-x-1/2 -mt-16 blur-2xl opacity-90 pointer-events-none -z-10\" aria-hidden=\"true\"><img src=\"https://cdn.sonr.build/images/page-illustration.svg\" class=\"max-w-none\" width=\"1440\" height=\"427\" alt=\"Page Illustration\"></div><div class=\"max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"pt-32 md:pt-40\"><!-- Section header --><div class=\"text-center pb-12 md:pb-20\"><div class=\"inline-flex font-medium bg-clip-text text-transparent bg-gradient-to-r from-purple-500 to-purple-200 pb-3\">Leaders love Sonr</div><h1 class=\"h1 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4\">Our Research into Identity</h1><div class=\"max-w-3xl mx-auto\"><p class=\"text-lg text-slate-400\">Sonr has been under development since 2021 and has grown to be a large project. We write about our platform structure with these documents.</p></div></div><!-- Customers grid --><div class=\"max-w-[352px] mx-auto sm:max-w-[728px] lg:max-w-none pb-12 md:pb-20\"><div class=\"grid gap-6 sm:grid-cols-2 lg:grid-cols-3 group [&amp;_*:nth-child(n+5):not(:nth-child(n+12))]:order-1 [&amp;_*:nth-child(n+10):not(:nth-child(n+11))]:!order-2\" data-highlighter><!-- Item #01 --><div><a href=\"customer.html\"><div class=\"relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] before:hover:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:_radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),theme(colors.slate.400),transparent)] after:group-hover:opacity-100 after:z-10 overflow-hidden\"><div class=\"relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"3\"></canvas></div><div class=\"flex items-center justify-center\"><img class=\"w-full h-full aspect-video object-cover\" src=\"https://cdn.sonr.build/images/customer-bg-01.png\" width=\"352\" height=\"198\" alt=\"Customer Background 01\" aria-hidden=\"true\"> <img class=\"absolute\" src=\"https://cdn.sonr.build/images/research-blockchain.svg\" width=\"110\" height=\"21\" alt=\"Customer 01\"></div></div></div></a></div><!-- Item #02 --><div><a href=\"customer.html\"><div class=\"relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] before:hover:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:_radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),theme(colors.slate.400),transparent)] after:group-hover:opacity-100 after:z-10 overflow-hidden\"><div class=\"relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"3\"></canvas></div><div class=\"flex items-center justify-center\"><img class=\"w-full h-full aspect-video object-cover\" src=\"https://cdn.sonr.build/images/customer-bg-02.png\" width=\"352\" height=\"198\" alt=\"Customer Background 02\" aria-hidden=\"true\"> <img class=\"absolute\" src=\"https://cdn.sonr.build/images/research-gov.svg\" width=\"110\" height=\"21\" alt=\"Customer 02\"></div></div></div></a></div><!-- Item #03 --><div><a href=\"customer.html\"><div class=\"relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] before:hover:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:_radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),theme(colors.slate.400),transparent)] after:group-hover:opacity-100 after:z-10 overflow-hidden\"><div class=\"relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"3\"></canvas></div><div class=\"flex items-center justify-center\"><img class=\"w-full h-full aspect-video object-cover\" src=\"https://cdn.sonr.build/images/customer-bg-03.png\" width=\"352\" height=\"198\" alt=\"Customer Background 03\" aria-hidden=\"true\"> <img class=\"absolute\" src=\"https://cdn.sonr.build/images/research-identity.svg\" width=\"110\" height=\"21\" alt=\"Customer 03\"></div></div></div></a></div><!-- Item #04 --><div><a href=\"customer.html\"><div class=\"relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] before:hover:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:_radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),theme(colors.slate.400),transparent)] after:group-hover:opacity-100 after:z-10 overflow-hidden\"><div class=\"relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"3\"></canvas></div><div class=\"flex items-center justify-center\"><img class=\"w-full h-full aspect-video object-cover\" src=\"https://cdn.sonr.build/images/customer-bg-04.png\" width=\"352\" height=\"198\" alt=\"Customer Background 04\" aria-hidden=\"true\"> <img class=\"absolute\" src=\"https://cdn.sonr.build/images/research-service.svg\" width=\"110\" height=\"21\" alt=\"Customer 04\"></div></div></div></a></div><!-- Testimonial #01 --><div class=\"flex flex-col items-center justify-center text-center p-4\"><p class=\"font-medium bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-3\"><span class=\"line-clamp-4\">“We struggled to bring all our conversations into one place until we found Sonr. The UI is very clean and we love the integration with Spark.”</span></p><div class=\"inline-flex mb-2\"></div><div class=\"text-sm font-medium text-slate-300\">Mike Hunt <span class=\"text-slate-700\">-</span> <a class=\"text-purple-500 hover:underline\" href=\"#0\">Thunderbolt</a></div></div><!-- Item #05 --><div><a href=\"customer.html\"><div class=\"relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] before:hover:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:_radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),theme(colors.slate.400),transparent)] after:group-hover:opacity-100 after:z-10 overflow-hidden\"><div class=\"relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"3\"></canvas></div><div class=\"flex items-center justify-center\"><img class=\"w-full h-full aspect-video object-cover\" src=\"https://cdn.sonr.build/images/customer-bg-05.png\" width=\"352\" height=\"198\" alt=\"Customer Background 05\" aria-hidden=\"true\"> <img class=\"absolute\" src=\"https://cdn.sonr.build/images/research-networking.svg\" width=\"110\" height=\"21\" alt=\"Customer 05\"></div></div></div></a></div><!-- Item #06 --><div><a href=\"customer.html\"><div class=\"relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] before:hover:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:_radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),theme(colors.slate.400),transparent)] after:group-hover:opacity-100 after:z-10 overflow-hidden\"><div class=\"relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"3\"></canvas></div><div class=\"flex items-center justify-center\"><img class=\"w-full h-full aspect-video object-cover\" src=\"https://cdn.sonr.build/images/customer-bg-06.png\" width=\"352\" height=\"198\" alt=\"Customer Background 06\" aria-hidden=\"true\"> <img class=\"absolute\" src=\"https://cdn.sonr.build/images/research-security.svg\" width=\"110\" height=\"21\" alt=\"Customer 06\"></div></div></div></a></div><!-- Item #07 --><div><a href=\"customer.html\"><div class=\"relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] before:hover:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:_radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),theme(colors.slate.400),transparent)] after:group-hover:opacity-100 after:z-10 overflow-hidden\"><div class=\"relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"3\"></canvas></div><div class=\"flex items-center justify-center\"><img class=\"w-full h-full aspect-video object-cover\" src=\"https://cdn.sonr.build/images/customer-bg-07.png\" width=\"352\" height=\"198\" alt=\"Customer Background 07\" aria-hidden=\"true\"> <img class=\"absolute\" src=\"https://cdn.sonr.build/images/research-economics.svg\" width=\"110\" height=\"21\" alt=\"Customer 07\"></div></div></div></a></div><!-- Item #08 --><div><a href=\"customer.html\"><div class=\"relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] before:hover:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:_radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),theme(colors.slate.400),transparent)] after:group-hover:opacity-100 after:z-10 overflow-hidden\"><div class=\"relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"3\"></canvas></div><div class=\"flex items-center justify-center\"><img class=\"w-full h-full aspect-video object-cover\" src=\"https://cdn.sonr.build/images/customer-bg-08.png\" width=\"352\" height=\"198\" alt=\"Customer Background 08\" aria-hidden=\"true\"> <img class=\"absolute\" src=\"https://cdn.sonr.build/images/research-staking.svg\" width=\"110\" height=\"21\" alt=\"Customer 08\"></div></div></div></a></div></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ctaSection() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"relative\"><!-- Particles animation --><div class=\"absolute left-1/2 -translate-x-1/2 top-0 -z-10 w-80 h-80 -mt-24\"><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"6\" data-particle-staticity=\"30\"></canvas></div></div><div class=\"max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"relative px-8 py-12 md:py-20 border-t border-b [border-image:linear-gradient(to_right,transparent,theme(colors.slate.800),transparent)1]\"><!-- Blurred shape --><div class=\"absolute top-0 -mt-24 left-1/2 -translate-x-1/2 ml-24 blur-2xl opacity-70 pointer-events-none -z-10\" aria-hidden=\"true\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"434\" height=\"427\"><defs><linearGradient id=\"bs4-a\" x1=\"19.609%\" x2=\"50%\" y1=\"14.544%\" y2=\"100%\"><stop offset=\"0%\" stop-color=\"#A855F7\"></stop> <stop offset=\"100%\" stop-color=\"#6366F1\" stop-opacity=\"0\"></stop></linearGradient></defs> <path fill=\"url(#bs4-a)\" fill-rule=\"evenodd\" d=\"m0 0 461 369-284 58z\" transform=\"matrix(1 0 0 -1 0 427)\"></path></svg></div><!-- Content --><div class=\"max-w-3xl mx-auto text-center\"><div><div class=\"inline-flex font-medium bg-clip-text text-transparent bg-gradient-to-r from-purple-500 to-purple-200 pb-3\">The security first platform</div></div><h2 class=\"h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4\">Supercharge your security</h2><p class=\"text-lg text-slate-400 mb-8\">All the lorem ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet.</p><div><a class=\"btn text-slate-900 bg-gradient-to-r from-white/80 via-white to-white/80 hover:bg-white transition duration-150 ease-in-out group\" href=\"#0\">Start Building <span class=\"tracking-normal text-purple-500 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1\">-&gt;</span></a></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
