package social

import (
	"sync"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/fetchers"
)

type Social struct {
	d         *data.Queries
	FunFact   *dtos.FunFact
	funFactMu sync.Mutex
	topics    []string
}

func New(d *data.Queries, f fetchers.Fetchers) *Social {
	s := &Social{
		d:      d,
		topics: getTopicsArr(),
	}
	if fetcher, ok := f["get_fun_fact"]; ok {
		go s.GetFunFactLoop(fetcher)
	}

	return s
}
