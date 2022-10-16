package utils

import (
	"net/url"
	"path"
)

// JoinPath concatenates baseUrl and the elements
// - check baseUrl format
// - concatenates baseUrl and the elements
func JoinPath(baseUrl string, elem ...string) (result string, err error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return
	}

	if len(elem) > 0 {
		elem = append([]string{url.Path}, elem...)
		url.Path = path.Join(elem...)
	}

	result = url.String()
	return
}
