#!/usr/bin/env python3
from tempfile import TemporaryDirectory
import pystache
import yaml
import shutil
import sys
import os
import re

try:
    # Import Manifest.yml
    current_dir = os.path.dirname(os.path.realpath(__file__))

    with open(os.path.join(current_dir, "../../Manifest.yml"), "r") as stream:
        manifest = yaml.safe_load(stream)

    bintray_url = manifest["global"]["demo_app"]["bintray_url"]
    bintray_package = manifest["ios_demo_app"]["package"]
    filename = manifest["ios_demo_app"]["filename"]
    group_id = manifest["global"]["group_id"]
    application_id = manifest["global"]["demo_app"]["application_id"]
    app_name = manifest["ios_demo_app"]["name"]

    # Get version from env (CI) or set to dev
    if "GOMOBILE_IPFS_VERSION" in os.environ:
        global_version = os.getenv("GOMOBILE_IPFS_VERSION")
    else:
        global_version = "0.0.42-dev"

    # Create temporary working directory
    with TemporaryDirectory() as temp_dir:
        # Generate plist file
        print("Generating podspec file")
        with open(os.path.join(current_dir, "plist_template"), "r") \
                as template:
            template_str = template.read()

        ipa_file = "%s-%s.ipa" % (filename, global_version)

        context = {
            "url": "%s/%s/%s/%s" % (
                bintray_url,
                bintray_package,
                global_version,
                ipa_file,
            ),
            "bundle-identifier": "%s.%s" % (group_id, application_id),
            "version": global_version,
            "title": app_name,
        }

        rendered = pystache.render(template_str, context)
        rendered_trimed = re.sub(r'\n\s*\n', "\n\n", rendered)

        plist_file = "%s-%s.plist" % (filename, global_version)
        with open(os.path.join(temp_dir, plist_file), "w") as output:
            output.write(rendered_trimed)

        # Check if ipa was generated
        ios_build_dir = os.path.join(
             os.path.dirname(os.path.dirname(current_dir)),
             "build/ios",
        )
        ios_build_dir_int_app_ipa = os.path.join(
            ios_build_dir,
            "intermediates/app/ipa/Example.ipa"
        )
        if not os.path.exists(ios_build_dir_int_app_ipa):
            raise Exception("%s does not exist" % ios_build_dir_int_app_ipa)

        # Create dest directory
        dest_dir = os.path.join(
            ios_build_dir,
            "app",
            global_version,
        )
        print("Creating destination directory: %s" % dest_dir)
        os.makedirs(dest_dir, exist_ok=True)

        # Copy generated artifacts to dest directory
        print("Copying artifacts to destination directory")
        shutil.copyfile(
            os.path.join(temp_dir, plist_file),
            os.path.join(dest_dir, plist_file),
        )
        shutil.copyfile(
            ios_build_dir_int_app_ipa,
            os.path.join(dest_dir, ipa_file),
        )

        print("Cocoapod formatting succeeded!")
except Exception as err:
    sys.stderr.write("Error: %s\n" % str(err))
    exit(1)
