#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SRC_DIR=$(dirname $DIR)

BUILD_CMD="sam build"
SERVER_CMD="sam local start-api"

rerun() {
    echo "Terminating existing server..."
    pkill -f "sam local start-api"
    
    echo "Building server..."
    if $BUILD_CMD; then
        echo "Build successful."
        # Run server
        echo "Starting server..."
        $SERVER_CMD &
    else
        echo "Build failed. Server not started."
    fi
    sleep 30
}

rerun

# Monitor file changes and rebuild
while true; do
    # Wait for file changes
    inotifywait -r -e modify,create,delete,move $SRC_DIR
    # Rebuild and run server
    rerun
done
