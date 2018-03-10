package transaction

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/CityOfZion/neo-go/pkg/util"
	"github.com/stretchr/testify/assert"
)

// Source of this TX: https://neotracker.io/tx/2c6a45547b3898318e400e541628990a07acb00f3b9a15a8e966ae49525304da
var rawTXClaim = "020004bc67ba325d6412ff4c55b10f7e9afb54bbb2228d201b37363c3d697ac7c198f70300591cd454d7318d2087c0196abfbbd1573230380672f0f0cd004dcb4857e58cbd010031bcfbed573f5318437e95edd603922a4455ff3326a979fdd1c149a84c4cb0290000b51eb6159c58cac4fe23d90e292ad2bcb7002b0da2c474e81e1889c0649d2c490000000001e72d286979ee6cb1b7e65dfddfb2e384100b8d148e7758de42e4168b71792c603b555f00000000005d9de59d99c0d1f6ed1496444473f4a0b538302f014140456349cec43053009accdb7781b0799c6b591c812768804ab0a0b56b5eae7a97694227fcd33e70899c075848b2cee8fae733faac6865b484d3f7df8949e2aadb232103945fae1ed3c31d778f149192b76734fcc951b400ba3598faa81ff92ebe477eacac"

func TestDecodeEncodeClaimTX(t *testing.T) {
	b, err := hex.DecodeString(rawTXClaim)
	if err != nil {
		t.Fatal(err)
	}
	tx := &Transaction{}
	if err := tx.DecodeBinary(bytes.NewReader(b)); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, tx.Type, ClaimType)
	assert.IsType(t, tx.Data, &ClaimTX{})
	claimTX := tx.Data.(*ClaimTX)
	assert.Equal(t, 4, len(claimTX.Claims))
	assert.Equal(t, 0, len(tx.Attributes))
	assert.Equal(t, 0, len(tx.Inputs))
	assert.Equal(t, 1, len(tx.Outputs))
	assert.Equal(t, "AQJseD8iBmCD4sgfHRhMahmoi9zvopG6yz", tx.Outputs[0].ScriptHash.Address())
	assert.Equal(t, "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7", tx.Outputs[0].AssetID.String())
	assert.Equal(t, tx.Outputs[0].Amount.String(), "0.06247739")
	invoc := "40456349cec43053009accdb7781b0799c6b591c812768804ab0a0b56b5eae7a97694227fcd33e70899c075848b2cee8fae733faac6865b484d3f7df8949e2aadb"
	verif := "2103945fae1ed3c31d778f149192b76734fcc951b400ba3598faa81ff92ebe477eacac"
	assert.Equal(t, 1, len(tx.Scripts))
	assert.Equal(t, invoc, hex.EncodeToString(tx.Scripts[0].InvocationScript))
	assert.Equal(t, verif, hex.EncodeToString(tx.Scripts[0].VerificationScript))

	buf := new(bytes.Buffer)
	if err := tx.EncodeBinary(buf); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, rawTXClaim, hex.EncodeToString(buf.Bytes()))
}

// Source of this TX: https://neotracker.io/tx/fe4b3af60677204c57e573a57bdc97bc5059b05ad85b1474f84431f88d910f64
var rawTXInvocation = "d101590400b33f7114839c33710da24cf8e7d536b8d244f3991cf565c8146063795d3b9b3cd55aef026eae992b91063db0db53c1087472616e7366657267c5cc1cb5392019e2cc4e6d6b5ea54c8d4b6d11acf166cb072961424c54f6000000000000000001206063795d3b9b3cd55aef026eae992b91063db0db0000014140c6a131c55ca38995402dff8e92ac55d89cbed4b98dfebbcb01acbc01bd78fa2ce2061be921b8999a9ab79c2958875bccfafe7ce1bbbaf1f56580815ea3a4feed232102d41ddce2c97be4c9aa571b8a32cbc305aa29afffbcae71b0ef568db0e93929aaac"

func TestDecodeEncodeInvocationTX(t *testing.T) {
	b, err := hex.DecodeString(rawTXInvocation)
	if err != nil {
		t.Fatal(err)
	}
	tx := &Transaction{}
	if err := tx.DecodeBinary(bytes.NewReader(b)); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, tx.Type, InvocationType)
	assert.IsType(t, tx.Data, &InvocationTX{})

	invocTX := tx.Data.(*InvocationTX)
	script := "0400b33f7114839c33710da24cf8e7d536b8d244f3991cf565c8146063795d3b9b3cd55aef026eae992b91063db0db53c1087472616e7366657267c5cc1cb5392019e2cc4e6d6b5ea54c8d4b6d11acf166cb072961424c54f6"
	assert.Equal(t, script, hex.EncodeToString(invocTX.Script))
	assert.Equal(t, util.Fixed8(0), invocTX.Gas)

	assert.Equal(t, 1, len(tx.Attributes))
	assert.Equal(t, 0, len(tx.Inputs))
	assert.Equal(t, 0, len(tx.Outputs))
	invoc := "40c6a131c55ca38995402dff8e92ac55d89cbed4b98dfebbcb01acbc01bd78fa2ce2061be921b8999a9ab79c2958875bccfafe7ce1bbbaf1f56580815ea3a4feed"
	verif := "2102d41ddce2c97be4c9aa571b8a32cbc305aa29afffbcae71b0ef568db0e93929aaac"
	assert.Equal(t, 1, len(tx.Scripts))
	assert.Equal(t, invoc, hex.EncodeToString(tx.Scripts[0].InvocationScript))
	assert.Equal(t, verif, hex.EncodeToString(tx.Scripts[0].VerificationScript))

	buf := new(bytes.Buffer)
	if err := tx.EncodeBinary(buf); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, rawTXInvocation, hex.EncodeToString(buf.Bytes()))
}