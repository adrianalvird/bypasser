package utils

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

// ConfigureProxyClient sets up an HTTP client to use a proxy.
func ConfigureProxyClient(proxyAddress string) (*http.Client, error) {
	proxyURL, err := url.Parse(proxyAddress)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	return client, nil
}
