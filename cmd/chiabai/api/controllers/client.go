package controllers

import (
	"sync"

	"gitlab.com/meta-node/meta-node/cmd/chiabai/config"
	log "github.com/sirupsen/logrus"
	controller_client "gitlab.com/meta-node/meta-node/cmd/client/controllers"
	"gitlab.com/meta-node/meta-node/pkg/bls"
	"gitlab.com/meta-node/meta-node/pkg/network"
	"gitlab.com/meta-node/meta-node/pkg/state"
	"github.com/gorilla/websocket"
)

type Client struct {
	ws     *websocket.Conn
	server *Server
	Caller CallData
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

func (client *Client) init() CallData {
	client.Caller = CallData{server: client.server, client: client}
	go client.handleMessage()
	log.Info("End init client")
	return client.Caller
}
func (client *Client) handleListen() {
	for {
		// Read in a new message as JSON and map it to a Message object
		var msg map[string]interface{}
		err := client.ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.ws.Close()
			break
		}
		// log.Info("Message from client: ", msg)
		client.handleCallChain(msg)
	}
}


// handle message struct tu chain tra ve va chuyen qua dang JSON gui toi cac client
func (client *Client) handleMessage() {
	for {
		msg := <-client.sendChan
		// msg1 := <-sendDataC
		log.Info(msg)
		err := client.ws.WriteJSON(msg)

		if err != nil {
			log.Printf("error: %v", err)
			client.ws.Close()
		}
	}
}
func (client *Client) handleCallChain(msg map[string]interface{}) {
	// handle call
	switch msg["command"] {
	case "connect-wallet":
		call:=msg["value"].(map[string]interface{}) 
		kq:=client.Caller.ConnectWallet(call)
		client.Caller.sentToClient("connect-wallet",kq)
	case "generate-keys":
		call:=msg["value"].(map[string]interface{}) 
		kq:=client.Caller.GeneratePlayerKeys(call)
		client.Caller.sentToClient("generate-keys",kq)
	// case "create-deck":
	// 	kq:=client.Caller.CreateDeck()
	// 	client.Caller.sentToClient("create-deck",kq)
	case "shuffle-and-encrypt-cards":
		call:=client.Caller.CreateDeck()
		deck:=client.Caller.ShuffleDeck(call)		
		call2:=msg["value"].(map[string]interface{}) 
		kq:=client.Caller.EncryptDeck(deck,call2)
		client.Caller.sentToClient("encrypt-cards",kq)
	case "decrypt-cards":
		call:=msg["value"].(map[string]interface{}) 
		kq:=client.Caller.DecryptDeck(call)
		client.Caller.sentToClient("decrypt-cards",kq)
	case "send-transaction":
		call:=msg["value"].(map[string]interface{}) 
		kq:=client.Caller.TryCall(call)
		client.Caller.sentToClient("send-transaction",kq)
	default:
		log.Warn("Require call not match: ", msg)
	}
}