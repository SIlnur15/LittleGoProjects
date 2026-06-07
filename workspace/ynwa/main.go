package main

import "time"

type Client struct {
	Transport     RoundTripper
	CheckRedirect func(req *Request, via []*Request) error
	Jar           CookieJar
	Timeout       time.Duration
}
