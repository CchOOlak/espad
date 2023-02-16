package ports

import "espad/internal/core/domain"

type UrlService interface {
	Create(originalUrl string, username string) (domain.Url, error)
	Get(shortenUrl string) (string, error)
}
