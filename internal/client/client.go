package client

import (
	rpcClient "github.com/agencyenterprise/gossip-host/internal/grpc/client"
	"github.com/agencyenterprise/gossip-host/pkg/logger"
)

func Send(msgLoc, peers string) error {
	msg, err := parseMessageFile(msgLoc)
	if err != nil {
		logger.Errorf("err parsing message file:\n%v", err)
		return err
	}

	peersArr := parsePeers(peers)
	for _, peer := range peersArr {
		if err := rpcClient.Send(peer, msg); err != nil {
			logger.Errorf("err sending message to %s:\n%v", peer, err)
			return err
		}
	}

	return nil
}
