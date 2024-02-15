package ip

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"wafy/internal/response"
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
	client *http.Client
}

func NewIpsumProvider(client *http.Client) *IpsumProvider {
	return &IpsumProvider{
		client: client,
	}
}

func (ip *IpsumProvider) GetAddresses(ctx context.Context) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ipsumBaseAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create response to fetch ipsum addresses: %w", err)
	}

	res, err := ip.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed do request using http client: %w", err)
	}

	if err := response.ValidateStatusCode(res); err != nil {
		return nil, fmt.Errorf("status code validation failed: %w", err)
	}

	addresses, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read ipsum response: %w", err)
	}

	return addressesToStringArray(addresses), nil
}

func addressesToStringArray(addressesBytes []byte) []string {
	addresses := make([]string, 0)
	lines := strings.Split(string(addressesBytes), "\n")

	for _, line := range lines {
		ipAddr, _, _ := strings.Cut(line, "\t")
		if net.ParseIP(ipAddr) == nil {
			continue
		}

		addresses = append(addresses, ipAddr)
	}

	return addresses
}
