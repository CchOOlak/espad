package urlrepo

import (
	"encoding/json"
	"espad/internal/core/domain"
	"espad/pkg/appErrors"

	"github.com/rs/zerolog/log"
)

type memstorage struct {
	// TODO: use Database (like MongoDB or Redis) for storage
	storage map[string][]byte
}

func NewMemstorage() *memstorage {
	return &memstorage{
		storage: map[string][]byte{},
	}
}

func (repo *memstorage) Get(key string) (domain.Url, error) {
	if value, ok := repo.storage[key]; ok {
		url := domain.Url{}
		err := json.Unmarshal(value, &url)
		if err != nil {
			log.Error().Msgf("fail to get value from memstorage: %v", err)
			return domain.Url{}, appErrors.Internal
		}

		return url, nil
	}

	return domain.Url{}, appErrors.NotFound
}

func (repo *memstorage) Save(url domain.Url) error {
	bytes, err := json.Marshal(url)
	if err != nil {
		log.Error().Msgf("url fails at marshal into json string: %v", err)
		return appErrors.Internal
	}

	repo.storage[url.Shorten] = bytes

	return nil
}
