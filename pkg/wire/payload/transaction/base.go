package transaction

import (
	"bytes"
	"io"

	"github.com/CityOfZion/neo-go/pkg/wire/payload/transaction/types"
	"github.com/CityOfZion/neo-go/pkg/wire/payload/transaction/version"
	"github.com/CityOfZion/neo-go/pkg/wire/util"
)

type encodeExclusiveFields func(bw *util.BinWriter)
type decodeExclusiveFields func(br *util.BinReader)

// Transactioner is the interface that will unite the
// transaction types. Each transaction will implement this interface
// and so wil be a transactioner
type Transactioner interface {
	Encode(io.Writer) error
	Decode(io.Reader) error
	BaseTx() *Base
	ID() (util.Uint256, error)
}

// Base transaction is the template for all other transactions
// It contains all of the shared fields between transactions and
// the additional encodeExclusive and decodeExclusive methods, which
// can be overwitten in the other transactions to encode the non shared fields
type Base struct {
	Type            types.TX
	Version         version.TX
	Inputs          []*Input
	Outputs         []*Output
	Attributes      []*Attribute
	Witnesses       []*Witness
	Hash            util.Uint256
	encodeExclusive encodeExclusiveFields
	decodeExclusive decodeExclusiveFields
}

func createBaseTransaction(typ types.TX, ver version.TX) *Base {
	return &Base{
		Type:       typ,
		Version:    ver,
		Inputs:     []*Input{},
		Outputs:    []*Output{},
		Attributes: []*Attribute{},
		Witnesses:  []*Witness{},
	}

}

// Decode implements the transactioner interface
func (b *Base) Decode(r io.Reader) error {
	br := &util.BinReader{R: r}
	return b.DecodePayload(br)
}

// Encode implements the transactioner interface
func (b *Base) Encode(w io.Writer) error {
	bw := &util.BinWriter{W: w}
	b.EncodePayload(bw)
	return bw.Err
}

//EncodePayload implements the Messager interface
func (b *Base) EncodePayload(bw *util.BinWriter) {
	b.encodeHashableFields(bw)

	lenWitnesses := uint64(len(b.Witnesses))
	bw.VarUint(lenWitnesses)

	for _, witness := range b.Witnesses {
		witness.Encode(bw)
	}
}

// DecodePayload implements the messager interface
func (b *Base) DecodePayload(br *util.BinReader) error {
	b.decodeHashableFields(br)

	lenWitnesses := br.VarUint()

	b.Witnesses = make([]*Witness, lenWitnesses)
	for i := 0; i < int(lenWitnesses); i++ {
		b.Witnesses[i] = &Witness{}
		b.Witnesses[i].Decode(br)
	}

	if br.Err != nil {
		return br.Err
	}

	return b.createHash()
}

func (b *Base) encodeHashableFields(bw *util.BinWriter) {
	b.Type.Encode(bw)
	b.Version.Encode(bw)

	b.encodeExclusive(bw)

	lenAttrs := uint64(len(b.Attributes))
	lenInputs := uint64(len(b.Inputs))
	lenOutputs := uint64(len(b.Outputs))

	bw.VarUint(lenAttrs)
	for _, attr := range b.Attributes {
		attr.Encode(bw)
	}

	bw.VarUint(lenInputs)
	for _, input := range b.Inputs {
		input.Encode(bw)
	}

	bw.VarUint(lenOutputs)
	for _, output := range b.Outputs {
		output.Encode(bw)
	}
}

func (b *Base) decodeHashableFields(br *util.BinReader) {
	b.Type.Decode(br)

	b.Version.Decode(br)

	b.decodeExclusive(br)

	lenAttrs := br.VarUint()
	b.Attributes = make([]*Attribute, lenAttrs)
	for i := 0; i < int(lenAttrs); i++ {

		b.Attributes[i] = &Attribute{}
		b.Attributes[i].Decode(br)
	}
	/*
		attrStr1 := hex.EncodeToString(b.Attributes[0].Data)
		attrStr2 := hex.EncodeToString(b.Attributes[1].Data)
		fmt.Println(attrStr1)
		fmt.Println(attrStr2)
	*/
	lenInputs := br.VarUint()

	b.Inputs = make([]*Input, lenInputs)
	for i := 0; i < int(lenInputs); i++ {
		b.Inputs[i] = &Input{}
		b.Inputs[i].Decode(br)
	}

	lenOutputs := br.VarUint()
	b.Outputs = make([]*Output, lenOutputs)
	for i := 0; i < int(lenOutputs); i++ {
		b.Outputs[i] = &Output{}
		b.Outputs[i].Decode(br)
	}

}

// AddInput adds an input to the transaction
func (b *Base) AddInput(i *Input) {
	b.Inputs = append(b.Inputs, i)
}

// AddOutput adds an output to the transaction
func (b *Base) AddOutput(o *Output) {
	b.Outputs = append(b.Outputs, o)
}

// AddAttribute adds an attribute to the transaction
func (b *Base) AddAttribute(a *Attribute) {
	b.Attributes = append(b.Attributes, a)
}

// AddWitness adds a witness object to the transaction
func (b *Base) AddWitness(w *Witness) {
	b.Witnesses = append(b.Witnesses, w)
}

func (b *Base) createHash() error {

	hash, err := util.CalculateHash(b.encodeHashableFields)
	b.Hash = hash
	return err
}

// ID returns the TXID of the transaction
func (b *Base) ID() (util.Uint256, error) {
	var emptyHash util.Uint256
	var err error
	if b.Hash == emptyHash {
		err = b.createHash()
	}
	return b.Hash, err
}

// Bytes returns the raw bytes of the tx
func (b *Base) Bytes() []byte {
	buf := new(bytes.Buffer)
	b.Encode(buf)
	return buf.Bytes()
}

// Size returns the size of the tx in number of bytes.
func (b Base) Size() int {
	return len(b.Bytes())
}

// UTXOs returns the outputs in the tx
func (b *Base) UTXOs() []*Output {
	return b.Outputs
}

// TXOs returns the inputs in the tx
func (b *Base) TXOs() []*Input {
	return b.Inputs
}

// Witness returns the witnesses in the tx
func (b *Base) Witness() []*Witness {
	return b.Witnesses
}

// TypeTx returns the Type in the tx
func (b Base) TypeTx() types.TX {
	return b.Type
}

// VersionTx returns the Version in the tx
func (b Base) VersionTx() version.TX {
	return b.Version
}

// Attrs returns the Attributes in the tx
func (b Base) Attrs() []*Attribute {
	return b.Attributes
}
