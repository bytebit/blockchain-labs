package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const (
	ProtocolVersion = "0.1.0"
	HandshakeMsg    = "handshake"
	AddrMsg         = "addr"
	PingMsg         = "ping"
	PongMsg         = "pong"
)

type Message struct {
	Type    string      `json:"type"`
	Version string      `json:"version"`
	Payload interface{} `json:"payload"`
}

type AddrPayload struct {
	Addrs []string `json:"addrs"`
}

type Node struct {
	mu            sync.RWMutex
	listenAddr    string
	peers         map[string]net.Conn
	knownAddrs    map[string]bool
	seedNodes     []string
}

func NewNode(listenAddr string, seedNodes []string) *Node {
	return &Node{
		listenAddr: listenAddr,
		peers:      make(map[string]net.Conn),
		knownAddrs: make(map[string]bool),
		seedNodes:  seedNodes,
	}
}

func (n *Node) Start() error {
	ln, err := net.Listen("tcp", n.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Printf("Node started on %s", n.listenAddr)
	n.knownAddrs[n.listenAddr] = true

	go n.acceptConnections(ln)
	go n.connectToSeedNodes()
	go n.periodicBroadcast()
	go n.printKnownNodes()

	select {}
}

func (n *Node) acceptConnections(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go n.handleConnection(conn, false)
	}
}

func (n *Node) connectToSeedNodes() {
	for _, seed := range n.seedNodes {
		if seed == n.listenAddr {
			continue
		}
		go n.connectToPeer(seed)
	}
}

func (n *Node) connectToPeer(addr string) {
	n.mu.Lock()
	if _, exists := n.peers[addr]; exists {
		n.mu.Unlock()
		return
	}
	n.mu.Unlock()

	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return
	}

	n.handleConnection(conn, true)
}

func (n *Node) handleConnection(conn net.Conn, outgoing bool) {
	remoteAddr := conn.RemoteAddr().String()

	n.mu.Lock()
	n.peers[remoteAddr] = conn
	n.mu.Unlock()

	defer func() {
		conn.Close()
		n.mu.Lock()
		delete(n.peers, remoteAddr)
		n.mu.Unlock()
	}()

	if outgoing {
		n.sendHandshake(conn)
	}

	decoder := json.NewDecoder(conn)
	for {
		var msg Message
		if err := decoder.Decode(&msg); err != nil {
			break
		}
		n.handleMessage(conn, msg)
	}
}

func (n *Node) sendHandshake(conn net.Conn) {
	msg := Message{
		Type:    HandshakeMsg,
		Version: ProtocolVersion,
		Payload: n.listenAddr,
	}
	json.NewEncoder(conn).Encode(msg)
}

func (n *Node) handleMessage(conn net.Conn, msg Message) {
	switch msg.Type {
	case HandshakeMsg:
		if addr, ok := msg.Payload.(string); ok {
			n.mu.Lock()
			n.knownAddrs[addr] = true
			n.mu.Unlock()
			n.sendAddrMessage(conn)
		}
	case AddrMsg:
		payloadBytes, _ := json.Marshal(msg.Payload)
		var payload AddrPayload
		json.Unmarshal(payloadBytes, &payload)
		n.mu.Lock()
		for _, addr := range payload.Addrs {
			n.knownAddrs[addr] = true
			if _, exists := n.peers[addr]; !exists && addr != n.listenAddr {
				go n.connectToPeer(addr)
			}
		}
		n.mu.Unlock()
	case PingMsg:
		n.sendPong(conn)
	case PongMsg:
	}
}

func (n *Node) sendAddrMessage(conn net.Conn) {
	n.mu.RLock()
	addrs := make([]string, 0, len(n.knownAddrs))
	for addr := range n.knownAddrs {
		addrs = append(addrs, addr)
	}
	n.mu.RUnlock()

	msg := Message{
		Type:    AddrMsg,
		Version: ProtocolVersion,
		Payload: AddrPayload{Addrs: addrs},
	}
	json.NewEncoder(conn).Encode(msg)
}

func (n *Node) sendPong(conn net.Conn) {
	msg := Message{
		Type:    PongMsg,
		Version: ProtocolVersion,
	}
	json.NewEncoder(conn).Encode(msg)
}

func (n *Node) broadcastAddrMessage() {
	n.mu.RLock()
	addrs := make([]string, 0, len(n.knownAddrs))
	for addr := range n.knownAddrs {
		addrs = append(addrs, addr)
	}
	peers := make([]net.Conn, 0, len(n.peers))
	for _, conn := range n.peers {
		peers = append(peers, conn)
	}
	n.mu.RUnlock()

	msg := Message{
		Type:    AddrMsg,
		Version: ProtocolVersion,
		Payload: AddrPayload{Addrs: addrs},
	}

	for _, conn := range peers {
		json.NewEncoder(conn).Encode(msg)
	}
}

func (n *Node) periodicBroadcast() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		n.broadcastAddrMessage()
	}
}

func (n *Node) printKnownNodes() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		n.mu.RLock()
		fmt.Printf("\n")
		fmt.Println("========================================")
		fmt.Printf("  Node: %s\n", n.listenAddr)
		fmt.Println("  Known P2P Network Nodes:")
		fmt.Println("========================================")
		i := 1
		for addr := range n.knownAddrs {
			fmt.Printf("  %d. %s\n", i, addr)
			i++
		}
		fmt.Println("========================================")
		fmt.Printf("  Total: %d node(s)\n", len(n.knownAddrs))
		fmt.Println("========================================")
		fmt.Printf("\n")
		n.mu.RUnlock()
	}
}

func main() {
	listenAddr := flag.String("addr", "127.0.0.1:8000", "Listen address")
	seeds := flag.String("seeds", "", "Comma-separated seed nodes")
	flag.Parse()

	var seedNodes []string
	if *seeds != "" {
		seedNodes = splitComma(*seeds)
	}

	log.SetOutput(os.Stdout)
	log.Printf("Starting P2P node on %s", *listenAddr)
	if len(seedNodes) > 0 {
		log.Printf("Seed nodes: %v", seedNodes)
	}

	node := NewNode(*listenAddr, seedNodes)
	if err := node.Start(); err != nil {
		log.Fatalf("Failed to start node: %v", err)
	}
}

func splitComma(s string) []string {
	var result []string
	var current string
	for _, c := range s {
		if c == ',' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(c)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}
