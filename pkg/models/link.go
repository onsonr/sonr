package models

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"regexp"

	"strings"

	"strconv"

	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// ************************* //
// ** URL Data Management ** //
// ************************* //

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
	r, err := NewDocument(html)
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

	r, err := NewDocument(string(html))
	if err != nil {
		return nil, err
	}

	info.Content = r.Text()

	return &info, nil
}

func GetPageInfoFromUrl(urlStr string) (*URLLink, *SonrError) {
	// Create Request
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, NewError(err, ErrorMessage_URL_HTTP_GET)
	}

	// Get Info
	info, err := GetPageInfoFromResponse(resp)
	if err != nil {
		return nil, NewError(err, ErrorMessage_URL_INFO_RESP)
	}

	// Set Link
	link := &URLLink{
		Url:         urlStr,
		Title:       info.Title,
		Type:        info.Type,
		Site:        info.Site,
		SiteName:    info.SiteName,
		Description: info.Description,
		Locale:      info.Locale,
	}

	// Get Images
	if info.Images != nil {
		for _, v := range info.Images {
			link.Images = append(link.Images, &URLLink_OpenGraphImage{
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
			link.Videos = append(link.Videos, &URLLink_OpenGraphVideo{
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
			link.Audios = append(link.Audios, &URLLink_OpenGraphAudio{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Type:      v.Type,
			})
		}
	}

	// Get Twitter
	if info.Twitter != nil {
		twitter := &URLLink_TwitterCard{
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
			Player: &URLLink_TwitterCard_Player{
				Url:    info.Twitter.Player.Url,
				Width:  int32(info.Twitter.Player.Width),
				Height: int32(info.Twitter.Player.Height),
				Stream: info.Twitter.Player.Stream,
			},
			Iphone: &URLLink_TwitterCard_IPhone{
				Name: info.Twitter.IPhone.Name,
				Id:   info.Twitter.IPhone.Id,
				Url:  info.Twitter.IPhone.Url,
			},
			Ipad: &URLLink_TwitterCard_IPad{
				Name: info.Twitter.IPad.Name,
				Id:   info.Twitter.IPad.Id,
				Url:  info.Twitter.IPad.Url,
			},
			GooglePlay: &URLLink_TwitterCard_GooglePlay{
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

var (
	Logger = log.New(ioutil.Discard, "[readability] ", log.LstdFlags)

	replaceBrsRegexp   = regexp.MustCompile(`(?i)(<br[^>]*>[ \n\r\t]*){2,}`)
	replaceFontsRegexp = regexp.MustCompile(`(?i)<(\/?)\s*font[^>]*?>`)

	blacklistCandidatesRegexp  = regexp.MustCompile(`(?i)popupbody`)
	okMaybeItsACandidateRegexp = regexp.MustCompile(`(?i)and|article|body|column|main|shadow`)
	unlikelyCandidatesRegexp   = regexp.MustCompile(`(?i)combx|comment|community|hidden|disqus|modal|extra|foot|header|menu|remark|rss|shoutbox|sidebar|sponsor|ad-break|agegate|pagination|pager|popup`)
	divToPElementsRegexp       = regexp.MustCompile(`(?i)<(a|blockquote|dl|div|img|ol|p|pre|table|ul)`)

	negativeRegexp = regexp.MustCompile(`(?i)combx|comment|com-|foot|footer|footnote|masthead|media|meta|outbrain|promo|related|scroll|shoutbox|sidebar|sponsor|shopping|tags|tool|widget`)
	positiveRegexp = regexp.MustCompile(`(?i)article|body|content|entry|hentry|main|page|pagination|post|text|blog|story`)

	stripCommentRegexp = regexp.MustCompile(`(?s)\<\!\-{2}.+?-{2}\>`)

	sentenceRegexp = regexp.MustCompile(`\.( |$)`)

	normalizeWhitespaceRegexp = regexp.MustCompile(`[\r\n\f]+`)
)

type candidate struct {
	selection *goquery.Selection
	score     float32
}

func (c *candidate) Node() *html.Node {
	return c.selection.Get(0)
}

type Document struct {
	input         string
	document      *goquery.Document
	content       string
	candidates    map[*html.Node]*candidate
	bestCandidate *candidate

	RemoveUnlikelyCandidates bool
	WeightClasses            bool
	CleanConditionally       bool
	BestCandidateHasImage    bool
	RetryLength              int
	MinTextLength            int
	RemoveEmptyNodes         bool
	WhitelistTags            []string
}

func NewDocument(s string) (*Document, error) {
	d := &Document{
		input:                    s,
		WhitelistTags:            []string{"div", "p"},
		RemoveUnlikelyCandidates: true,
		WeightClasses:            true,
		CleanConditionally:       true,
		RetryLength:              250,
		MinTextLength:            25,
		RemoveEmptyNodes:         true,
	}
	err := d.initializeHtml(s)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Document) initializeHtml(s string) error {
	// replace consecutive <br>'s with p tags
	s = replaceBrsRegexp.ReplaceAllString(s, "</p><p>")

	// replace font tags
	s = replaceFontsRegexp.ReplaceAllString(s, `<${1}span>`)

	// manually strip regexps since html parser seems to miss some
	s = stripCommentRegexp.ReplaceAllString(s, "")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		return err
	}

	// if no body (like from a redirect or empty string)
	if doc.Find("body").Length() == 0 {
		s = "<body/>"
		return d.initializeHtml(s)
	}

	d.document = doc
	return nil
}

func (d *Document) Content() string {
	if d.content == "" {
		d.prepareCandidates()

		article := d.getArticle()
		articleText := d.sanitize(article)

		length := len(strings.TrimSpace(articleText))
		if length < d.RetryLength {
			retry := true

			if d.RemoveUnlikelyCandidates {
				d.RemoveUnlikelyCandidates = false
			} else if d.WeightClasses {
				d.WeightClasses = false
			} else if d.CleanConditionally {
				d.CleanConditionally = false
			} else {
				d.content = articleText
				retry = false
			}

			if retry {
				Logger.Printf("Retrying with length %d < retry length %d\n", length, d.RetryLength)
				err := d.initializeHtml(d.input)
				if err != nil {
					return ""
				}
				articleText = d.Content()
			}
		}

		d.content = articleText
	}

	return d.content
}

func (d *Document) Text() string {
	s := d.Content() // Genergate doc

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		return ""
	}

	return strings.Trim(doc.Text(), " ")
}

func (d *Document) prepareCandidates() {
	// noscript might be valid, but probably not so we'll just remove it
	d.document.Find("script, style,noscript").Each(func(i int, s *goquery.Selection) {
		removeNodes(s)
	})

	if d.RemoveUnlikelyCandidates {
		d.removeUnlikelyCandidates()
	}

	d.transformMisusedDivsIntoParagraphs()
	d.scoreParagraphs(d.MinTextLength)
	d.selectBestCandidate()
}

func (d *Document) selectBestCandidate() {
	var best *candidate

	for _, c := range d.candidates {
		if best == nil {
			best = c
		} else if best.score < c.score {
			best = c
		}
	}

	if best == nil {
		best = &candidate{d.document.Find("body"), 0}
	}

	d.bestCandidate = best
}

func (d *Document) getArticle() string {
	output := bytes.NewBufferString("<div>")

	siblingScoreThreshold := float32(math.Max(10, float64(d.bestCandidate.score*.2)))

	d.bestCandidate.selection.Siblings().Union(d.bestCandidate.selection).Each(func(i int, s *goquery.Selection) {
		append := false
		n := s.Get(0)

		if n == d.bestCandidate.Node() {
			append = true
		} else if c, ok := d.candidates[n]; ok && c.score >= siblingScoreThreshold {
			append = true
		}

		if s.Is("p") {
			linkDensity := d.getLinkDensity(s)
			content := s.Text()
			contentLength := len(content)

			if contentLength >= 80 && linkDensity < .25 {
				append = true
			} else if contentLength < 80 && linkDensity == 0 {
				append = sentenceRegexp.MatchString(content)
			}
		}

		if append {
			tag := "div"
			if s.Is("p") {
				tag = n.Data
			}

			html, _ := s.Html()
			fmt.Fprintf(output, "<%s>%s</%s>", tag, html, tag)
		}
	})

	output.Write([]byte("</div>"))

	return output.String()
}

func (d *Document) removeUnlikelyCandidates() {
	d.document.Find("*").Not("html,body").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		id, _ := s.Attr("id")

		str := class + id

		if blacklistCandidatesRegexp.MatchString(str) || (unlikelyCandidatesRegexp.MatchString(str) && !okMaybeItsACandidateRegexp.MatchString(str)) {
			Logger.Printf("Removing unlikely candidate - %s\n", str)
			removeNodes(s)
		}
	})
}

func (d *Document) transformMisusedDivsIntoParagraphs() {
	d.document.Find("div").Each(func(i int, s *goquery.Selection) {
		html, err := s.Html()
		if err != nil {
			Logger.Printf("Unable to transform div to p %s\n", err)
			return
		}

		// transform <div>s that do not contain other block elements into <p>s
		if !divToPElementsRegexp.MatchString(html) {
			class, _ := s.Attr("class")
			id, _ := s.Attr("id")
			Logger.Printf("Altering div(#%s.%s) to p\n", id, class)

			node := s.Get(0)
			node.Data = "p"
		}
	})
}

func (d *Document) scoreParagraphs(minimumTextLength int) {
	candidates := make(map[*html.Node]*candidate)

	d.document.Find("p,td").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		// if this paragraph is less than x chars, don't count it
		if len(text) < minimumTextLength {
			return
		}

		parent := s.Parent()
		parentNode := parent.Get(0)

		grandparent := parent.Parent()
		var grandparentNode *html.Node
		if grandparent.Length() > 0 {
			grandparentNode = grandparent.Get(0)
		}

		if _, ok := candidates[parentNode]; !ok {
			candidates[parentNode] = d.scoreNode(parent)
		}
		if grandparentNode != nil {
			if _, ok := candidates[grandparentNode]; !ok {
				candidates[grandparentNode] = d.scoreNode(grandparent)
			}
		}

		contentScore := float32(1.0)
		contentScore += float32(strings.Count(text, ",") + 1)
		contentScore += float32(math.Min(float64(int(len(text)/100.0)), 3))

		candidates[parentNode].score += contentScore
		if grandparentNode != nil {
			candidates[grandparentNode].score += contentScore / 2.0
		}
	})

	// scale the final candidates score based on link density. Good content
	// should have a relatively small link density (5% or less) and be mostly
	// unaffected by this operation
	for _, candidate := range candidates {
		candidate.score = candidate.score * (1 - d.getLinkDensity(candidate.selection))
	}

	d.candidates = candidates
}

func (d *Document) getLinkDensity(s *goquery.Selection) float32 {
	linkLength := len(s.Find("a").Text())
	textLength := len(s.Text())

	if textLength == 0 {
		return 0
	}

	return float32(linkLength) / float32(textLength)
}

func (d *Document) classWeight(s *goquery.Selection) int {
	weight := 0
	if !d.WeightClasses {
		return weight
	}

	class, _ := s.Attr("class")
	id, _ := s.Attr("id")

	if class != "" {
		if negativeRegexp.MatchString(class) {
			weight -= 25
		}

		if positiveRegexp.MatchString(class) {
			weight += 25
		}
	}

	if id != "" {
		if negativeRegexp.MatchString(id) {
			weight -= 25
		}

		if positiveRegexp.MatchString(id) {
			weight += 25
		}
	}

	return weight
}

func (d *Document) scoreNode(s *goquery.Selection) *candidate {
	contentScore := d.classWeight(s)
	if s.Is("div") {
		contentScore += 5
	} else if s.Is("blockquote,form") {
		contentScore = 3
	} else if s.Is("th") {
		contentScore -= 5
	}

	return &candidate{s, float32(contentScore)}
}

func (d *Document) sanitize(article string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(article))
	if err != nil {
		Logger.Println("Unable to create document", err)
		return ""
	}

	s := doc.Find("body")
	s.Find("h1,h2,h3,h4,h5,h6").Each(func(i int, header *goquery.Selection) {
		if d.classWeight(header) < 0 || d.getLinkDensity(header) > 0.33 {
			removeNodes(header)
		}
	})

	s.Find("input,select,textarea,button,object,iframe,embed").Each(func(i int, s *goquery.Selection) {
		removeNodes(s)
	})

	if d.RemoveEmptyNodes {
		s.Find("p").Each(func(i int, s *goquery.Selection) {
			html, _ := s.Html()
			if len(strings.TrimSpace(html)) == 0 {
				removeNodes(s)
			}
		})
	}

	d.cleanConditionally(s, "table,ul,div")

	// we'll sanitize all elements using a whitelist
	replaceWithWhitespace := map[string]bool{
		"br":         true,
		"hr":         true,
		"h1":         true,
		"h2":         true,
		"h3":         true,
		"h4":         true,
		"h5":         true,
		"h6":         true,
		"dl":         true,
		"dd":         true,
		"ol":         true,
		"li":         true,
		"ul":         true,
		"address":    true,
		"blockquote": true,
		"center":     true,
	}

	whitelist := make(map[string]bool)
	for _, tag := range d.WhitelistTags {
		tag = strings.ToLower(tag)
		whitelist[tag] = true
		delete(replaceWithWhitespace, tag)
	}

	var text string

	s.Find("*").Each(func(i int, s *goquery.Selection) {
		if text != "" {
			return
		}

		// only look at element nodes
		node := s.Get(0)
		if node.Type != html.ElementNode {
			return
		}

		// if element is in whitelist, delete all its attributes
		if _, ok := whitelist[node.Data]; ok {
			node.Attr = make([]html.Attribute, 0)
		} else {
			if _, ok := replaceWithWhitespace[node.Data]; ok {
				// just replace with a text node and add whitespace
				node.Data = fmt.Sprintf(" %s ", s.Text())
				node.Type = html.TextNode
				node.FirstChild = nil
				node.LastChild = nil
			} else {
				if node.Parent == nil {
					text = s.Text()
					return
				} else {
					// replace node with children
					replaceNodeWithChildren(node)
				}
			}
		}
	})

	if text == "" {
		text, _ = doc.Html()
	}

	return normalizeWhitespaceRegexp.ReplaceAllString(text, "\n")
}

