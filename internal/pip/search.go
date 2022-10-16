package pip

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/0x00-ketsu/lazypip/internal/utils"
	"github.com/PuerkitoBio/goquery"
)

func Search(query string, page int) ([]Package, string, error) {
	body, reqURL, err := doQuery(query, page)
	if err != nil {
		return nil, reqURL, err
	}

	packages, err := parseHTML(body)
	return packages, reqURL, err
}

// Parse html, extract package info and return
func parseHTML(body []byte) ([]Package, error) {
	var packages []Package

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		logger.Error(fmt.Sprintf("[pip search] %v", err.Error()))
		return nil, err
	}

	doc.Find("a[class*=\"package-snippet\"]").Each(func(i int, s *goquery.Selection) {
		name := s.Find("span[class*=\"package-snippet__name\"]").Text()
		version := s.Find("span[class*=\"package-snippet__version\"]").Text()
		released := s.Find("span[class*=\"package-snippet__created\"]").Find("time").Text()
		desc := s.Find("p[class*=\"package-snippet__description\"]").Text()
		link, _ := utils.JoinPath(conf.Pip.IndexURL, s.AttrOr("href", ""))

		packages = append(packages, Package{
			Name:        strings.TrimSpace(name),
			Version:     strings.TrimSpace(version),
			Released:    strings.TrimSpace(released),
			Description: strings.TrimSpace(desc),
			Link:        link,
		})
	})

	return packages, nil
}

// Send query request to Pypi search URL
// Return response Body & request URL
func doQuery(query string, page int) ([]byte, string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, conf.Pip.SearchURL, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("[pip search] %v", err.Error()))
		return nil, "", err
	}

	q := req.URL.Query()
	q.Add("q", query)
	q.Add("page", strconv.Itoa(page))

	req.URL.RawQuery = q.Encode()
	logger.Info(fmt.Sprintf("[pip search] send request: %v", req.URL))

	resp, err := client.Do(req)
	if err != nil {
		msg := "[pip search] Errored when sending request to the pypi search URL"
		logger.Error(msg)
		return nil, "", errors.New(msg)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("[pip search] Status code error: %d", resp.StatusCode)
		logger.Error(msg)
		return nil, "", errors.New(msg)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("[pip search] Read response body failed: %v", err.Error()))
		return nil, "", errors.New("[pip search] Read response body failed")
	}

	return respBody, req.URL.String(), nil
}
