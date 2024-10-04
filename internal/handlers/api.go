package handlers

import (
    "encoding/json"
    "net/http"
    "whatsapp-bot/pkg/whatsapp"
)

type SendMessageRequest struct {
    To      string `json:"to"`
    Message string `json:"message"`
}

// SendMessageHandler handles HTTP requests to send WhatsApp messages
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
    var req SendMessageRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = whatsapp.SendMessage(req.To, req.Message)
    if err != nil {
        http.Error(w, "Failed to send message", http.StatusInternalServerError)
        return
    }

    // write head and send message to the number 
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Message sent successfully!"))
}
