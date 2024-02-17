package server

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

type Job struct {
	Title, Description, Link string
	Date                     time.Time
}

type Server struct {
	ntfyTopic     string
	upworkRssUrls []string
	feedParser    *gofeed.Parser
}

func (s *Server) ParseRssFeed(rssUrl string) []Job {
	feed, err := s.feedParser.ParseURL(rssUrl)
	if err != nil {
		panic(err)
	}
	jobs := make([]Job, len(feed.Items))
	for _, f := range feed.Items {
		cur := Job{
			Title:       f.Title,
			Description: f.Description,
			Link:        f.Link,
			Date:        *f.PublishedParsed,
		}
		jobs = append(jobs, cur)
	}
	return jobs
}

func (s *Server) Notify(job Job) {
	fmt.Printf("Title: %s\n", job.Title)
}

func (s *Server) Fetch() {
	for _, url := range s.upworkRssUrls {
		jobs := s.ParseRssFeed(url)
		for _, job := range jobs {
			s.Notify(job)
		}
	}
}

func (s *Server) Run() {
	s.Fetch()
}

func NewServer(ntfyTopic string, upworkRssUrls []string) *Server {
	return &Server{ntfyTopic: ntfyTopic, upworkRssUrls: upworkRssUrls, feedParser: gofeed.NewParser()}
}
