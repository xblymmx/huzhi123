package utils

import (
	"net/url"
	"fmt"
)

// convert absolute url to relative url
func RelativeUrlToAbsoluteUrl(curURL string, baseURL string) (string, error) {
	curURLData, err := url.Parse(curURL)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	baseURLData, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	curURLData = baseURLData.ResolveReference(curURLData)
	return curURLData.String(), nil
}