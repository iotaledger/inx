// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::Error;
use crate::proto;
use bee_block_stardust as stardust;
use stardust::output::RentStructureBuilder;

impl From<proto::RentStructure> for stardust::output::RentStructure {
    fn from(value: proto::RentStructure) -> Self {
        RentStructureBuilder::new()
            .byte_cost(value.v_byte_cost as u32)
            .data_factor(value.v_byte_factor_data as u8)
            .key_factor(value.v_byte_factor_key as u8)
            .finish()
    }
}

/// The [`ProtocolParameters`] type.
#[derive(Clone, Debug, PartialEq)]
pub struct ProtocolParameters {
    /// The protocol version of the network.
    pub version: u32,
    /// The name of the network.
    pub network_name: String,
    /// The human-readable part of the bech32 format.
    pub bech32_hrp: String,
    /// Minimum required PoW score.
    pub min_pow_score: u32,
    /// The below max depth (BMD) parameter of the tip selection algorithm.
    pub below_max_depth: u32,
    /// Defines the parameters for the byte cost calculation
    pub rent_structure: stardust::output::RentStructure,
    /// The overall token supply.
    pub token_supply: u64,
}

impl TryFrom<proto::ProtocolParameters> for ProtocolParameters {
    type Error = Error;

    fn try_from(value: proto::ProtocolParameters) -> Result<Self, Error> {
        Ok(Self {
            version: value.version,
            network_name: value.network_name,
            bech32_hrp: value.bech32_hrp,
            min_pow_score: value.min_po_w_score,
            below_max_depth: value.below_max_depth,
            rent_structure: value
                .rent_structure
                .ok_or(Error::MissingField("rent_structure"))?
                .into(),
            token_supply: value.token_supply,
        })
    }
}
