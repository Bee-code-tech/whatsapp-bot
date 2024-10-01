package whatsapp

import (
    "context"
    "fmt"

    "go.mau.fi/whatsmeow"
    "go.mau.fi/whatsmeow/store/sqlstore"
    waLog "go.mau.fi/whatsmeow/util/log"
    "go.mau.fi/whatsmeow/types"
    waProto "go.mau.fi/whatsmeow/binary/proto"
    _ "github.com/jackc/pgx/v4/stdlib" // PostgreSQL driver for WhatsMeow session store
    "google.golang.org/protobuf/proto"

    "github.com/skip2/go-qrcode" // For generating a QR code image file
)

var client *whatsmeow.Client

// InitializeWhatsAppClient initializes the WhatsMeow client and handles login via QR code
func InitializeWhatsAppClient() error {
    // Use PostgreSQL for session storage
    connString := "user=doctorbee password=baba2003 host=localhost dbname=whatsapp-bot sslmode=disable"
    container, err := sqlstore.New("pgx", connString, nil)
    if err != nil {
        return fmt.Errorf("failed to initialize PostgreSQL session store: %v", err)
    }

    // Get device from session container
    device, err := container.GetFirstDevice()
    if err != nil {
        return fmt.Errorf("failed to get device: %v", err)
    }

    // WhatsMeow provides its own logger
    logger := waLog.Stdout("WhatsApp", "INFO", true)

    client = whatsmeow.NewClient(device, logger)

    // Check if the client is logged in
    if client.Store.ID == nil {
        // Not logged in, generate a QR code for login
        qrChan, _ := client.GetQRChannel(context.Background())
        err = client.Connect()
        if err != nil {
            return fmt.Errorf("failed to connect client: %v", err)
        }

        for evt := range qrChan {
            if evt.Event == "code" {
                // Save the QR code as an image file (whatsapp-qr.png) with adjusted size
                fmt.Println("Saving QR code to 'whatsapp-qr.png'...")

                // Generate a smaller QR code image (200x200 with high error correction)
                err := qrcode.WriteFile(evt.Code, qrcode.High, 200, "whatsapp-qr.png")
                if err != nil {
                    return fmt.Errorf("failed to generate QR code image: %v", err)
                }

                fmt.Println("QR code saved successfully. Open 'whatsapp-qr.png' and scan it with WhatsApp.")
                break // Once QR code is saved, stop further processing
            }
        }
    } else {
        err = client.Connect()
        if err != nil {
            return fmt.Errorf("failed to connect client: %v", err)
        }
        fmt.Println("WhatsApp client connected successfully!")
    }

    return nil
}

// SendMessage sends a WhatsApp message to a specified number
func SendMessage(to string, message string) error {
    jid, err := types.ParseJID(to)
    if err != nil {
        return fmt.Errorf("invalid JID: %v", err)
    }

    // Create a text message to send
    textMessage := &waProto.Message{
        Conversation: proto.String(message),
    }

    // Send the message using a background context
    _, err = client.SendMessage(context.Background(), jid, textMessage)
    if err != nil {
        return fmt.Errorf("failed to send message: %v", err)
    }

    return nil
}
