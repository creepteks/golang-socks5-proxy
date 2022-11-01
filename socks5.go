package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"golang.org/x/net/proxy"
)

func main() {
	proxyUrl := "127.0.0.1:1080"
	dialer, err := proxy.SOCKS5("tcp", proxyUrl, nil, proxy.Direct)
	dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
		return dialer.Dial(network, address)
	}
	transport := &http.Transport{DialContext: dialContext,
		DisableKeepAlives: true}
	cl := &http.Client{Transport: transport}

	resp, err := cl.Get("https://wtfismyip.com/json")
	if err != nil {
		// TODO handle me
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	// TODO work with the response
	if err != nil {
		fmt.Println("body read failed")
	}
	fmt.Println(string(body))
}
