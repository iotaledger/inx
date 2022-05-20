use std::error::Error;

use futures::StreamExt;
use inx::proto::{self};
use tokio;

const INX_ADDRESS: &str = "localhost:9029";

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let mut inx = inx::client::InxClient::connect(INX_ADDRESS).await?;

    // Listen to the milestones from the node.
    let mut milestone_stream = inx
        .listen_to_confirmed_milestone(proto::NoParams {})
        .await?
        .into_inner();

    while let Some(Ok(proto_milestone)) = milestone_stream.next().await {
        // Parse the `inx::proto::Milestone` into an `inx::Milestone`.
        let milestone: inx::Milestone = proto_milestone.try_into()?;

        // Create a request for the current milestone.
        let request = inx::MilestoneRequest::MilestoneIndex(milestone.milestone_info.milestone_index);

        println!("Fetch cone of milestone {}", milestone.milestone_info.milestone_index);

        // Listen to messages in the past cone of a milestone.
        let mut cone_stream = inx
            .read_milestone_cone(proto::MilestoneRequest::from(request))
            .await?
            .into_inner();

        // Keep track of the number of blocks.
        let mut count = 0usize;

        while let Some(Ok(proto_block_metadata)) = cone_stream.next().await {

            let block_metadata: inx::BlockWithMetadata = proto_block_metadata.try_into()?;

            println!("Block {}: {:#?}", block_metadata.metadata.block_id, block_metadata.block);
            count += 1;
        }

        println!("Milestone {} contained {count} blocks", milestone.milestone_info.milestone_id);
    }

    Ok(())
}
