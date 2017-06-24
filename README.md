# DEXtrading
### Just some ideas I'm putting down before coding


## Build:
```
export GOPATH=$(pwd)
go get github.com/ethereum/go-ethereum
go get github.com/rafaeldias/async

go get github.com/aayanl/DEXtrading

cd src/github.com/aayanl/DEXtrading

go build main.go
```


## Specs:

* #### p2p:
  * Hubs listing peers: Anyone can host these.
  * Peers connect to hub and gets an id (optional)
  * Peers can connect to other peers via ids by querying the hub which ip has the id
  * Peers can also directly connect to other peers
  * Peers can list what they're interested in and send that data to hub (optional)

* #### Trade:
  * TBD


* #### Langauge:
  * <s>Python and NodeJS are good options, I really want to do a project in Go though.</s>
  * It shall be go!
