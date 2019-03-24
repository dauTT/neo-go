package transaction

import (
	"bytes"

	"github.com/CityOfZion/neo-go/pkg/wire/util"
)

// Input represents a Transaction input.
type Input struct {
	// The hash of the previous transaction.
	PrevHash util.Uint256

	// The index of the previous transaction.
	PrevIndex uint16
}

//NewInput returns a transaction input object
func NewInput(prevHash util.Uint256, prevIndex uint16) *Input {
	return &Input{
		prevHash,
		prevIndex,
	}
}

// Encode encodes the given input into a binary writer
func (i *Input) Encode(bw *util.BinWriter) {
	bw.Write(i.PrevHash)
	bw.Write(i.PrevIndex)
}

// Decode decodes a binary reader into an input object
func (i *Input) Decode(br *util.BinReader) {
	br.Read(&i.PrevHash)
	br.Read(&i.PrevIndex)
}

// Bytes returns the raw bytes of the Input.
func (i *Input) Bytes() []byte {
	buf := new(bytes.Buffer)
	bw := &util.BinWriter{W: buf}
	i.Encode(bw)
	return buf.Bytes()
}

// Size returns the size of the Input in number of bytes.
func (i Input) Size() int {
	return len(i.Bytes())
}
