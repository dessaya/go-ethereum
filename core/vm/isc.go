package vm

import "math/big"

type ISCMagicContract interface {
	Run(evm *EVM, caller ContractRef, input []byte, value *big.Int, gas uint64, readOnly bool) ([]byte, uint64, error)
}
