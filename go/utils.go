// nolint: nosnakecase // generated GRPC code uses snake case
package inx

import (
	"github.com/iotaledger/hive.go/ierrors"
	"github.com/iotaledger/hive.go/lo"
	"github.com/iotaledger/hive.go/serializer/v2/serix"
	iotago "github.com/iotaledger/iota.go/v4"
	"github.com/iotaledger/iota.go/v4/api"
)

func blockIDsFromSlice(slice []*BlockId) iotago.BlockIDs {
	result := make([]iotago.BlockID, len(slice))
	for i := range slice {
		result[i] = slice[i].Unwrap()
	}

	return result
}

// Block

func WrapBlock(block *iotago.Block) (*RawBlock, error) {
	bytes, err := block.API.Encode(block)
	if err != nil {
		return nil, err
	}

	return &RawBlock{
		Data: bytes,
	}, nil
}

func (x *RawBlock) UnwrapBlock(apiProvider iotago.APIProvider) (*iotago.Block, error) {
	return lo.DropCount(iotago.BlockFromBytes(apiProvider)(x.GetData()))
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

func (x *Block) UnwrapBlock(apiProvider iotago.APIProvider) (*iotago.Block, error) {
	return x.GetBlock().UnwrapBlock(apiProvider)
}

func (x *Block) MustUnwrapBlock(apiProvider iotago.APIProvider) *iotago.Block {
	msg, err := x.GetBlock().UnwrapBlock(apiProvider)
	if err != nil {
		panic(err)
	}

	return msg
}

func WrapRootBlocks(rootBlocks map[iotago.BlockID]iotago.CommitmentID) *RootBlocksResponse {
	rootBlocksWrapped := make([]*RootBlock, 0, len(rootBlocks))
	for blockID, commitmentID := range rootBlocks {
		rootBlocksWrapped = append(rootBlocksWrapped, &RootBlock{
			BlockId:      NewBlockId(blockID),
			CommitmentId: NewCommitmentId(commitmentID),
		})
	}

	return &RootBlocksResponse{
		RootBlocks: rootBlocksWrapped,
	}
}

func (x *RootBlocksResponse) Unwrap() (map[iotago.BlockID]iotago.CommitmentID, error) {
	rootBlockUnwrapped := make(map[iotago.BlockID]iotago.CommitmentID)

	for _, rootBlock := range x.RootBlocks {
		rootBlockUnwrapped[rootBlock.GetBlockId().Unwrap()] = rootBlock.GetCommitmentId().Unwrap()
	}

	return rootBlockUnwrapped, nil
}

func WrapSlotIndex(slot iotago.SlotIndex) *SlotIndex {
	return &SlotIndex{
		Index: uint32(slot),
	}
}

func (x *SlotIndex) Unwrap() iotago.SlotIndex {
	return iotago.SlotIndex(x.GetIndex())
}

// BlockMetadata

func WrapBlockState(state api.BlockState) BlockMetadata_BlockState {
	return BlockMetadata_BlockState(state)
}

func WrapBlockFailureReason(reason api.BlockFailureReason) BlockMetadata_BlockFailureReason {
	return BlockMetadata_BlockFailureReason(reason)
}

func (x BlockMetadata_BlockState) Unwrap() api.BlockState {
	return api.BlockState(x)
}

func (x BlockMetadata_BlockFailureReason) Unwrap() api.BlockFailureReason {
	return api.BlockFailureReason(x)
}

func WrapBlockMetadata(blockMetadata *api.BlockMetadataResponse) (*BlockMetadata, error) {
	return &BlockMetadata{
		BlockId:             NewBlockId(blockMetadata.BlockID),
		BlockState:          WrapBlockState(blockMetadata.BlockState),
		BlockFailureReason:  WrapBlockFailureReason(blockMetadata.BlockFailureReason),
		TransactionMetadata: WrapTransactionMetadata(blockMetadata.TransactionMetadata),
	}, nil
}

func (x *BlockMetadata) Unwrap() (*api.BlockMetadataResponse, error) {
	return &api.BlockMetadataResponse{
		BlockID:             x.GetBlockId().Unwrap(),
		BlockState:          x.GetBlockState().Unwrap(),
		BlockFailureReason:  x.GetBlockFailureReason().Unwrap(),
		TransactionMetadata: x.GetTransactionMetadata().Unwrap(),
	}, nil
}

// TransactionMetadata

func WrapTransactionState(state api.TransactionState) TransactionMetadata_TransactionState {
	return TransactionMetadata_TransactionState(state)
}

func WrapTransactionFailureReason(reason api.TransactionFailureReason) TransactionMetadata_TransactionFailureReason {
	return TransactionMetadata_TransactionFailureReason(reason)
}

func (x TransactionMetadata_TransactionState) Unwrap() api.TransactionState {
	return api.TransactionState(x)
}

func (x TransactionMetadata_TransactionFailureReason) Unwrap() api.TransactionFailureReason {
	return api.TransactionFailureReason(x)
}

func WrapTransactionMetadata(transactionMetadata *api.TransactionMetadataResponse) *TransactionMetadata {
	if transactionMetadata == nil {
		return nil
	}

	return &TransactionMetadata{
		TransactionId:            NewTransactionId(transactionMetadata.TransactionID),
		TransactionState:         WrapTransactionState(transactionMetadata.TransactionState),
		TransactionFailureReason: WrapTransactionFailureReason(transactionMetadata.TransactionFailureReason),
	}
}

func (x *TransactionMetadata) Unwrap() *api.TransactionMetadataResponse {
	if x == nil {
		return nil
	}

	return &api.TransactionMetadataResponse{
		TransactionID:            x.GetTransactionId().Unwrap(),
		TransactionState:         x.GetTransactionState().Unwrap(),
		TransactionFailureReason: x.GetTransactionFailureReason().Unwrap(),
	}
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

func (x *LedgerOutput) UnwrapOutputIDProof(api iotago.API) (*iotago.OutputIDProof, error) {
	return x.GetOutputIdProof().Unwrap(api)
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

func WrapOutputIDProof(proof *iotago.OutputIDProof) (*RawOutputIDProof, error) {
	bytes, err := proof.Bytes()
	if err != nil {
		return nil, err
	}

	return &RawOutputIDProof{
		Data: bytes,
	}, nil
}

func (x *RawOutputIDProof) Unwrap(api iotago.API) (*iotago.OutputIDProof, error) {
	data := x.GetData()
	if len(data) == 0 {
		return nil, ierrors.New("invalid output ID proof length")
	}

	return lo.DropCount(iotago.OutputIDProofFromBytes(api)(data))
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

func (x *NodeConfiguration) APIProvider() *iotago.EpochBasedProvider {
	// Create a new api provider that uses the protocol parameters of the node
	apiProvider := iotago.NewEpochBasedProvider()

	for _, rawParams := range x.GetProtocolParameters() {
		startEpoch, protoParams, err := rawParams.Unwrap()
		if err != nil {
			panic(err)
		}

		apiProvider.AddProtocolParametersAtEpoch(protoParams, startEpoch)
	}

	return apiProvider
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

func WrapPayload(block iotago.ApplicationPayload, api iotago.API) (*RawPayload, error) {
	bytes, err := api.Encode(block)
	if err != nil {
		return nil, err
	}

	return &RawPayload{
		Data: bytes,
	}, nil
}

func (x *RawPayload) Unwrap(api iotago.API, opts ...serix.Option) (iotago.ApplicationPayload, error) {
	var payload iotago.ApplicationPayload
	if _, err := api.Decode(x.GetData(), &payload, opts...); err != nil {
		return nil, err
	}

	return payload, nil
}

// BoolResponse

func WrapBoolResponse(value bool) *BoolResponse {
	return &BoolResponse{Value: value}
}
