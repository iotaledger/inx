// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#[deny(warnings, missing_docs)]

#[cfg(feature = "types")]
mod types;

/// The raw protobuf definitions.
pub mod proto {
    tonic::include_proto!("inx");
}

/// Rexports of [`tonic`] types.
pub mod tonic {
    pub use tonic;
}

pub use self::proto::{inx_client as client, inx_server as server};

#[cfg(feature = "types")]
pub use self::types::{Error, Milestone, Message};
