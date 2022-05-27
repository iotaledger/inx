// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![deny(warnings, missing_docs)]

//! This crate provides bindings for the IOTA node extensions (INX).

#[cfg(feature = "types")]
mod types;

/// The raw protobuf definitions.
pub mod proto {
    #![allow(missing_docs)]
    tonic::include_proto!("inx");
}

/// Rexports of [`tonic`] types.
pub mod tonic {
    pub use tonic::{
        transport::{Channel, Error},
        Code, Request, Response, Status,
    };
}

pub use self::proto::{inx_client as client, inx_server as server};

#[cfg(feature = "types")]
pub use self::types::*;

impl proto::MilestoneRequest {
    /// Creates a [`MilestoneRequest`](proto::MilestoneRequest) from an milestone index.
    pub fn from_index(milestone_index: u32) -> Self {
        Self {
            milestone_index,
            milestone_id: None,
        }
    }

    /// Creates a [`MilestoneRequest`](proto::MilestoneRequest) from a [`MilestoneId`](proto::MilestoneId).
    pub fn from_id(milestone_id: proto::MilestoneId) -> Self {
        Self {
            milestone_index: 0,
            milestone_id: Some(milestone_id),
        }
    }
}
