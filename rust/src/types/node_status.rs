// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::{Error, Milestone};
use crate::proto;

/// The [`NodeStatus`] type.
#[derive(PartialEq, Debug)]
pub struct NodeStatus {
    pub is_healthy: bool,
    pub latest_milestone: Milestone,
    pub confirmed_milestone: Milestone,
    pub pruning_index: u32,
    pub ledger_index: u32,
}

impl TryFrom<proto::NodeStatus> for NodeStatus {
    type Error = Error;

    fn try_from(value: proto::NodeStatus) -> Result<Self, Self::Error> {
        Ok(NodeStatus {
            is_healthy: value.is_healthy,
            latest_milestone: value
                .latest_milestone
                .ok_or(Error::MissingField("latest_milestone"))?
                .try_into()?,
            confirmed_milestone: value
                .confirmed_milestone
                .ok_or(Error::MissingField("confirmed_milestone"))?
                .try_into()?,
            pruning_index: value.pruning_index,
            ledger_index: value.ledger_index,
        })
    }
}
