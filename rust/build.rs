// Copyright 2020-2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::env;

fn main() -> Result<(), std::io::Error> {
    let out_dir = env::var("OUT_DIR").unwrap();
    let manifest_dir = env::var("CARGO_MANIFEST_DIR").unwrap();
    std::fs::copy(
        format!("{manifest_dir}/../proto/inx.proto"),
        format!("{out_dir}/inx.proto"),
    )?;
    tonic_build::compile_protos(format!("{out_dir}/inx.proto"))?;
    Ok(())
}
