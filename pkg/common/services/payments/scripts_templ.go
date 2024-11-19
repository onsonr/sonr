// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793

package payments

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

var paymentsHandle = templ.NewOnceHandle()

// Payment types
type PaymentMethodData struct {
	SupportedMethods string         `json:"supportedMethods"`
	Data             map[string]any `json:"data,omitempty"`
}

type PaymentItem struct {
	Label  string `json:"label"`
	Amount Money  `json:"amount"`
}

type Money struct {
	Currency string `json:"currency"`
	Value    string `json:"value"` // Decimal as string for precision
}

type PaymentOptions struct {
	RequestPayerName  bool `json:"requestPayerName,omitempty"`
	RequestPayerEmail bool `json:"requestPayerEmail,omitempty"`
	RequestPayerPhone bool `json:"requestPayerPhone,omitempty"`
	RequestShipping   bool `json:"requestShipping,omitempty"`
}

type PaymentDetails struct {
	Total           PaymentItem      `json:"total"`
	DisplayItems    []PaymentItem    `json:"displayItems,omitempty"`
	ShippingOptions []ShippingOption `json:"shippingOptions,omitempty"`
}

type ShippingOption struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Amount   Money  `json:"amount"`
	Selected bool   `json:"selected,omitempty"`
}

type PaymentRequest struct {
	MethodData []PaymentMethodData `json:"methodData"`
	Details    PaymentDetails      `json:"details"`
	Options    PaymentOptions      `json:"options,omitempty"`
}

// Base payments script template
func PaymentsScripts() templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script type=\"text/javascript\">\n            // Check if Payment Request API is supported\n            function isPaymentRequestSupported() {\n                return window.PaymentRequest !== undefined;\n            }\n\n            // Create and show payment request\n            async function showPaymentRequest(request) {\n                try {\n                    const paymentMethods = request.methodData;\n                    const details = request.details;\n                    const options = request.options || {};\n\n                    const paymentRequest = new PaymentRequest(\n                        paymentMethods,\n                        details,\n                        options\n                    );\n\n                    // Handle shipping address changes if shipping is requested\n                    if (options.requestShipping) {\n                        paymentRequest.addEventListener('shippingaddresschange', event => {\n                            event.updateWith(Promise.resolve(details));\n                        });\n                    }\n\n                    // Handle shipping option changes\n                    if (details.shippingOptions && details.shippingOptions.length > 0) {\n                        paymentRequest.addEventListener('shippingoptionchange', event => {\n                            event.updateWith(Promise.resolve(details));\n                        });\n                    }\n\n                    const response = await paymentRequest.show();\n                    \n                    // Create response object\n                    const result = {\n                        methodName: response.methodName,\n                        details: response.details,\n                    };\n\n                    if (options.requestPayerName) {\n                        result.payerName = response.payerName;\n                    }\n                    if (options.requestPayerEmail) {\n                        result.payerEmail = response.payerEmail;\n                    }\n                    if (options.requestPayerPhone) {\n                        result.payerPhone = response.payerPhone;\n                    }\n                    if (options.requestShipping) {\n                        result.shippingAddress = response.shippingAddress;\n                        result.shippingOption = response.shippingOption;\n                    }\n\n                    // Complete the payment\n                    await response.complete('success');\n\n                    // Dispatch success event\n                    window.dispatchEvent(new CustomEvent('paymentComplete', {\n                        detail: result\n                    }));\n\n                } catch (err) {\n                    // Dispatch error event\n                    window.dispatchEvent(new CustomEvent('paymentError', {\n                        detail: err.message\n                    }));\n                }\n            }\n\n            // Abort payment request\n            function abortPaymentRequest() {\n                if (window.currentPaymentRequest) {\n                    window.currentPaymentRequest.abort();\n                }\n            }\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = paymentsHandle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// Template for initiating payment request
func StartPayment(request PaymentRequest) templ.Component {
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
		templ_7745c5c3_Err = PaymentsScripts().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n        (async () => {\n            try {\n                if (!isPaymentRequestSupported()) {\n                    throw new Error(\"Payment Request API is not supported in this browser\");\n                }\n                const request = { templ.JSONString(request) };\n                await showPaymentRequest(request);\n            } catch (err) {\n                window.dispatchEvent(new CustomEvent('paymentError', {\n                    detail: err.message\n                }));\n            }\n        })();\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
