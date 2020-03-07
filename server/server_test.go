package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handler(t *testing.T) {
	type args struct {
		response http.ResponseWriter
		request  *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Handle with Success",
			args: args{
				response: httptest.NewRecorder(),
				request:  httptest.NewRequest("GET", "/service/ready", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.response, tt.args.request)
		})
	}
}

func Test_readyHandler(t *testing.T) {
	type args struct {
		response http.ResponseWriter
		request  *http.Request
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "Handle with Success",
			args: args{
				response: httptest.NewRecorder(),
				request:  httptest.NewRequest("GET", "/service/ready", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readyHandler(tt.args.response, tt.args.request)
		})
	}
}

func TestStartServer(t *testing.T) {
	type args struct {
		config Configuration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name: "Start Server with RootPath not found",
			args: args{
				config: Configuration{
					DefaultResource: "/index.html",
					Hostname:        "127.0.0.1",
					Port:            "8080",
					RootPath:        "./tests/app",
				},
			},
			wantErr: true,
		},
		{
			name: "Start Server with default resource not found",
			args: args{
				config: Configuration{
					DefaultResource: "/fake.html",
					Hostname:        "127.0.0.1",
					Port:            "8080",
					RootPath:        "../internal/test-app",
				},
			},
			wantErr: true,
		},
		{
			name: "Start Server with forbiden port",
			args: args{
				config: Configuration{
					DefaultResource: "/index.html",
					Hostname:        "127.0.0.1",
					Port:            "80",
					RootPath:        "../internal/test-app",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartServer(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("StartServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
