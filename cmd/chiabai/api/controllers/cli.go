package controllers

import (
	"sync"

	// log "github.com/sirupsen/logrus"
	"gitlab.com/meta-node/meta-node/cmd/chiabai/config"
	controller_client "gitlab.com/meta-node/meta-node/cmd/client/controllers"
	"gitlab.com/meta-node/meta-node/pkg/bls"
	"gitlab.com/meta-node/meta-node/pkg/network"
	"gitlab.com/meta-node/meta-node/pkg/state"
)

type Cli struct {
	// ws     *websocket.Conn
	server *Server
	// Caller CallData
	sync.Mutex
	sendChan                 chan Message1
	keyPairMap               map[string]*bls.KeyPair
	config                   *config.Config
	messageSenderMap         map[string]network.IMessageSender
	connectionsManager       network.IConnectionsManager
	tcpServerMap             map[string]network.ISocketServer
	transactionControllerMap map[string]controller_client.ITransactionController
	accountStateChan         chan state.IAccountState
	
}

// func (cli *Cli) init() CallData {
// 	cli.Caller = CallData{server: cli.server, cli: cli}
// 	// go client.handleMessage()
// 	log.Info("End init client")
// 	return cli.Caller
// }




