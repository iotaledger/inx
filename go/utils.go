package inx

import (
	"github.com/pkg/errors"

	"github.com/iotaledger/hive.go/serializer/v2/serix"
	iotago "github.com/iotaledger/iota.go/v4"
)

func blockIDsFromSlice(slice []*BlockId) iotago.BlockIDs {
	result := make([]iotago.BlockID, len(slice))
	for i := range slice {
		result[i] = slice[i].Unwrap()
	}

	return result
}

// Block

func WrapBlock(block *iotago.Block, api iotago.API) (*RawBlock, error) {
	bytes, err := api.Encode(block)
	if err != nil {
		return nil, err
	}

	return &RawBlock{
		Data: bytes,
	}, nil
}

func (x *RawBlock) UnwrapBlock(api iotago.API, opts ...serix.Option) (*iotago.Block, error) {
	block := &iotago.Block{}
	if _, err := api.Decode(x.GetData(), block, opts...); err != nil {
		return nil, err
	}

	return block, nil
}

func (x *BlockId) Unwrap() iotago.BlockID {
	if len(x.GetId()) != iotago.BlockIDLength {
		return iotago.EmptyBlockID()
	}

	return iotago.BlockID(x.GetId())
}

func (x *Block) UnwrapBlockID() iotago.BlockID {
	return x.GetBlockId().Unwrap()
}

func (x *Block) UnwrapBlock(api iotago.API, opts ...serix.Option) (*iotago.Block, error) {
	return x.GetBlock().UnwrapBlock(api, opts...)
}

func (x *Block) MustUnwrapBlock(api iotago.API, opts ...serix.Option) *iotago.Block {
	msg, err := x.GetBlock().UnwrapBlock(api, opts...)
	if err != nil {
		panic(err)
	}

	return msg
}

func (x *BlockMetadata) UnwrapBlockID() iotago.BlockID {
	return x.GetBlockId().Unwrap()
}

func (x *BlockMetadata) UnwrapParents() iotago.BlockIDs {
	return blockIDsFromSlice(x.GetParents())
}

func (x *BlockWithMetadata) UnwrapBlock(api iotago.API, opts ...serix.Option) (*iotago.Block, error) {
	return x.GetBlock().UnwrapBlock(api, opts...)
}

// Ledger

func (x *TransactionId) Unwrap() iotago.TransactionID {
	id := iotago.TransactionID{}
	if len(x.GetId()) != iotago.TransactionIDLength {
		return iotago.TransactionID{}
	}
	copy(id[:], x.GetId())

	return id
}

func (x *OutputId) Unwrap() iotago.OutputID {
	id := iotago.OutputID{}
	if len(x.GetId()) != iotago.OutputIDLength {
		return iotago.OutputID{}
	}
	copy(id[:], x.GetId())

	return id
}

func (x *LedgerOutput) UnwrapOutputID() iotago.OutputID {
	return x.OutputId.Unwrap()
}

func (x *LedgerOutput) UnwrapBlockID() iotago.BlockID {
	return x.BlockId.Unwrap()
}

func (x *LedgerOutput) UnwrapOutput(api iotago.API, opts ...serix.Option) (iotago.Output, error) {
	return x.GetOutput().Unwrap(api, opts...)
}

func (x *LedgerOutput) MustUnwrapOutput(api iotago.API, opts ...serix.Option) iotago.Output {
	output, err := x.UnwrapOutput(api, opts...)
	if err != nil {
		panic(err)
	}

	return output
}

func WrapOutput(output iotago.Output, api iotago.API) (*RawOutput, error) {
	bytes, err := api.Encode(output)
	if err != nil {
		return nil, err
	}

	return &RawOutput{
		Data: bytes,
	}, nil
}

func (x *RawOutput) Unwrap(api iotago.API, opts ...serix.Option) (iotago.Output, error) {
	data := x.GetData()
	if len(data) == 0 {
		return nil, errors.New("invalid output length")
	}

	var output iotago.TxEssenceOutput
	_, err := api.Decode(data, &output, opts...)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (x *LedgerSpent) UnwrapTransactionIDSpent() iotago.TransactionID {
	return x.GetTransactionIdSpent().Unwrap()
}

// Milestones

func (x *CommitmentId) Unwrap() iotago.CommitmentID {
	if len(x.GetId()) != iotago.CommitmentIDLength {
		return iotago.CommitmentID{}
	}

	return iotago.CommitmentID(x.GetId())
}

func (x *Commitment) UnwrapCommitment(api iotago.API, opts ...serix.Option) (*iotago.Commitment, error) {
	return x.GetCommitment().Unwrap(api, opts...)
}

func (x *RawCommitment) Unwrap(api iotago.API, opts ...serix.Option) (*iotago.Commitment, error) {
	commitment := &iotago.Commitment{}
	_, err := api.Decode(x.GetData(), commitment, opts...)
	if err != nil {
		return nil, err
	}

	return commitment, nil
}

// Tips

func (x *TipsResponse) UnwrapTips() iotago.BlockIDs {
	return blockIDsFromSlice(x.GetTips())
}
