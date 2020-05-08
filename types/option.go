package types

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ContractDeployOption for setting option when deploying contract
type ContractDeployOption UnsignedTransactionBase

// ContractMethodCallOption for setting option when call contract method
type ContractMethodCallOption struct {
	From         *Address
	Nonce        *hexutil.Big
	GasPrice     *hexutil.Big
	Gas          *hexutil.Big
	Value        *hexutil.Big
	StorageLimit *hexutil.Big
	ChainID      *hexutil.Big
	Epoch        *Epoch
}

// ContractMethodSendOption for setting option when call contract method
type ContractMethodSendOption UnsignedTransactionBase

// CallRequest represents a request to execute contract.
type CallRequest struct {
	From         *Address     `json:"from,omitempty"`
	To           *Address     `json:"to,omitempty"`
	GasPrice     *hexutil.Big `json:"gasPrice,omitempty"`
	Gas          *hexutil.Big `json:"gas,omitempty"`
	Value        *hexutil.Big `json:"value,omitempty"`
	Data         string       `json:"data,omitempty"`
	Nonce        *hexutil.Big `json:"nonce,omitempty"`
	StorageLimit *hexutil.Big `json:"storageLimit,omitempty"`
}

// FillByUnsignedTx fills CallRequest fields by tx
func (cq *CallRequest) FillByUnsignedTx(tx *UnsignedTransaction) {
	if tx != nil {
		cq.From = tx.From
		cq.To = tx.To
		cq.GasPrice = tx.GasPrice
		cq.Value = tx.Value
		cq.StorageLimit = tx.StorageLimit

		if tx.Gas != nil {
			cq.Gas = tx.Gas
		}

		_data := "0x" + hex.EncodeToString(tx.Data)
		cq.Data = _data

		if tx.Nonce != nil {
			cq.Nonce = tx.Nonce
		}
	}
}

// FillByCallOption fills CallRequest fields by
func (cq *CallRequest) FillByCallOption(option *ContractMethodCallOption) {
	if option != nil {
		cq.From = option.From
		cq.GasPrice = option.GasPrice
		cq.Gas = option.Gas
		cq.Value = option.Value
		cq.Nonce = option.Nonce
		cq.StorageLimit = option.StorageLimit
	}
}
