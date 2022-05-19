package inx

import (
	iotago "github.com/iotaledger/iota.go/v3"
)

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

func NewOutputId(outputID iotago.OutputID) *OutputId {
	id := &OutputId{
		Id: make([]byte, iotago.OutputIDLength),
	}
	copy(id.Id, outputID[:])
	return id
}

func NewMilestoneId(milestoneID iotago.MilestoneID) *MilestoneId {
	id := &MilestoneId{
		Id: make([]byte, iotago.MilestoneIDLength),
	}
	copy(id.Id, milestoneID[:])
	return id
}

func NewMilestoneInfo(milestoneID iotago.MilestoneID, index uint32, timestamp uint32) *MilestoneInfo {
	return &MilestoneInfo{
		MilestoneId:        NewMilestoneId(milestoneID),
		MilestoneIndex:     index,
		MilestoneTimestamp: timestamp,
	}
}

func NewProtocolParameters(protoParas *iotago.ProtocolParameters) *ProtocolParameters {
	return &ProtocolParameters{
		Version:       uint32(protoParas.Version),
		NetworkName:   protoParas.NetworkName,
		Bech32HRP:     string(protoParas.Bech32HRP),
		MinPoWScore:   float32(protoParas.MinPoWScore),
		BelowMaxDepth: uint32(protoParas.BelowMaxDepth),
		RentStructure: &RentStructure{
			VByteCost:       protoParas.RentStructure.VByteCost,
			VByteFactorData: uint64(protoParas.RentStructure.VBFactorData),
			VByteFactorKey:  uint64(protoParas.RentStructure.VBFactorKey),
		},
		TokenSupply: protoParas.TokenSupply,
	}
}
