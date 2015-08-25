package gothumbor_test

import (
	"github.com/globocom/gothumbor"
	"testing"
	"strings"
)

const MYKEY = "my-security-key"
const IMAGEURL = "my.server.com/some/path/to/image.jpg"
const WIDTH = 300
const HEIGHT = 200
const ENCRYPTED_URL = "/8ammJH8D-7tXy6kU3lTvoXlhu4o=/300x200/my.server.com/some/path/to/image.jpg"


func TestGetUrlPartialWithDefautWidthAndHeight(t *testing.T){
	thumborOptions := gothumbor.ThumborOptions{}
	url, err := gothumbor.GetUrlParts(IMAGEURL, thumborOptions)
	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url", err)
	}
	if url != strings.Join([]string{"0x0", IMAGEURL}, "/") {
		t.Error("Got an unxpected partial url:", url)
	}
}

func TestGetUrlPartialOnlyWithWidthAndHeight(t *testing.T){
	thumborOptions := gothumbor.ThumborOptions{Width: WIDTH, Height: HEIGHT}
	url, err := gothumbor.GetUrlParts(IMAGEURL, thumborOptions)
	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url", err)
	}
	if url != strings.Join([]string{"300x200", IMAGEURL}, "/") {
		t.Error("Got an unxpected partial url:", url)
	}
}

func TestGetUrlUnderSpec1(t *testing.T) {
	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries

	thumborOptions := gothumbor.ThumborOptions{Width: WIDTH, Height: HEIGHT}

	newUrl, err := gothumbor.GetThumborPath(MYKEY, IMAGEURL, thumborOptions)
	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url", err)
	}
	if newUrl != ENCRYPTED_URL {
		t.Error("Got an unxpected thumbor url:", newUrl)
	}
}
