package transaction

import (
	"bytes"

	"github.com/CityOfZion/neo-go/pkg/wire/util"
)

//Witness represents a Witness object in a neo transaction
type Witness struct {
	InvocationScript   []byte
	VerificationScript []byte
}

// Encode encodes a Witness into a binary writer
func (s *Witness) Encode(bw *util.BinWriter) error {

	bw.VarUint(uint64(len(s.InvocationScript)))
	bw.Write(s.InvocationScript)

	bw.VarUint(uint64(len(s.VerificationScript)))
	bw.Write(s.VerificationScript)

	return bw.Err
}

// Decode decodes a binary reader into a Witness object
func (s *Witness) Decode(br *util.BinReader) error {

	lenb := br.VarUint()
	s.InvocationScript = make([]byte, lenb)
	br.Read(s.InvocationScript)

	lenb = br.VarUint()
	s.VerificationScript = make([]byte, lenb)
	br.Read(s.VerificationScript)

	return br.Err
}

// Bytes returns the Byte representation of Witness.
func (s *Witness) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	bbuf := &util.BinWriter{W: buf}
	err := s.Encode(bbuf)
	return buf.Bytes(), err
}

// Size returns the size of the Block in number of bytes.
func (b *Witness) Size() (int, error) {
	bb, err := b.Bytes()
	if err != nil {
		return 0, err
	}
	return len(bb), nil
}
