func (r *Room) addClient(c *Client) {
	r.cMu.Lock()
	defer r.cMu.Unlock()
	if p, ok := r.game.Players[c.user.Uuid]; ok {
		c.player = p
	} else {
		c.player.user = c.user
		if !r.game.IsGame {
			r.game.Players[c.user.Uuid] = &Player{
				user: c.user,
			}
		}
	}
	r.clients[c.user.Uuid] = c
}
func (r *Room) removeClient(c *Client) {
	r.cMu.Lock()
	defer r.cMu.Unlock()
	if _, ok := r.game.Players[c.user.Uuid]; ok {
		if !r.game.IsGame {
			delete(r.game.Players, c.user.Uuid)
		}
	}
}