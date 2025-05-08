# TinyProxy Example

## Setup Instructions

### Server
- Run the server to create an endpoint for testing the proxy:
  ```
  cd server
  go run main.go
  ```
  This will start a server on port 8080 with an `/ipcheck` endpoint.

### Proxy Server
- Use Debian or Ubuntu 20.04+ for the proxy server
- Install TinyProxy: `sudo apt-get install tinyproxy -y`
- Edit the `config/tinyproxy.conf` file with your desired settings
    - You can add specific ranges on the allow rule to finetune access
- Configure ufw to allow traffic on the specific port of the forward proxy
    - In this example, `sudo ufw allow from any to any port 11111 proto tcp` will work
    - You can configure ip access on ufw OR tinyproxy
- Run TinyProxy in debug mode: `tinyproxy -d -c tinyproxy.conf`

### Client
- Run the client to test the connection through the proxy:
  ```
  cd client
  go run main.go
  ```
  This will connect to the server through the configured proxy.