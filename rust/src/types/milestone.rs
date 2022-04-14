// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::proto;
use super::Error;

use bee_message_stardust as stardust;

/// The [`Milestone`] type.
#[derive(PartialEq, Debug)]
pub struct Milestone {
    /// The milestone index.
    pub milestone_index: u32,
    /// The timestamp of the milestone.
    pub milestone_timestamp: u32,
    /// The [`MessageId`](stardust::MessageId) of the milestone.
    pub message_id: stardust::MessageId,
    /// The [`MilestoneId`](stardust::payload::milestone::MilestoneId) of the milestone.
    pub milestone_id: stardust::payload::milestone::MilestoneId,
}

impl TryFrom<proto::MilestoneId> for stardust::payload::milestone::MilestoneId {
    type Error = Error;

    fn try_from(value: proto::MilestoneId) -> Result<Self, Self::Error> {
        let bytes: [u8; stardust::payload::milestone::MilestoneId::LENGTH] =
            value.id.try_into().map_err(|_| Error::InvalidBufferLength)?;
        Ok(stardust::payload::milestone::MilestoneId::from(bytes))
    }
}

impl TryFrom<proto::Milestone> for Milestone {
    type Error = Error;

    fn try_from(value: proto::Milestone) -> Result<Self, Self::Error> {
        Ok(Milestone {
            milestone_index: value.milestone_index,
            milestone_timestamp: value.milestone_timestamp,
            message_id: value.message_id.ok_or(Error::MissingField("message_id"))?.try_into()?,
            milestone_id: value
                .milestone_id
                .ok_or(Error::MissingField("milestone_id"))?
                .try_into()?,
        })
    }
}
