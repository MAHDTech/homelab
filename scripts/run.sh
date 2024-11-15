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

function dotenv() {

	if [[ -f ".env" ]]; then
		message "Sourcing .env file..."
		# shellcheck disable=SC1091
		source .env
	else
		message "No .env file found, skipping..."
	fi

	return 0

}

function dependencies() {

	# Make sure pulumi is installed.
	message "Checking for Pulumi installation..."
	pulumi version >/dev/null 2>&1 || {
		error "Pulumi is not installed!"
		return 1
	}

	# Extract the Go package name.
	message "Extracting Go package name..."
	GO_PACKAGE_NAME=$(go list -m)

	# Go get all package dependencies.
	message "Go getting package dependencies..."
	go get -v "${GO_PACKAGE_NAME}" || {
		error "Failed to go get package dependencies!"
		return 1
	}

	# Tidy the Go modules.
	message "Tidying Go modules..."
	go mod tidy || {
		error "Failed to tidy Go modules!"
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
	header "OK: Starting preview..."

	message "Pulumi previewing..."
	pulumi preview \
		--refresh \
		--diff \
		--show-replacement-steps \
		--logtostderr \
		--verbose=3 || {
		error "Failed to Pulumi preview!"
		return 1
	}
}

function run() {
	header "OK: Starting update..."

	message "Pulumi updating..."
	pulumi update \
		--refresh \
		--skip-preview \
		--show-replacement-steps \
		--logtostderr \
		--verbose=3 || {
		error "Failed to Pulumi update!"
		return 1
	}

}

function main() {

	# Source the dotenv if it exists.
	dotenv || {
		error "Failed to source dotenv!"
		return 1
	}

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

header "OK: Running script..."

main || {
	header "ERROR: Script failed! Review the output for more information."
	exit 1
}

header "OK: Script completed successfully!"
exit 0
