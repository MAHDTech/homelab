#!/usr/bin/env bash

clear

set -euo pipefail

##################################################
# CONSTANTS
##################################################

readonly DOTENV_FILE_COMMON=".env"

readonly EXTERNAL_DEPS=(
	jq
)

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

function dotenv() {

	local ENV_FILE="$1"

	if [[ -f $ENV_FILE ]]; then
		message "Sourcing dotenv file $ENV_FILE..."
		# shellcheck disable=SC1090
		source "$ENV_FILE"
	else
		message "No dotenv file found named $ENV_FILE, skipping..."
	fi

	return 0

}

function dependencies() {
	header "✨ TASK: Checking dependencies..."

	# Go dependencies
	go_dependencies || {
		error "Failed to check Go dependencies!"
		return 1
	}

	# External dependencies
	external_dependencies || {
		error "Failed to check external dependencies!"
		return 1
	}

	return 0

}

function go_dependencies() {
	header "✨ TASK: Checking dependencies..."

	# Extract the Go package name.
	message "Extracting Go package name..."
	GO_PACKAGE_NAME=$(go list -m)
	if [[ ${GO_PACKAGE_NAME:-EMPTY} == "EMPTY" ]]; then
		error "Failed to extract Go package name!"
		return 1
	fi

	# Go get all package dependencies.
	message "Go getting package dependencies..."
	go get -v "${GO_PACKAGE_NAME}" || {
		error "Failed to go get package dependencies!"
		return 1
	}

	# Download and Tidy the Go modules.
	message "Downloading and Tidying Go modules..."
	go mod tidy || {
		error "Failed to tidy Go modules!"
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

function external_dependencies() {
	header "✨ TASK: Checking external dependencies..."

	# Make sure pulumi is installed.
	message "Checking for Pulumi installation..."
	pulumi version >/dev/null 2>&1 || {
		error "Pulumi is not installed!"
		return 1
	}

	for DEP in "${EXTERNAL_DEPS[@]}"; do
		message "Checking for $DEP installation..."
		$DEP --version >/dev/null 2>&1 || {
			error "$DEP is not installed!"
			return 1
		}
	done

	return 0
}

function select_stack() {
	header "✨ TASK: Selecting stack..."

	message "Pulumi stack selection..."
	pulumi stack select || {
		error "Failed to select stack!"
		return 1
	}

	PULUMI_STACK=$(pulumi stack ls --json | jq -r '.[] | select(.current == true) | .name')

	message "Selected stack: $PULUMI_STACK"
	dotenv ".env.${PULUMI_STACK}" || {
		error "Failed to source dotenv file for stack ${PULUMI_STACK} from file .env.${PULUMI_STACK}!"
		return 1
	}

	return 0
}

function dry_run() {
	header "✨ TASK: Starting preview..."

	message "Pulumi previewing..."
	pulumi preview \
		--refresh \
		--show-replacement-steps \
		--logtostderr \
		--verbose=3 || {
		error "Failed to Pulumi preview!"
		return 1
	}
}

function run() {
	header "✨ TASK: Starting update..."

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

	# Source the common dotenv if it exists.
	dotenv "${DOTENV_FILE_COMMON}" || {
		error "Failed to source dotenv!"
		return 1
	}

	# Make sure the dependencies are installed.
	dependencies || {
		error "Failed to check dependencies!"
		return 1
	}

	# Ask the user to select the stack.
	select_stack || {
		error "Failed to select stack!"
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

header "✨ TASK: Running script..."

main || {
	header "💥 ERROR: Script failed! Review the output for more information."
	exit 1
}

header "✨ TASK: Script completed successfully!"
exit 0
