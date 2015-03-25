// Generated by protocol_builder
// Do not edit

package protocol

import (
	"io"
)

func (l *LoginStart) id() int { return 0 }
func (l *LoginStart) write(ww io.Writer) (err error) {
	if err = writeString(ww, l.Username); err != nil {
		return
	}
	return
}
func (l *LoginStart) read(rr io.Reader) (err error) {
	if l.Username, err = readString(rr); err != nil {
		return
	}
	return
}

func (e *EncryptionResponse) id() int { return 1 }
func (e *EncryptionResponse) write(ww io.Writer) (err error) {
	if err = writeVarInt(ww, VarInt(len(e.SharedSecret))); err != nil {
		return
	}
	if _, err = ww.Write(e.SharedSecret); err != nil {
		return
	}
	if err = writeVarInt(ww, VarInt(len(e.VerifyToken))); err != nil {
		return
	}
	if _, err = ww.Write(e.VerifyToken); err != nil {
		return
	}
	return
}
func (e *EncryptionResponse) read(rr io.Reader) (err error) {
	var tmp0 VarInt
	if tmp0, err = readVarInt(rr); err != nil {
		return
	}
	e.SharedSecret = make([]byte, tmp0)
	if _, err = rr.Read(e.SharedSecret); err != nil {
		return
	}
	var tmp1 VarInt
	if tmp1, err = readVarInt(rr); err != nil {
		return
	}
	e.VerifyToken = make([]byte, tmp1)
	if _, err = rr.Read(e.VerifyToken); err != nil {
		return
	}
	return
}

func init() {
	packetCreator[Login][serverbound][0] = func() Packet { return &LoginStart{} }
	packetCreator[Login][serverbound][1] = func() Packet { return &EncryptionResponse{} }
}