package main

import (
	"encoding/json"
	"fmt"
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

func get_state_root(base string, state string) StateRoot {

	var endpoint = fmt.Sprintf("eth/v1/beacon/states/%s/root", state)
	url := base + endpoint

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var root StateRoot
	json.Unmarshal([]byte(body), &root)
	return root
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
		log.Print(err)
	}

	var peer PeerDescription
	json.Unmarshal([]byte(body), &peer)

	return peer

}
