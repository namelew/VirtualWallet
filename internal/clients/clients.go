package clients

type Client struct {
	ID     uint64
	Amount float64
}

func (c *Client) Transfer(target uint64, value float64) {
}
