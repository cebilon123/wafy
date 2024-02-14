package ip

import (
	"fmt"
	"io"
	"sync"

	"github.com/twmb/murmur3"
)

// Checker should be implemented by all the structs
// that can check if given ip address is malicious/suspicious.
type Checker interface {
	// CheckIPAddress checks given ip address for being malicious. Returns true
	// if ip address is suspicious with the message.
	CheckIPAddress(ipAddr string) (bool, string)
}

// Murmur3SyncSafeChecker is a checker that is safe for
// the concurrent access.
type Murmur3SyncSafeChecker struct {
	// ipHashMap is and hash map containing all the ip addresses
	ipHashMap *sync.Map
	// seed is the seed used to hash entries in the ipHashMap
	seed      uint32
	logWriter io.Writer
}

func NewSyncSafeChecker(ipHashMap *sync.Map, seed uint32, logWriter io.Writer) *Murmur3SyncSafeChecker {
	return &Murmur3SyncSafeChecker{
		ipHashMap: ipHashMap,
		seed:      seed,
		logWriter: logWriter,
	}
}

func (s *Murmur3SyncSafeChecker) CheckIPAddress(ipAddr string) (bool, string) {
	hash := murmur3.SeedNew32(s.seed)
	_, err := hash.Write([]byte(ipAddr))
	if err != nil {
		_, _ = s.logWriter.Write([]byte(fmt.Sprintf("error while writing hash: %s", err.Error())))
	}
	sum := hash.Sum32()

	_, ok := s.ipHashMap.Load(sum)
	if ok {
		return ok, fmt.Sprintf("ip address: '%s' found in the malicious database", ipAddr)
	}

	return false, ""
}
