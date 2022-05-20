use std::error::Error;

use futures::StreamExt;
use inx::{
    client::InxClient,
    proto::{MessageFilter, MilestoneRequest, NoParams},
    tonic::{Channel, Status}, Milestone,
};
use tokio;

const INX_ADDRESS: &str = "localhost:9029";

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let inx = InxClient::connect(INX_ADDRESS).await?;

    let milestone_stream = inx.listen_to_confirmed_milestone(NoParams {}).await?.into_inner();

    while let Some(Ok(proto_milestone)) = milestone_stream.next().await {

        let milestone: Milestone = proto_milestone.try_into()?;

        let request = MilestoneRequest::MilestoneIndex(milestone.milestone_info.milestone_index);

        let cone_stream = inx.read_milestone_cone(request.into()).await?;


    }

    Ok(())
}
