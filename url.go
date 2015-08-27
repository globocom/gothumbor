package gothumbor

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

type ThumborOptions struct {
	Width  int
	Height int
	Smart  bool
}

func GetThumborPath(key, imageURL string, options ThumborOptions) (url string, err error) {
	var partial string
	if partial, err = GetURLParts(imageURL, options); err != nil {
		return
	}

	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(partial))
	message := hash.Sum(nil)
	url = base64.URLEncoding.EncodeToString(message)
	url = strings.Join([]string{url, partial}, "/")
	return url, err
}

func GetURLParts(imageURL string, options ThumborOptions) (urlPartial string, err error) {

	var parts []string
	parts = append(parts, fmt.Sprintf("%dx%d", options.Width, options.Height))

	if options.Smart {
		parts = append(parts, "smart")
	}

	parts = append(parts, imageURL)
	urlPartial = strings.Join(parts, "/")
	return urlPartial, err
}
