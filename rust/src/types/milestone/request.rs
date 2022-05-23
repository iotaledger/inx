// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::proto;

impl<T> From<T> for proto::MilestoneRequest
where
    T: Into<u32>,
{
    fn from(value: T) -> Self {
        Self {
            milestone_index: value.into(),
            milestone_id: None,
        }
    }
}

impl From<proto::MilestoneId> for proto::MilestoneRequest {
    fn from(value: proto::MilestoneId) -> Self {
        Self {
            milestone_index: 0,
            milestone_id: Some(value),
        }
    }
}
