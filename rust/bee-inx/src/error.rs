// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

/// Represents [`Error`](Error) s that happened during conversion.
#[allow(missing_docs)]
#[derive(Clone, Debug, PartialEq, thiserror::Error)]
pub enum Error {
    #[error(transparent)]
    InxError(bee_block_stardust::error::inx::InxError)
}
