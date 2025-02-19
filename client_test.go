package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

type FutureResp struct {
	Data []struct {
		Date      int64  `json:"date"`
		Price     string `json:"price"`
		Amount    string `json:"amount"`
		Tid       int64  `json:"tid"`
		Type      string `json:"type"`
		TradeType string `json:"trade_type"`
	} `json:"data"`
	No      int    `json:"no"`
	Channel string `json:"channel"`
}

type TestReceiver struct{}

func (r *TestReceiver) Unmarshal(frameType FrameType, reader io.Reader) (any, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if bytes[0] != '{' {
		return string(bytes), nil
	}
	resp := new(FutureResp)
	return resp, json.Unmarshal(bytes, resp)
}

func (r *TestReceiver) OnMessage(msg any) {
	fmt.Printf("收到消息: %+v\n", msg)
}

func TestClient(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	client, err := NewClient(ctx, "wss://api.bw6.com/websocket", new(TestReceiver),
		WithHeartbeatInterval(time.Second),
		WithAutoReConnect(),
		WithHeartbeatHandler(func(c *Client) {
			c.SendMessage([]byte("ping"))
		}),
	)
	require.NoError(t, err)

	require.NoError(t, client.SendJson(struct {
		Event   string `json:"event,omitempty"`
		Channel string `json:"channel,omitempty"`
	}{"addChannel", "btcusdt_trades"}))

	//time.AfterFunc(time.Second*10, func() {
	//	cancel()
	//})
	_ = cancel
	select {}
}

func TestMockServer(t *testing.T) {
	http.HandleFunc("/ws", serveWS)
	time.AfterFunc(time.Second, func() {
		fmt.Println("ws served")
	})
	http.ListenAndServe(":9999", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s\n", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}

type ReconnectReceiver struct {
	ReaderReceiver
}

func (ReconnectReceiver) OnMessage(v any) {
	b, err := io.ReadAll(v.(io.Reader))
	if err != nil {
		fmt.Println("receive:", err)
		return
	}

	fmt.Println("rece:", string(b))
}

func TestReconnect(t *testing.T) {
	NewClient(
		context.Background(),
		"ws://127.0.0.1:9999/ws",
		new(ReconnectReceiver),
		WithAutoReConnect(),
		WithHeartbeatInterval(time.Second),
		WithHeartbeatHandler(func(c *Client) {
			c.SendMessage([]byte("ping"))
		}),
		WithConnected(func(c *Client) {
			c.SendMessage([]byte("first connected"))
		}),
		WithReconnected(func(c *Client) {
			c.SendMessage([]byte("after reconnected"))
		}),
	)
	select {}
}
