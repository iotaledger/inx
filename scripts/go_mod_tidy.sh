#!/bin/bash

pushd ./../go
go mod tidy
popd

pushd ./../examples/go
go mod tidy
popd
