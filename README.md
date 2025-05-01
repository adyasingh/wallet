# Wallet Management System

## Overview
The Wallet Management System is a RESTful API that allows users to manage their wallets. Users can deposit, withdraw, transfer money, check their wallet balance, and view their transaction history.

## Features
- **Deposit Money**: Users can deposit money into their wallets.
- **Withdraw Money**: Users can withdraw money from their wallets.
- **Transfer Money**: Users can transfer money between wallets.
- **Check Wallet Balance**: Users can check the current balance of their wallets.
- **View Transaction History**: Users can view the history of transactions associated with their wallets.

## Project Structure
```
wallet-app
├── cmd
│   └── main.go
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   ├── balance.go
│   │   │   ├── deposit.go
│   │   │   ├── transaction_history.go
│   │   │   ├── transfer.go
│   │   │   └── withdraw.go
│   │   └── routes.go
│   ├── models
│   │   ├── transaction.go
│   │   └── wallet.go
│   ├── services
│   │   ├── balance_service.go
│   │   ├── deposit_service.go
│   │   ├── transaction_service.go
│   │   ├── transfer_service.go
│   │   └── withdraw_service.go
│   └── storage
│       ├── database.go
│       └── memory_storage.go
├── config
│   └── config.go
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites
- Go 1.16 or later
- A working database (if using the database storage option)

### Installation
1. Clone the repository:
   ```
   git clone <repository-url>
   cd wallet-app
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Application
To run the application, execute the following command:
```
go run cmd/main.go
```

### API Endpoints
- **POST /deposit**: Deposit money into a wallet.
- **POST /withdraw**: Withdraw money from a wallet.
- **POST /transfer**: Transfer money between wallets.
- **GET /balance/{walletID}**: Check the balance of a wallet.
- **GET /transactions/{walletID}**: View the transaction history of a wallet.

## License
This project is licensed under the MIT License. See the LICENSE file for details.