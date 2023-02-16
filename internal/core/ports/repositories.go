package ports

import "espad/internal/core/domain"

type UrlRepository interface {
	Save(domain.Url) error
	Get(shortenUrl string) (domain.Url, error)
}
