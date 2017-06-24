package main

import (
    "github.com/aayanl/DEXtrading/peer"

    "github.com/rafaeldias/async"
)



func main() {
    async.Parallel(async.MapTasks{
        "peer": func() {
            peer.StartPeer();
        },
    })
}
