### Deployment steps

The server is run and tested on:
Go version: 1.15
OS version: macOS Mojave 10.14.6 

We could run by pulling the code, installing the go dependencies and then running 
```bash
  go run -mod=vendor main.go
```
The configuration specifics like the `PORT` and `PRIVATE_KEY` could be `exported` as environment variables.
```bash
export PORT=3000
export PRIVATE_KEY="-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIEFMEZrmlYxczXKFxIlNvNGR5JQvDhTkLovJYxwQd3ua
-----END PRIVATE KEY-----"
```
If no environment variable is set, by default the signature server runs in the PORT 3000 with the PRIVATE_KEY as mentioned above.

A dockerfile with a docker-compose is also provided so that the server could be run and tested in different machines. To run using the docker compose,

```bash
  docker-compose up
```

# API

##### GET/public_key: 

Returns a JSON object containing the publickey of the daemon key.
Response
```json
{"public_key": "27smkc6v97lNodf7C7aKN/J/Z641/8+BBFXmqfAoGCQ="}` 
```

##### PUT/transaction:

Takes a blob of data (arbitrary bytes)representing the transaction data in the form of a base64 string, and remembers it in memory. Returns a random, unique identifier for the transaction.

Request
```json
{"txn": "/+ABAgM="}
```
Response
```json
{"id": "7d8b3649-a48e-4169-819f-82bceb6d4613"}
```

##### POST/signature:

 Takes a list of transaction identifiers,and builds a JSON array of strings containing the base64-encoded transaction blobs indicated by the given identifiers. It signs this array (serialised as JSON without any whitespace) using the daemon private key. Finally, it returns the array that was signed, as well as the signature as a base64 string.

Request
```json
{"ids": ["7d8b3649-a48e-4169-819f-82bceb6d4613"]}
```
Response
```json
{"message": ["/+ABAgM="], "signature": "3/s5I9s67uwb1typ6JFXpVTqkRquE7QOFGH4cbg7YU9l2t8Ik1rzjp3i1vvle05WuRPoXi1WDrv/yomFdj wOAg=="}
```

The server is multi-threaded. `http.ListenAndServe()` spins new go routine for every new request. By default, `GOMAXPROCS` is set to the number of cores of the CPU. Hence when spammed with requests, all cores of the machine is used.