package clienthttp

type (
	LomsHttpClient struct {
		addr string
	}
)

func NewLomsHttpClient(addr string) *LomsHttpClient {
	return &LomsHttpClient{
		addr: addr,
	}
}
