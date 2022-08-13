package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type ConfigFile struct {
	Schema string `json:"$schema"`
	Global struct {
		FabricVersion string `json:"fabricVersion"`
		Tls           bool   `json:"tls"`
		Monitoring    struct {
			Loglevel string `json:"loglevel"`
		} `json:"monitoring"`
	} `json:"global"`
	Orgs []struct {
		Organization struct {
			Name    string `json:"name"`
			Domain  string `json:"domain"`
			MspName string `json:"mspName,omitempty"`
		} `json:"organization"`
		Orderers []struct {
			GroupName string `json:"groupName"`
			Prefix    string `json:"prefix"`
			Type      string `json:"type"`
			Instances int    `json:"instances"`
		} `json:"orderers,omitempty"`
		Ca struct {
			Prefix string `json:"prefix"`
		} `json:"ca,omitempty"`
		Peer struct {
			Prefix    string `json:"prefix"`
			Instances int    `json:"instances"`
			Db        string `json:"db"`
		} `json:"peer,omitempty"`
	} `json:"orgs"`
	Channels   []Channels   `json:"channels"`
	ChainCodes []ChainCodes `json:"chaincodes"`
}

type Organization struct {
	Name  string   `json:"name"`
	Peers []string `json:"peers"`
}

type Channels struct {
	Name string         `json:"name"`
	Orgs []Organization `json:"orgs"`
}

type ChainCodes struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Lang        string `json:"lang"`
	Channel     string `json:"channel"`
	Init        string `json:"init"`
	Endorsement string `json:"endorsement"`
	Directory   string `json:"directory"`
}

func GenerateConfigFile(config ConfigFile) error {
	config.Schema = "https://github.com/hyperledger/releases/download/1.1.0/schema.json"
	config.Global.Tls = true

	mapping, err := json.Marshal(config)
	if err != nil {
		return errors.New("failed to mapping json")
	}

	err = ioutil.WriteFile("sample.json", mapping, 0644)
	if err != nil {
		return errors.New("failed to write file")
	}
	
	return nil
}
