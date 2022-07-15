// Copyright 2020-2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

fn main() -> Result<(), std::io::Error> {
    tonic_build::compile_protos("../proto/inx.proto")
}
