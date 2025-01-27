package quiz

type Social struct {
	m          *Manager
	topicIndex int
}

func NewSocial(m *Manager) *Social {
	return &Social{
		m:          m,
		topicIndex: 0,
	}
}
