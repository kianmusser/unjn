/*
Copyright © 2024 Kian Musser
Freely available under the MIT license
*/
package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/k3a/html2text"
	"github.com/mmcdole/gofeed"
)

// A JSON payload to send to ntfy.sh
type NtfySubmission struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
	Title   string `json:"title"`
}

// A parsed job description from Upwork
type Job struct {
	Title       string
	Description string
	Link        string
	Date        time.Time
}

type Server struct {
	ntfyTopic     string
	upworkRssUrls []string
	feedParser    *gofeed.Parser
	seenJobs      map[string]bool
}

func (s *Server) ParseRssFeed(rssUrl string) []Job {
	feed, err := s.feedParser.ParseURL(rssUrl)
	if err != nil {
		log.Printf("error retrieving RSS feed: %s\n", err)
		return []Job{}
	}

	jobs := make([]Job, len(feed.Items))
	for i, f := range feed.Items {
		cur := Job{
			Title:       f.Title,
			Description: html2text.HTML2Text(f.Description),
			Link:        f.Link,
			Date:        *f.PublishedParsed,
		}
		jobs[i] = cur
	}
	return jobs
}

func (s *Server) Notify(job Job) {
	req := NtfySubmission{Topic: s.ntfyTopic, Title: job.Title, Message: job.Description}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("https://ntfy.sh", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func (s *Server) Fetch() {
	log.Println("fetching...")
	for _, url := range s.upworkRssUrls {
		jobs := s.ParseRssFeed(url)
		for _, job := range jobs {
			if !s.seenJobs[job.Link] {
				s.seenJobs[job.Link] = true
				s.Notify(job)
			}
		}
	}
}

func (s *Server) Run() {
	ticker := time.NewTicker(10 * time.Minute)
	s.Fetch()
	for {
		<-ticker.C
		s.Fetch()
	}
}

func NewServer(ntfyTopic string, upworkRssUrls []string) *Server {
	return &Server{
		ntfyTopic:     ntfyTopic,
		upworkRssUrls: upworkRssUrls,
		feedParser:    gofeed.NewParser(),
		seenJobs:      make(map[string]bool),
	}
}
