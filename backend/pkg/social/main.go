package social

import (
	"context"
	"sync"
	"time"

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

func (s *Social) GetFunFactLoop(f fetchers.Fetchable) {
	i := 0
	for {
		topic := s.topics[i]
		content, err := f.Fetch(context.Background(), topic)
		if err != nil {
			continue
		}
		s.funFactMu.Lock()
		s.FunFact = &dtos.FunFact{
			Topic:   topic,
			Content: content,
		}
		//
		if i == 280 {
			i = 0
		} else {
			i++
		}
		s.funFactMu.Unlock()
		time.Sleep(time.Hour * 24)
	}
}
