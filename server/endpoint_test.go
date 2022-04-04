package server

import (
	"net/url"
	"reflect"
	"testing"
)

func TestEndPoint(t *testing.T) {
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "grpc://127.0.0.1?isSecure=false",
			args: args{NewEndpoint("grpc", "127.0.0.1", false)},
			want: false,
		},
		{
			name: "grpc://127.0.0.1?isSecure=true",
			args: args{NewEndpoint("http", "127.0.0.1", true)},
			want: true,
		},
		{
			name: "grpc://127.0.0.1",
			args: args{NewEndpoint("grpc", "localhost", false)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndpointIsSecure(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEndpoint(t *testing.T) {
	type args struct {
		scheme   string
		host     string
		isSecure bool
	}
	tests := []struct {
		name string
		args args
		want *url.URL
	}{
		{
			name: "https://github.com/raylin666/go-utils/",
			args: args{"https", "github.com/raylin666/go-utils/", false},
			want: &url.URL{Scheme: "https", Host: "github.com/raylin666/go-utils/"},
		},
		{
			name: "http://www.ls331.com/",
			args: args{"http", "www.ls331.com/", true},
			want: &url.URL{Scheme: "http", Host: "www.ls331.com/", RawQuery: "isSecure=true"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEndpoint(tt.args.scheme, tt.args.host, tt.args.isSecure); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseEndpoint(t *testing.T) {
	type args struct {
		endpoints []string
		scheme    string
		isSecure  bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "raylin666",
			args:    args{endpoints: []string{"https://github.com/raylin666/go-utils?isSecure=true"}, scheme: "https", isSecure: true},
			want:    "github.com",
			wantErr: false,
		},
		{
			name:    "test",
			args:    args{endpoints: []string{"http://www.ls331.com/"}, scheme: "http", isSecure: true},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseEndpoint(tt.args.endpoints, tt.args.scheme, tt.args.isSecure)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseEndpoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}
