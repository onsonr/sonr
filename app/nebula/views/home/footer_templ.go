// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Footer() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Site footer --><footer><div class=\"max-w-5xl mx-auto px-4 sm:px-6\"><!-- Top area: Blocks --><div class=\"grid sm:grid-cols-12 gap-8 py-8 md:py-12 border-t border-zinc-200\"><!-- 1st block --><div class=\"sm:col-span-6 md:col-span-3 lg:col-span-6 max-sm:order-1 flex flex-col\"><div class=\"mb-4\"><!-- Logo --><a class=\"flex items-center justify-center bg-white w-8 h-8 rounded shadow-sm shadow-zinc-950/20\" href=\"/\"><img src=\"https://cdn.sonr.id/logo-zinc.svg\" width=\"24\" height=\"24\" alt=\"Logo\"></a></div><div class=\"grow text-sm text-zinc-500\">&copy; diDAO DUNA. All rights reserved.</div><!-- Social links --><ul class=\"flex space-x-4 mt-4 mb-1\"><li><a class=\"flex justify-center items-center text-zinc-700 hover:text-zinc-900 transition\" href=\"#0\" aria-label=\"Twitter\"><svg class=\"fill-current\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\"><path d=\"m7.063 3 3.495 4.475L14.601 3h2.454l-5.359 5.931L18 17h-4.938l-3.866-4.893L4.771 17H2.316l5.735-6.342L2 3h5.063Zm-.74 1.347H4.866l8.875 11.232h1.36L6.323 4.347Z\"></path></svg></a></li><li><a class=\"flex justify-center items-center text-zinc-700 hover:text-zinc-900 transition\" href=\"#0\" aria-label=\"Medium\"><svg class=\"fill-current\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\"><path d=\"M17 2H3a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm-1.708 3.791-.858.823a.251.251 0 0 0-.1.241V12.9a.251.251 0 0 0 .1.241l.838.823v.181h-4.215v-.181l.868-.843c.085-.085.085-.11.085-.241V7.993L9.6 14.124h-.329l-2.81-6.13V12.1a.567.567 0 0 0 .156.472l1.129 1.37v.181h-3.2v-.181l1.129-1.37a.547.547 0 0 0 .146-.472V7.351A.416.416 0 0 0 5.683 7l-1-1.209V5.61H7.8l2.4 5.283 2.122-5.283h2.971l-.001.181Z\"></path></svg></a></li><li><a class=\"flex justify-center items-center text-zinc-700 hover:text-zinc-900 transition\" href=\"#0\" aria-label=\"Telegram\"><svg class=\"fill-current\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\"><path d=\"M17.968 3.276a.338.338 0 0 0-.232-.253 1.192 1.192 0 0 0-.63.045S3.087 8.106 2.286 8.664c-.172.121-.23.19-.259.272-.138.4.293.573.293.573l3.613 1.177a.388.388 0 0 0 .183-.011c.822-.519 8.27-5.222 8.7-5.38.068-.02.118 0 .1.049-.172.6-6.606 6.319-6.64 6.354a.138.138 0 0 0-.05.118l-.337 3.528s-.142 1.1.956 0a30.66 30.66 0 0 1 1.9-1.738c1.242.858 2.58 1.806 3.156 2.3a1 1 0 0 0 .732.283.825.825 0 0 0 .7-.622S17.894 5.292 17.98 3.909c.008-.135.021-.217.021-.317a1.177 1.177 0 0 0-.032-.316Z\"></path></svg></a></li></ul></div><!-- 2nd block --><div class=\"sm:col-span-6 md:col-span-3 lg:col-span-2\"></div><!-- 3rd block --><div class=\"sm:col-span-6 md:col-span-3 lg:col-span-2\"><h6 class=\"text-sm text-zinc-800 font-medium mb-2\">Resources</h6><ul class=\"text-sm space-y-2\"><li><a class=\"text-zinc-500 hover:text-zinc-900 transition\" href=\"#0\">Community</a></li><li><a class=\"text-zinc-500 hover:text-zinc-900 transition\" href=\"#0\">Documentation</a></li><li><a class=\"text-zinc-500 hover:text-zinc-900 transition\" href=\"#0\">Privacy policy</a></li></ul></div><!-- 4th block --><div class=\"sm:col-span-6 md:col-span-3 lg:col-span-2\"><h6 class=\"text-sm text-zinc-800 font-medium mb-2\">Legals</h6><ul class=\"text-sm space-y-2\"><li><a class=\"text-zinc-500 hover:text-zinc-900 transition\" href=\"#0\">About the DAO</a></li><li><a class=\"text-zinc-500 hover:text-zinc-900 transition\" href=\"#0\">Privacy policy</a></li><li><a class=\"text-zinc-500 hover:text-zinc-900 transition\" href=\"https://brandfetch.io/sonr.io\">Brand Kit</a></li></ul></div></div></div></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
