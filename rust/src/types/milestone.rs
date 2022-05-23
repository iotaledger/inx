// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::Error;
use crate::proto;

use bee_block_stardust as stardust;
use packable::PackableExt;

#[allow(missing_docs)]
#[derive(PartialEq, Debug)]
pub enum MilestoneRequest {
    MilestoneIndex(u32),
    MilestoneId(stardust::payload::milestone::MilestoneId),
}

/// Convenience method to create a [`proto::MilestoneRangeRequest`]. All the milestone ranges are inclusive.
pub enum MilestoneRangeRequest {
    /// Returns a stream that starts now and continues indefinetly.
    UntilForever,
    /// Returns a stream that starts now and continues until a given milestone index.
    UntilMilestoneIndex(u32),
    /// Returns a stream that starts and end with given milestone indices.
    FromUntilMilestoneIndex(u32, u32),
    /// Returns a stream that starts with a given given milestone index and continues indefinetly.
    FromMilestoneIndex(u32),
}

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
            milestone: value.milestone.ok_or(Error::MissingField("milestone"))?.try_into()?,
        })
    }
}

impl From<MilestoneRequest> for proto::MilestoneRequest {
    fn from(value: MilestoneRequest) -> Self {
        match value {
            MilestoneRequest::MilestoneIndex(milestone_index) => proto::MilestoneRequest {
                milestone_index,
                milestone_id: None,
            },
            MilestoneRequest::MilestoneId(milestone_id) => proto::MilestoneRequest {
                milestone_id: Some(milestone_id.into()),
                milestone_index: 0,
            },
        }
    }
}

impl From<MilestoneRangeRequest> for proto::MilestoneRangeRequest {
    fn from(value: MilestoneRangeRequest) -> Self {
        match value {
            MilestoneRangeRequest::UntilForever => proto::MilestoneRangeRequest {
                start_milestone_index: 0,
                end_milestone_index: 0,
            },
            MilestoneRangeRequest::UntilMilestoneIndex(end_milestone_index) => proto::MilestoneRangeRequest {
                start_milestone_index: 0,
                end_milestone_index,
            },
            MilestoneRangeRequest::FromMilestoneIndex(start_milestone_index) => proto::MilestoneRangeRequest {
                start_milestone_index,
                end_milestone_index: 0,
            },
            MilestoneRangeRequest::FromUntilMilestoneIndex(start_milestone_index, end_milestone_index) => {
                proto::MilestoneRangeRequest {
                    start_milestone_index,
                    end_milestone_index,
                }
            }
        }
    }
}
