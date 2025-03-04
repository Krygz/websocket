# WebSocket in Golang

This is a simple WebSocket implementation in Golang, consisting of a client and a server. The server listens for WebSocket connections, receives messages, and responds with a timestamp. The client connects to the server, sends messages from the terminal, and receives responses.

---

## Features
- Simple WebSocket server using `gorilla/websocket`
- Client that connects and communicates with the server
- Sends and receives messages in real-time
- Includes timestamped responses from the server

---

## Installation & Setup

### Prerequisites
- **Go 1.18+** installed
- **Gorilla WebSocket** package installed:
  ```bash
  go get -u github.com/gorilla/websocket
  ```

### Clone the Repository
```bash
git clone https://github.com/yourusername/websocket-golang.git
cd websocket-golang
```

---

## Running the WebSocket Server
Run the following command to start the WebSocket server:
```bash
go run server.go
```

The server will start listening on port `3000`:
```bash
Starting server on: 3000
```

---

## Running the WebSocket Client
Run the following command to start the client:
```bash
go run client.go
```

Upon execution, the client will connect to the server:
```bash
Connecting to ws://localhost:3000/ws
Type something...
```

Type a message and press **Enter** to send it. The server will respond with the same message along with a timestamp.

**Example Interaction:**
```bash
You: Hello Server!
Server: Hello Server! , at 2024-03-04 12:30:00 +0000 UTC
```

---

## How It Works
### Server:
1. Upgrades incoming HTTP connections to WebSocket.
2. Reads messages from the client.
3. Prints received messages.
4. Sends the same message back with a timestamp.

### Client:
1. Connects to the WebSocket server.
2. Listens for messages from the terminal.
3. Sends messages to the server.
4. Displays the server's response.

---

## Notes
- Ensure port **3000** is not blocked by a firewall.
- If running the server on a remote machine, replace `localhost` with the server's IP in `client.go`.
- The `gorilla/websocket` package is used for handling WebSocket connections efficiently.
