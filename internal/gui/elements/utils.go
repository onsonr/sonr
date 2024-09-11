package elements

import (
	"strings"

	"github.com/a-h/templ"
)

type Icon interface {
	Render() templ.Component
}

type Variant interface {
	Attributes() templ.Attributes
}

func clsxMerge(variants ...Variant) templ.Attributes {
	combinedAttrs := templ.Attributes{}
	var classElements []string

	for _, variant := range variants {
		attrs := variant.Attributes()
		if class, ok := attrs["class"].(string); ok {
			classElements = append(classElements, strings.Fields(class)...)
		}
		for key, value := range attrs {
			if key != "class" {
				combinedAttrs[key] = value
			}
		}
	}

	if len(classElements) > 0 {
		combinedAttrs["class"] = strings.Join(classElements, " ")
	}
	return combinedAttrs
}
