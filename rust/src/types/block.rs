// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::Error;
use crate::proto;

use bee_block_stardust as stardust;
use packable::PackableExt;

/// The [`Block`] type.
#[derive(Clone, Debug, PartialEq)]
pub struct Block {
    /// The [`BlockId`](stardust::BlockId) of the block.
    pub block_id: stardust::BlockId,
    /// The complete [`Block`](stardust::Block).
    pub block: stardust::Block,
    /// The raw bytes of the block.
    pub raw: Vec<u8>,
}

/// The [`BlockWithMetadata`] type.
#[derive(Clone, Debug, PartialEq)]
pub struct BlockWithMetadata {
    /// The [`Metadata`](crate::BlockMetadata) of the block.
    pub metadata: crate::BlockMetadata,
    /// The complete [`Block`](stardust::Block).
    pub block: stardust::Block,
    /// The raw bytes of the block.
    pub raw: Vec<u8>,
}

impl TryFrom<proto::BlockWithMetadata> for BlockWithMetadata {
    type Error = Error;

    fn try_from(value: proto::BlockWithMetadata) -> Result<Self, Self::Error> {
        let raw = value.block.ok_or(Error::MissingField("block"))?;
        Ok(BlockWithMetadata {
            metadata: value.metadata.ok_or(Error::MissingField("metadata"))?.try_into()?,
            block: raw.clone().try_into()?,
            raw: raw.data,
        })
    }
}

impl From<stardust::BlockId> for proto::BlockId {
    fn from(value: stardust::BlockId) -> Self {
        Self {
            id: value.pack_to_vec(),
        }
    }
}

impl TryFrom<proto::BlockId> for stardust::BlockId {
    type Error = Error;

    fn try_from(value: proto::BlockId) -> Result<Self, Self::Error> {
        let bytes: [u8; stardust::BlockId::LENGTH] = value.id.try_into().map_err(|_| Error::InvalidBufferLength)?;
        Ok(stardust::BlockId::from(bytes))
    }
}

impl TryFrom<proto::TransactionId> for stardust::payload::transaction::TransactionId {
    type Error = Error;

    fn try_from(value: proto::TransactionId) -> Result<Self, Self::Error> {
        let bytes: [u8; stardust::payload::transaction::TransactionId::LENGTH] =
            value.id.try_into().map_err(|_| Error::InvalidBufferLength)?;
        Ok(stardust::payload::transaction::TransactionId::from(bytes))
    }
}

impl TryFrom<proto::Block> for Block {
    type Error = Error;

    fn try_from(value: proto::Block) -> Result<Self, Self::Error> {
        let raw = value.block.ok_or(Error::MissingField("block"))?;
        Ok(Block {
            block_id: value.block_id.ok_or(Error::MissingField("block_id"))?.try_into()?,
            block: raw.clone().try_into()?,
            raw: raw.data,
        })
    }
}

impl TryFrom<proto::Block> for (Block, Vec<u8>) {
    type Error = Error;

    fn try_from(value: proto::Block) -> Result<Self, Self::Error> {
        if let Some(proto::RawBlock { data }) = value.block.clone() {
            Ok((value.try_into()?, data))
        } else {
            Err(Error::MissingField("block"))
        }
    }
}

macro_rules! impl_try_from_raw {
    ($from_type:ty, $to_type:ty) => {
        impl TryFrom<$from_type> for $to_type {
            type Error = Error;

            fn try_from(value: $from_type) -> Result<Self, Self::Error> {
                Self::unpack_verified(value.data).map_err(|e| Error::PackableError(format!("{e}")))
            }
        }
    };
}

impl_try_from_raw!(proto::RawBlock, stardust::Block);
impl_try_from_raw!(proto::RawOutput, stardust::output::Output);
