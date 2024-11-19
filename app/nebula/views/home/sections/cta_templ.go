// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package sections

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import models "github.com/onsonr/sonr/app/nebula/models"

// ╭───────────────────────────────────────────────────────────╮
// │                         Data Model                        │
// ╰───────────────────────────────────────────────────────────╯

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

// ╭───────────────────────────────────────────────────────────╮
// │                  Render Section View                      │
// ╰───────────────────────────────────────────────────────────╯
func CallToAction() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section><div class=\"py-12 md:py-20\"><div class=\"max-w-5xl mx-auto px-4 sm:px-6\"><div class=\"relative max-w-3xl mx-auto text-center pb-12 md:pb-16\"><div class=\"inline-flex items-center justify-center w-20 h-20 bg-white rounded-xl shadow-md mb-8 relative before:absolute before:-top-12 before:w-52 before:h-52 before:bg-zinc-900 before:opacity-[.08] before:rounded-full before:blur-3xl before:-z-10\"><a href=\"index.html\"><img src=\"https://cdn.sonr.id/logo-zinc.svg\" width=\"60\" height=\"60\" alt=\"Logo\"></a></div><h2 class=\"font-inter-tight text-3xl md:text-4xl font-bold text-zinc-900 mb-4\">Take control of your Identity <em class=\"relative not-italic inline-flex justify-center items-end\">today <svg class=\"absolute fill-zinc-300 w-[calc(100%+1rem)] -z-10\" xmlns=\"http://www.w3.org/2000/svg\" width=\"120\" height=\"10\" viewBox=\"0 0 120 10\" aria-hidden=\"true\" preserveAspectRatio=\"none\"><path d=\"M118.273 6.09C79.243 4.558 40.297 5.459 1.305 9.034c-1.507.13-1.742-1.521-.199-1.81C39.81-.228 79.647-1.568 118.443 4.2c1.63.233 1.377 1.943-.17 1.89Z\"></path></svg></em></h2><p class=\"text-lg text-zinc-500 mb-8\">Sonr removes creative distances by connecting beginners, pros, and every team in between. Are you ready to start your journey?</p><div class=\"max-w-xs mx-auto sm:max-w-none sm:inline-flex sm:justify-center space-y-4 sm:space-y-0 sm:space-x-4\"><div><a class=\"btn text-zinc-100 bg-zinc-900 hover:bg-zinc-800 w-full shadow\" href=\"request-demo.html\">Register</a></div><div><a class=\"btn text-zinc-600 bg-white hover:text-zinc-900 w-full shadow\" href=\"#0\">Log in</a></div></div></div><!-- Clients --><div class=\"text-center\"><ul class=\"inline-flex flex-wrap items-center justify-center -m-2 [mask-image:linear-gradient(to_right,transparent_8px,_theme(colors.white/.7)_64px,_theme(colors.white)_50%,_theme(colors.white/.7)_calc(100%-64px),_transparent_calc(100%-8px))]\"><li class=\"m-2 p-4 relative rounded-lg border border-transparent [background:linear-gradient(theme(colors.zinc.50),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box]\"><svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" aria-label=\"Fido Alliance\" height=\"40\" width=\"40\"><path d=\"M7.849 7.513a1.085 1.085 0 1 0 1.085 1.086v-.001c0-.599-.486-1.085-1.085-1.085zM4.942 10.553v1.418H6.89v4.793h.704V14.04h.509v2.724h.71v-6.211H4.941zM14.122 11.089H14.1c-.287-.416-.862-.702-1.639-.702-1.489 0-2.797 1.224-2.786 3.319 0 1.936 1.181 3.201 2.659 3.201.797 0 1.56-.361 1.935-1.04l.117.893h1.669V7.651h-1.934zm0 2.904c0 .158-.012.313-.034.465l.002-.017c-.11.532-.574.925-1.13.925h-.014.001c-.797 0-1.318-.659-1.318-1.723 0-.978.446-1.767 1.329-1.767.606 0 1.022.437 1.138.947.014.09.023.194.023.3l-.001.054v-.003zM4.802 8.89l.475-1.6a2.914 2.914 0 0 0-.384-.101l-.019-.003a3.654 3.654 0 0 0-.829-.092 3.73 3.73 0 0 0-1.084.159l.027-.007a2.022 2.022 0 0 0-.38.153l.011-.005a2.624 2.624 0 0 0-.663.475c-.5.49-.754 1.155-.754 1.975v.708H-.001v1.418h1.199v4.793h1.921V11.97h1.199v-1.416H3.119v-.75a1.019 1.019 0 0 1 .23-.713l-.001.002a.736.736 0 0 1 .063-.062l.001-.001s.414-.41 1.389-.14zM20.306 10.388c-2.01 0-3.327 1.286-3.327 3.307s1.393 3.212 3.213 3.212c1.664 0 3.276-1.04 3.276-3.327-.002-1.874-1.267-3.192-3.162-3.192zm-.063 5.126c-.832 0-1.276-.797-1.276-1.871 0-.915.361-1.861 1.276-1.861.871 0 1.234.936 1.234 1.851 0 1.137-.482 1.882-1.234 1.882zM22.493 9.761h.232v.589h.14v-.589h.231v-.117h-.603v.117zM23.799 9.644l-.182.505-.181-.505h-.203v.707h.13V9.78l.198.571h.113l.195-.571v.571h.13v-.707h-.201z\"></path></svg><!-- \t\t\t\t\t\t\t<svg role=\"img\" viewBox=\"0 0 24 24\" width=\"40\" height=\"40\" aria-label=\"IPFS\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M12 0L1.608 6v12L12 24l10.392-6V6zm-1.073 1.445h.001a1.8 1.8 0 002.138 0l7.534 4.35a1.794 1.794 0 000 .403l-7.535 4.35a1.8 1.8 0 00-2.137 0l-7.536-4.35a1.795 1.795 0 000-.402zM21.324 7.4c.109.08.226.147.349.201v8.7a1.8 1.8 0 00-1.069 1.852l-7.535 4.35a1.8 1.8 0 00-.349-.2l-.009-8.653a1.8 1.8 0 001.07-1.851zm-18.648.048l7.535 4.35a1.8 1.8 0 001.069 1.852v8.7c-.124.054-.24.122-.349.202l-7.535-4.35a1.8 1.8 0 00-1.069-1.852v-8.7c.124-.054.24-.122.35-.202z\"></path></svg> --></li><li class=\"m-2 p-4 relative rounded-lg border border-transparent [background:linear-gradient(theme(colors.zinc.50),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box]\"><svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" height=\"40\" width=\"40\" aria-label=\"Ethereum\"><path d=\"M11.944 17.97L4.58 13.62 11.943 24l7.37-10.38-7.372 4.35h.003zM12.056 0L4.69 12.223l7.365 4.354 7.365-4.35L12.056 0z\"></path></svg></li><li class=\"m-2 p-4 relative rounded-lg border border-transparent [background:linear-gradient(theme(colors.zinc.50),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box]\"><svg role=\"img\" viewBox=\"0 0 24 24\" width=\"40\" height=\"40\" aria-label=\"Bitcoin\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M23.638 14.904c-1.602 6.43-8.113 10.34-14.542 8.736C2.67 22.05-1.244 15.525.362 9.105 1.962 2.67 8.475-1.243 14.9.358c6.43 1.605 10.342 8.115 8.738 14.548v-.002zm-6.35-4.613c.24-1.59-.974-2.45-2.64-3.03l.54-2.153-1.315-.33-.525 2.107c-.345-.087-.705-.167-1.064-.25l.526-2.127-1.32-.33-.54 2.165c-.285-.067-.565-.132-.84-.2l-1.815-.45-.35 1.407s.975.225.955.236c.535.136.63.486.615.766l-1.477 5.92c-.075.166-.24.406-.614.314.015.02-.96-.24-.96-.24l-.66 1.51 1.71.426.93.242-.54 2.19 1.32.327.54-2.17c.36.1.705.19 1.05.273l-.51 2.154 1.32.33.545-2.19c2.24.427 3.93.257 4.64-1.774.57-1.637-.03-2.58-1.217-3.196.854-.193 1.5-.76 1.68-1.93h.01zm-3.01 4.22c-.404 1.64-3.157.75-4.05.53l.72-2.9c.896.23 3.757.67 3.33 2.37zm.41-4.24c-.37 1.49-2.662.735-3.405.55l.654-2.64c.744.18 3.137.524 2.75 2.084v.006z\"></path></svg></li><li class=\"m-2 p-4 relative rounded-lg border border-transparent [background:linear-gradient(theme(colors.zinc.50),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box]\"><svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" width=\"40\" height=\"40\" aria-label=\"Solana\"><path d=\"m23.8764 18.0313-3.962 4.1393a.9201.9201 0 0 1-.306.2106.9407.9407 0 0 1-.367.0742H.4599a.4689.4689 0 0 1-.2522-.0733.4513.4513 0 0 1-.1696-.1962.4375.4375 0 0 1-.0314-.2545.4438.4438 0 0 1 .117-.2298l3.9649-4.1393a.92.92 0 0 1 .3052-.2102.9407.9407 0 0 1 .3658-.0746H23.54a.4692.4692 0 0 1 .2523.0734.4531.4531 0 0 1 .1697.196.438.438 0 0 1 .0313.2547.4442.4442 0 0 1-.1169.2297zm-3.962-8.3355a.9202.9202 0 0 0-.306-.2106.941.941 0 0 0-.367-.0742H.4599a.4687.4687 0 0 0-.2522.0734.4513.4513 0 0 0-.1696.1961.4376.4376 0 0 0-.0314.2546.444.444 0 0 0 .117.2297l3.9649 4.1394a.9204.9204 0 0 0 .3052.2102c.1154.049.24.0744.3658.0746H23.54a.469.469 0 0 0 .2523-.0734.453.453 0 0 0 .1697-.1961.4382.4382 0 0 0 .0313-.2546.4444.4444 0 0 0-.1169-.2297zM.46 6.7225h18.7815a.9411.9411 0 0 0 .367-.0742.9202.9202 0 0 0 .306-.2106l3.962-4.1394a.4442.4442 0 0 0 .117-.2297.4378.4378 0 0 0-.0314-.2546.453.453 0 0 0-.1697-.196.469.469 0 0 0-.2523-.0734H4.7596a.941.941 0 0 0-.3658.0745.9203.9203 0 0 0-.3052.2102L.1246 5.9687a.4438.4438 0 0 0-.1169.2295.4375.4375 0 0 0 .0312.2544.4512.4512 0 0 0 .1692.196.4689.4689 0 0 0 .2518.0739z\"></path></svg><!-- <svg width=\"40\" height=\"40\" viewBox=\"0 0 164 164\" xmlns=\"http://www.w3.org/2000/svg\" aria-label=\"Sonr\"> --><!-- \t<path d=\"M71.8077 133.231C74.5054 135.928 78.1636 137.443 81.978 137.443C85.7924 137.443 89.4506 135.928 92.1483 133.231L133.219 92.1638C135.909 89.4654 137.42 85.8102 137.42 81.9998C137.42 78.1895 135.909 74.5345 133.219 71.8361L112.886 51.5272L131.665 32.7499L152.031 53.1143C159.696 60.7963 164 71.2046 164 82.0559C164 92.9072 159.696 103.315 152.031 110.997L110.95 152.065C107.154 155.869 102.642 158.883 97.6739 160.931C92.7059 162.98 87.3809 164.023 82.0071 164L82.0052 164C76.622 164.019 71.2886 162.969 66.3145 160.91C61.3405 158.852 56.8247 155.826 53.0294 152.009L53.0289 152.008L48.7187 147.699L67.4974 128.921L71.8077 133.231Z\"></path> --><!-- \t<path d=\"M110.95 11.9912L115.26 16.3011L96.481 35.0785L92.1707 30.7685C89.4731 28.072 85.8148 26.5572 82.0004 26.5572C78.186 26.5572 74.5277 28.072 71.8301 30.7685L30.7597 71.8359C29.4247 73.1706 28.3658 74.7552 27.6433 76.4991C26.9208 78.2431 26.549 80.1122 26.549 81.9999C26.549 83.8876 26.9208 85.7567 27.6433 87.5007C28.3658 89.2446 29.4247 90.8292 30.7597 92.1639L51.1256 112.528L32.3138 131.306L11.9923 110.941C8.19043 107.141 5.17433 102.629 3.1167 97.6635C1.05907 92.6976 0 87.3751 0 81.9999C0 76.6247 1.05907 71.3022 3.1167 66.3363C5.17433 61.3705 8.19021 56.8587 11.9921 53.0586L53.0625 11.9912C56.8629 8.18964 61.3751 5.17395 66.3413 3.11647C71.3075 1.05899 76.6304 0 82.006 0C87.3816 0 92.7045 1.05899 97.6707 3.11647C102.637 5.17395 107.149 8.18964 110.95 11.9912Z\"></path> --><!-- \t<path d=\"M55.603 76.6744L76.6993 55.5798C79.6327 52.6465 84.3888 52.6465 87.3223 55.5797L108.419 76.6744C111.352 79.6077 111.352 84.3634 108.419 87.2966L87.3223 108.391C84.3888 111.325 79.6327 111.325 76.6993 108.391L55.603 87.2966C52.6696 84.3634 52.6696 79.6077 55.603 76.6744Z\"></path> --><!-- </svg> --></li><li class=\"m-2 p-4 relative rounded-lg border border-transparent [background:linear-gradient(theme(colors.zinc.50),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box]\"><svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" width=\"40\" height=\"40\" aria-label=\"Meta\"><path d=\"M6.915 4.03c-1.968 0-3.683 1.28-4.871 3.113C.704 9.208 0 11.883 0 14.449c0 .706.07 1.369.21 1.973a6.624 6.624 0 0 0 .265.86 5.297 5.297 0 0 0 .371.761c.696 1.159 1.818 1.927 3.593 1.927 1.497 0 2.633-.671 3.965-2.444.76-1.012 1.144-1.626 2.663-4.32l.756-1.339.186-.325c.061.1.121.196.183.3l2.152 3.595c.724 1.21 1.665 2.556 2.47 3.314 1.046.987 1.992 1.22 3.06 1.22 1.075 0 1.876-.355 2.455-.843a3.743 3.743 0 0 0 .81-.973c.542-.939.861-2.127.861-3.745 0-2.72-.681-5.357-2.084-7.45-1.282-1.912-2.957-2.93-4.716-2.93-1.047 0-2.088.467-3.053 1.308-.652.57-1.257 1.29-1.82 2.05-.69-.875-1.335-1.547-1.958-2.056-1.182-.966-2.315-1.303-3.454-1.303zm10.16 2.053c1.147 0 2.188.758 2.992 1.999 1.132 1.748 1.647 4.195 1.647 6.4 0 1.548-.368 2.9-1.839 2.9-.58 0-1.027-.23-1.664-1.004-.496-.601-1.343-1.878-2.832-4.358l-.617-1.028a44.908 44.908 0 0 0-1.255-1.98c.07-.109.141-.224.211-.327 1.12-1.667 2.118-2.602 3.358-2.602zm-10.201.553c1.265 0 2.058.791 2.675 1.446.307.327.737.871 1.234 1.579l-1.02 1.566c-.757 1.163-1.882 3.017-2.837 4.338-1.191 1.649-1.81 1.817-2.486 1.817-.524 0-1.038-.237-1.383-.794-.263-.426-.464-1.13-.464-2.046 0-2.221.63-4.535 1.66-6.088.454-.687.964-1.226 1.533-1.533a2.264 2.264 0 0 1 1.088-.285z\"></path></svg></li><li class=\"m-2 p-4 relative rounded-lg border border-transparent [background:linear-gradient(theme(colors.zinc.50),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box]\"><svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" width=\"40\" height=\"40\" aria-label=\"Apple\"><path d=\"M12.152 6.896c-.948 0-2.415-1.078-3.96-1.04-2.04.027-3.91 1.183-4.961 3.014-2.117 3.675-.546 9.103 1.519 12.09 1.013 1.454 2.208 3.09 3.792 3.039 1.52-.065 2.09-.987 3.935-.987 1.831 0 2.35.987 3.96.948 1.637-.026 2.676-1.48 3.676-2.948 1.156-1.688 1.636-3.325 1.662-3.415-.039-.013-3.182-1.221-3.22-4.857-.026-3.04 2.48-4.494 2.597-4.559-1.429-2.09-3.623-2.324-4.39-2.376-2-.156-3.675 1.09-4.61 1.09zM15.53 3.83c.843-1.012 1.4-2.427 1.245-3.83-1.207.052-2.662.805-3.532 1.818-.78.896-1.454 2.338-1.273 3.714 1.338.104 2.715-.688 3.559-1.701\"></path></svg></li></ul></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate