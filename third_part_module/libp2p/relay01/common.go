package main

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
)

func NewHost() host.Host {
	options := make([]libp2p.Option, 0, 10)

	if h, err := libp2p.New(options...); err == nil {
		return h
	} else {
		panic(err)
	}
	return nil
}

func NewRelay() (host.Host, *relay.Relay, error) {
	h, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	r, err := relay.New(h)
	if err != nil {
		return nil, nil, err
	}
	return h, r, err
}

func NewNode() (host.Host, error) {
	h, err := libp2p.New(
		libp2p.NoListenAddrs,
		libp2p.EnableRelay(),
	)
	return h, err
}
