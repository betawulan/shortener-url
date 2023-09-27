package service

import (
	"math/rand"
	"time"
)

type shortenUrl struct {
	URLs          map[string]string
	ClickCount    map[string]int
	Expiry        map[string]time.Time
	DefaultExpiry time.Duration
}

func (s *shortenUrl) Generate(url, expire string) string {
	originalURL := url
	expiryParam := expire

	if expiryParam != "" {
		expiry, err := time.Parse(time.RFC3339, expiryParam)
		if err != nil {
			return ""
		}

		s.Expiry[originalURL] = expiry
	} else {
		expiry := time.Now().Add(s.DefaultExpiry)
		s.Expiry[originalURL] = expiry
	}

	shortURL := generateShortURL()
	s.URLs[shortURL] = originalURL
	s.ClickCount[shortURL] = 0

	return shortURL
}

func generateShortURL() string {
	const letters = "abcdefgqewrrbghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func NewShortenUrlService(defaultExpiry time.Duration) ShortenUrlService {
	return &shortenUrl{
		URLs:          make(map[string]string),
		ClickCount:    make(map[string]int),
		Expiry:        make(map[string]time.Time),
		DefaultExpiry: defaultExpiry,
	}
}
