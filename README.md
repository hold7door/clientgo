# clientgo
A very simple http client

## Usage

```
 clientgo URL [flags]

Flags:
  -d, --data string         data to be sent as the request body
  -H, --headers strings     custom headers headers to be sent with the request, headers are separated by "," as in "HeaderName: Header content,OtherHeader: Some other value"
  -h, --help                help for clientgo
  -k, --insecure            allows insecure server connections over HTTPS
  -m, --method string       HTTP method to be used for the request (default "GET")
  -u, --user-agent string   the user agent to be used for requests (default "clientgo")
  ```
  
  ## Build binary
  
  To build a binary for your platform, run - 
  
  ```
  # building the program for intel macs
GOOS=darwin GOARCH=amd64 go build -o clientgo cmd/clientgo/main.go 

# building the program for M1 macs
GOOS=darwin GOARCH=amd64 go build -o clientgo cmd/clientgo/main.go 

# building the program for 64 bits amd/intel linux
GOOS=linux GOARCH=amd64 go build -o clientgo cmd/clientgo/main.go 
  ```
