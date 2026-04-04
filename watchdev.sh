#!/bin/bash

# Watch for changes in the project and reload the server

inotifywait -m -r -e close_write -e create -e delete www | while read -r path action file; do
    echo $action: $path$file
done
