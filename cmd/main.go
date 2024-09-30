package main

import (
    "log"
    "whatsapp-bot/pkg/database"
    "whatsapp-bot/pkg/whatsapp"
)

func main() {
    // Connect to PostgreSQL
    err := database.ConnectPostgres()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer database.ClosePostgres()

    // Start WhatsApp client
    err = whatsapp.InitializeWhatsAppClient()
    if err != nil {
        log.Fatalf("Failed to initialize WhatsApp client: %v", err)
    }

    // Keep the application running
    select {}
}
