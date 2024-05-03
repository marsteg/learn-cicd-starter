package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}
	x := make(http.Header)
	x.Add("Authorization", "ApiKey 123")
	y := make(http.Header)
	y.Add("Authorization", "ApiKey lalala")

	tests := []test{
		{input: x, want: "123"},
		{input: y, want: "lalala"},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
