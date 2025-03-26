package p2p

// Peer is an interface that represents the remte node
type Peer interface{}

// Transport is anything that handles the communication
// between the nodes in the network.
// form (TCP, UDP, websockets)
type Transport interface {
	ListenAndAccept() error
}
