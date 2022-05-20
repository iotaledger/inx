// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use thiserror::Error;

/// Represents [`Error`]s that happened during conversion.
#[allow(missing_docs)]
#[derive(PartialEq, Debug, Error)]
pub enum Error {
    #[error("missing field: {0}")]
    MissingField(&'static str),
    #[error("invalid field: {0}")]
    InvalidField(&'static str),
    #[error("invalid buffer length")]
    InvalidBufferLength,
    #[error("packable error: {0}")]
    PackableError(String),
}
