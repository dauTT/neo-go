package transaction

import (
	"bytes"

	"github.com/CityOfZion/neo-go/pkg/wire/util"
)

// Output represents a transaction output in the neo-network
type Output struct {
	// The NEO asset id used in the transaction.
	AssetID util.Uint256

	// Amount of AssetType send or received.
	Amount int64

	// The address of the remittee.
	ScriptHash util.Uint160
}

//NewOutput returns an output object
func NewOutput(assetID util.Uint256, Amount int64, ScriptHash util.Uint160) *Output {
	return &Output{
		assetID,
		Amount,
		ScriptHash,
	}
}

// Encode encodes the Output into a binary writer
func (o *Output) Encode(bw *util.BinWriter) {
	bw.Write(o.AssetID)
	bw.Write(o.Amount)
	bw.Write(o.ScriptHash)
}

// Decode decodes a binary reader into an output object
func (o *Output) Decode(br *util.BinReader) {
	br.Read(&o.AssetID)
	br.Read(&o.Amount)
	br.Read(&o.ScriptHash)
}

// Bytes returns the raw bytes of the Output.
func (o *Output) Bytes() []byte {
	buf := new(bytes.Buffer)
	bw := &util.BinWriter{W: buf}
	o.Encode(bw)
	return buf.Bytes()
}

// Size returns the size of the Output in number of bytes.
func (o Output) Size() int {
	return len(o.Bytes())
}
