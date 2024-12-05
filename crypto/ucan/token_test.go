package ucan_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan"
)

const (
	keyOneBase64Std = "CAASqAkwggSkAgEAAoIBAQChp1HiZxTsLQCaHmW3/cc2ZDZpgLwn5o1/nZPgqT7SyXHP5bn7GQMG3kPEQWcl4nhtLX9hkrBEskHrdIlqp9zXFMwBfat+qfzCylGC/QBDF7wT9umLd7nbq7pAxQXteXgntt2Zhg4gE/kEk7vIyL+P9KpWJZ/yjpykgsDC7NPnrr8qZBo2tL0F4w+33nZhEx7Pp7Rnaq22JM8rF+NHCgSkUh63lp7Vhwm9PQoGtt0XTnEKxrMQnUme/IhGNxs84RphxHc5+nW6jYjgm5bcJonGyPU7bq+v51Mr2Ol4RT3L9ZNJgz0SWTSmAtiBLx2ryLrTjmDPSvN7wLm9sWEdWmRVAgMBAAECggEBAJMumrl+jWgz2TZ5sreBEp6NQ5VvpuDVY8PrnzaQIikdTMizK1BaB417VUwdGGM//dG5+R7HxkHl42sT4gH/8GzL/Krm1vwunXplZy3SWSi9NXsf9qgLTGebxasvOCRt0l6mesFLcxT12ma2c+VuEixp4aUqAKWB/1Ex03wm0RFBcSttPHe5ODW8Eaz+ZU8cpObEcZdCIPVxeWqLVdkAImOmsknL0EAxP8Wo/V6Rh5Cg4PnwnfJiQ45C+m6h7NTIw0H4UOncv7EBABra6LqF6Uoda9vmv8CpwaXwR557DPchQglFjtm48jWGeVKO3Zyutizu420eRrFZ0GmJo5flvkkCgYEA0SLysOZNxDgjYA0ihVYL6UbCvYUSADuDyTWREOUiRfmxAmS1xN9o7fieCJnA4aAAnSugtT2BI7HEqT1lLz0YF8NRDKL07TNbkmNLIHXBbXA5saf10N2juhflfIm5/b/W9lC3QsngMR27J25Ztqof6Ur36bIKJ6Y6XvYdlkkZkc8CgYEAxeCHUWMvtHtBID9ZOtrZRNhNJ/uz+2rzVSPd6ZdhEUWsvv/0p7JXmSAp2eoJDDKHeSnVxcxQMqhq0/edUSSzSvDpWha8UU4N8hRpu+M0XZNke0ijhpK6NIqNHPvZdsyFD0VR1Vaj2Ruy+pzih6PhqSnn2ZwvpQJAwBnqc2VCJJsCgYAkQr33hAbpxZ4EkmJw4elwye8L8x2a4rbH1TzQxBm8Lj3Nn26Qsve7gwbLkPULabWRirXzlrVkXfcuLNH1bc9Wl2vfGAYFdokjCYpGF4SxF+s47VlGnJc9tdT5UdvorjF0RaxwrRXtDi2b+Zsee8LKrU/sugzesQif3GZm30fKqwKBgQCQHwHP+HMFfAQqLZma8UzwBK7loUEsrHAAoff+K8CKKPoxvxD9lzqQD8oLqpbeaGsdh6fowe/jhaERM7dEI3vm6GK9t/N/MF+d4tpD+67nPPQhiv13haTTodo3swNnsHx1a+K3hLwf5DnOqLehXW59nET+zPAyudpZUEbft2+eYwKBgCMS6SitXwa2UjFNgkMAaOeJjkjnUKcr1tO/zPtaYPugKgkMQB890q4dcq5rnG2onhJ7hkoMwcrFugbD2nub9AIkaMc6Y46jyh2mSeA0337MpoMp99Jmp2/B1rouYo4IRS25b7jk22yjV8ARCzsxFVQxEwA1Lg8YpaXaifuI+/2O"
	keyTwoBase64Std = "CAASqAkwggSkAgEAAoIBAQDE6jTAXSR6TbBb2rfkosYR0XIrmR9sH/0HJI75xq55DIJGBVcl0ki+9nKLUOCi/pC487BP4ZzSsdTctThrINbIYEu3xe1lggFwNvzFlSag8sc/F97nooGbXpXIBUngVIVqdUT8idNAPZYiO1fEnxcb51hGP9K7h1cnFAfwzsKqJhs/BF6EioS7l0uXlKU6BPoRLVcjGOYtSJAgHOkhlaa+ISeDHK5Iy8ggoasZ0lezHKW9A3PZ5vrIA+PwiJSRihTMm966kzyizDY3/hI5tdEfJYJMsdqyklyaSsq55RY62otsL6wLKjGctyPV30ZLUSdI2kZuyO2w/ok95GOMJqWPAgMBAAECggEBAKwX89pajNLGqubcE/MhvvE7lwg7XpbkrgJcFQh+d2UbZY9Eg5FuYl1ijWDsYiaRTHIXp3NoveH1wQ7S4mfd31hnsEUAGiWopREpPWiAna3z/+ZIOms+Pv9Gfqi81n/T3nXX317GJXXzXQ61xlL0pwGgAioDBW0XLzfb7cSrLr373N5BQp6j/Et8C6oWnT48LOBr7TIE3unjVC/g64LjRJ24Ry5XZAJno8kUjvX/qg6LB2sqLLbA6R8FeCwszRCXTBzIdTbadX4FkPKpI7NrfCfOLG/L+Zf2LJEfAw/vBi+F7jBq4Rt9qvyjJj7/1oJrJ6QGESmFdbPGX+HgbwF0a0ECgYEA+t7IJKQf8deZceMDJ/jucy1ZSJwW9iv9CYMt6rxCBFNaOpEk1E+zP6Kz5wsqPXUSf5dTpDBH/vm0D255rmwVz0Tpx7xRKokvAYcjTexZehxzhqPdZymcPRZIF9V7Iyln9qUAJuLUf2MwnMbMl69/Poz/dq+iK3+HpnLnxg1RBRcCgYEAyPD8l9h8C7Jp1uOPb65IVwbm7dZJ7e/My3vjsoVewcoHZA6P4N21MPllI61VRZr3Z1uMDhbwmUvp2wtOTtXuAImERs2PFszG9NQQyNvRqNMzMFH+MOcXiQ3/Ws9zfK8reKmufD5ZNQe+HcgWLfIuy7iOt0p7xggTYGewZVlYnkkCgYBPWkzAmlGoc+QLjB0hdbInKH8HYqg4se1WJvJNP8M6DwuJXwPhTFyMknCJcpSn3/I7/aftVYBQfLeh8fX3YCT97PRtw3mBFOeTeiWGrm4XHAzG1+peiWDsSbIAJ/zNQHmsIMENi85fhQaJcLCiglajeIIODrwjOjG0SsBZezjXfQKBgHWl+w5glsg2bpd9ZsbJsNsbVGveMizYYPymjbtBMSifQ9KGYCEVTffdnSTVYH6/a6kdRZQeREJM2x//r5qi0JWJ7mOSCPwda0N/QlCHu2pwNaFN8FjrhLEe++pMWd6fpQEhv+JIkuxkmyBOvQWrrVBjv1N7jZp1sfqY2wOL20HZAoGBANMrTbs2YZW+Jy3WmGEm/MMI5VKr1ajbyJoFAEOVEggSAchI3B/E9HhTN7vOC44WMHgOVgCdZfqQoRjfJIOqvcVzPsFXUV8hi0kCsk6+RWMuiIwW1LDGi/LXJT6gkkfzPs4gAyCJ3tEmPWkOlgtORhb1zLRJsuudeHst5Of12A/b"
)

