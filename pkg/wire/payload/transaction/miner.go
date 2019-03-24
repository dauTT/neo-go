package transaction

import (
	"github.com/CityOfZion/neo-go/pkg/wire/payload/transaction/types"
	"github.com/CityOfZion/neo-go/pkg/wire/payload/transaction/version"
	"github.com/CityOfZion/neo-go/pkg/wire/util"
)

//Miner represents a miner transaction on the neo network
type Miner struct {
	*Base
	Nonce uint32
}

//NewMiner returns a miner transaction
func NewMiner(ver version.TX) *Miner {
	basicTrans := createBaseTransaction(types.Miner, ver)

	Miner := &Miner{}
	Miner.Base = basicTrans
	Miner.encodeExclusive = Miner.encodeExcl
	Miner.decodeExclusive = Miner.decodeExcl
	return Miner
}

func (m *Miner) encodeExcl(bw *util.BinWriter) {
	bw.Write(m.Nonce)
}

func (m *Miner) decodeExcl(br *util.BinReader) {
	br.Read(&m.Nonce)
}

// BaseTx returns the Base field of the Miner transaction.
func (m Miner) BaseTx() *Base {
	return m.Base
}
