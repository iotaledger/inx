package main

import (
	"context"
	"encoding/json"
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

	// Listen to all confirmed milestones
	stream, err := client.ListenToConfirmedCommitments(context.Background(), &inx.CommitmentRangeRequest{})
	if err != nil {
		panic(err)
	}
	for {
		commitmentAndParams, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		api := iotago.V3API(&iotago.ProtocolParameters{})
		protoParams := &iotago.ProtocolParameters{}
		if _, err := api.Decode(commitmentAndParams.GetCurrentProtocolParameters().GetParams(), protoParams); err != nil {
			panic(err)
		}

		// Fetch all messages included by this milestone
		if err := fetchCommitmentCone(client, commitmentAndParams.GetCommitment(), iotago.V3API(protoParams)); err != nil {
			panic(err)
		}
	}
}

func fetchCommitmentCone(client inx.INXClient, commitment *inx.Commitment, api iotago.API) error {
	index := commitment.GetCommitmentInfo().GetCommitmentIndex()
	fmt.Printf("Fetch cone of milestone %d\n", index)
	req := &inx.CommitmentRequest{
		CommitmentIndex: index,
	}
	stream, err := client.ReadCommitmentCone(context.Background(), req)
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
		block, err := payload.UnwrapBlock(api)
		if err != nil {
			return err
		}

		blockID := payload.GetMetadata().UnwrapBlockID()

		// Serialize the Block to JSON
		jsonBlock, err := json.MarshalIndent(block, "", "  ")
		if err != nil {
			return err
		}

		// Print the JSON representation of the block
		fmt.Printf("Block: %s\n%s\n\n", blockID.ToHex(), string(jsonBlock))
		count++
	}
	fmt.Printf("Milestone %d contained %d blocks\n", index, count)
	return nil
}
