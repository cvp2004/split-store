package p2p

// Peer is an interface that represents a remote node.
type Peer interface {
	Close() error
}

// Transport is anything that handles communication
// between nodes in the network. This can be of the form TCP, UDP, WebSockets, etc.
type Transport interface {
	listenAndAccept() error
	Consume() <-chan RPC
}
