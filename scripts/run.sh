#!/usr/bin/env bash

clear

set -euo pipefail

##################################################
# FUNCTIONS
##################################################

function header() {
    echo "----------------------------------------"
    echo "$*"
    echo "----------------------------------------"
}

function message() {
    echo "$*"
}

function error() {
    echo "ERROR: $*"
}

function dependencies() {

    # Make sure pulumi is installed.
    message "Checking for Pulumi installation..."
    pulumi version > /dev/null 2>&1 || {
        error "Pulumi is not installed!"
        return 1
    }

    # Extract the Go package name.
    message "Extracting Go package name..."
    GO_PACKAGE_NAME=$(go list -m)

    # Go get all package dependencies.
    message "Go getting package dependencies..."
    go get $GO_PACKAGE_NAME || {
        error "Failed to go get package dependencies!"
        return 1
    }

    # Download the Go modules.
    message "Downloading Go modules..."
    go mod download || {
        error "Failed to download Go modules!"
        return 1
    }

    # Vendor the Go modules for offline use.
    message "Vendoring Go modules..."
    go mod vendor || {
        error "Failed to vendor Go modules!"
        return 1
    }

    return 0
}

function dry_run() {
    header "Dry run..."

    message "Pulumi previewing..."
    pulumi preview \
        --refresh \
        --diff \
        --show-replacement-steps \
        --verbose=3 || {
        error "Failed to Pulumi preview!"
        return 1
    }
}

function run() {
    header "Run..."

    message "Pulumi updating..."
    pulumi update \
        --refresh \
        --skip-preview \
        --show-replacement-steps \
        --verbose=3 || {
        error "Failed to Pulumi update!"
        return 1
    }

}

function main() {

    # Make sure the dependencies are installed.
    dependencies || {
        error "Failed to check dependencies!"
        return 1
    }

    # Run the script (Dry run)
	dry_run || {
        error "Failed to run (dry run)!"
        return 1
    }

    # Run the script (Actual run)
	run || {
        error "Failed to run!"
        return 1
    }

    return 0
}

##################################################
# MAIN
##################################################

header "Running script..."

main || {
    error "Failed to run script!"
    exit 1
}

header "Script completed successfully!"
exit 0
