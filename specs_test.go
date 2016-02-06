package gothumbor_test

import (
	"testing"

	"github.com/globocom/gothumbor"
)

const myKey = "my-security-key"
const width = 300
const height = 200
const imageURL = "my.server.com/some/path/to/image.jpg"
const widthHeightEncryptedURL = "8ammJH8D-7tXy6kU3lTvoXlhu4o=/300x200/my.server.com/some/path/to/image.jpg"
const metaEncryptedURL = "Ps3ORJDqxlSQ8y00T29GdNAh2CY=/meta/my.server.com/some/path/to/image.jpg"
const smartEncryptedURL = "-2NHpejRK2CyPAm61FigfQgJBxw=/smart/my.server.com/some/path/to/image.jpg"
const fitInEncryptedURL = "uvLnA6TJlF-Cc-L8z9pEtfasO3s=/fit-in/my.server.com/some/path/to/image.jpg"
const filtersEncrytedURL = "H49B0suv7d2eZqsvy9oR9fsgeSM=/filters:quality(20):brightness(10)/my.server.com/some/path/to/image.jpg"

func TestGetUrlScenario1(t *testing.T) {
	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries#scenario-1---signing-of-a-known-url-results

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
	//For spec 3: https://github.com/thumbor/thumbor/wiki/Libraries#scenario-3---thumbor-matching-of-signature-with-my-library-signature-with-meta
	thumborOptions := gothumbor.ThumborOptions{Meta: true}
	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)

	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
	}

	if newURL != metaEncryptedURL {
		t.Error("Got an unxpected thumbor path:", newURL)
	}
}

func TestGetUrlScenario4(t *testing.T) {
	//For spec 4: https://github.com/thumbor/thumbor/wiki/Libraries#scenario-4---thumbor-matching-of-signature-with-my-library-signature-with-smart

	thumborOptions := gothumbor.ThumborOptions{Smart: true}
	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)

	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
	}

	if newURL != smartEncryptedURL {
		t.Error("Got an unxpected thumbor path:", newURL)
	}
}

func TestGetUrlScenario5(t *testing.T) {
	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries#scenario-5---thumbor-matching-of-signature-with-my-library-signature-with-fit-in

	thumborOptions := gothumbor.ThumborOptions{FitIn: true}
	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)

	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
	}

	if newURL != fitInEncryptedURL {
		t.Error("Got an unxpected thumbor path:", newURL)
	}
}

func TestGetUrlScenario6(t *testing.T) {
	//For spec 1: https://github.com/thumbor/thumbor/wiki/Libraries#scenario-6---thumbor-matching-of-signature-with-my-library-signature-with-filters

	thumborOptions := gothumbor.ThumborOptions{Filters: []string{"quality(20)", "brightness(10)"}}
	newURL, err := gothumbor.GetCryptedThumborPath(myKey, imageURL, thumborOptions)

	if err != nil {
		t.Errorf("Got an error when tried to generate the thumbor url:%s", err)
	}

	if newURL != filtersEncrytedURL {
		t.Error("Got an unxpected thumbor path:", newURL)
	}
}
