# Kilimo-Chain

Kilimo-Chain is a web application for managing agricultural supply chains. It allows users to register as farmers or cooperates, log in, and interact with a blockchain to add transactions and view the blockchain.

## Features

- User registration for farmers and cooperates
- User login for farmers and cooperates
- Dashboard displaying user information
- Adding transactions to the blockchain
- Viewing the blockchain

## Project Structure

- `main.go`: The main application file that sets up routes and starts the server
- `User`: Struct for user information
- `blockchain`: Package for managing blockchain data
- `static`: Directory for static files (CSS, JavaScript, images)

## Prerequisites

- Go (version 1.16+)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/kilimo-chain.git
cd kilimo-chain
```

## Run the server:

```bash
go run main.go
```

## Usage

1. Sign up as a farmer or cooperate by navigating to /signup.
2. Log in by navigating to /signin.
3. After logging in, access the dashboard at /dashboard.

## Contributing
- Contributions are welcome! Please open an issue or submit a pull request.


## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
