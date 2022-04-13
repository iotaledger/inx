// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::proto;

use bee_message as stardust;
use packable::{PackableExt};

/// Represents [`Error`]s that happened during conversion.
#[allow(missing_docs)]
#[derive(PartialEq, Debug)]
pub enum Error {
    MissingField(&'static str),
    InvalidField(&'static str),
    InvalidBufferLength,
    PackableError,
}

/// The [`Message`] type.
pub struct Message {
    /// The [`MessageId`](stardust::MessageId) of the message.
    pub message_id: stardust::MessageId,
    /// The complete [`Message`](stardust::Message).
    pub message: stardust::Message,
}

/// The [`Milestone`] type.
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

impl TryFrom<proto::MessageId> for stardust::MessageId {
    type Error = Error;

    fn try_from(value: proto::MessageId) -> Result<Self, Self::Error> {
        let bytes: [u8; stardust::MessageId::LENGTH] = value.id.try_into().map_err(|_| Error::InvalidBufferLength)?;
        Ok(stardust::MessageId::from(bytes))
    }
}

impl TryFrom<proto::RawMessage> for stardust::Message {
    type Error = Error;

    fn try_from(value: proto::RawMessage) -> Result<Self, Self::Error> {
        stardust::Message::unpack_verified(value.data).map_err(|_| Error::PackableError)
    }
}

