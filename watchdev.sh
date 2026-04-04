#!/bin/bash

# dnf install inotify-tools
inotifywait -m -r -e close_write -e create -e delete www | while read -r path action file; do
    echo $action: $path$file
    curl -X POST http://localhost:8080/hot/push -d "$action: $path$file"
done
# also, find way to wire this into make up
