package peer

import (
	"fmt"

	"github.com/aayanl/DEXtrading/logging"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
)

const messageId = 0
type Message string


func StartPeer() {
	nodekey, _ := crypto.GenerateKey()

	config := p2p.Config {
		MaxPeers:   2,                            //One connection to peer for trade, another for the server hub.
		PrivateKey: nodekey,
		Name:       "DEX node",
		ListenAddr: ":20200",
		Protocols: []p2p.Protocol {
			protoSendChall(),
			protoConnectHub(),
		},
	}

	srv := p2p.Server {
        	Config:     config,
    	}

	if err := srv.Start(); err != nil {
		fmt.Println(err);
	} else {
		logging.Write("Node started", true);
		fmt.Println(srv.Self());
	}

	select {}
};

func protoSendChall() p2p.Protocol {
	return p2p.Protocol{
		Name:    "SendChallenge",
		Version: 1,
		Length:  1,
		Run:     SendChallenge,
	}
};

func protoConnectHub() p2p.Protocol {
	return p2p.Protocol{
		Name:    "ConnectHub",
		Version: 1,
		Length:  1,
		Run:     ConnectHub,
	}
}
