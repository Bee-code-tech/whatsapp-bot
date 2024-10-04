package database

import (
    "context"
    "log"
    "github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

// ConnectPostgres establishes a connection to the PostgreSQL database
func ConnectPostgres() error {
    var err error
    conn, err = pgx.Connect(context.Background(), "postgres://doctorbee:baba2003@localhost:5432/whatsapp-bot")
    if err != nil {
        return err
    }

    log.Println("Connected to PostgreSQL!")
    return nil
}

// ClosePostgres closes the connection to PostgreSQL
func ClosePostgres() {
    conn.Close(context.Background())
}

// StoreSession stores the WhatsApp session data in the database
func StoreSession(userID int, session []byte) error {
    _, err := conn.Exec(context.Background(), "INSERT INTO whatsapp_sessions (user_id, session_data) VALUES ($1, $2)", userID, session)
    return err
}

// GetSession retrieves the WhatsApp session data from the database..
func GetSession(userID int) ([]byte, error) {
    var session []byte
    err := conn.QueryRow(context.Background(), "SELECT session_data FROM whatsapp_sessions WHERE user_id=$1", userID).Scan(&session)
    return session, err
}
