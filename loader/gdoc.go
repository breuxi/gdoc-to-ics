package loader

import (
	"errors"
	"io"
	"mime"
	"net/http"
	"strings"
	"time"
)

const GDOC_URL = "https://docs.google.com/spreadsheets/d/%id%/export?format=csv&gid=%sheet_id%"

type GDoc struct {
	Id         string
	Sheet_Id   string
	Content    string
	Filename   string
	Created_at time.Time
}

var cache = map[string]GDoc{}

func GetGDocCSV(id string, sheet_id string) (*GDoc, error) {
	cache_entry, cache_ok := cache[id+"_"+sheet_id]

	if cache_ok {
		diff := time.Now().Sub(cache_entry.Created_at)
		if diff.Minutes() < 1 {
			return &cache_entry, nil
		}
	}

	replacer := strings.NewReplacer("%id%", id, "%sheet_id%", sheet_id)
	url := replacer.Replace(GDOC_URL)

	resp, err := http.Get(url)

	if err != nil {
		if cache_ok {
			return &cache_entry, nil
		} else {
			return nil, err
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			if cache_ok {
				return &cache_entry, nil
			} else {
				return nil, err
			}
		}

		bodyString := string(bodyBytes)

		var filename string
		if cd := resp.Header.Get("Content-Disposition"); cd != "" {
			if _, params, err := mime.ParseMediaType(cd); err == nil {
				filename = params["filename"]
			}
		}

		filename_parts := strings.Split(filename, " ")
		if len(filename_parts) > 0 && len(filename_parts) > 2 {
			filename = strings.Join(filename_parts[:len(filename_parts)-2], " ")
		}

		doc := &GDoc{Id: id, Sheet_Id: sheet_id, Content: bodyString, Filename: filename, Created_at: time.Now()}

		cache[id+"_"+sheet_id] = *doc

		return doc, nil

	} else {
		if cache_ok {
			return &cache_entry, nil
		} else {
			return nil, errors.New(resp.Status)
		}
	}
}
