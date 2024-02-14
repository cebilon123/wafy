package threat

import (
	"net/http"
	"wafy/internal/ip"
)

type Result struct {
	IsThreat bool
	Message  string
}

type Threat interface {
	Check(r *http.Request) (*Result, error)
}

type IPThreat struct {
	checker ip.Checker
}

func NewIPThreat(checker ip.Checker) *IPThreat {
	return &IPThreat{checker: checker}
}

func (it *IPThreat) Check(r *http.Request) (*Result, error) {
	it.checker.CheckIPAddress(r.RemoteAddr)
}
