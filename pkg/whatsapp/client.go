package whatsapp

import (
    "go.mau.fi/whatsmeow"
    "go.mau.fi/whatsmeow/store/sqlstore"
    _ "github.com/mattn/go-sqlite3"
)


var client *whatsmeow.Client

// InitializeWhatsAppClient initializes the WhatsMeow client and handles login via QR code
func InitializeWhatsAppClient() error {
    // Load session from the database (replace SQLite with PostgreSQL)
    sessionStore, err := sqlstore.New("sqlite3", "file:whatsapp-session.db?_foreign_keys=on", log.Default())
    if err != nil {
        return fmt.Errorf("failed to initialize session store: %v", err)
    }

    client = whatsmeow.NewClient(sessionStore, log.Default())

    // Check if the client is logged in
    if client.Store.ID == nil {
        qrChan, _ := client.GetQRChannel(context.Background())
        err = client.Connect()
        if err != nil {
            return fmt.Errorf("failed to connect client: %v", err)
        }

        // Generate QR code for login
        for evt := range qrChan {
            if evt.Event == "code" {
                fmt.Println("Scan this QR code in WhatsApp:", evt.Code)
            } else {
                fmt.Println("Login event:", evt.Event)
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
    jid, err := whatsmeow.ParseJID(to)
    if err != nil {
        return fmt.Errorf("invalid JID: %v", err)
    }

    textMessage := &whatsmeow.Message{
        Conversation: message,
    }

    _, err = client.SendMessage(jid, "", textMessage)
    if err != nil {
        return fmt.Errorf("failed to send message: %v", err)
    }

    return nil
}
