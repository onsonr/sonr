#!/usr/bin/env python3
import yaml
import sys
import os

try:
    if len(sys.argv) != 2:
        raise Exception("This script requires exactly 1 argument")

    # Import Manifest.yml
    current_dir = os.path.dirname(os.path.realpath(__file__))

    with open(os.path.join(current_dir, "../../Manifest.yml"), "r") as stream:
        manifest = yaml.safe_load(stream)

    keys = sys.argv[1].split('.')
    result = manifest

    for key in keys:
        result = result[key]

    if type(result) is str \
            or type(result) is bool \
            or type(result) is int:
        print(result)
    else:
        print(yaml.dump(result, default_flow_style=False))

except Exception as err:
    sys.stderr.write("Error: %s\n" % str(err))
    exit(1)
