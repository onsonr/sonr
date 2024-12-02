#!/usr/bin/env bash

set -e

ROOT_DIR=$(git rev-parse --show-toplevel)

# Package the PKL projects
bunx pkl project package $ROOT_DIR/pkl/*/

# Process each directory in .out
for dir in .out/*/; do
    # Get the folder name and version
    folder=$(basename "$dir")
    version=$(echo "$folder" | grep -o '@.*' | sed 's/@//')
    new_folder=$(echo "$folder" | sed 's/@[0-9.]*$//')
    
    # Create new directory without version
    mkdir -p ".out/$new_folder/$version"
    
    # Copy contents to versioned subdirectory
    cp -r "$dir"* ".out/$new_folder/$version/"
    
    # Find and copy only .pkl files from the original package
    pkg_dir="$ROOT_DIR/pkl/$new_folder"
    if [ -d "$pkg_dir" ]; then
        # Copy only .pkl files to version directory
        find "$pkg_dir" -name "*.pkl" -exec cp {} ".out/$new_folder/$version/" \;
    fi
    
    # Remove old versioned directory
    rm -rf "$dir"
    
    # Upload to R2 with new structure
    rclone copy ".out/$new_folder" "r2:pkljar/$new_folder"
done

# Cleanup .out directory
rm -rf .out

# Handle static files
rclone copy $ROOT_DIR/static "r2:nebula"
rm -rf $ROOT_DIR/static
