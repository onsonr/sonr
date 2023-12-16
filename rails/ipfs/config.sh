#!bin/bash

ipfs config --json Gateway '{
        "HTTPHeaders": {
            "Access-Control-Allow-Origin": [
                "*"
            ],
        },
        "RootRedirect": "",
        "Writable": false,
        "PathPrefixes": [
            "/blog",
            "/refs"
        ],
        "APICommands": [],
        "NoFetch": false,
        "NoDNSLink": false,
        "PublicGateways": {
            "dweb.link": {
                "NoDNSLink": false,
                "Paths": [
                    "/ipfs",
                    "/ipns",
                    "/api"
                ],
                "UseSubdomains": true
            },
            "gateway.ipfs.io": {
                "NoDNSLink": false,
                "Paths": [
                    "/ipfs",
                    "/ipns",
                    "/api"
                ],
                "UseSubdomains": false
            },
            "ipfs.io": {
                "NoDNSLink": false,
                "Paths": [
                    "/ipfs",
                    "/ipns",
                    "/api"
                ],
                "UseSubdomains": false
            }
        }
    }'
