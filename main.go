package main

import (
	"fmt"
)

func main() {
	var base_url string = "http://localhost:8003/"
	var direction string = "outbound"
	var peer_state string = "connected"
	var block_state string = "finalized"

	peer := get_peers(base_url, peer_state, direction)

	if len(peer.Data) >= 1 {
		fmt.Println(peer.Data[0].LastSeenP2PAddress)
	} else {
		fmt.Println(fmt.Sprintf("no peers with state = %s and direction = %s", peer_state, direction))
	}

	genesis := get_genesis(base_url)
	fmt.Println(genesis.Data.GenesisValidatorsRoot)

	state_root := get_state_root(base_url, block_state)
	fmt.Println(state_root.Data)

}
