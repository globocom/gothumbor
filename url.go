package gothumbor

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

type ThumborOptions struct {
	Width   int
	Height  int
	Smart   bool
	FitIn   bool
	Filters []string
	Flip    bool
	Flop    bool
	Meta    bool
}

func GetCryptedThumborPath(key, imageURL string, options ThumborOptions) (url string, err error) {
	var partial string
	if partial, err = GetThumborPath(imageURL, options); err != nil {
		return
	}
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(partial))
	message := hash.Sum(nil)
	url = base64.URLEncoding.EncodeToString(message)
	url = strings.Join([]string{url, partial}, "/")
	return
}

func GetThumborPath(imageURL string, options ThumborOptions) (path string, err error) {
	if path, err = getURLParts(imageURL, options); err != nil {
		return
	}
	return
}

func getURLParts(imageURL string, options ThumborOptions) (urlPartial string, err error) {

	var parts []string

	partialObject, err := url.Parse(imageURL)
	if err != nil {
		return "", err
	}
	newImageURL := partialObject.String()
	if options.Meta {
		parts = append(parts, "meta")
	}

	if options.Height != 0 || options.Width != 0 || options.Flip || options.Flop {
		flip := ""
		flop := ""

		if options.Flip {
			flip = "-"
		}
		if options.Flop {
			flop = "-"
		}
		parts = append(parts, fmt.Sprintf("%s%dx%s%d", flip, options.Width, flop, options.Height))
	}

	filters := []string{}
	for _, value := range options.Filters {
		filters = append(filters, value)
	}

	if options.Smart {
		parts = append(parts, "smart")
	}

	if options.FitIn {
		parts = append(parts, "fit-in")
	}

	if len(options.Filters) > 0 {
		filtersValue := strings.Join(filters, ":")
		parts = append(parts, "filters:"+filtersValue)
	}

	parts = append(parts, newImageURL)
	urlPartial = strings.Join(parts, "/")

	return
}
