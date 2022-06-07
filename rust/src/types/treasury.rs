// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::Error;
use crate::proto;

use bee_block_stardust as stardust;

/// Represents a treasury output.
#[allow(missing_docs)]
#[derive(Clone, Debug, PartialEq)]
pub struct TreasuryOutput {
    pub milestone_id: stardust::payload::milestone::MilestoneId,
    pub amount: u64,
}

impl TryFrom<proto::TreasuryOutput> for TreasuryOutput {
    type Error = Error;

    fn try_from(value: proto::TreasuryOutput) -> Result<Self, Self::Error> {
        Ok(TreasuryOutput {
            milestone_id: value
                .milestone_id
                .ok_or(Error::MissingField("milestone_id"))?
                .try_into()?,
            amount: value.amount,
        })
    }
}

/// Represents an update to the treasury.
#[allow(missing_docs)]
#[derive(Clone, Debug, PartialEq)]
pub struct TreasuryUpdate {
    pub milestone_index: u32,
    pub created: TreasuryOutput,
    pub consumed: TreasuryOutput,
}

impl TryFrom<proto::TreasuryUpdate> for TreasuryUpdate {
    type Error = Error;

    fn try_from(value: proto::TreasuryUpdate) -> Result<Self, Self::Error> {
        Ok(Self {
            milestone_index: value.milestone_index,
            created: value.created.ok_or(Error::MissingField("created"))?.try_into()?,
            consumed: value.consumed.ok_or(Error::MissingField("consumed"))?.try_into()?,
        })
    }
}
