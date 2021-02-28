package host

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"strings"

	"strconv"

	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	readability "github.com/julianshen/go-readability"
	md "github.com/sonr-io/core/internal/models"
)

var (
	ErrorType = errors.New("Should not be non-ptr or nil")
)

type OgImage struct {
	Url       string `meta:"og:image,og:image:url"`
	SecureUrl string `meta:"og:image:secure_url"`
	Width     int    `meta:"og:image:width"`
	Height    int    `meta:"og:image:height"`
	Type      string `meta:"og:image:type"`
}

type OgVideo struct {
	Url       string `meta:"og:video,og:video:url"`
	SecureUrl string `meta:"og:video:secure_url"`
	Width     int    `meta:"og:video:width"`
	Height    int    `meta:"og:video:height"`
	Type      string `meta:"og:video:type"`
}

type OgAudio struct {
	Url       string `meta:"og:audio,og:audio:url"`
	SecureUrl string `meta:"og:audio:secure_url"`
	Type      string `meta:"og:audio:type"`
}

type TwitterCard struct {
	Card        string `meta:"twitter:card"`
	Site        string `meta:"twitter:site"`
	SiteId      string `meta:"twitter:site:id"`
	Creator     string `meta:"twitter:creator"`
	CreatorId   string `meta:"twitter:creator:id"`
	Description string `meta:"twitter:description"`
	Title       string `meta:"twitter:title"`
	Image       string `meta:"twitter:image,twitter:image:src"`
	ImageAlt    string `meta:"twitter:image:alt"`
	Url         string `meta:"twitter:url"`
	Player      struct {
		Url    string `meta:"twitter:player"`
		Width  int    `meta:"twitter:width"`
		Height int    `meta:"twitter:height"`
		Stream string `meta:"twitter:stream"`
	}
	IPhone struct {
		Name string `meta:"twitter:app:name:iphone"`
		Id   string `meta:"twitter:app:id:iphone"`
		Url  string `meta:"twitter:app:url:iphone"`
	}
	IPad struct {
		Name string `meta:"twitter:app:name:ipad"`
		Id   string `meta:"twitter:app:id:ipad"`
		Url  string `meta:"twitter:app:url:ipad"`
	}
	Googleplay struct {
		Name string `meta:"twitter:app:name:googleplay"`
		Id   string `meta:"twitter:app:id:googleplay"`
		Url  string `meta:"twitter:app:url:googleplay"`
	}
}

type PageInfo struct {
	Title       string `meta:"og:title"`
	Type        string `meta:"og:type"`
	Url         string `meta:"og:url"`
	Site        string `meta:"og:site"`
	SiteName    string `meta:"og:site_name"`
	Description string `meta:"og:description"`
	Locale      string `meta:"og:locale"`
	Images      []*OgImage
	Videos      []*OgVideo
	Audios      []*OgAudio
	Twitter     *TwitterCard
	Content     string
}

func GetPageDataFromHtml(html []byte, data interface{}) error {
	buf := bytes.NewBuffer(html)
	doc, err := goquery.NewDocumentFromReader(buf)

	if err != nil {
		return err
	}

	return GetPageData(doc, data)
}

func GetPageData(doc *goquery.Document, data interface{}) error {
	doc = goquery.CloneDocument(doc)
	return getPageData(doc, data)
}

func GetPageInfo(doc *goquery.Document) (*PageInfo, error) {
	info := PageInfo{}
	err := GetPageData(doc, &info)

	if err != nil {
		return nil, err
	}

	html, _ := doc.Html()
	r, err := readability.NewDocument(html)
	if err != nil {
		return nil, err
	}

	info.Content = r.Text()
	return &info, nil
}

func GetPageInfoFromResponse(response *http.Response) (*PageInfo, error) {
	info := PageInfo{}
	html, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = GetPageDataFromHtml(html, &info)

	if err != nil {
		return nil, err
	}

	r, err := readability.NewDocument(string(html))
	if err != nil {
		return nil, err
	}

	info.Content = r.Text()

	return &info, nil
}

