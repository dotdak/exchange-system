package gateway

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_getOpenAPIHandler(t *testing.T) {
	tests := []struct {
		name string
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOpenAPIHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOpenAPIHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	type args struct {
		dialAddr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args.dialAddr); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
