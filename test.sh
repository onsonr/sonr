#!/bin/sh

curl -fsS https://pkgx.sh | sh

pkgx +gh +gum +docker +git +make +bash which node  #=> /tmp/pkgx.sh/nodejs.org/v16/bin/node

gum confirm "Are you sure you want to deploy?" && gh release create v1.0.0
