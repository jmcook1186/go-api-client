package main

type PeerDescription struct {
	Data []struct {
		PeerID             string      `json:"peer_id"`
		Enr                interface{} `json:"enr"`
		LastSeenP2PAddress string      `json:"last_seen_p2p_address"`
		State              string      `json:"state"`
		Direction          string      `json:"direction"`
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}