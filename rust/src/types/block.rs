// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::Error;
use crate::proto;

use bee_block_stardust as stardust;
use packable::PackableExt;

/// The [`Block`] type.
#[derive(PartialEq, Debug)]
pub struct Block {
    /// The [`BlockId`](stardust::BlockId) of the block.
    pub block_id: stardust::BlockId,
    /// The complete [`Block`](stardust::Block).
    pub block: stardust::Block,
}

/// The [`BlockWithMetadata`] type.
#[derive(PartialEq, Debug)]
pub struct BlockWithMetadata {
    /// The [`Metadata`](crate::BlockMetadata) of the block.
    pub metadata: crate::BlockMetadata,
    /// The complete [`Block`](stardust::Block).
    pub block: stardust::Block,
}

impl TryFrom<proto::BlockWithMetadata> for BlockWithMetadata {
    type Error = Error;

    fn try_from(value: proto::BlockWithMetadata) -> Result<Self, Self::Error> {
        let metadata = value.metadata.ok_or(Error::MissingField("metadata"))?.try_into()?;
        let block = value.block.ok_or(Error::MissingField("block"))?.try_into()?;

        Ok(BlockWithMetadata { metadata, block })
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

impl TryFrom<proto::Block> for Block {
    type Error = Error;

    fn try_from(value: proto::Block) -> Result<Self, Self::Error> {
        Ok(Block {
            block_id: value.block_id.ok_or(Error::MissingField("block_id"))?.try_into()?,
            block: value.block.ok_or(Error::MissingField("block"))?.try_into()?,
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

impl TryFrom<proto::RawBlock> for stardust::Block {
    type Error = Error;

    fn try_from(value: proto::RawBlock) -> Result<Self, Self::Error> {
        stardust::Block::unpack_verified(value.data).map_err(|e| Error::PackableError(format!("{e}")))
    }
}
