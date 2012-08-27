// Simple gmail ATOM parser
package rss

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Entry struct {
	Title    string `xml:"title"`
	Summary  string `xml:"summary"`
	Modified string `xml:"modified"`
	Author   Author `xml:"author"`
}

// Parses modification time string to time structure
func (e *Entry) ModifiedTime() (time.Time, error) {
	return time.Parse(time.RFC3339, e.Modified)
}

// Returns a list of Entry objects by parsing the url atom feed
func Read(url, user, pass string) ([]*Entry, error) {
	client := new(http.Client)

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Received bad status code: %v", resp.StatusCode))
	}
	defer resp.Body.Close()

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return unmarshal(text)
}

func unmarshal(text []byte) (es []*Entry, err error) {
	var feed struct {
		Entries []*Entry `xml:"entry"`
	}
	err = xml.Unmarshal(text, &feed)
	if err != nil {
		return nil, err
	}

	return feed.Entries, nil
}
