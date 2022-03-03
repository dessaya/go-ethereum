package vm

import "github.com/ethereum/go-ethereum/common"

// ISCAddress is the address of the special ISC system contract
var ISCAddress = common.HexToAddress("0x1074")

type ISCContract interface {
	Run(evm *EVM, caller ContractRef, input []byte, gas uint64, readOnly bool) ([]byte, uint64)
}
