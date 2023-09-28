package Functions

import (
	"net/http"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CustomTransport struct {
	Base    http.RoundTripper
	Headers map[string]string
}

func Address(s string) *string {
	caser := cases.Title(language.Und)
	upper := caser.String(s)
	return &upper
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add custom headers to the request.
	for key, value := range t.Headers {
		req.Header.Add(key, value)
	}

	// Use the underlying transport to perform the actual HTTP request.
	return t.Base.RoundTrip(req)
}

func CustomHTTPClient(token string) *http.Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &CustomTransport{
			Base:    http.DefaultTransport,
			Headers: map[string]string{"Authorization": "Bearer " + token},
		},
	}

	return client
}
