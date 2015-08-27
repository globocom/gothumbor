package gothumbor

import (
	"strings"
	"testing"
)

const IMAGEURL = "my.server.com/some/path/to/image.jpg"
const WIDTH = 300
const HEIGHT = 200
const ENCRYPTEDURL = "8ammJH8D-7tXy6kU3lTvoXlhu4o=/300x200/my.server.com/some/path/to/image.jpg"
const UNSAFEURL = "/300x200/my.server.com/some/path/to/image.jpg"

func TestGetUrlPartialWithWidthAndHeight(t *testing.T) {
	thumborOptions := ThumborOptions{Width: 1, Height: 1, Smart: false}
	url, err := getURLParts(IMAGEURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}

	urlE := strings.Join([]string{"1x1", IMAGEURL}, "/")
	if url != urlE {
		t.Errorf("Got an unxpected partial url: %s != %s", url, urlE)
	}
}

func TestGetUrlPartialWithSmart(t *testing.T) {
	thumborOptions := ThumborOptions{Width: 1, Height: 1, Smart: true}
	url, err := getURLParts(IMAGEURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}

	urlE := strings.Join([]string{"1x1", "smart", IMAGEURL}, "/")
	if url != urlE {
		t.Errorf("Got an unxpected partial url: %s != %s", url, urlE)
	}
}

func TestGetUrlPartialOnlyWithWidthAndHeight(t *testing.T) {
	thumborOptions := ThumborOptions{Width: WIDTH, Height: HEIGHT}
	url, err := getURLParts(IMAGEURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}
	if url != strings.Join([]string{"300x200", IMAGEURL}, "/") {
		t.Error("Got an unxpected partial path:", url)
	}
}


