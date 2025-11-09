package main

import (
	"fmt"
	"sort"

	"github.com/google/uuid"
)

type NodeDetails struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var nodeRegistry = map[string]NodeDetails{}

var hashKeyToNodeNameMap = map[string]string{}
var availableServers = []string{}

func AddNode(details NodeDetails) NodeDetails {
	details.Id = uuid.New().String()
	hashKey := getHashValue(details.Id)
	_, exists := hashKeyToNodeNameMap[hashKey]
	if exists {
		return NodeDetails{}
	}
	fmt.Println("hashKey: " + hashKey)
	nodeRegistry[hashKey] = details
	availableServers = append(availableServers, hashKey)
	sort.Strings(availableServers)
	return details
}

func getTargetNodeAddress(key string) string {
	var hash = getHashValue(key)
	nodeIdHash := findTargetNodeIdHash(hash)
	return nodeRegistry[nodeIdHash].Address
}

func findTargetNodeIdHash(keyHash string) string {
	for i := 0; i < len(availableServers); i++ {
		if availableServers[i] >= keyHash {
			return availableServers[i]
		}
	}
	return availableServers[0]
}

func getNodeRegistry() []NodeDetails {
	var nodes []NodeDetails

	for _, details := range nodeRegistry {
		nodes = append(nodes, details)
	}
	return nodes
}

func removeNode(nodeId string) {
	hashKey := getHashValue(nodeId)
	removeKeyFromAvailableServers(hashKey)
}

func removeKeyFromAvailableServers(key string) {

	for i, k := range availableServers {
		if k == key {
			// Remove element at index i
			availableServers = append(availableServers[:i], availableServers[i+1:]...)
			break
		}
	}
}
