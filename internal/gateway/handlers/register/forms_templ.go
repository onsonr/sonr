// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package register

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/onsonr/sonr/internal/gateway/handlers/register/ui"
	"github.com/onsonr/sonr/pkg/common/styles/layout"
)

func formCreateProfile(action string, method string, data CreateProfileData) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form action=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.SafeURL(action)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" method=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(method)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/gateway/handlers/register/forms.templ`, Line: 9, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><sl-card class=\"card-form gap-6 w-full max-w-xl\"><div slot=\"header\"><div class=\"w-full py-1\"><sl-progress-bar value=\"50\"></sl-progress-bar></div></div><div class=\"grid grid-cols-1 lg:grid-cols-2 gap-4\"><sl-input name=\"first_name\" placeholder=\"Satoshi\" type=\"text\" label=\"First Name\" required autofocus></sl-input> <sl-input name=\"last_name\" placeholder=\"N\" maxlength=\"1\" type=\"text\" label=\"Last Initial\"></sl-input></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = layout.Spacer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<sl-input name=\"handle\" placeholder=\"digitalgold\" type=\"text\" label=\"Handle\" minlength=\"4\" maxlength=\"12\" required><div slot=\"prefix\"><sl-icon name=\"at-sign\" library=\"sonr\"></sl-icon></div></sl-input>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = layout.Spacer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<sl-range name=\"is_human\" label=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(data.IsHumanLabel())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/gateway/handlers/register/forms.templ`, Line: 27, Col: 56}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" help-text=\"Prove you are a human.\" min=\"0\" max=\"9\" step=\"1\"></sl-range><div slot=\"footer\"><sl-button href=\"/\" outline><sl-icon slot=\"prefix\" name=\"arrow-left\" library=\"sonr\"></sl-icon> Cancel</sl-button> <sl-button type=\"submit\">Next <sl-icon slot=\"suffix\" name=\"arrow-right\" library=\"sonr\"></sl-icon></sl-button></div><style>\n\t\t\t\t.card-form {\n\t\t\t\t\tmargin: 1rem;\n\t\t\t\t}\n\t\t\t\t.card-form [slot='footer'] {\n\t\t\t\t\tdisplay: flex;\n\t\t\t\t\tjustify-content: space-between;\n\t\t\t\t\talign-items: center;\n\t\t\t\t\tflex-wrap: wrap;\n\t\t\t\t\tgap: 1rem;\n\t\t\t\t}\n\t\t\t\t@media (max-width: 640px) {\n\t\t\t\t\t.card-form {\n\t\t\t\t\t\tmargin: 0.5rem;\n\t\t\t\t\t}\n\t\t\t\t\t.card-form [slot='footer'] {\n\t\t\t\t\t\tflex-direction: column-reverse;\n\t\t\t\t\t\twidth: 100%;\n\t\t\t\t\t}\n\t\t\t\t\t.card-form [slot='footer'] sl-button {\n\t\t\t\t\t\twidth: 100%;\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t</style></sl-card></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func formRegisterPasskey(action, method string, data RegisterPasskeyData) templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form action=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 templ.SafeURL = templ.SafeURL(action)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" method=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(method)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/gateway/handlers/register/forms.templ`, Line: 67, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" id=\"passkey-form\"><input type=\"hidden\" name=\"credential\" id=\"credential-data\" required> <sl-card class=\"card-form gap-4 max-w-xl\"><div slot=\"header\"><div class=\"w-full py-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ui.ProfileCard(data.Address, data.Name, data.Handle, data.CreationBlock).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ui.SelectCoins().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div slot=\"footer\" class=\"space-y-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ui.InputPasskey(data.Address, data.Handle, data.Challenge).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<sl-button href=\"/\" style=\"width: 100%;\" outline><sl-icon slot=\"prefix\" name=\"x-lg\"></sl-icon> Cancel</sl-button></div><style>\n  \t\t.card-form [slot='footer'] {\n    \t\tjustify-content: space-evenly;\n    \t\talign-items: center;\n  \t\t}\n\t\t</style></sl-card></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate