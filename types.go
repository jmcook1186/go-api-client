package main

type GenesisData struct {
	Data struct {
		GenesisTime           string `json:"genesis_time"`
		GenesisValidatorsRoot string `json:"genesis_validators_root"`
		GenesisForkVersion    string `json:"genesis_fork_version"`
	} `json:"data"`
}

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
