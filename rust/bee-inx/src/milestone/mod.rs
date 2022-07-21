// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

mod range;

use super::Error;
use crate::proto;

use bee_block_stardust as stardust;
use packable::PackableExt;

/// The [`MilestoneInfo`] type.
#[derive(Clone, Debug, PartialEq)]
pub struct MilestoneInfo {
    /// The [`MilestoneId`](stardust::payload::milestone::MilestoneId) of the milestone.
    pub milestone_id: Option<stardust::payload::milestone::MilestoneId>,
    /// The milestone index.
    pub milestone_index: u32,
    /// The timestamp of the milestone.
    pub milestone_timestamp: u32,
}

/// The [`Milestone`] type.
#[derive(Clone, Debug, PartialEq)]
pub struct Milestone {
    /// Information about the milestone.
    pub milestone_info: MilestoneInfo,
    /// The raw bytes of the milestone.
    pub milestone: Option<stardust::payload::MilestonePayload>,
}

// Remove
impl TryFrom<proto::MilestoneId> for stardust::payload::milestone::MilestoneId {
    type Error = Error;

    fn try_from(value: proto::MilestoneId) -> Result<Self, Self::Error> {
        let bytes: [u8; stardust::payload::milestone::MilestoneId::LENGTH] =
            value.id.try_into().map_err(|_| Error::InvalidBufferLength)?;
        Ok(stardust::payload::milestone::MilestoneId::from(bytes))
    }
}

// Remove
impl From<stardust::payload::milestone::MilestoneId> for proto::MilestoneId {
    fn from(value: stardust::payload::milestone::MilestoneId) -> Self {
        proto::MilestoneId {
            id: value.pack_to_vec(),
        }
    }
}

impl TryFrom<proto::MilestoneInfo> for MilestoneInfo {
    type Error = Error;

    fn try_from(value: proto::MilestoneInfo) -> Result<Self, Self::Error> {
        Ok(MilestoneInfo {
            milestone_id: value.milestone_id.map(TryInto::try_into).transpose()?,
            milestone_index: value.milestone_index,
            milestone_timestamp: value.milestone_timestamp,
        })
    }
}

// Remove
impl TryFrom<proto::RawMilestone> for stardust::payload::MilestonePayload {
    type Error = Error;

    fn try_from(value: proto::RawMilestone) -> Result<Self, Self::Error> {
        let payload = stardust::payload::Payload::unpack_verified(value.data)
            .map_err(|e| Error::PackableError(format!("{e}")))?;

        match payload {
            stardust::payload::Payload::Milestone(payload) => Ok(*payload),
            _ => Err(Error::InvalidField("milestone (wrong payload type)")),
        }
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
            milestone: value.milestone.map(TryInto::try_into).transpose()?,
        })
    }
}
