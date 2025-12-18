# Wallet Service

Test assignment: REST service for wallet operations.

## Stack
- Go
- PostgreSQL
- Docker / Docker Compose

## API

### POST /api/v1/wallet
```json
{
  "walletId": "UUID",
  "operationType": "DEPOSIT | WITHDRAW",
  "amount": 1000
}

#Run
docker compose up -d
go run ./cmd/app
