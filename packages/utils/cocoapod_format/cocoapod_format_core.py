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

    core_name = manifest["go_core"]["ios"]["name"]
    description = manifest["go_core"]["ios"]["summary"]
    website = manifest["global"]["github"]["url"]
    platform = manifest["global"]["ios"]["platform"]
    bintray_url = manifest["global"]["ios"]["bintray_url"]
    bintray_package = manifest["go_core"]["ios"]["package"]
    licenses = manifest["global"]["licenses"]
    developers = manifest["global"]["developers"]

    # Get version from env (CI) or set to dev
    if "GOMOBILE_IPFS_VERSION" in os.environ:
        global_version = os.getenv("GOMOBILE_IPFS_VERSION")
    else:
        global_version = "0.0.42-dev"

    # Create temporary working directory
    with TemporaryDirectory() as temp_dir:
        # Generate podspec file
        print("Generating podspec file")
        with open(os.path.join(current_dir, "podspec_template"), "r") \
                as template:
            template_str = template.read()

        context = {
            "name": core_name,
            "version": global_version,
            "summary": description,
            "homepage": website,
            "license": {"type": '', "text": ''},
            "authors": [],
            "platform": platform,
            "source": "%s/%s/%s/%s-%s.zip" % (
                bintray_url,
                bintray_package,
                global_version,
                bintray_package,
                global_version,
            ),
            "framework": "Core.framework",
        }

        i = 0
        for license in licenses:
            if license["short_name"] and license["url"]:
                if i > 0:
                    context["license"]["type"] += " / "
                    context["license"]["text"] += " - "
                context["license"]["type"] += license["short_name"]
                context["license"]["text"] += license["url"]
                i += 1

        i = 0
        for developer in developers:
            if developer["name"] and developer["email"]:
                context["authors"].append({
                    "name": developer["name"],
                    "email": developer["email"],
                    "has_comma": True if i > 0 else False
                })
                i += 1

        rendered = pystache.render(template_str, context)
        rendered_trimed = re.sub(r'\n\s*\n', "\n\n", rendered).replace(
            "\n\nend",
            "\nend",
        )

        podspec_file = "%s.podspec" % core_name
        with open(os.path.join(temp_dir, podspec_file), "w") as output:
            output.write(rendered_trimed)

        # Lint generated podspec
        if os.system("cd %s && pod spec lint --quick" % temp_dir):
            exit(1)

        # Generate zip archive containing framework
        print("Generating framework zip file")
        ios_build_dir = os.path.join(
            os.path.dirname(os.path.dirname(current_dir)),
            "build/ios",
        )
        ios_build_dir_int = os.path.join(ios_build_dir, "intermediates/core")
        ios_build_dir_ccp = os.path.join(ios_build_dir, "cocoapods")
        zip_file = "%s-%s.zip" % (bintray_package, global_version)

        if os.system("cd %s && zip -ry %s ." %
                     (ios_build_dir_int, os.path.join(temp_dir, zip_file))):
            exit(1)

        # Create dest directory
        dest_dir = os.path.join(
            ios_build_dir_ccp,
            bintray_package,
            global_version,
        )
        print("Creating destination directory: %s" % dest_dir)
        os.makedirs(dest_dir, exist_ok=True)

        # Copy generated artifacts to dest directory
        print("Copying artifacts to destination directory")
        shutil.copyfile(
            os.path.join(temp_dir, podspec_file),
            os.path.join(dest_dir, podspec_file),
        )
        shutil.copyfile(
            os.path.join(temp_dir, zip_file),
            os.path.join(dest_dir, zip_file),
        )

        print("Cocoapod formatting succeeded!")
except Exception as err:
    sys.stderr.write("Error: %s\n" % str(err))
    exit(1)
