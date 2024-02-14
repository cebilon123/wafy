package ip

import (
	"context"
	"fmt"
	"net/http"
)

const (
	ipsumBaseAddress = "https://raw.githubusercontent.com/stamparm/ipsum/master/ipsum.txt"
)

// Provider should be implemented by the ip provider
// that provides slice of the ip addresses.
type Provider interface {
	GetAddresses(ctx context.Context) ([]string, error)
}

type IpsumProvider struct {
}

func NewIpsumProvider() *IpsumProvider {
	return &IpsumProvider{}
}

func (i *IpsumProvider) GetAddresses(ctx context.Context) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ipsumBaseAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request to fetch ipsum addresses: %w", err)
	}

	//TODO implement me
	panic("implement me")
}
