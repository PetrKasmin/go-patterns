package pkg

type Creator struct {
	State string
}

func (c *Creator) Create() *Snapshot {
	return &Snapshot{
		state: c.State,
	}
}

func (c *Creator) Restore(s *Snapshot) {
	c.State = s.GetSavedState()
}

func (c *Creator) SetState(s string) {
	c.State = s
}

func (c *Creator) GetState() string {
	return c.State
}
