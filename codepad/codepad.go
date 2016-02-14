package codepad

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

const CODEPAD_URL = "http://codepad.org"

var re = regexp.MustCompile(CODEPAD_URL + "/([[:alnum:]]+)")

// Extract codepad ID from the given URL.
func GetID(url string) (string, error) {
	var match []string

	match = re.FindStringSubmatch(url)
	if len(match) != 2 {
		return "", errors.New("invalid codepad url")
	}
	return match[1], nil
}

// Get the data from the paste ID.
// This returns the entire data as a string.
func Get(id string) (string, error) {
	var err error
	var reader io.Reader
	var body []byte

	reader, err = GetReader(id)
	if err != nil {
		return "", err
	}

	body, err = ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("read server response error:", err)
		return "", err
	}
	return string(body), nil
}

// Get the reader for the paste ID.
// The caller can then read from this reader, when reading is done,
// the caller need to close the reader.
func GetReader(id string) (io.Reader, error) {
	var err error
	var resp *http.Response
	var url string

	url = fmt.Sprintf("%s/%s/raw.*", CODEPAD_URL, id)

	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("error while get paste from codepad:", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Println("server returned an error:", resp.Status)
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

// Send the data string to paste server.
// Returns the ID of this paste if succeeds.
func Put(data string) (string, error) {
	return "", errors.New("Not implemented")
}