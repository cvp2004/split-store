package p2p

import (
	"bufio"
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, rpc *RPC) error {

	bufReader := bufio.NewReader(r)
	line, err := bufReader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return err
	}
	rpc.Payload = line

	return nil
}
