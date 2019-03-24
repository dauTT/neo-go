package transaction

import (
	"errors"

	"github.com/CityOfZion/neo-go/pkg/wire/payload/transaction/types"
	"github.com/CityOfZion/neo-go/pkg/wire/payload/transaction/version"
	"github.com/CityOfZion/neo-go/pkg/wire/util"
	"github.com/CityOfZion/neo-go/pkg/wire/util/fixed8"
)

//Invocation represents an invocation transaction on the neo network
type Invocation struct {
	*Base
	Script []byte
	Gas    fixed8.Fixed8
}

//NewInvocation returns an invocation transaction
func NewInvocation(ver version.TX) *Invocation {
	basicTrans := createBaseTransaction(types.Invocation, ver)

	invocation := &Invocation{}
	invocation.Base = basicTrans
	invocation.encodeExclusive = invocation.encodeExcl
	invocation.decodeExclusive = invocation.decodeExcl
	return invocation
}

func (i *Invocation) encodeExcl(bw *util.BinWriter) {
	bw.VarUint(uint64(len(i.Script)))
	bw.Write(i.Script)

	switch i.Version {
	case 0:
		i.Gas = fixed8.Fixed8(0)
	case 1:
		bw.Write(&i.Gas)
	default:
		bw.Write(&i.Gas)
	}

	return
}

func (i *Invocation) decodeExcl(br *util.BinReader) {

	lenScript := br.VarUint()
	i.Script = make([]byte, lenScript)
	br.Read(&i.Script)

	switch i.Version {
	case 0:
		i.Gas = fixed8.Fixed8(0)
	case 1:
		br.Read(&i.Gas)
	default:
		br.Err = errors.New("invalid Version Number for Invocation Transaction")
	}
	return
}

// BaseTx returns the Base field of the Invocation transaction.
func (i Invocation) BaseTx() *Base {
	return i.Base
}
