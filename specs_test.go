package gothumbor_test

import (
	"github.com/globocom/gothumbor"
	"testing"
)

const MYKEY = "my-security-key"
const WIDTH = 300
const HEIGHT = 200
const ENCRYPTEDURL = "8ammJH8D-7tXy6kU3lTvoXlhu4o=/300x200/my.server.com/some/path/to/image.jpg"
const UNSAFEURL = "/300x200/my.server.com/some/path/to/image.jpg"
const IMAGEURL = "my.server.com/some/path/to/image.jpg"

func TestGetUrlUnderSpec1(t *testing.T) {
	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries

	thumborOptions := gothumbor.ThumborOptions{Width: WIDTH, Height: HEIGHT}
	newURL, err := gothumbor.GetCryptedThumborPath(MYKEY, IMAGEURL, thumborOptions)

	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
	}

	if newURL != ENCRYPTEDURL {
		t.Error("Got an unxpected thumbor path:", newURL)
	}
}
