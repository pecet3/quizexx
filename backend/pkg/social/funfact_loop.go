package social

import (
	"context"
	"time"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/fetchers"
	"github.com/pecet3/quizex/pkg/logger"
)

func (s *Social) GetFunFactLoop(f fetchers.Fetchable) {
	i := 1
	for {
		// dev, delete in prod
		time.Sleep(time.Hour * 24)
		topic := s.topics[i]
		currTopic, _ := s.d.GetCurrentFunFact(context.Background())
		if topic == currTopic.Topic {
			i++
			continue
		}
		str, err := f.Fetch(context.Background(), topic)
		content := str.(string)
		if err != nil {
			continue
		}
		logger.Debug(content)
		_, err = s.d.AddFunFact(context.Background(), data.AddFunFactParams{
			Topic:   topic,
			Content: content,
		})
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
