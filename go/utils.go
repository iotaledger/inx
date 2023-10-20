package inx

import (
	"github.com/iotaledger/hive.go/ierrors"
	"github.com/iotaledger/hive.go/lo"
	"github.com/iotaledger/hive.go/serializer/v2/serix"
	iotago "github.com/iotaledger/iota.go/v4"
	"github.com/iotaledger/iota.go/v4/nodeclient/apimodels"
)

func blockIDsFromSlice(slice []*BlockId) iotago.BlockIDs {
	result := make([]iotago.BlockID, len(slice))
	for i := range slice {
		result[i] = slice[i].Unwrap()
	}

	return result
}

// Block

func WrapBlock(block *iotago.ProtocolBlock) (*RawBlock, error) {
	bytes, err := block.API.Encode(block)
	if err != nil {
		return nil, err
	}

	return &RawBlock{
		Data: bytes,
	}, nil
}

func (x *RawBlock) UnwrapBlock(apiProvider iotago.APIProvider) (*iotago.ProtocolBlock, error) {
	return lo.DropCount(iotago.ProtocolBlockFromBytes(apiProvider)(x.GetData()))
}

func (x *BlockId) Unwrap() iotago.BlockID {
	if len(x.GetId()) != iotago.BlockIDLength {
		return iotago.EmptyBlockID
	}

	return iotago.BlockID(x.GetId())
}

func (x *Block) UnwrapBlockID() iotago.BlockID {
	return x.GetBlockId().Unwrap()
}

func (x *Block) UnwrapBlock(apiProvider iotago.APIProvider) (*iotago.ProtocolBlock, error) {
	return x.GetBlock().UnwrapBlock(apiProvider)
}

func (x *Block) MustUnwrapBlock(apiProvider iotago.APIProvider) *iotago.ProtocolBlock {
	msg, err := x.GetBlock().UnwrapBlock(apiProvider)
	if err != nil {
		panic(err)
	}

	return msg
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
		return nil, ierrors.New("invalid output length")
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

// Commitment

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

// ProtocolParameters

func WrapProtocolParameters(startEpoch iotago.EpochIndex, params iotago.ProtocolParameters) (*RawProtocolParameters, error) {
	bytes, err := params.Bytes()
	if err != nil {
		return nil, err
	}

	return &RawProtocolParameters{
		ProtocolVersion: uint32(params.Version()),
		StartEpoch:      uint32(startEpoch),
		Params:          bytes,
	}, nil
}

func (x *RawProtocolParameters) Unwrap() (iotago.EpochIndex, iotago.ProtocolParameters, error) {
	params, _, err := iotago.ProtocolParametersFromBytes(x.Params)
	if err != nil {
		return 0, nil, err
	}

	return iotago.EpochIndex(x.StartEpoch), params, nil
}

// BlockMetadata

func WrapBlockState(state apimodels.BlockState) BlockMetadata_BlockState {
	return BlockMetadata_BlockState(state)
}

func WrapBlockFailureReason(reason apimodels.BlockFailureReason) BlockMetadata_BlockFailureReason {
	return BlockMetadata_BlockFailureReason(reason)
}

func WrapTransactionState(state apimodels.TransactionState) BlockMetadata_TransactionState {
	return BlockMetadata_TransactionState(state)
}

func WrapTransactionFailureReason(reason apimodels.TransactionFailureReason) BlockMetadata_TransactionFailureReason {
	return BlockMetadata_TransactionFailureReason(reason)
}

func (x BlockMetadata_BlockState) Unwrap() apimodels.BlockState {
	return apimodels.BlockState(x)
}

func (x BlockMetadata_TransactionState) Unwrap() apimodels.TransactionState {
	return apimodels.TransactionState(x)
}

func (x BlockMetadata_BlockFailureReason) Unwrap() apimodels.BlockFailureReason {
	return apimodels.BlockFailureReason(x)
}

func (x BlockMetadata_TransactionFailureReason) Unwrap() apimodels.TransactionFailureReason {
	return apimodels.TransactionFailureReason(x)
}

// Block Issuance

func (x *TipsResponse) UnwrapStrongTips() iotago.BlockIDs {
	return blockIDsFromSlice(x.GetStrongTips())
}

func (x *TipsResponse) UnwrapWeakTips() iotago.BlockIDs {
	return blockIDsFromSlice(x.GetWeakTips())
}

func (x *TipsResponse) UnwrapShallowLikeTips() iotago.BlockIDs {
	return blockIDsFromSlice(x.GetShallowLikeTips())
}

// Payload

func WrapPayload(block iotago.BlockPayload, api iotago.API) (*RawPayload, error) {
	bytes, err := api.Encode(block)
	if err != nil {
		return nil, err
	}

	return &RawPayload{
		Data: bytes,
	}, nil
}

func (x *RawPayload) Unwrap(api iotago.API, opts ...serix.Option) (iotago.BlockPayload, error) {
	var payload iotago.BlockPayload
	if _, err := api.Decode(x.GetData(), &payload, opts...); err != nil {
		return nil, err
	}

	return payload, nil
}

// BoolResponse

func WrapBoolResponse(value bool) *BoolResponse {
	return &BoolResponse{Value: value}
}
