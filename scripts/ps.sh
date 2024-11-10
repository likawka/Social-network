#!/bin/bash

# Directories to exclude
EXCLUDE_DIRS="node_modules"

clear

# Generate the tree structure and filter out excluded directories
tree -L 10 -I  "$EXCLUDE_DIRS" > project_structure.txt