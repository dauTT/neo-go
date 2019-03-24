package transaction

import (
	"github.com/CityOfZion/neo-go/pkg/wire/payload/transaction/types"
	"github.com/CityOfZion/neo-go/pkg/wire/payload/transaction/version"
	"github.com/CityOfZion/neo-go/pkg/wire/util"
)

//Contract represents a contract transaction on the neo network
type Contract struct {
	*Base
}

//NewContract returns a contract transaction
func NewContract(ver version.TX) *Contract {
	basicTrans := createBaseTransaction(types.Contract, ver)

	contract := &Contract{
		basicTrans,
	}
	contract.encodeExclusive = contract.encodeExcl
	contract.decodeExclusive = contract.decodeExcl
	return contract
}

func (c *Contract) encodeExcl(bw *util.BinWriter) {}

func (c *Contract) decodeExcl(br *util.BinReader) {}

// BaseTx returns the Base field of the Contract.
func (c Contract) BaseTx() *Base {
	return c.Base
}
