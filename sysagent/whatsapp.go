package sysagent

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type WhatsApp struct {
	client *whatsmeow.Client
}

// NewWhatsApp returns an initalized WhatsApp object.
func NewWhatsApp() *WhatsApp {
	return &WhatsApp{
		client: nil,
	}
}

// Login logs on to the whatsapp service.
// If login failed channel is nil with error set.
// If login success channel is nil and error is nil
// If login needs QR auth, channel is set for QRChannelItem and error is nil
func (w *WhatsApp) Login(debug bool) (<-chan whatsmeow.QRChannelItem, error) {
	var dbLog, clientLog waLog.Logger
	if debug {
		dbLog = waLog.Stdout("Database", "DEBUG", true)
		clientLog = waLog.Stdout("Client", "DEBUG", true)
	}

	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		return nil, err
	}

	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid)
	// or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return nil, err
	}
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(w.eventHandler)
	w.client = client

	// Already authenticated & logged in, just connect.
	if client.Store.ID != nil {
		if err := client.Connect(); err != nil {
			return nil, err
		}
		return nil, nil // Login completed.
	}

	// Not logged in or authenticated. Need to get a QR code and authenticate.
	qrChan, err := client.GetQRChannel(context.Background())
	if err != nil {
		return nil, err
	}
	if err = client.Connect(); err != nil {
		return nil, err
	}

	// See https://pkg.go.dev/go.mau.fi/whatsmeow@v0.0.0-20241121132808-ae900cb6bee4#QRChannelItem
	// on how to intepret channel results.
	return qrChan, nil
}

// SendMessage sends a message to the destination number (eg 16025550044,
// 91999935090)
func (w *WhatsApp) SendMessage(dest, msg string) error {
	if w.client == nil || !w.client.IsLoggedIn() {
		return fmt.Errorf("cannot send message without logging in first")
	}
	jid, err := types.ParseJID(dest + "@s.whatsapp.net")
	if err != nil {
		return err
	}
	if _, err = w.client.SendMessage(context.Background(), jid,
		&waProto.Message{Conversation: proto.String(msg)}); err != nil {
		return err
	}
	return nil
}

// Disconnect disconnects session but remains authenticated on the server.
func (w *WhatsApp) Disconnect() {
	if w.client != nil {
		w.client.Disconnect()
	}
}

// Logout logs out of whats app and unauthenicates session on servers.
// only needed if you want another device to login.
func (w *WhatsApp) Logout() error {
	if w.client == nil || !w.client.IsLoggedIn() {
		return nil
	}
	if err := w.client.Logout(); err != nil {
		w.client.Disconnect()
		if err := w.client.Store.Delete(); err != nil {
			return err
		}
	}
	return nil
}

// eventHandler handles any whatsapp events from server.
func (w *WhatsApp) eventHandler(evt interface{}) {
	// TODO: For now do nothing. For future use.
	/*	switch v := evt.(type) {
		case *events.Message:
			fmt.Println("Received a message!", v.Message.GetConversation())
			fmt.Println("Sender" + v.Info.Sender.String())
		} */
}
