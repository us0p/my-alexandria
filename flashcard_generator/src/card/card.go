package card

type Card struct {
	Question string
	Answer   string
	ID       string
	Tags     []string
}

func (c *Card) AddTag(tags ...string) {
	for _, tag := range tags {
		if tag == "" {
			continue
		}
		c.Tags = append(c.Tags, tag)
	}
}
