package main

import (
	"bufio"
	"context"
	"flag"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"log/slog"
	"time"
)

func main() {
	var (
		timeout     *string
		serviceName *string
	)
	timeout = flag.String("timeout", "30", "timeout")
	serviceName = flag.String("service-name", "chat", "service name")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	defer h.Close()

	if len(*serviceName) == 0 {
		*serviceName = "chat"
	}

	_ctx := context.WithValue(ctx, "discovery_type", "mDNS")
	go DiscoverPeersWithMDNS(_ctx, h, *serviceName)

	protocolID := protocol.ID("/chat/1.0.0")
	h.SetStreamHandler(protocolID, func(s network.Stream) {
		defer s.Close()

		_, err = s.Write([]byte("hello\n"))
		if err != nil {
			panic(err)
		}
		slog.Info("[handler] Write a message", "src_peer_id", h.ID().String(), "dst_peer_id", s.Conn().RemotePeer())

		var r = bufio.NewReader(s)

		got, err := r.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		slog.Info("[handler] Read a message", "src_peer_id", h.ID().String(), "dst_peer_id", s.Conn().RemotePeer(), "message", string(got))

	})

	if len(*timeout) == 0 {
		*timeout = "20"
	}
	d, _ := time.ParseDuration(*timeout + "s")
	timeoutC := time.After(d)

	for {
		select {
		case <-timeoutC:
			return
		case peerID := <-mainChan:
			func() {
				stream, err := h.NewStream(ctx, peerID, protocolID)
				if err != nil {
					panic(err)
				}
				defer stream.Close()
				slog.Info("NewStream", "src_peer_id", h.ID().String(), "dst_peer_id", peerID)

				time.Sleep(time.Second)

				_, err = stream.Write([]byte("hello\n"))
				if err != nil {
					panic(err)
				}
				slog.Info("Write a message", "src_peer_id", h.ID().String(), "dst_peer_id", peerID)

				var r = bufio.NewReader(stream)

				got, err := r.ReadBytes('\n')
				if err != nil {
					panic(err)
				}
				slog.Info("Read a message", "src_peer_id", h.ID().String(), "dst_peer_id", peerID)
				println(string(got))
			}()
		}
	}
}
