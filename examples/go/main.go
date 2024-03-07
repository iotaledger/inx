package main

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	inx "github.com/iotaledger/inx/go"
	iotago "github.com/iotaledger/iota.go/v4"
)

const (
	INXAddress = "localhost:9029"
)

func main() {
	// Establish a connection to the node over INX
	conn, err := grpc.Dial(INXAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	// Create a new INX client
	client := inx.NewINXClient(conn)

	config, err := client.ReadNodeConfiguration(context.Background(), &inx.NoParams{})
	if err != nil {
		panic(err)
	}

	// Create a new api provider that uses the protocol parameters of the node
	apiProvider := config.APIProvider()

	// Listen to all commitments
	stream, err := client.ListenToCommitments(context.Background(), &inx.SlotRangeRequest{})
	if err != nil {
		panic(err)
	}
	for {
		payload, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		// Fetch all blocks included by this commitment
		if err := fetchCommitmentBlocks(client, payload.GetCommitmentId().Unwrap(), apiProvider); err != nil {
			panic(err)
		}
	}
}

func fetchCommitmentBlocks(client inx.INXClient, commitmentID iotago.CommitmentID, apiProvider iotago.APIProvider) error {
	slot := commitmentID.Slot()
	fmt.Printf("Fetch accepted blocks included in commitment %d\n", slot)
	req := &inx.SlotRequest{
		Slot: uint32(slot),
	}
	stream, err := client.ReadAcceptedBlocks(context.Background(), req)
	if err != nil {
		return err
	}
	var count int
	for {
		payload, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// We are done
				break
			}
			return err
		}

		// Deserialize the raw bytes into an iota.go Block
		block, err := payload.GetBlock().UnwrapBlock(apiProvider)
		if err != nil {
			return err
		}

		blockID := payload.GetMetadata().GetBlockId().Unwrap()

		// Serialize the Block to JSON
		jsonBlock, err := apiProvider.APIForSlot(blockID.Slot()).JSONEncode(block)
		if err != nil {
			return err
		}

		// Print the JSON representation of the block
		fmt.Printf("Block: %s\n%s\n\n", blockID.ToHex(), string(jsonBlock))
		count++
	}
	fmt.Printf("Commitment %d contained %d accepted blocks\n", slot, count)
	return nil
}
