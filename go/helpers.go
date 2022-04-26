package inx

import (
	iotago "github.com/iotaledger/iota.go/v3"
)

func NewMessageId(messageID iotago.MessageID) *MessageId {
	id := &MessageId{
		Id: make([]byte, len(messageID)),
	}
	copy(id.Id, messageID[:])
	return id
}

func NewMessageWithBytes(messageID iotago.MessageID, data []byte) *Message {
	m := &Message{
		MessageId: NewMessageId(messageID),
		Message: &RawMessage{
			Data: make([]byte, len(data)),
		},
	}
	copy(m.Message.Data, data)
	return m
}

func NewOutputId(outputID *iotago.OutputID) *OutputId {
	id := &OutputId{
		Id: make([]byte, len(outputID)),
	}
	copy(id.Id, outputID[:])
	return id
}

func NewMilestoneId(milestoneID iotago.MilestoneID) *MilestoneId {
	id := &MilestoneId{
		Id: make([]byte, len(milestoneID)),
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
		Version:     uint32(protoParas.Version),
		NetworkName: protoParas.NetworkName,
		Bech32HRP:   string(protoParas.Bech32HRP),
		MinPowScore: float32(protoParas.MinPowScore),
		RentStructure: &RentStructure{
			VByteCost:       protoParas.RentStructure.VByteCost,
			VByteFactorData: uint64(protoParas.RentStructure.VBFactorData),
			VByteFactorKey:  uint64(protoParas.RentStructure.VBFactorKey),
		},
		TokenSupply: protoParas.TokenSupply,
	}
}
