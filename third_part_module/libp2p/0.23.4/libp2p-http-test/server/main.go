package main

import (
	"encoding/hex"
	"fmt"
	"github.com/libp2p/go-libp2p"
	gostream "github.com/libp2p/go-libp2p-gostream"
	p2phttp "github.com/libp2p/go-libp2p-http"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	myProtocol protocol.ID = "/testiti-test"
)

func main() {
	configDir := "./config"
	port := "0"

	var cryptoFileDir string
	if len(os.Args) > 1 {
		cryptoFileDir = configDir + "/client"
	} else {
		cryptoFileDir = configDir + "/server"
		port = "51080"
	}

	if _, err := os.Stat(cryptoFileDir); os.IsNotExist(err) {
		if err = os.MkdirAll(cryptoFileDir, 0755); err != nil {
			panic(err)
		}
	}
	var privKey crypto.PrivKey
	keyFile := cryptoFileDir + "/private.key"
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		privKey, _, err = crypto.GenerateKeyPair(crypto.RSA, 2048)
		if err != nil {
			panic(err)
		}
		func() {
			f, err := os.OpenFile(keyFile, os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			keyData, err := privKey.Raw()
			if err != nil {
				panic(err)
			}

			if _, err := hex.NewEncoder(f).Write(keyData); err != nil {
				panic(err)
			}
		}()
	} else {
		func() {
			f, err := os.Open(keyFile)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			b, err := io.ReadAll(f)
			if err != nil {
				panic(err)
			}
			var keyResult = make([]byte, len(b))
			n, err := hex.Decode(keyResult, b)
			privKey, err = crypto.UnmarshalRsaPrivateKey(keyResult[:n])
			if err != nil {
				panic(err)
			}
		}()

	}

	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/"+port),
		libp2p.Identity(privKey),
	)
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 1 {
		httpServer(host)
	} else {
		httpClient(host, os.Args[1])
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("System shutting down ...")
	if err := host.Close(); err != nil {
		panic(err)
	}
}

func httpServer(host host.Host) {
	go func() {
		listener, err := gostream.Listen(host, myProtocol)
		if err != nil {
			panic(err)
		}
		defer listener.Close()

		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hi!"))
		})
		server := &http.Server{}
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()
	peerInfo := peer.AddrInfo{
		host.ID(), host.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("httpServer address:", addrs[0])
}

func httpClient(host host.Host, addr string) {
	addrInfo, err := peer.AddrInfoFromString(addr)
	if err != nil {
		panic(err)
	}
	host.Peerstore().AddAddrs(addrInfo.ID, addrInfo.Addrs, peerstore.PermanentAddrTTL)

	tr := &http.Transport{}
	tr.RegisterProtocol("libp2p", p2phttp.NewTransport(host, p2phttp.ProtocolOption(myProtocol)))
	client := &http.Client{Transport: tr}

	requestUrl := fmt.Sprintf("libp2p://%s/hello", addrInfo.ID.String())
	res, err := client.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %#v", res)
}
