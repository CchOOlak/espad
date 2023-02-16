package urlhdl

import (
	"espad/internal/core/domain"
	"fmt"
)

type BodyCreate struct {
	OriginalUrl string `json:"url"`
	Username    string `json:"username"`
}

// TODO: use environment variables instead
const (
	PROTOCOL = "https"
	HOSTNAME = "shUrl.co"
)

func MakeResponseUrl(url domain.Url) domain.Url {
	return domain.Url{
		Original: url.Original,
		Shorten:  MakeFullUrl(url.Shorten),
		Username: url.Username,
	}
}

func MakeFullUrl(path string) string {
	return fmt.Sprintf("%s://%s/%s", PROTOCOL, HOSTNAME, path)
}
