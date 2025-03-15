package p2p

import "net"

// Message represents any arbitrary data that is being dent over the
// each Transport between two nodes in the network.
type Message struct {
	Payload []byte
	From    net.Addr
}
