package urlsrv

import (
	"errors"
	"espad/internal/core/domain"
	"espad/internal/core/ports"
	"espad/pkg/appErrors"
	"espad/pkg/hash"
	"net/url"
)

type service struct {
	urlRepository ports.UrlRepository
	hashGenerator hash.HashGenerator
}

func New(urlRepository ports.UrlRepository, hashGenerator hash.HashGenerator) *service {
	return &service{
		urlRepository: urlRepository,
		hashGenerator: hashGenerator,
	}
}

func (srv *service) Create(originalUrl string, username string) (domain.Url, error) {
	// verify originalURL
	if !isUrl(originalUrl) {
		return domain.Url{}, appErrors.InvalidInput
	}

	shortenUrl := srv.hashGenerator.GetHash(originalUrl + username)
	// check if exist
	if u, err := srv.urlRepository.Get(shortenUrl); err == nil {
		return u, nil
	}

	u := domain.NewUrl(originalUrl, shortenUrl, username)
	if err := srv.urlRepository.Save(u); err != nil {
		return domain.Url{}, appErrors.Internal
	}
	return u, nil
}

func (srv *service) Get(shortenUrl string) (domain.Url, error) {
	u, err := srv.urlRepository.Get(shortenUrl)
	if err != nil {
		if errors.Is(err, appErrors.NotFound) {
			return domain.Url{}, appErrors.NotFound
		}
		return domain.Url{}, appErrors.Internal
	}
	return u, nil
}

func isUrl(input string) bool {
	if _, err := url.ParseRequestURI(input); err != nil {
		return false
	}
	return true
}