var (
	keyOne crypto.PrivKey
	keyTwo crypto.PrivKey
)

func init() {
	var err error
	keyOneBytes, err := base64.StdEncoding.DecodeString(keyOneBase64Std)
	if err != nil {
		panic(err)
	}
	if keyOne, err = crypto.UnmarshalPrivateKey(keyOneBytes); err != nil {
		panic(err)
	}
	keyTwoBytes, err := base64.StdEncoding.DecodeString(keyOneBase64Std)
	if err != nil {
		panic(err)
	}
	if keyTwo, err = crypto.UnmarshalPrivateKey(keyTwoBytes); err != nil {
		panic(err)
	}
}

func TestPrivKeySource(t *testing.T) {
	source, err := ucan.NewPrivKeySource(keyOne)
	if err != nil {
		t.Fatal(err)
	}

	didStr, err := ucan.DIDStringFromPublicKey(keyOne.GetPublic())
	if err != nil {
		t.Fatal(err)
	}

	caps := ucan.NewNestedCapabilities("SUPER_USER", "OVERWRITE", "SOFT_DELETE", "REVISE", "CREATE")
	att := ucan.Attenuations{
		{caps.Cap("SUPER_USER"), ucan.NewStringLengthResource("api", "*")},
		{caps.Cap("SUPER_USER"), ucan.NewStringLengthResource("dataset", "b5:world_bank_population:*")},
	}
	zero := time.Time{}

	root, err := source.NewOriginToken(didStr, att, nil, zero, zero)
	if err != nil {
		t.Fatal(err)
	}

	expect := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsInVjdiI6IjAuNy4wIn0.eyJhdWQiOiJkaWQ6a2V5OnoyTUd3NGdrODRVU290YVdmNEFrSjgzRGNucmZnR2FjZUY4NktRWFJZTWZRN3hxblVGRVREN3RGNGprTkd2QUV0dlZ6Rm5RdTkydEZ3ZkNyb2ZWZ0c2WnkyZVlKZURSckVVZ1c5WXZCazJXSHVUUnJFZ0FTZnBxR2tKUkdnR3FiNmNidjFxR2N2Nm1FSFF5YlR5M2JOeXFxd3VRWXA3c3BuSlFObVJhUGk0Z3F1czVEWnVOc2NRZjFSMXhCdVN6WHk1YnNCaFlSZzdFcDZmNkJad0U2cHZCamt5ellqWERVYlF2a0VpeG0zcHR3RTRnZ2VZU21oUXFxbU1ac1lwcG1lY1VFMjhuTTdFekx2Q3hRZEZ1QndXZ2U3QURVYzdxVGYxeXNpUzl1YXdOTnA1aER2aHl2cXRDaWg3a3FvTHVzTGVnd2pHZTJTcDhDcUZmdUNRNWgxdHh4WHozdEdtRGZEUDE3Nm15R1htc3R0eDV5MjVTOXpwejg1ZEc4WnRrRnZ4bXNjOFltaXZlRUMycWFkY2FrWEoiLCJpc3MiOiJkaWQ6a2V5OnoyTUd3NGdrODRVU290YVdmNEFrSjgzRGNucmZnR2FjZUY4NktRWFJZTWZRN3hxblVGRVREN3RGNGprTkd2QUV0dlZ6Rm5RdTkydEZ3ZkNyb2ZWZ0c2WnkyZVlKZURSckVVZ1c5WXZCazJXSHVUUnJFZ0FTZnBxR2tKUkdnR3FiNmNidjFxR2N2Nm1FSFF5YlR5M2JOeXFxd3VRWXA3c3BuSlFObVJhUGk0Z3F1czVEWnVOc2NRZjFSMXhCdVN6WHk1YnNCaFlSZzdFcDZmNkJad0U2cHZCamt5ellqWERVYlF2a0VpeG0zcHR3RTRnZ2VZU21oUXFxbU1ac1lwcG1lY1VFMjhuTTdFekx2Q3hRZEZ1QndXZ2U3QURVYzdxVGYxeXNpUzl1YXdOTnA1aER2aHl2cXRDaWg3a3FvTHVzTGVnd2pHZTJTcDhDcUZmdUNRNWgxdHh4WHozdEdtRGZEUDE3Nm15R1htc3R0eDV5MjVTOXpwejg1ZEc4WnRrRnZ4bXNjOFltaXZlRUMycWFkY2FrWEoiLCJhdHQiOlt7ImFwaSI6IioiLCJjYXAiOiJTVVBFUl9VU0VSIn0seyJjYXAiOiJTVVBFUl9VU0VSIiwiZGF0YXNldCI6ImI1OndvcmxkX2JhbmtfcG9wdWxhdGlvbjoqIn1dfQ.ggJJ8fovwKVbq_hLb0sVKlvkFRcKJWb-55RgIagjzO9D9I0pZ5g3elkOCkf1XA60_32fE-K_Dj97np6-BQx_F5u_PofYi6C0BYfR4YATA5jzymT7nBLrdOzhr0Vh2FChYoyM82CT4tkOGSlnvaoCgB4OLMOzlvpWmbfRi5GB3d6yAfdhBGsDuAJ9L7V_N5WBFpjarSu2xAYhVUhFM8IbILHMipXWRd2sP-_DzxmsneS7qu9fxAk35DDC_8cONPp1ZrFUw-VhGts0jYOVuExl_ZWqE3y2XJLtiperxFWO45q59WdULrn15ciQjRhn2O-QmXHmOxLNZMCI6lbFFiCHWg`
	if expect != root.Raw {
		t.Errorf("token mismatch. expected: %q.\ngot: %q", expect, root.Raw)
	}

	att = ucan.Attenuations{
		{caps.Cap("OVERWRITE"), ucan.NewStringLengthResource("dataset", "b5:world_bank_population:*")},
	}

	derivedToken, err := source.NewAttenuatedToken(root, didStr, att, nil, zero, zero)
	if err != nil {
		t.Fatal(err)
	}

	cidStr := mustCidString(t, derivedToken)
	expectCID := "bafkreifhpoxctmbmvocdevfbmio6cpzltwauesyyjycipnylocoykwghzu"

	if expectCID != cidStr {
		t.Errorf("derived token CID mismatch. expected: %q.\ngot: %q", expectCID, cidStr)
	}

	// tokenWithExpiryString, err := tokens.CreateToken(pro, time.Hour)
	// expect = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOi02MjEzNTU5MzIwMCwic3ViIjoiUW1lTDJtZFZrYTFlYWhLRU5qZWhLNnRCeGtrcGs1ZE5RMXFNY2dXaTdIcmI0QiIsInVzZXJuYW1lIjoiZG91ZyJ9.d7XPhsj7hkyxg1JzC59hfu90RYem5q6Pie-ofJhdlGk_sY5bH8gcqG90LndMh4_LglEvtrwf_SVFcM1b78qhNon_Yo91kG_K_MmyExa-AlpY65Ji_kpRWcnI8hl-mxrZ2MzxPjvAEOa6c80DUWgTFKlkrgf9RnZlqq-nHnxHHXbVKYI3girsDgWynaIhR53yMBDIhbTCZaQ8XKtU_Pr0L1dJAW7YvOo2H01VM4LI_UQqhCmEbTnQX1Zee0tg88IMzLl7WsdNNOzUsf7dCYWGerLtzxGbxR0wweXbqVJBlzIl0Upke8-FBuZIbcdGSniy4DX643KrNnp_FnzQ8oBHTA`
	// if expect != tokenWithExpiryString {
	// 	t.Errorf("token mismatch. expected: %q.\ngot: %q", expect, tokenWithExpiryString)
	// }
}

