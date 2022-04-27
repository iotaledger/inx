// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::Error;
use crate::proto;

use bee_message_stardust as stardust;
use packable::PackableExt;

/// The [`MilestoneInfo`] type.
#[derive(PartialEq, Debug)]
pub struct MilestoneInfo {
    /// The [`MilestoneId`](stardust::payload::milestone::MilestoneId) of the milestone.
    pub milestone_id: stardust::payload::milestone::MilestoneId,
    /// The milestone index.
    pub milestone_index: u32,
    /// The timestamp of the milestone.
    pub milestone_timestamp: u32,
}

/// The [`Milestone`] type.
#[derive(PartialEq, Debug)]
pub struct Milestone {
    /// Information about the milestone.
    pub milestone_info: MilestoneInfo,
    /// The raw bytes of the milestone.
    pub milestone: stardust::payload::MilestonePayload,
}

impl TryFrom<proto::MilestoneId> for stardust::payload::milestone::MilestoneId {
    type Error = Error;

    fn try_from(value: proto::MilestoneId) -> Result<Self, Self::Error> {
        let bytes: [u8; stardust::payload::milestone::MilestoneId::LENGTH] =
            value.id.try_into().map_err(|_| Error::InvalidBufferLength)?;
        Ok(stardust::payload::milestone::MilestoneId::from(bytes))
    }
}

impl TryFrom<proto::MilestoneInfo> for MilestoneInfo {
    type Error = Error;

    fn try_from(value: proto::MilestoneInfo) -> Result<Self, Self::Error> {
        Ok(MilestoneInfo {
            milestone_id: value
                .milestone_id
                .ok_or(Error::MissingField("milestone_id"))?
                .try_into()?,
            milestone_index: value.milestone_index,
            milestone_timestamp: value.milestone_timestamp,
        })
    }
}

impl TryFrom<proto::RawMilestone> for stardust::payload::MilestonePayload {
    type Error = Error;

    fn try_from(value: proto::RawMilestone) -> Result<Self, Self::Error> {
        stardust::payload::MilestonePayload::unpack_verified(value.data)
            .map_err(|e| Error::PackableError(format!("{e}")))
    }
}

impl TryFrom<proto::Milestone> for Milestone {
    type Error = Error;

    fn try_from(value: proto::Milestone) -> Result<Self, Self::Error> {
        Ok(Milestone {
            milestone_info: value
                .milestone_info
                .ok_or(Error::MissingField("milestone_info"))?
                .try_into()?,
            milestone: value.milestone.ok_or(Error::MissingField("milestone"))?.try_into()?,
        })
    }
}
