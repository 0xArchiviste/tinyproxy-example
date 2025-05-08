# TinyProxy Example

## Setup Instructions

### Server
- Run the server to create an endpoint for testing the proxy:
  ```
  cd server
  go run main.go
  ```
  This will start a server on port 8080 with an `/ipcheck` endpoint.

### Proxy
- Use Debian or Ubuntu 20.04+ for the proxy server
- Install TinyProxy: `apt-get install tinyproxy -y`
- Edit the `config/tinyproxy.conf` file with your desired settings
- Run TinyProxy in debug mode: `tinyproxy -d -c tinyproxy.conf`

### Client
- Run the client to test the connection through the proxy:
  ```
  cd client
  go run main.go
  ```
  This will connect to the server through the configured proxy.