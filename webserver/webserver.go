package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/cskr/pubsub"
	"github.com/dh1tw/remoteAudio/comms"
	"github.com/dh1tw/remoteAudio/events"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Hub struct {
	muClients    sync.Mutex
	clients      map[*Client]bool
	broadcast    chan []byte
	addClient    chan *Client
	removeClient chan *Client
	events       *pubsub.PubSub
	muAppState   sync.RWMutex
	appState     ApplicationState
}

var hub = Hub{
	broadcast:    make(chan []byte),
	addClient:    make(chan *Client),
	removeClient: make(chan *Client),
	muClients:    sync.Mutex{},
	clients:      make(map[*Client]bool),
	events:       nil,
	muAppState:   sync.RWMutex{},
	appState:     ApplicationState{},
}

type WebServerSettings struct {
	Events *pubsub.PubSub
}

type ApplicationState struct {
	ConnectionStatus bool   `json:"connectionStatus"`
	ServerAudioOn    bool   `json:"serverAudioOn"`
	Ptt              bool   `json:"ptt"`
	TxUser           string `json:"txUser"`
	Ping             int    `json:"ping"`
}

type ClientMessage struct {
	RequestServerAudioOn *bool `json:"serverAudioOn, omitempty"`
	SetPtt               *bool `json:"ptt, omitempty"`
}

func (hub *Hub) sendMsg() {
	hub.muAppState.RLock()
	data, err := json.Marshal(hub.appState)
	if err != nil {
		log.Println(err)
	}
	hub.muAppState.RUnlock()
	hub.muClients.Lock()

	for client := range hub.clients {

		client.send <- data
	}
	hub.muClients.Unlock()
}

func (hub *Hub) handleClientMsg(data []byte) {
	msg := ClientMessage{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println("Webserver: unable to unmarshal ClientMessage")
	}

	if msg.SetPtt != nil {
		hub.muAppState.Lock()
		hub.appState.Ptt = *msg.SetPtt
		hub.muAppState.Unlock()
		hub.events.Pub(*msg.SetPtt, events.RxAudioOn)
		hub.sendMsg()
	}

	if msg.RequestServerAudioOn != nil {
		hub.events.Pub(*msg.RequestServerAudioOn, events.RxAudioOn)
		hub.sendMsg()
	}
}

func (hub *Hub) start() {

	connectionStatusCh := hub.events.Sub(events.MqttConnStatus)
	rxAudioOnCh := hub.events.Sub(events.RxAudioOn)
	txUserCh := hub.events.Sub(events.TxUser)

	for {
		select {
		case ev := <-rxAudioOnCh:
			hub.muAppState.Lock()
			hub.appState.ServerAudioOn = ev.(bool)
			hub.muAppState.Unlock()
			hub.sendMsg()

		case ev := <-txUserCh:
			hub.muAppState.Lock()
			hub.appState.TxUser = ev.(string)
			hub.muAppState.Unlock()
			hub.sendMsg()

		case ev := <-connectionStatusCh:
			cs := ev.(int)
			hub.muAppState.Lock()
			if cs == comms.CONNECTED {
				hub.appState.ConnectionStatus = true
			} else {
				hub.appState.ConnectionStatus = false
			}
			hub.muAppState.Unlock()
			hub.sendMsg()

		case client := <-hub.addClient:
			log.Println("WebSocket connected")
			hub.muClients.Lock()
			hub.clients[client] = true
			hub.muClients.Unlock()
			hub.sendMsg() // should be send only to connecting client

		case client := <-hub.removeClient:
			log.Println("WebSocket disconnected")
			hub.muClients.Lock()
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}
			hub.muClients.Unlock()

		}
	}
}

type Client struct {
	ws   *websocket.Conn
	send chan []byte
}

func (c *Client) write() {
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.ws.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) read() {
	defer func() {
		hub.removeClient <- c
		c.ws.Close()
	}()

	for {
		_, data, err := c.ws.ReadMessage()
		fmt.Println("got msg")
		if err != nil {
			break
		}
		hub.handleClientMsg(data)
	}
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)

	if err != nil {
		http.NotFound(res, req)
		return
	}

	client := &Client{
		ws:   conn,
		send: make(chan []byte),
	}

	hub.addClient <- client

	go client.write()
	go client.read()
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/index.html")
}

func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func Webserver(s WebServerSettings) {

	hub.events = s.Events

	go hub.start()

	http.Handle("/static/",
		noDirListing(http.FileServer(http.Dir("html/"))))

	http.HandleFunc("/ws", wsPage)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