func TestED25519PrivKeySource(t *testing.T) {
	keyOne, _, err := crypto.GenerateKeyPair(crypto.Ed25519, 123)
	if err != nil {
		t.Fatal(err)
	}

	source, err := ucan.NewPrivKeySource(keyOne)
	if err != nil {
		t.Fatal(err)
	}

	audienceDID, err := ucan.DIDStringFromPublicKey(keyOne.GetPublic())
	if err != nil {
		panic(err)
	}

	zero := time.Time{}

	// create a root UCAN
	origin, err := source.NewOriginToken(audienceDID, nil, nil, zero, zero)
	if err != nil {
		panic(err)
	}

	if _, err = origin.CID(); err != nil {
		t.Fatal(err)
	}
}

func TestTokenSource(t *testing.T) {
	// ucan_spec.AssertTokenSourceSpec(t, func(ctx context.Context) ucan.TokenSource {
	// 	source, err := ucan.NewPrivKeyTokenSource(peerInfo.PrivKey)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return source
	// })
}

func TestTokenParse(t *testing.T) {
	raw := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsInVjdiI6IjAuNy4wIn0.eyJhdWQiOiJkaWQ6a2V5OnoyTUd3NGdrODRVU290YVdmNEFrSjgzRGNucmZnR2FjZUY4NktRWFJZTWZRN3hxblVGRVREN3RGNGprTkd2QUV0dlZ6Rm5RdTkydEZ3ZkNyb2ZWZ0c2WnkyZVlKZURSckVVZ1c5WXZCazJXSHVUUnJFZ0FTZnBxR2tKUkdnR3FiNmNidjFxR2N2Nm1FSFF5YlR5M2JOeXFxd3VRWXA3c3BuSlFObVJhUGk0Z3F1czVEWnVOc2NRZjFSMXhCdVN6WHk1YnNCaFlSZzdFcDZmNkJad0U2cHZCamt5ellqWERVYlF2a0VpeG0zcHR3RTRnZ2VZU21oUXFxbU1ac1lwcG1lY1VFMjhuTTdFekx2Q3hRZEZ1QndXZ2U3QURVYzdxVGYxeXNpUzl1YXdOTnA1aER2aHl2cXRDaWg3a3FvTHVzTGVnd2pHZTJTcDhDcUZmdUNRNWgxdHh4WHozdEdtRGZEUDE3Nm15R1htc3R0eDV5MjVTOXpwejg1ZEc4WnRrRnZ4bXNjOFltaXZlRUMycWFkY2FrWEoiLCJpc3MiOiJkaWQ6a2V5OnoyTUd3NGdrODRVU290YVdmNEFrSjgzRGNucmZnR2FjZUY4NktRWFJZTWZRN3hxblVGRVREN3RGNGprTkd2QUV0dlZ6Rm5RdTkydEZ3ZkNyb2ZWZ0c2WnkyZVlKZURSckVVZ1c5WXZCazJXSHVUUnJFZ0FTZnBxR2tKUkdnR3FiNmNidjFxR2N2Nm1FSFF5YlR5M2JOeXFxd3VRWXA3c3BuSlFObVJhUGk0Z3F1czVEWnVOc2NRZjFSMXhCdVN6WHk1YnNCaFlSZzdFcDZmNkJad0U2cHZCamt5ellqWERVYlF2a0VpeG0zcHR3RTRnZ2VZU21oUXFxbU1ac1lwcG1lY1VFMjhuTTdFekx2Q3hRZEZ1QndXZ2U3QURVYzdxVGYxeXNpUzl1YXdOTnA1aER2aHl2cXRDaWg3a3FvTHVzTGVnd2pHZTJTcDhDcUZmdUNRNWgxdHh4WHozdEdtRGZEUDE3Nm15R1htc3R0eDV5MjVTOXpwejg1ZEc4WnRrRnZ4bXNjOFltaXZlRUMycWFkY2FrWEoiLCJhdHQiOlt7ImFwaSI6IioiLCJjYXAiOiJTVVBFUl9VU0VSIn0seyJjYXAiOiJTVVBFUl9VU0VSIiwiZGF0YXNldCI6ImI1OndvcmxkX2JhbmtfcG9wdWxhdGlvbjoqIn1dfQ.ggJJ8fovwKVbq_hLb0sVKlvkFRcKJWb-55RgIagjzO9D9I0pZ5g3elkOCkf1XA60_32fE-K_Dj97np6-BQx_F5u_PofYi6C0BYfR4YATA5jzymT7nBLrdOzhr0Vh2FChYoyM82CT4tkOGSlnvaoCgB4OLMOzlvpWmbfRi5GB3d6yAfdhBGsDuAJ9L7V_N5WBFpjarSu2xAYhVUhFM8IbILHMipXWRd2sP-_DzxmsneS7qu9fxAk35DDC_8cONPp1ZrFUw-VhGts0jYOVuExl_ZWqE3y2XJLtiperxFWO45q59WdULrn15ciQjRhn2O-QmXHmOxLNZMCI6lbFFiCHWg`

	caps := ucan.NewNestedCapabilities("SUPER_USER", "OVERWRITE", "SOFT_DELETE", "REVISE", "CREATE")

	ac := func(m map[string]interface{}) (ucan.Attenuation, error) {
		var (
			cap string
			rsc ucan.Resource
		)
		for key, vali := range m {
			val, ok := vali.(string)
			if !ok {
				return ucan.Attenuation{}, fmt.Errorf(`expected attenuation value to be a string`)
			}

			if key == ucan.CapKey {
				cap = val
			} else {
				rsc = ucan.NewStringLengthResource(key, val)
			}
		}

		return ucan.Attenuation{
			Rsc: rsc,
			Cap: caps.Cap(cap),
		}, nil
	}

	store := ucan.NewMemTokenStore()
	p := ucan.NewTokenParser(ac, ucan.StringDIDPubKeyResolver{}, store.(ucan.CIDBytesResolver))
	_, err := p.ParseAndVerify(context.Background(), raw)
	if err != nil {
		t.Error(err)
	}
}

func mustCidString(t *testing.T, tok *ucan.Token) string {
	t.Helper()
	id, err := tok.CID()
	if err != nil {
		t.Fatal(err)
	}
	return id.String()
}
