package inx

import (
	iotago "github.com/iotaledger/iota.go/v4"
)

//nolint:revive,stylecheck
func NewBlockId(blockID iotago.BlockID) *BlockId {
	id := &BlockId{
		Id: make([]byte, iotago.BlockIDLength),
	}
	copy(id.GetId(), blockID[:])

	return id
}

//nolint:revive
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
	copy(m.GetBlock().GetData(), data)

	return m
}

//nolint:revive,stylecheck // this name is auto generated
func NewTransactionId(transactionID iotago.TransactionID) *TransactionId {
	id := &TransactionId{
		Id: make([]byte, iotago.TransactionIDLength),
	}
	copy(id.GetId(), transactionID[:])

	return id
}

//nolint:revive,stylecheck // this name is auto generated
func NewOutputId(outputID iotago.OutputID) *OutputId {
	id := &OutputId{
		Id: make([]byte, iotago.OutputIDLength),
	}
	copy(id.GetId(), outputID[:])

	return id
}

//nolint:revive,stylecheck // this name is auto generated
func NewCommitmentId(commitmentID iotago.CommitmentID) *CommitmentId {
	id := &CommitmentId{
		Id: make([]byte, iotago.CommitmentIDLength),
	}
	copy(id.GetId(), commitmentID[:])

	return id
}

func NewAccountInfoRequest(accountID iotago.AccountID, slot iotago.SlotIndex) *AccountInfoRequest {
	accountInfoRequest := &AccountInfoRequest{
		AccountId:   make([]byte, iotago.AccountIDLength),
		AccountSlot: uint32(slot),
	}
	copy(accountInfoRequest.GetAccountId(), accountID[:])

	return accountInfoRequest
}

func NewCommitmentWithBytes(commitmentID iotago.CommitmentID, data []byte) *Commitment {
	c := &Commitment{
		CommitmentId: NewCommitmentId(commitmentID),
		Commitment: &RawCommitment{
			Data: make([]byte, len(data)),
		},
	}
	copy(c.GetCommitment().GetData(), data)

	return c
}
