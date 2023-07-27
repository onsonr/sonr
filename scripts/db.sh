#!/bin/bash

# remove existing data
rm -rf ~/data

IceFireDB start -a $SONR_ICEFIRE_HOST
