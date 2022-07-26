package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/iotaledger/hive.go/serializer/v2"
	inx "github.com/iotaledger/inx/go"
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
	stream, err := client.ListenToConfirmedMilestones(context.Background(), &inx.MilestoneRangeRequest{})
	if err != nil {
		panic(err)
	}
	for {
		milestoneAndParams, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		// Fetch all messages included by this milestone
		if err := fetchMilestoneCone(client, milestoneAndParams.GetMilestone()); err != nil {
			panic(err)
		}
	}
}

func fetchMilestoneCone(client inx.INXClient, milestone *inx.Milestone) error {
	index := milestone.GetMilestoneInfo().GetMilestoneIndex()
	fmt.Printf("Fetch cone of milestone %d\n", index)
	req := &inx.MilestoneRequest{
		MilestoneIndex: index,
	}
	stream, err := client.ReadMilestoneCone(context.Background(), req)
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
		block, err := payload.UnwrapBlock(serializer.DeSeriModeNoValidation, nil)
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
