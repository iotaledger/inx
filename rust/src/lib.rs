// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

//! This crate provides bindings for the IOTA node extensions (INX).

/// The raw protobuf definitions.
pub mod proto {
    #![allow(missing_docs)]
    #![allow(clippy::derive_partial_eq_without_eq)]
    tonic::include_proto!("inx");
}

/// Re-exports of [`tonic`] types.
pub mod tonic {
    pub use tonic::*;
}

pub use self::proto::inx_client as client;
