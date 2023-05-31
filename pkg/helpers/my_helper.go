package helpers

import (
	"fmt"
	"github.com/GRbit/go-pcre"
	"github.com/gosimple/slug"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"html"
)

func AutoLinkedTags(tags []string, content string, typeID *int64) string {
	var findTag []string
	var replaceTag []string

	for key, name := range tags {
		search := html.UnescapeString(name)
		findStringTitle := fmt.Sprintf(`\b(%s)\b|<a[^>]*>[^>]*>(*SKIP)(*F)|<img[^>]*>(*SKIP)(*F)/i`, cases.Title(language.Indonesian).String(search))
		if key == 1 {
			findTag = append(findTag, findStringTitle)
			if *typeID == 2 {
				replaceTag = append(replaceTag, "<a rel='dofollow' target='_blank' href='https://www.sportstars.id/tag/"+slug.Make(name)+"'>"+cases.Title(language.Indonesian).String(name)+"</a>")
			} else if *typeID == 34 {
				replaceTag = append(replaceTag, "<a rel='dofollow' target='_blank' href='https://www.idxchannel.com/tag/"+slug.Make(name)+"'>"+cases.Title(language.Indonesian).String(name)+"</a>")
			} else if *typeID == 35 {
				replaceTag = append(replaceTag, "<a rel='dofollow' target='_blank' href='https://www.celebrities.id/tag/"+slug.Make(name)+"'>"+cases.Title(language.Indonesian).String(name)+"</a>")
			} else {
				replaceTag = append(replaceTag, "<a rel='dofollow' href='https://www.inews.id/tag/"+slug.Make(name)+"'>"+cases.Title(language.Indonesian).String(name)+"</a>")
			}
		} else {
			findTag = append(findTag, findStringTitle)
			replaceTag = append(replaceTag, "<a rel='dofollow' href='https://www.inews.id/tag/"+slug.Make(name)+"'>"+cases.Title(language.Indonesian).String(name)+"</a>")
		}
	}

	return PregReplace(content, findTag, replaceTag)
}

func PregReplace(str string, original []string, replacement []string) string {

	for i, toreplace := range original {
		regex := pcre.MustCompile(toreplace, 0)
		str = regex.ReplaceAllString(str, replacement[i], 0)
	}

	return str
}
