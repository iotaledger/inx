// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::Error;
use crate::proto;

use bee_message_stardust as stardust;
use packable::PackableExt;

/// The [`Message`] type.
#[derive(PartialEq, Debug)]
pub struct Message {
    /// The [`MessageId`](stardust::MessageId) of the message.
    pub message_id: stardust::MessageId,
    /// The complete [`Message`](stardust::Message).
    pub message: stardust::Message,
}

impl From<stardust::MessageId> for proto::MessageId {
    fn from(value: stardust::MessageId) -> Self {
        Self {
            id: value.pack_to_vec(),
        }
    }
}

impl TryFrom<proto::MessageId> for stardust::MessageId {
    type Error = Error;

    fn try_from(value: proto::MessageId) -> Result<Self, Self::Error> {
        let bytes: [u8; stardust::MessageId::LENGTH] = value.id.try_into().map_err(|_| Error::InvalidBufferLength)?;
        Ok(stardust::MessageId::from(bytes))
    }
}

impl TryFrom<proto::Message> for Message {
    type Error = Error;

    fn try_from(value: proto::Message) -> Result<Self, Self::Error> {
        Ok(Message {
            message_id: value.message_id.ok_or(Error::MissingField("message_id"))?.try_into()?,
            message: value.message.ok_or(Error::MissingField("message"))?.try_into()?,
        })
    }
}

impl TryFrom<proto::Message> for (Message, Vec<u8>) {
    type Error = Error;

    fn try_from(value: proto::Message) -> Result<Self, Self::Error> {
        if let Some(proto::RawMessage { data }) = value.message.clone() {
            Ok((value.try_into()?, data))
        } else {
            Err(Error::MissingField("message"))
        }
    }
}

impl TryFrom<proto::RawMessage> for stardust::Message {
    type Error = Error;

    fn try_from(value: proto::RawMessage) -> Result<Self, Self::Error> {
        stardust::Message::unpack_verified(value.data).map_err(|e| Error::PackableError(format!("{e}")))
    }
}
