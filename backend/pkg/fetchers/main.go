package fetchers

import "context"

type Fetchable interface {
	Fetch(ctx context.Context, i interface{}) (string, error)
}

type Fetchers map[string]Fetchable

func New() Fetchers {
	f := make(Fetchers)
	f["test_game_content"] = Test{}
	f["gpt4omini_game_content"] = FetcherGPT4ominiGameContent{}
	return f
}
