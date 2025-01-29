package fetchers

import "context"

type Fetchable interface {
	Fetch(ctx context.Context, i interface{}) (interface{}, error)
}

type Fetchers map[string]Fetchable

func New() Fetchers {
	f := make(Fetchers)
	f["test_game_content"] = Test{}
	f["gpt4omini_game_content"] = FetcherGPT4ominiGameContent{}
	f["get_fun_fact"] = FetcherFunFact{}
	return f
}
