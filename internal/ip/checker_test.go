package ip

import (
	"io"
	"os"
	"sync"
	"testing"

	"github.com/twmb/murmur3"
)

func TestMurmur3SyncSafeChecker_CheckIPAddress(t *testing.T) {
	addr := "127.0.0.1"
	seed := uint32(1234)

	hash := murmur3.SeedNew32(seed)
	hash.Write([]byte(addr))

	ipHashMap := &sync.Map{}
	ipHashMap.Store(hash.Sum32(), addr)

	type fields struct {
		ipHashMap *sync.Map
		seed      uint32
		logWriter io.Writer
	}
	type args struct {
		ipAddr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "address exists, error returned",
			fields: fields{
				ipHashMap: ipHashMap,
				seed:      seed,
				logWriter: os.Stdin,
			},
			args: args{
				ipAddr: addr,
			},
			wantErr: true,
		},
		{
			name: "address not found, nil returned",
			fields: fields{
				ipHashMap: ipHashMap,
				seed:      seed,
				logWriter: os.Stdin,
			},
			args: args{
				ipAddr: "192.168.100.100",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Murmur3SyncSafeChecker{
				ipHashMap: tt.fields.ipHashMap,
				seed:      tt.fields.seed,
				logWriter: tt.fields.logWriter,
			}
			if err := s.CheckIPAddress(tt.args.ipAddr); (err != nil) != tt.wantErr {
				t.Errorf("CheckIPAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
