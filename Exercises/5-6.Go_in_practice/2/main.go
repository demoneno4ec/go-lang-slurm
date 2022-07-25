package main

const (
	noData       = "" // TODO: код писать здесь
	dataTemplate = "" // TODO: код писать здесь
)

func main() {
	// TODO: код писать здесь
}

type Client struct {
	u string
}

func NewClient(u string) *Client {
	return &Client{u: u}
}

func (c *Client) data() string {
	// TODO: код писать здесь
}
