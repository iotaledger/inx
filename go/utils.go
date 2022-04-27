package inx

import (
	"github.com/pkg/errors"

	"github.com/iotaledger/hive.go/serializer/v2"
	iotago "github.com/iotaledger/iota.go/v3"
)

// Node

func (x *NodeConfiguration) UnwrapProtocolParameters() *iotago.ProtocolParameters {
	return x.GetProtocolParameters().Unwrap()
}

func (x *ProtocolParameters) Unwrap() *iotago.ProtocolParameters {
	return &iotago.ProtocolParameters{
		Version:     byte(x.GetVersion()),
		NetworkName: x.GetNetworkName(),
		Bech32HRP:   iotago.NetworkPrefix(x.GetBech32HRP()),
		MinPoWScore: float64(x.GetMinPoWScore()),
		RentStructure: iotago.RentStructure{
			VByteCost:    x.GetRentStructure().GetVByteCost(),
			VBFactorData: iotago.VByteCostFactor(x.GetRentStructure().GetVByteFactorData()),
			VBFactorKey:  iotago.VByteCostFactor(x.GetRentStructure().GetVByteFactorKey()),
		},
		TokenSupply: x.GetTokenSupply(),
	}
}

// Message

func WrapMessage(msg *iotago.Message) (*RawMessage, error) {
	bytes, err := msg.Serialize(serializer.DeSeriModeNoValidation, nil)
	if err != nil {
		return nil, err
	}
	return &RawMessage{
		Data: bytes,
	}, nil
}

func (x *RawMessage) UnwrapMessage(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (*iotago.Message, error) {
	msg := &iotago.Message{}
	if _, err := msg.Deserialize(x.GetData(), deSeriMode, protoParas); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *MessageId) Unwrap() iotago.MessageID {
	id := iotago.MessageID{}
	if len(x.GetId()) != iotago.MessageIDLength {
		return id
	}
	copy(id[:], x.GetId())
	return id
}

func (x *Message) UnwrapMessageID() iotago.MessageID {
	return x.GetMessageId().Unwrap()
}

func (x *Message) UnwrapMessage(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (*iotago.Message, error) {
	return x.GetMessage().UnwrapMessage(deSeriMode, protoParas)
}

func (x *Message) MustUnwrapMessage(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) *iotago.Message {
	msg, err := x.GetMessage().UnwrapMessage(deSeriMode, protoParas)
	if err != nil {
		panic(err)
	}
	return msg
}

func (x *MessageMetadata) UnwrapMessageID() iotago.MessageID {
	return x.GetMessageId().Unwrap()
}

// Ledger

func (x *OutputId) Unwrap() *iotago.OutputID {
	if len(x.GetId()) != iotago.OutputIDLength {
		return nil
	}
	id := &iotago.OutputID{}
	copy(id[:], x.GetId())
	return id
}

func (x *LedgerOutput) UnwrapOutputID() *iotago.OutputID {
	return x.OutputId.Unwrap()
}

func (x *LedgerOutput) UnwrapMessageID() iotago.MessageID {
	return x.MessageId.Unwrap()
}

func (x *LedgerOutput) UnwrapOutput(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) (iotago.Output, error) {
	data := x.GetOutput()
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

func (x *LedgerOutput) MustUnwrapOutput(deSeriMode serializer.DeSerializationMode, protoParas *iotago.ProtocolParameters) iotago.Output {
	output, err := x.UnwrapOutput(deSeriMode, protoParas)
	if err != nil {
		panic(err)
	}
	return output
}

func (x *LedgerSpent) UnwrapTransactionIDSpent() *iotago.TransactionID {
	if len(x.GetTransactionIdSpent()) != iotago.TransactionIDLength {
		return nil
	}
	id := &iotago.TransactionID{}
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
