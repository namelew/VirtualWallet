package clients

type Client struct {
	ID     uint64
	Name   string
	Amount float64
}

func (c *Client) Transfer(target uint64, value float64) {
}
