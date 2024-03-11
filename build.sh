#!/bin/bash

# Create the "out" folder if it doesn't exist
mkdir -p out

# List of target OS and architectures
targets=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "windows/amd64"
    "darwin/arm64"
    "windows/arm64"
)

# Loop through the targets and build the sylote
for target in "${targets[@]}"; do
    # Split the target into OS and architecture
    IFS="/" read -r os arch <<< "$target"

    # Set the GOOS and GOARCH environment variables
    export GOOS="$os"
    export GOARCH="$arch"

    # Build the sylote and export it to the "out" folder
    go build -o "out/sylote-$os-$arch"

    # Clear the environment variables
    unset GOOS
    unset GOARCH
done

lipo -create -output out/sylote-darwin-universal out/sylote-darwin-amd64 out/sylote-darwin-arm64