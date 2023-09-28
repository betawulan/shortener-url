package service_test

import (
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

	t.Run("failed", func(t *testing.T) {
		oriURL := "https://www.youtube.com/watch?v=53_Nt9f-Wrk&ab_channel=McDonaldsMalaysia"

		srv := service.NewShortenUrlService(1 * time.Hour)
		shortURL := srv.Generate(oriURL, "")
		originalURL, err := srv.Redirect("random")

		require.Empty(t, originalURL, shortURL)
		require.Error(t, err)
	})
}
