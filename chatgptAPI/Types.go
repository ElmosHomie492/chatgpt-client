package chatgptAPI

import "github.com/go-resty/resty/v2"

type OpenAIClient struct {
	Client *resty.Client
}
