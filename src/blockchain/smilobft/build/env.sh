#!/bin/sh

set -e

if [ ! -f "src/blockchain/smilobft/build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
ethdir="$workspace/src"
if [ ! -L "$ethdir/go-didux" ]; then
    mkdir -p "$ethdir"
    cd "$ethdir"
    ln -s ../../../. go-didux
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$ethdir/go-didux"
PWD="$ethdir/go-didux"

# Launch the arguments with the configured environment.
export GO111MODULE=off
exec "$@"
