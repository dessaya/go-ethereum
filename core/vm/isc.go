package vm

type ISCMagicContract interface {
	Run(evm *EVM, caller ContractRef, input []byte, gas uint64, readOnly bool) ([]byte, uint64, error)
}
