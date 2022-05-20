package inx

import (
	"github.com/pkg/errors"

	"github.com/iotaledger/hive.go/serializer/v2"
	iotago "github.com/iotaledger/iota.go/v3"
)

func blockIDsFromSlice(slice []*BlockId) iotago.BlockIDs {
	result := make([]iotago.BlockID, len(slice))
	for i := range slice {
		result[i] = slice[i].Unwrap()
	}
	return result
}

// Node

func (x *NodeConfiguration) UnwrapProtocolParameters() *iotago.ProtocolParameters {
	return x.GetProtocolParameters().Unwrap()
}

func (x *ProtocolParameters) Unwrap() *iotago.ProtocolParameters {
	return &iotago.ProtocolParameters{
		Version:       byte(x.GetVersion()),
		NetworkName:   x.GetNetworkName(),
		Bech32HRP:     iotago.NetworkPrefix(x.GetBech32HRP()),
		MinPoWScore:   float64(x.GetMinPoWScore()),
		BelowMaxDepth: uint16(x.GetBelowMaxDepth()),
		RentStructure: iotago.RentStructure{
			VByteCost:    x.GetRentStructure().GetVByteCost(),
			VBFactorData: iotago.VByteCostFactor(x.GetRentStructure().GetVByteFactorData()),
			VBFactorKey:  iotago.VByteCostFactor(x.GetRentStructure().GetVByteFactorKey()),
		},
		TokenSupply: x.GetTokenSupply(),
	}
}

// Block

func WrapBlock(msg *iotago.Block) (*RawBlock, error) {
	bytes, err := msg.Serialize(serializer.DeSeriModeNoValidation, nil)
	if err != nil {
		return nil, err
	}
	return &RawBlock{
		Data: bytes,
	}, nil
}

func (x *RawBlock) UnwrapBlock(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (*iotago.Block, error) {
	msg := &iotago.Block{}
	if _, err := msg.Deserialize(x.GetData(), deSeriMode, protoParas); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *BlockId) Unwrap() iotago.BlockID {
	id := iotago.BlockID{}
	if len(x.GetId()) != iotago.BlockIDLength {
		return id
	}
	copy(id[:], x.GetId())
	return id
}

func (x *Block) UnwrapBlockID() iotago.BlockID {
	return x.GetBlockId().Unwrap()
}

func (x *Block) UnwrapBlock(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (*iotago.Block, error) {
	return x.GetBlock().UnwrapBlock(deSeriMode, protoParas)
}

func (x *Block) MustUnwrapBlock(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) *iotago.Block {
	msg, err := x.GetBlock().UnwrapBlock(deSeriMode, protoParas)
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

func (x *BlockWithMetadata) UnwrapBlock(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (*iotago.Block, error) {
	return x.GetBlock().UnwrapBlock(deSeriMode, protoParas)
}

// Ledger

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

func (x *LedgerOutput) UnwrapOutput(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (iotago.Output, error) {
	return x.GetOutput().Unwrap(deSeriMode, protoParas)
}

func (x *LedgerOutput) MustUnwrapOutput(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) iotago.Output {
	output, err := x.UnwrapOutput(deSeriMode, protoParas)
	if err != nil {
		panic(err)
	}
	return output
}

func WrapOutput(output iotago.Output) (*RawOutput, error) {
	bytes, err := output.Serialize(serializer.DeSeriModeNoValidation, nil)
	if err != nil {
		return nil, err
	}
	return &RawOutput{
		Data: bytes,
	}, nil
}

func (x *RawOutput) Unwrap(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (iotago.Output, error) {
	data := x.GetData()
	if len(data) == 0 {
		return nil, errors.New("invalid output length")
	}

	output, err := iotago.OutputSelector(uint32(data[0]))
	if err != nil {
		return nil, err
	}

	_, err = output.Deserialize(data, deSeriMode, protoParas)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (x *LedgerSpent) UnwrapTransactionIDSpent() iotago.TransactionID {
	id := iotago.TransactionID{}
	if len(x.GetTransactionIdSpent()) != iotago.TransactionIDLength {
		return id
	}
	copy(id[:], x.GetTransactionIdSpent())
	return id
}

func (x *TreasuryOutput) UnwrapMilestoneID() iotago.MilestoneID {
	return x.GetMilestoneId().Unwrap()
}

// Milestones

func (x *MilestoneId) Unwrap() iotago.MilestoneID {
	id := iotago.MilestoneID{}
	if len(x.GetId()) != iotago.MilestoneIDLength {
		return id
	}
	copy(id[:], x.GetId())
	return id
}

func (x *Milestone) UnwrapMilestone(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (*iotago.Milestone, error) {
	return x.GetMilestone().Unwrap(deSeriMode, protoParas)
}

func (x *RawMilestone) Unwrap(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (*iotago.Milestone, error) {
	milestone := &iotago.Milestone{}
	_, err := milestone.Deserialize(x.GetData(), deSeriMode, protoParas)
	if err != nil {
		return nil, err
	}
	return milestone, nil
}

func WrapReceipt(receipt *iotago.ReceiptMilestoneOpt) (*RawReceipt, error) {
	bytes, err := receipt.Serialize(serializer.DeSeriModeNoValidation, nil)
	if err != nil {
		return nil, err
	}
	return &RawReceipt{
		Data: bytes,
	}, nil
}

func (x *RawReceipt) UnwrapReceipt(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (*iotago.ReceiptMilestoneOpt, error) {
	r := &iotago.ReceiptMilestoneOpt{}
	if _, err := r.Deserialize(x.GetData(), deSeriMode, protoParas); err != nil {
		return nil, err
	}
	return r, nil
}

func (x *WhiteFlagRequest) UnwrapParents() iotago.BlockIDs {
	return blockIDsFromSlice(x.GetParents())
}

// Tips

func (x *TipsResponse) UnwrapTips() iotago.BlockIDs {
	return blockIDsFromSlice(x.GetTips())
}
