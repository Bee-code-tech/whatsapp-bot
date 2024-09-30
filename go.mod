module whatsapp-bot

go 1.19
require go.mau.fi/whatsmeow v0.3.0


require (
    github.com/jackc/pgx/v4 v4.13.0
    github.com/tulir/whatsmeow v0.3.0
    github.com/mattn/go-sqlite3 v1.14.7 // You can replace this when using PostgreSQL for sessions
)