func GetPageInfoFromUrl(urlStr string) (*md.URLLink, error) {
	// Create Request
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}

	// Get Info
	info, err := GetPageInfoFromResponse(resp)
	if err != nil {
		return nil, err
	}

	// Set Link
	link := &md.URLLink{
		Link:        urlStr,
		Title:       info.Title,
		Type:        info.Type,
		Url:         info.Url,
		Site:        info.Site,
		SiteName:    info.SiteName,
		Description: info.Description,
		Locale:      info.Locale,
	}

	// Get Images
	if info.Images != nil {
		for _, v := range info.Images {
			link.Images = append(link.Images, &md.URLLink_OpenGraphImage{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Width:     int32(v.Width),
				Height:    int32(v.Height),
				Type:      v.Type,
			})
		}
	}

	// Get Videos
	if info.Videos != nil {
		for _, v := range info.Videos {
			link.Videos = append(link.Videos, &md.URLLink_OpenGraphVideo{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Width:     int32(v.Width),
				Height:    int32(v.Height),
				Type:      v.Type,
			})
		}
	}

	// Get Audios
	if info.Audios != nil {
		for _, v := range info.Videos {
			link.Audios = append(link.Audios, &md.URLLink_OpenGraphAudio{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Type:      v.Type,
			})
		}
	}

	// Get Twitter
	if info.Twitter != nil {
		twitter := &md.URLLink_TwitterCard{
			Card:        info.Twitter.Card,
			Site:        info.Twitter.Site,
			SiteId:      info.Twitter.SiteId,
			Creator:     info.Twitter.Creator,
			CreatorId:   info.Twitter.CreatorId,
			Description: info.Twitter.Description,
			Title:       info.Twitter.Title,
			Image:       info.Twitter.Image,
			ImageAlt:    info.Twitter.ImageAlt,
			Url:         info.Twitter.Url,
			Player: &md.URLLink_TwitterCard_Player{
				Url:    info.Twitter.Player.Url,
				Width:  int32(info.Twitter.Player.Width),
				Height: int32(info.Twitter.Player.Height),
				Stream: info.Twitter.Player.Stream,
			},
			Iphone: &md.URLLink_TwitterCard_IPhone{
				Name: info.Twitter.IPhone.Name,
				Id:   info.Twitter.IPhone.Id,
				Url:  info.Twitter.IPhone.Url,
			},
			Ipad: &md.URLLink_TwitterCard_IPad{
				Name: info.Twitter.IPad.Name,
				Id:   info.Twitter.IPad.Id,
				Url:  info.Twitter.IPad.Url,
			},
			GooglePlay: &md.URLLink_TwitterCard_GooglePlay{
				Name: info.Twitter.Googleplay.Name,
				Id:   info.Twitter.Googleplay.Id,
				Url:  info.Twitter.Googleplay.Url,
			},
		}
		link.Twitter = twitter
	}

	// Return Link
	return link, nil
}

func getPageData(doc *goquery.Document, data interface{}) error {
	var rv reflect.Value
	var ok bool
	if rv, ok = data.(reflect.Value); !ok {
		rv = reflect.ValueOf(data)
	}

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return ErrorType
	}

	rt := rv.Type()

	for i := 0; i < rv.Elem().NumField(); i++ {
		fv := rv.Elem().Field(i)
		field := rt.Elem().Field(i)

		switch fv.Type().Kind() {
		case reflect.Ptr:
			if fv.IsNil() {
				fv.Set(reflect.New(fv.Type().Elem()))
			}
			e := getPageData(doc, fv)

			if e != nil {
				return e
			}
		case reflect.Struct:
			e := getPageData(doc, fv.Addr())

			if e != nil {
				return e
			}
		case reflect.Slice:
			if fv.IsNil() {
				fv.Set(reflect.MakeSlice(fv.Type(), 0, 0))
			}

			switch field.Type.Elem().Kind() {
			case reflect.Struct:
				last := reflect.New(field.Type.Elem())
				for {
					data := reflect.New(field.Type.Elem())
					e := getPageData(doc, data.Interface())

					if e != nil {
						return e
					}

					//Ugly solution (I can't remove nodes. Why?)
					if !reflect.DeepEqual(last.Elem().Interface(), data.Elem().Interface()) {
						fv.Set(reflect.Append(fv, data.Elem()))
						last.Elem().Set(data.Elem())

					} else {
						break
					}
				}
			case reflect.Ptr:
				last := reflect.New(field.Type.Elem().Elem())
				for {
					data := reflect.New(field.Type.Elem().Elem())
					e := getPageData(doc, data.Interface())

					if e != nil {
						return e
					}

					//Ugly solution (I can't remove nodes. Why?)
					if !reflect.DeepEqual(last.Elem().Interface(), data.Elem().Interface()) {
						fv.Set(reflect.Append(fv, data))
						last.Elem().Set(data.Elem())

					} else {
						break
					}
				}
			default:
				if tag, ok := field.Tag.Lookup("meta"); ok {
					tags := strings.Split(tag, ",")

					for _, t := range tags {
						contents := []reflect.Value{}

						processMeta := func(idx int, sel *goquery.Selection) {
							if c, existed := sel.Attr("content"); existed {
								if field.Type.Elem().Kind() == reflect.String {
									contents = append(contents, reflect.ValueOf(c))
								} else {
									i, e := strconv.Atoi(c)

									if e == nil {
										contents = append(contents, reflect.ValueOf(i))
									}
								}

								fv.Set(reflect.Append(fv, contents...))
							}
						}

						doc.Find(fmt.Sprintf("meta[property=\"%s\"]", t)).Each(processMeta)

						doc.Find(fmt.Sprintf("meta[name=\"%s\"]", t)).Each(processMeta)

						fv = reflect.Append(fv, contents...)
					}
				}
			}
		default:
			if tag, ok := field.Tag.Lookup("meta"); ok {

				tags := strings.Split(tag, ",")

				content := ""
				existed := false
				sel := (*goquery.Selection)(nil)
				for _, t := range tags {
					if sel = doc.Find(fmt.Sprintf("meta[property=\"%s\"]", t)).First(); sel.Size() > 0 {
						content, existed = sel.Attr("content")
					}

					if !existed {
						if sel = doc.Find(fmt.Sprintf("meta[name=\"%s\"]", t)).First(); sel.Size() > 0 {
							content, existed = sel.Attr("content")
						}
					}

					if existed {
						if fv.Type().Kind() == reflect.String {
							fv.Set(reflect.ValueOf(content))
						} else if fv.Type().Kind() == reflect.Int {
							if i, e := strconv.Atoi(content); e == nil {
								fv.Set(reflect.ValueOf(i))
							}
						}
						break
					}
				}
			}
		}
	}
	return nil
}
