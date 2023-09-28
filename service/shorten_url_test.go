package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/betawulan/shortener-url/service"
)

func Test_ShortenURL_Generate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		oriURL := "https://www.youtube.com/watch?v=53_Nt9f-Wrk&ab_channel=McDonaldsMalaysia"

		srv := service.NewShortenUrlService(1 * time.Hour)
		shortURL := srv.Generate(oriURL, "")
		originalURL, err := srv.Redirect(shortURL)

		require.Equal(t, originalURL, oriURL)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		oriURL := "https://www.youtube.com/watch?v=53_Nt9f-Wrk&ab_channel=McDonaldsMalaysia"

		srv := service.NewShortenUrlService(1 * time.Hour)
		shortURL := srv.Generate(oriURL, "")
		originalURL, err := srv.Redirect("random")

		require.Empty(t, originalURL, shortURL)
		require.Error(t, err)
	})

	t.Run("success with expiry date", func(t *testing.T) {
		oriURL := "https://www.youtube.com/watch?v=53_Nt9f-Wrk&ab_channel=McDonaldsMalaysia&expiry=2023-11-11T00:00:00Z"

		srv := service.NewShortenUrlService(1 * time.Hour)
		shortUrl := srv.Generate(oriURL, "2023-11-11T00:00:00Z")
		originalURL, err := srv.Redirect(shortUrl)
		require.NoError(t, err)
		require.Equal(t, originalURL, oriURL)
	})
}

func Test_ShortenUrl_GetClickCount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		oriURL := "https://www.youtube.com/watch?v=53_Nt9f-Wrk&ab_channel=McDonaldsMalaysia&expiry=2023-11-11T00:00:00Z"

		srv := service.NewShortenUrlService(1 * time.Hour)
		shortUrl := srv.Generate(oriURL, "2023-11-11T00:00:00Z")
		_, _ = srv.Redirect(shortUrl)
		_, _ = srv.Redirect(shortUrl)
		_, _ = srv.Redirect(shortUrl)
		click := srv.GetClickCount(shortUrl)

		require.Equal(t, 3, click)
	})
}

func Test_ShortenUrl_Redirect(t *testing.T) {
	t.Run("url was expired", func(t *testing.T) {
		oriURL := "https://www.youtube.com/watch?v=53_Nt9f-Wrk&ab_channel=McDonaldsMalaysia"

		srv := service.NewShortenUrlService(1 * time.Nanosecond)
		shortUrl := srv.Generate(oriURL, "2023-09-09T00:00:00Z")
		_, err := srv.Redirect(shortUrl)

		require.Error(t, err)
		require.Equal(t, errors.New("url was expired"), err)
	})

}

func Test_ShortenUrl_Sort(t *testing.T) {
	t.Run("sort by asc", func(t *testing.T) {
		firstOriURL := "https://www.youtube.com/watch?v=S71uY8wPUyM&ab_channel=America%27sGotTalent"
		secondOriURL := "https://www.youtube.com/watch?v=53_Nt9f-Wrk&ab_channel=McDonaldsMalaysia"

		srv := service.NewShortenUrlService(1. * time.Hour)

		firstShortUrl := srv.Generate(firstOriURL, "")
		secondShortUrl := srv.Generate(secondOriURL, "")

		_, err := srv.Redirect(firstShortUrl)
		_, err = srv.Redirect(firstShortUrl)
		_, err = srv.Redirect(secondShortUrl)
		require.NoError(t, err)

		urlsSort, err := srv.Sort("asc")

		require.Equal(t, urlsSort, []string{secondShortUrl, firstShortUrl})
	})
}
