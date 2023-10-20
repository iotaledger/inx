package inx

import (
	iotago "github.com/iotaledger/iota.go/v4"
)

// nolint:revive,stylecheck // this name is auto generated
func NewBlockId(blockID iotago.BlockID) *BlockId {
	id := &BlockId{
		Id: make([]byte, iotago.BlockIDLength),
	}
	copy(id.Id, blockID[:])

	return id
}

func NewBlockIds(blockIDs iotago.BlockIDs) []*BlockId {
	result := make([]*BlockId, len(blockIDs))
	for i := range blockIDs {
		result[i] = NewBlockId(blockIDs[i])
	}

	return result
}

func NewBlockWithBytes(blockID iotago.BlockID, data []byte) *Block {
	m := &Block{
		BlockId: NewBlockId(blockID),
		Block: &RawBlock{
			Data: make([]byte, len(data)),
		},
	}
	copy(m.Block.Data, data)

	return m
}

// nolint:revive,stylecheck // this name is auto generated
func NewTransactionId(transactionID iotago.TransactionID) *TransactionId {
	id := &TransactionId{
		Id: make([]byte, iotago.TransactionIDLength),
	}
	copy(id.Id, transactionID[:])

	return id
}

// nolint:revive,stylecheck // this name is auto generated
func NewOutputId(outputID iotago.OutputID) *OutputId {
	id := &OutputId{
		Id: make([]byte, iotago.OutputIDLength),
	}
	copy(id.Id, outputID[:])

	return id
}

// nolint:revive,stylecheck // this name is auto generated
func NewCommitmentId(commitmentID iotago.CommitmentID) *CommitmentId {
	id := &CommitmentId{
		Id: make([]byte, iotago.CommitmentIDLength),
	}
	copy(id.Id, commitmentID[:])

	return id
}

// nolint:revive,stylecheck // this name is auto generated
func NewAccountInfoRequest(accountID iotago.AccountID, slot iotago.SlotIndex) *AccountInfoRequest {
	accountInfoRequest := &AccountInfoRequest{
		AccountId:   make([]byte, iotago.CommitmentIDLength),
		AccountSlot: uint32(slot),
	}
	copy(accountInfoRequest.AccountId, accountID[:])

	return accountInfoRequest
}

func NewCommitmentWithBytes(commitmentID iotago.CommitmentID, data []byte) *Commitment {
	c := &Commitment{
		CommitmentId: NewCommitmentId(commitmentID),
		Commitment: &RawCommitment{
			Data: make([]byte, len(data)),
		},
	}
	copy(c.Commitment.Data, data)

	return c
}
