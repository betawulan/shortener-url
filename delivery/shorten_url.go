package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/shortener-url/service"
)

type shortenDelivery struct {
	service service.ShortenUrlService
}

func (s shortenDelivery) Shorten(c echo.Context) error {
	originalURL := c.QueryParam("url")
	expiryParam := c.QueryParam("expiry")

	shortURL := s.service.Generate(originalURL, expiryParam)

	return c.JSON(http.StatusOK, map[string]string{"shortenedURL": shortURL})
}

func (s shortenDelivery) Redirect(c echo.Context) error {
	shortURL := c.Param("shortURL")

	originalURL, err := s.service.Redirect(shortURL)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.Redirect(http.StatusSeeOther, originalURL)
}

func (s *shortenDelivery) GetClickCount(c echo.Context) error {
	shortURL := c.Param("shortURL")
	clickCount := s.service.GetClickCount(shortURL)

	return c.JSON(http.StatusOK, map[string]int{"clickCount": clickCount})
}

func (s *shortenDelivery) Sort(c echo.Context) error {
	sortType := c.QueryParam("sortType")

	sortedURLList, err := s.service.Sort(sortType)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, sortedURLList)
}

func NewShortenUrlDelivery(e *echo.Echo, urlService service.ShortenUrlService) {
	delivery := shortenDelivery{service: urlService}

	e.GET("/shorten", delivery.Shorten)
	e.GET("/:shortURL", delivery.Redirect)
	e.GET("/:shortURL/clicks", delivery.GetClickCount)
	e.GET("/sort", delivery.Sort)
}
