# Create the "out" folder if it doesn't exist
New-Item -ItemType Directory -Force -Path "out"

# List of target OS and architectures
$targets = @(
    "linux/amd64",
    "linux/arm64",
    "darwin/amd64",
    "windows/amd64",
    "darwin/arm64",
    "windows/arm64",
)

# Loop through the targets and build the app
foreach ($target in $targets) {
    # Split the target into OS and architecture
    $os, $arch = $target -split "/"

    # Set the GOOS and GOARCH environment variables
    $env:GOOS = $os
    $env:GOARCH = $arch

    # Build the app and export it to the "out" folder
    go build -o "out/app-$os-$arch"

    # Clear the environment variables
    Remove-Item -Path "Env:GOOS"
    Remove-Item -Path "Env:GOARCH"
}