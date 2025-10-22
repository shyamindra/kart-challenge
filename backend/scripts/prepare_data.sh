#!/bin/bash

set -e

DATA_DIR="$(dirname "$0")"/../data

mkdir -p "$DATA_DIR"

FILES=(
    "https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase1.gz"
    "https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase2.gz"
    "https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase3.gz"
)

for FILE_URL in "${FILES[@]}"; do
    FILE_NAME=$(basename "$FILE_URL")
    GZ_PATH="$DATA_DIR/$FILE_NAME"
    TXT_PATH="${GZ_PATH%.gz}"

    if [ ! -f "$TXT_PATH" ]; then
        echo "Downloading $FILE_NAME..."
        curl -o "$GZ_PATH" "$FILE_URL"
        echo "Decompressing $FILE_NAME..."
        gunzip -k "$GZ_PATH"
        rm "$GZ_PATH"
    else
        echo "$TXT_PATH already exists, skipping download and decompression."
    fi
done

echo "Data preparation complete."
