package service

type ShortenUrlService interface {
	Generate(url, expire string) string
	Redirect(shortenUrl string) (string, error)
}