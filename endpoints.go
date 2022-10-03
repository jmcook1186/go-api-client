package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//beacon namespace

func get_genesis(base string) GenesisData {
	var endpoint = "eth/v1/beacon/genesis"
	url := base + endpoint

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var genesis GenesisData
	json.Unmarshal([]byte(body), &genesis)
	return genesis
}



// node namespace
func get_peers(base string, state string, direction string) PeerDescription {

	var endpoint string = "eth/v1/node/peers/"
	url := base + endpoint + "?" + "state=" + state + "&direction=" + direction

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var peer PeerDescription
	json.Unmarshal([]byte(body), &peer)
	return peer

}
