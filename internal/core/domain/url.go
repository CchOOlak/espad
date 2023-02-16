package domain

type Url struct {
	Original string `json:"original"`
	Shorten  string `json:"shorten"`
	Username string `json:"username"`
}

func NewUrl(originalUrl string, shortenUrl string, username string) Url {
	return Url{
		Original: originalUrl,
		Shorten:  shortenUrl,
		Username: username,
	}
}
