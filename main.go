package main

import (
  "fmt"
)

func main(){
    var base_url string = "http://localhost:8003/"
    var direction string = "outbound"
    var state string = "connected"

    peer := get_peers(base_url, state, direction)
    fmt.Println(peer.Data[0].PeerID)
}

