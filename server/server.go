package server

import "github.com/mmcdole/gofeed"

type Server struct {
	ntfyTopic     string
	upworkRssUrls []string
	feedParser    *gofeed.Parser
}

func (s *Server) ParseRssFeed(rssUrl) {
	s.feedParser.ParseURL()
}

func (s *Server) Run() {
}

func NewServer(ntfyTopic string, upworkRssUrls []string) *Server {
	return &Server{ntfyTopic: ntfyTopic, upworkRssUrls: upworkRssUrls, feedParser: gofeed.NewParser()}
}
