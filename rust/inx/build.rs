// Copyright 2020-2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::env;

fn main() -> Result<(), std::io::Error> {
    let manifest_dir = env::var("CARGO_MANIFEST_DIR").unwrap();
    tonic_build::compile_protos(format!("{manifest_dir}/proto/inx.proto"))?;
    Ok(())
}
