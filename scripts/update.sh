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
	echo "✅ $*"
}

function error() {
	echo "💥 ERROR: $*"
}

function main() {

	message "Tidying Go modules..."
	go mod tidy || {
		error "Failed to tidy Go modules!"
		return 1
	}

	message "Updating Go dependencies..."
	go get -v -u ./... || {
		error "Failed to update Go dependencies!"
		return 1
	}

	message "Vendoring Go dependencies..."
	go mod vendor || {
		error "Failed to vendor Go dependencies!"
		return 1
	}

	return 0
}

##################################################
# MAIN
##################################################

header "✨ TASK: Running script..."

main || {
	header "💥 ERROR: Script failed! Review the output for more information."
	exit 1
}

header "✨ TASK: Script completed successfully!"
exit 0
