package p2p

// HandshakeFunc is a
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
