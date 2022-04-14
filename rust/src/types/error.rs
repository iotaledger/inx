// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

/// Represents [`Error`]s that happened during conversion.
#[allow(missing_docs)]
#[derive(PartialEq, Debug)]
pub enum Error {
    MissingField(&'static str),
    InvalidField(&'static str),
    InvalidBufferLength,
    PackableError(String),
}
