// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::proto;

use std::ops::{Range, RangeFrom, RangeFull, RangeInclusive, RangeTo};

impl<T> From<Range<T>> for proto::MilestoneRangeRequest
where
    T: Into<u32>,
{
    fn from(value: Range<T>) -> Self {
        proto::MilestoneRangeRequest {
            start_milestone_index: value.start.into(),
            end_milestone_index: value.end.into() - 1, // `proto::MilestoneRangeRequest` is inclusive
        }
    }
}

impl<T> From<RangeInclusive<T>> for proto::MilestoneRangeRequest
where
    T: Into<u32>,
{
    fn from(value: RangeInclusive<T>) -> Self {
        let (from, to) = value.into_inner();
        proto::MilestoneRangeRequest {
            start_milestone_index: from.into(),
            end_milestone_index: to.into(),
        }
    }
}

impl<T> From<RangeFrom<T>> for proto::MilestoneRangeRequest
where
    T: Into<u32>,
{
    fn from(value: RangeFrom<T>) -> Self {
        proto::MilestoneRangeRequest {
            start_milestone_index: value.start.into(),
            end_milestone_index: 0,
        }
    }
}

impl<T> From<RangeTo<T>> for proto::MilestoneRangeRequest
where
    T: Into<u32>,
{
    fn from(value: RangeTo<T>) -> Self {
        proto::MilestoneRangeRequest {
            start_milestone_index: 0,
            end_milestone_index: value.end.into(),
        }
    }
}

impl From<RangeFull> for proto::MilestoneRangeRequest {
    fn from(_: RangeFull) -> Self {
        proto::MilestoneRangeRequest {
            start_milestone_index: 0,
            end_milestone_index: 0,
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    struct MilestoneIndex(u32);

    impl From<MilestoneIndex> for u32 {
        fn from(value: MilestoneIndex) -> Self {
            value.0
        }
    }

    #[test]
    fn exclusive() {
        let from = MilestoneIndex(17);
        let to = MilestoneIndex(43);
        let range = proto::MilestoneRangeRequest::from(from..to);
        assert_eq!(
            range,
            proto::MilestoneRangeRequest {
                start_milestone_index: 17,
                end_milestone_index: 42
            }
        );
    }
}
