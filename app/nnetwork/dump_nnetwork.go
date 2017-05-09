package nnetwork

import "bytes"
import "encoding/gob"
import "log"

func (this *NNetwork) toBytes () []byte {
	var buf bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&buf) // Will write to network.

	// Encode (send) some values.
	err := enc.Encode(*this)

	if err != nil {
		log.Panic(err)
	}

	return buf.Bytes()
}

func fromBytes (b []byte) NNetwork {
	// Decode (receive) and print the values.
	var network NNetwork

	buf := bytes.NewReader(b)
	dec := gob.NewDecoder(buf) // Will read from network.
	err := dec.Decode(&network)
	if err != nil {
		log.Panic(err)
	}

	return network
}
