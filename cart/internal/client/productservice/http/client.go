package clienthttp

type (
	ProductServiceHttpClient struct {
		addr  string
		token string
	}
)

func NewProductServiceHttpClient(addr string, token string) *ProductServiceHttpClient {
	return &ProductServiceHttpClient{
		addr:  addr,
		token: token,
	}
}
