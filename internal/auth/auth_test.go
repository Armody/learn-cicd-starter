package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	h1 := http.Header{}
	http.Header.Add(h1, "Authorization", "ApiKey 123456")
	h2 := http.Header{}
	h3 := http.Header{}
	http.Header.Add(h3, "Authorization", "123456")

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "regular",
			args: args{
				headers: h1,
			},
			want:    "12345",
			wantErr: false,
		},
		{
			name: "no-auth",
			args: args{
				headers: h2,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed auth",
			args: args{
				headers: h3,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
