# Golnet

**golnet** is a simple telnet client for the game *7 Days to Die*. It allows you to send commands to a server without receiving any response messages.
This tool is perfect for server administrators who need to automate or send commands to the server without worrying about feedback.

## Features

- Send commands to a 7 Days to Die server via Telnet.
- No response messages are returned (ideal for server management).
- Lightweight and easy to use.

## Compilation

To compile **golnet**, use the following command:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o golnet -ldflags="-w -s" golnet.go
```

# Usage
To run **golnet**, use the following command structure:
```bash
./golnet <IP> <PORT> <PASSWORD> '<COMMAND>'
```

# Example
Here’s an example of how to send a command to a server:
```bash
./golnet 127.0.0.1 8081 you_password 'say "[FF0000]Checking the server label[-] without panic."'
```

This will send a message to the server with the text ```Checking the server label``` in red color (```[FF0000]```).

# Acknowledgements
golnet was created with the help of ChatGPT, a powerful AI tool developed by OpenAI, to assist in generating the structure and logic of the program.
