# Wallet Management System

## Overview
The Wallet Management System is a RESTful API that allows users to manage their wallets. Users can deposit, withdraw, transfer money, check their wallet balance, and view their transaction history.

## Features
- **Deposit Money**: Users can deposit money into their wallets.
- **Withdraw Money**: Users can withdraw money from their wallets.
- **Transfer Money**: Users can transfer money across wallets.
- **Check Wallet Balance**: Users can check the current balance of their wallets.
- **View Transaction History**: Users can view the history of transactions associated with their wallets.


## Getting Started

This application is dockerised and ready with some seeded data containing 2 users and wallets.

### Prerequisites
- Docker: https://www.docker.com/products/docker-desktop/
- Postman: https://www.postman.com/downloads/

### Installation
1. Clone the repository:
   ```
   > git clone https://github.com/adyasingh/wallet
   > cd wallet
   ```

2. Build and run docker images:
   ```
   > docker-compose up -d --build
   ```

### Calling APIs

To call the APIs, import the Postman collection, which includes the following endpoints:

- **GET `/api/wallet/:id/balance`**
- **GET `/api/wallet/:id/transactions`**
- **POST `/api/wallet/:id/deposit`**
- **POST `/api/wallet/:id/withdraw`**
- **POST `/api/wallet/:id/transfer `**

## Testing

To run the tests, you can exec into the application container. This application utilizes a test database to conduct comprehensive tests. While some projects may utilize database mocking, I have always found it more useful and robust to test the service using a real test database.

Approximately one-third of the project development time was dedicated to writing these tests, ensuring robust functionality and reliability.

1. Exec into the container:
   ```
   > docker exec -it app bash
   ```
1. Running the tests
   ```
   > go test ./...
   ```

## Scope 

This project implements API business logic effectively. To make it production-ready, consider the following enhancements:

- **Monitoring and Logging:** Implement robust monitoring and logging solutions to track application performance, errors, and usage patterns. This will aid in identifying issues and improving reliability.

- **User Authentication:** Incorporate user authentication mechanisms to ensure secure access to the API. Options include OAuth, JWT, or API keys, depending on the use case.

- **Service Authentication:** Introduce service-to-service authentication to secure interactions between different components of the application. This helps prevent unauthorized access and strengthens security.

- **Secret Management:** Utilize secret management tools to securely store and manage sensitive information such as API keys, database credentials, and configuration settings.

- **Deployments and Environments:** Establish a streamlined deployment process, including CI/CD pipelines, to facilitate smooth transitions between development, staging, and production environments.

- **Enhanced API Response Data:** Improve the API responses by including metadata and additional context where necessary. This can enhance the usability of the API for clients.

- **Cleaner data seeding:** The fixtures can be migrated to an alternate endpoint or sql file for a cleaner main.go

- **Extended database:** Enhance the database schemas and names for tracking and feature extension.

There are, of course, several extensions to the current features that would add significant value and prepare this project for production. These include, but are not limited to:
Frontend pages
- Ability to onboard new users
- Ability to update and manage existing users
- Ability to onboard new wallets
- Ability to update and manage existing wallets
- Introduce other transaction types like refunds, chargebacks, etc.
- Multi-currency support

## Reviewing the Project
To effectively review this project, follow these steps:

- **Explore API Endpoints:** Begin by examining the Postman collection or using curl commands to interact with the APIs. This will help you understand the behavior and responses of each endpoint.

- **Review Test Cases:** Next, go through the test cases to familiarize yourself with the covered use cases. This will provide insights into the expected functionality and edge cases of the application.

- **Examine Models and Repositories:** Review the models and repository packages to understand the database structure and how the application interacts with the database layer. This is crucial for grasping data management.

- **Understand Business Logic:** The service package contains the core business logic of the application. Reviewing this will help you comprehend how data is processed and manipulated.

- **Review the errors:** The errors package contains all the errors that this application might return. It will give you an idea of the assumptions of the system.

- **Analyze HTTP Handlers:** Finally, examine the handler package, which manages HTTP interactions. This will give you an overview of how the application handles incoming requests and sends responses.

