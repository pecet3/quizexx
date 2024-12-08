package magic_link

import (
	"log"
	"time"
)

type tmpSection struct {
	Id                int
	LastNewSession    time.Time
	NewSessionCounter int
}

type tmpSessions = map[int]*tmpSection

func (as *MagicLink) newTmpSection(uId int) *tmpSection {
	return &tmpSection{
		LastNewSession:    time.Now(),
		NewSessionCounter: 1,
		Id:                uId,
	}
}

func (as *MagicLink) getTmpSection(id int) (*tmpSection, bool) {
	as.tMu.Lock()
	defer as.tMu.Unlock()
	su, exists := as.tmpSessions[id]
	if !exists {
		return nil, false
	}
	return su, true
}

func (as *MagicLink) addTmpSection(su *tmpSection) error {
	as.tMu.Lock()
	defer as.tMu.Unlock()
	as.tmpSessions[su.Id] = su
	log.Println(as.tmpSessions)
	return nil
}

func (as *MagicLink) removeTmpSection(uId int) {
	as.tMu.Lock()
	defer as.tMu.Unlock()
	delete(as.tmpSessions, uId)
}

func (as *MagicLink) updateTmpSection(su *tmpSection) {
	as.removeTmpSection(su.Id)
	err := as.addTmpSection(su)
	if err != nil {
		log.Println(err)
	}
	log.Println("Updated su", su)
}
