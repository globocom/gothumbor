package gothumbor_test

import (
	"github.com/globocom/gothumbor"
	"testing"
)

const myKey = "my-security-key"
const width = 300
const height = 200
const imageURL = "my.server.com/some/path/to/image.jpg"
const widthHeightEncryptedURL = "8ammJH8D-7tXy6kU3lTvoXlhu4o=/300x200/my.server.com/some/path/to/image.jpg"
const metaEncryptedURL = "Ps3ORJDqxlSQ8y00T29GdNAh2CY=/meta/my.server.com/some/path/to/image.jpg"

func TestGetUrlScenario1(t *testing.T) {
	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries

	thumborOptions := gothumbor.ThumborOptions{Width: width, Height: height}
	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)

	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
	}

	if newURL != widthHeightEncryptedURL {
		t.Error("Got an unxpected thumbor path:", newURL)
	}
}

func TestGetUrlScenario3(t *testing.T) {
	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries
	thumborOptions := gothumbor.ThumborOptions{Meta: true}
	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)

	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
	}

	if newURL != metaEncryptedURL {
		t.Error("Got an unxpected thumbor path:", newURL)
	}
}

//func TestGetUrlScenario4(t *testing.T) {
//	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries
//
//	thumborOptions := gothumbor.ThumborOptions{Width: width, Height: height}
//	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)
//
//	if err != nil {
//		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
//	}
//
//	if newURL != encryptedURL {
//		t.Error("Got an unxpected thumbor path:", newURL)
//	}
//}
//
//func TestGetUrlScenario5(t *testing.T) {
//	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries
//
//	thumborOptions := gothumbor.ThumborOptions{Width: width, Height: height}
//	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)
//
//	if err != nil {
//		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
//	}
//
//	if newURL != encryptedURL {
//		t.Error("Got an unxpected thumbor path:", newURL)
//	}
//}
//
//func TestGetUrlScenario6(t *testing.T) {
//	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries
//
//	thumborOptions := gothumbor.ThumborOptions{Width: width, Height: height}
//	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)
//
//	if err != nil {
//		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
//	}
//
//	if newURL != encryptedURL {
//		t.Error("Got an unxpected thumbor path:", newURL)
//	}
//}

