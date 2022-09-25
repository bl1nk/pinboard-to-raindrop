package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
)

type PinboardBookmark struct {
	Href        string `json:"href"`
	Description string `json:"description"`
	Extended    string `json:"extended"`
	Time        string `json:"time"`
	ToRead      string `json:"toread"`
	Tags        string `json:"tags"`
}

var (
	input         string
	output        string
	defaultFolder string
	unreadFolder  string
)

func main() {
	flag.StringVar(&input, "input", "", "JSON export of Pinboard bookmarks")
	flag.StringVar(&output, "output", "", "CSV file to write to")
	flag.StringVar(&defaultFolder, "default-folder", "Imported", "Folder to put imported bookmarks in")
	flag.StringVar(&unreadFolder, "unread-folder", "To Read", "Folder to put unread imported bookmarks in")
	flag.Parse()

	if input == "" {
		log.Fatal("Must set -input")
	}
	if output == "" {
		log.Fatal("Must set -output")
	}

	if err := run(input, output); err != nil {
		log.Fatal(err)
	}
}

// run takes a input file and output file and returns an error if conversion
// from pinboard bookmarks.json to csv fails.
func run(input, output string) error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()

	var bookmarks []PinboardBookmark
	if err = json.NewDecoder(f).Decode(&bookmarks); err != nil {
		return err
	}

	var records [][]string
	for _, bookmark := range bookmarks {
		records = append(records, convert(bookmark))
	}

	o, err := os.Create(output)
	if err != nil {
		return err
	}
	defer o.Close()

	w := csv.NewWriter(o)
	if err = w.Write([]string{"url", "folder", "title", "description", "tags", "created"}); err != nil {
		return err
	}
	if err = w.WriteAll(records); err != nil {
		return err
	}

	return nil
}

// convert converts a PinboardBookmark to a CSV record to be imported on
// Raindrop.io
func convert(b PinboardBookmark) []string {
	url := b.Href
	folder := defaultFolder
	if b.ToRead == "yes" {
		folder = unreadFolder
	}
	title := b.Description
	description := b.Extended
	tags := strings.Join(strings.Split(b.Tags, " "), ", ")
	created := b.Time
	return []string{url, folder, title, description, tags, created}
}
