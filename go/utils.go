package inx

import (
	"github.com/pkg/errors"

	"github.com/iotaledger/hive.go/serializer/v2"
	iotago "github.com/iotaledger/iota.go/v3"
)

// Node

func (x *ProtocolParameters) NetworkPrefix() iotago.NetworkPrefix {
	return iotago.NetworkPrefix(x.GetBech32HRP())
}

// Message

func WrapMessage(msg *iotago.Message) (*RawMessage, error) {
	bytes, err := msg.Serialize(serializer.DeSeriModeNoValidation, iotago.ZeroRentParas)
	if err != nil {
		return nil, err
	}
	return &RawMessage{
		Data: bytes,
	}, nil
}

func (x *RawMessage) UnwrapMessage(deSeriMode serializer.DeSerializationMode) (*iotago.Message, error) {
	msg := &iotago.Message{}
	if _, err := msg.Deserialize(x.GetData(), deSeriMode, iotago.ZeroRentParas); err != nil {
		return nil, err
	}
	return msg, nil
}

func (x *MessageId) Unwrap() iotago.MessageID {
	id := iotago.MessageID{}
	copy(id[:], x.GetId())
	return id
}

func (x *Message) UnwrapMessageID() iotago.MessageID {
	return x.GetMessageId().Unwrap()
}

func (x *Message) UnwrapMessage(deSeriMode serializer.DeSerializationMode) (*iotago.Message, error) {
	return x.GetMessage().UnwrapMessage(deSeriMode)
}

func (x *Message) MustUnwrapMessage(deSeriMode serializer.DeSerializationMode) *iotago.Message {
	msg, err := x.GetMessage().UnwrapMessage(deSeriMode)
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

func (x *LedgerOutput) UnwrapOutput(deSeriMode serializer.DeSerializationMode) (iotago.Output, error) {
	data := x.GetOutput()
	if len(data) == 0 {
		return nil, errors.New("invalid output length")
	}

	output, err := iotago.OutputSelector(uint32(data[0]))
	if err != nil {
		return nil, err
	}

	_, err = output.Deserialize(data, deSeriMode, iotago.ZeroRentParas)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (x *LedgerOutput) MustUnwrapOutput(deSeriMode serializer.DeSerializationMode) iotago.Output {
	output, err := x.UnwrapOutput(deSeriMode)
	if err != nil {
		panic(err)
	}
	return output
}

func (x *LedgerSpent) UnwrapTransactionIDSpent() *iotago.TransactionID {
	id := &iotago.TransactionID{}
	copy(id[:], x.GetTransactionIdSpent())
	return id
}

// Milestones

func WrapReceipt(receipt *iotago.Receipt) (*RawReceipt, error) {
	bytes, err := receipt.Serialize(serializer.DeSeriModeNoValidation, iotago.ZeroRentParas)
	if err != nil {
		return nil, err
	}
	return &RawReceipt{
		Data: bytes,
	}, nil
}

func (x *RawReceipt) UnwrapReceipt(deSeriMode serializer.DeSerializationMode) (*iotago.Receipt, error) {
	r := &iotago.Receipt{}
	if _, err := r.Deserialize(x.GetData(), deSeriMode, iotago.ZeroRentParas); err != nil {
		return nil, err
	}
	return r, nil
}