func (d *Document) cleanConditionally(s *goquery.Selection, selector string) {
	if !d.CleanConditionally {
		return
	}

	s.Find(selector).Each(func(i int, s *goquery.Selection) {
		node := s.Get(0)
		weight := float32(d.classWeight(s))
		contentScore := float32(0)

		if c, ok := d.candidates[node]; ok {
			contentScore = c.score
		}

		if weight+contentScore < 0 {
			removeNodes(s)
			Logger.Printf("Conditionally cleaned %s%s with weight %f and content score %f\n", node.Data, getName(s), weight, contentScore)
			return
		}

		text := s.Text()
		if strings.Count(text, ",") < 10 {
			counts := map[string]int{
				"p":     s.Find("p").Length(),
				"img":   s.Find("img").Length(),
				"li":    s.Find("li").Length() - 100,
				"a":     s.Find("a").Length(),
				"embed": s.Find("embed").Length(),
				"input": s.Find("input").Length(),
			}

			contentLength := len(strings.TrimSpace(text))
			linkDensity := d.getLinkDensity(s)
			remove := false
			reason := ""

			if counts["img"] > counts["p"] {
				reason = "too many images"
				remove = true
			} else if counts["li"] > counts["p"] && !s.Is("ul,ol") {
				reason = "more <li>s than <p>s"
				remove = true
			} else if counts["input"] > int(counts["p"]/3.0) {
				reason = "less than 3x <p>s than <input>s"
				remove = true
			} else if contentLength < d.MinTextLength && (counts["img"] == 0 || counts["img"] > 2) {
				reason = "too short content length without a single image"
				remove = true
			} else if weight < 25 && linkDensity > 0.2 {
				reason = fmt.Sprintf("too many links for its weight (%f)", weight)
				remove = true
			} else if weight >= 25 && linkDensity > 0.5 {
				reason = fmt.Sprintf("too many links for its weight (%f)", weight)
				remove = true
			} else if (counts["embed"] == 1 && contentLength < 75) || counts["embed"] > 1 {
				reason = "<embed>s with too short a content length, or too many <embed>s"
				remove = true
			}

			if remove {
				Logger.Printf("Conditionally cleaned %s%s with weight %f and content score %f because it has %s\n", node.Data, getName(s), weight, contentScore, reason)
				removeNodes(s)
			}
		}
	})
}

func getName(s *goquery.Selection) string {
	class, _ := s.Attr("class")
	id, _ := s.Attr("id")

	return fmt.Sprintf("#%s.%s", id, class)
}

func removeNodes(s *goquery.Selection) {
	s.Each(func(i int, s *goquery.Selection) {
		parent := s.Parent()
		if parent.Length() == 0 {

		} else {
			parent.Get(0).RemoveChild(s.Get(0))
		}
	})
}

func replaceNodeWithChildren(n *html.Node) {
	var next *html.Node
	parent := n.Parent

	for c := n.FirstChild; c != nil; c = next {
		next = c.NextSibling
		n.RemoveChild(c)

		parent.InsertBefore(c, n)
	}

	parent.RemoveChild(n)
}
