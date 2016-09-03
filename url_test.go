package gothumbor

import (
	"fmt"
	"strings"
	"testing"
)

const (
	baseImageURL = "my.server.com/some/path/to/image.jpg"
	imageURL     = "http://" + baseImageURL
	width        = 300
	height       = 200
	encryptedURL = "8ammJH8D-7tXy6kU3lTvoXlhu4o=/300x200/my.server.com/some/path/to/image.jpg"
	unsafeURL    = "/300x200/my.server.com/some/path/to/image.jpg"
)

func TestGetUrlPartialWithCrop(t *testing.T) {

	thumborOptions := ThumborOptions{Left: 1, Top: 2, Right: 3, Bottom: 4}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}

	urlE := strings.Join([]string{"1x2:3x4", imageURL}, "/")
	if url != urlE {
		t.Errorf("Got an unxpected partial url: %s != %s", url, urlE)
	}
}

func TestGetUrlPartialWithWidthAndHeight(t *testing.T) {
	thumborOptions := ThumborOptions{Width: 1, Height: 1, Smart: false}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}

	urlE := strings.Join([]string{"1x1", imageURL}, "/")
	if url != urlE {
		t.Errorf("Got an unxpected partial url: %s != %s", url, urlE)
	}
}

func TestGetUrlPartsShouldMaintainURL(t *testing.T) {
	thumborOptions := ThumborOptions{Width: 1, Height: 1}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}
	if !strings.HasSuffix(url, imageURL) {
		t.Errorf("Got an unxpected partial url: %s != %s", url, imageURL)
	}
}

func TestGetUrlPartialWithSmart(t *testing.T) {
	thumborOptions := ThumborOptions{Width: 1, Height: 1, Smart: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}

	urlE := strings.Join([]string{"1x1", "smart", imageURL}, "/")
	if url != urlE {
		t.Errorf("Got an unxpected partial url: %s != %s", url, urlE)
	}
}

func TestGetUrlPartialWithSmartTopFallback(t *testing.T) {
	thumborOptions := ThumborOptions{Width: 1, Height: 1, VAlign: "top", Smart: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}

	urlE := strings.Join([]string{"1x1", "top", "smart", imageURL}, "/")
	if url != urlE {
		t.Errorf("Got an unxpected partial url: %s != %s", url, urlE)
	}
}

func TestGetUrlPartialOnlyWithWidthAndHeight(t *testing.T) {
	thumborOptions := ThumborOptions{Width: width, Height: height}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}
	if url != strings.Join([]string{"300x200", imageURL}, "/") {
		t.Error("Got an unxpected partial path:", url)
	}
}

func TestEscapeURLByRFC3986(t *testing.T) {
	thumborOptions := ThumborOptions{}
	url, err := getURLParts("/a-path with spaces.jpg", thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to process a path with spaces", err)
	}
	if url != "/a-path%20with%20spaces.jpg" {
		t.Error("Got an unxpected partial path:", url)
	}
}

func TestFitInParameter(t *testing.T) {
	thumborOptions := ThumborOptions{FitIn: true}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if !strings.Contains(url, "fit-in") {
		t.Errorf("url doesn't have a fit-in parameter")
	}

}

func TestOneFilterParameter(t *testing.T) {
	filter := "max-age(360000)"
	filters := []string{filter}
	thumborOptions := ThumborOptions{Width: 200, Height: 300, Filters: filters}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if !strings.Contains(url, "filters:"+filter) {
		t.Errorf("url doesn't have a filters parameter")
	}

}

func TestTwoFiltersParameter(t *testing.T) {
	firstFilter := "max-age(360000)"
	secondFilter := "grayscale()"
	filters := []string{firstFilter, secondFilter}
	thumborOptions := ThumborOptions{Width: 200, Height: 300, Filters: filters}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if !strings.Contains(url, "filters:"+firstFilter+":"+secondFilter) {
		t.Errorf("url doesn't have a first filter parameter")
	}
}

func TestFlipWithoutAnyOtherParameter(t *testing.T) {
	thumborOptions := ThumborOptions{Flip: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if url != strings.Join([]string{"-0x0", imageURL}, "/") {
		t.Errorf("url %q is not fliped", url)
	}
}

func TestFlopWithoutAnyOtherParameter(t *testing.T) {
	thumborOptions := ThumborOptions{Flop: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if url != strings.Join([]string{"0x-0", imageURL}, "/") {
		t.Errorf("url %q is not floped", url)
	}
}

func TestFlipFlopWithoutAnyOtherParameter(t *testing.T) {
	thumborOptions := ThumborOptions{Flop: true, Flip: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if url != strings.Join([]string{"-0x-0", imageURL}, "/") {
		t.Errorf("url %q is not flipfloped", url)
	}
}

func TestFlipFlopWithWidthAndHeigh(t *testing.T) {
	thumborOptions := ThumborOptions{Flop: true, Flip: true, Height: 500, Width: 400}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if url != strings.Join([]string{"-400x-500", imageURL}, "/") {
		t.Errorf("url %q is not flipfloped", url)
	}
}

func TestFiltersAndSmartCombinatiotn(t *testing.T) {
	firstFilter := "max-age(360000)"
	secondFilter := "grayscale()"
	filters := []string{firstFilter, secondFilter}
	thumborOptions := ThumborOptions{Smart: true, Width: 200, Height: 300, Filters: filters}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	filtersPosition := strings.Index(url, "filters:"+firstFilter+":"+secondFilter)
	smartPosition := strings.Index(url, "smart")
	if filtersPosition < smartPosition {
		t.Errorf("Filters parameter should be before smart option")
	}
}

func TestFiltersAndFitInCombinatiotn(t *testing.T) {
	firstFilter := "max-age(360000)"
	secondFilter := "grayscale()"
	filters := []string{firstFilter, secondFilter}
	thumborOptions := ThumborOptions{FitIn: true, Width: 200, Height: 300, Filters: filters}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	filtersPosition := strings.Index(url, "filters:"+firstFilter+":"+secondFilter)
	fitInPosition := strings.Index(url, "fit-in")
	if filtersPosition < fitInPosition {
		t.Errorf("Filters parameter should be after fit-in option")
	}
}

func TestSetUpVerticalAlignment(t *testing.T) {
	thumborOptions := ThumborOptions{VAlign: "top", Width: 200, Height: 300}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}
	if !strings.Contains(url, "top") {
		t.Errorf("Vertical alignment is not applied on the url")
	}
}

func TestSetUpVerticalAlignmentBeforeWidthAndWidth(t *testing.T) {

	thumborOptions := ThumborOptions{VAlign: "bottom", Width: 200, Height: 300}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}
	WidthHeightPosition := strings.Index(url, "200x300")
	VAlignPosition := strings.Index(url, "bottom")
	shiftedPosition := WidthHeightPosition + len("200x300") + 1
	if VAlignPosition != (shiftedPosition) {
		fmt.Println("url : ", url)
		fmt.Println("widthxheight position: ", shiftedPosition)
		fmt.Println("valign position: ", VAlignPosition)
		t.Errorf("Valign parameter should be 2 characters before width and height option")
	}
}
