package main

import (
	"context"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"log/slog"
)

var ConnectedHosts = make(map[peer.ID]bool)
var mainChan = make(chan peer.ID)

func DiscoverPeersWithMDNS(ctx context.Context, h host.Host, serviceName string) {
	slog.Info("Starting discovery with mDNS", "peer_id:", h.ID().String(), "ns:", serviceName)
	peerChan := initMDNS(h, serviceName)
	for {
		select {
		case <-ctx.Done():
			return
		case p := <-peerChan:
			if p.ID == h.ID() {
				continue
			}
			if tryConnect(ctx, h, p) {
				mainChan <- p.ID
			}
		}
	}
}

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

func (n *discoveryNotifee) HandlePeerFound(p peer.AddrInfo) {
	n.PeerChan <- p
}

// Initialize the MDNS service
func initMDNS(peerhost host.Host, serviceName string) chan peer.AddrInfo {
	// register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	ser := mdns.NewMdnsService(peerhost, serviceName, n)
	if err := ser.Start(); err != nil {
		panic(err)
	}
	return n.PeerChan
}

func tryConnect(ctx context.Context, h host.Host, p peer.AddrInfo) bool {
	if h.Network().Connectedness(p.ID) == network.Connected {
		slog.Info("Connected", "src_peer_id", h.ID().String(), "dst_peer_id", p.ID.String(), "discovery_type", ctx.Value("discovery_type"))
		if _, ok := ConnectedHosts[p.ID]; !ok {
			ConnectedHosts[p.ID] = true
		}
		return true
	}
	err := h.Connect(ctx, p)
	if err != nil {
		slog.Error("Connect to node failed.", "src_peer_id", h.ID().String(), "dst_peer_id", p.ID.String(), "error", err, "discovery_type", ctx.Value("discovery_type"))
		return false
	}
	ConnectedHosts[p.ID] = true
	slog.Info("Connected", "src_peer_id", h.ID().String(), "dst_peer_id", p.ID.String(), "discovery_type", ctx.Value("discovery_type"))
	return true
}
