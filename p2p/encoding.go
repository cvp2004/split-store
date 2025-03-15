package p2p

import (
	"bufio"
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *Message) error {

	bufReader := bufio.NewReader(r)
	line, err := bufReader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return err
	}
	msg.Payload = line

	return nil
}
