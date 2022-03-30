// Copyright 2020-2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

pub mod proto {
    tonic::include_proto!("inx");
}

pub use self::proto::{inx_client as client, inx_server as server};
