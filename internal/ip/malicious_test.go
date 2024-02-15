package ip

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestIpsumProvider_GetAddresses(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				client: http.DefaultClient,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := &IpsumProvider{
				client: tt.fields.client,
			}
			got, err := ip.GetAddresses(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddresses() got = %v, want %v", got, tt.want)
			}
		})
	}
}
