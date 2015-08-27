package gothumbor_test

import (
	"github.com/globocom/gothumbor"
	"testing"
)

const myKey = "my-security-key"
const width = 300
const height = 200
const encryptedURL = "8ammJH8D-7tXy6kU3lTvoXlhu4o=/300x200/my.server.com/some/path/to/image.jpg"
const unsafeURL = "/300x200/my.server.com/some/path/to/image.jpg"
const imageURL = "my.server.com/some/path/to/image.jpg"

func TestGetUrlUnderSpec1(t *testing.T) {
	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries

	thumborOptions := gothumbor.ThumborOptions{Width: width, Height: height}
	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)

	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
	}

	if newURL != encryptedURL {
		t.Error("Got an unxpected thumbor path:", newURL)
	}
}
