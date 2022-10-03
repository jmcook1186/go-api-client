# Go API client

This is entirely for myself to improve my Go and should not be considered in any way suitable for any downstream usage.

The client makes http requests to a beacon client. I have been using Lighthouse local testnet.

To run this 

```sh
cd lighthouse
nvm use --lts
./scripts/local_testnet/start_local_testnet.sh 
```

The run the client

```sh
go run main.go endpoints.go types.go 
```

