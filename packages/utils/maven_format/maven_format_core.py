#!/usr/bin/env python3
from tempfile import TemporaryDirectory
from zipfile import ZipFile
import pystache
import yaml
import shutil
import os
import sys
import re
import subprocess

try:
    # Import Manifest.yml
    current_dir = os.path.dirname(os.path.realpath(__file__))

    with open(os.path.join(current_dir, "../../Manifest.yml"), "r") as stream:
        manifest = yaml.safe_load(stream)

    target_sdk = manifest["global"]["android"]["target_sdk_version"]
    artifact_id = manifest["go_core"]["android"]["artifact_id"]
    group_id = manifest["global"]["group_id"]
    packaging = manifest["global"]["android"]["packaging"]
    core_name = manifest["go_core"]["android"]["name"]
    description = manifest["go_core"]["android"]["description"]
    website = manifest["global"]["github"]["url"]
    licenses = manifest["global"]["licenses"]
    developers = manifest["global"]["developers"]
    scm_conn = manifest["global"]["android"]["scm"]["connection"]
    scm_dev_conn = manifest["global"]["android"]["scm"]["developer_connection"]
    scm_url = manifest["global"]["android"]["scm"]["url"]

    # Get version from env (CI) or set to dev
    if "GOMOBILE_IPFS_VERSION" in os.environ:
        global_version = os.getenv("GOMOBILE_IPFS_VERSION")
    else:
        global_version = "0.0.42-dev"

    # Check if ANDROID_HOME is set in env
    if "ANDROID_HOME" not in os.environ:
        raise Exception("ANDROID_HOME must be set in environment")

    # Check if required targetSdkVersion is installed
    target_sdk_path = os.path.join(
        os.getenv("ANDROID_HOME"),
        "platforms",
        "android-%d" % target_sdk,
        "android.jar",
    )

    if os.path.exists(target_sdk_path) is False:
        raise Exception(
            "Android required targetSDKVersion (%d) not found (tried: %s)" %
            (target_sdk, target_sdk_path)
        )

    # Create temporary working directory
    with TemporaryDirectory() as temp_dir:
        # Define target filenames for formatted artifacts
        pom_art = "%s-%s.pom" % (artifact_id, global_version)
        javadoc_art = "%s-%s-javadoc.jar" % (artifact_id, global_version)
        sources_art = "%s-%s-sources.jar" % (artifact_id, global_version)
        package_art = "%s-%s.aar" % (artifact_id, global_version)

        # Copy intermediate artifacts and rename them
        print("Copying intermediate artifacts to tmp working directory")
        android_build_dir = os.path.join(
            os.path.dirname(os.path.dirname(current_dir)),
            "build/android",
        )
        android_build_dir_mav = os.path.join(android_build_dir, "maven")
        android_build_dir_int_core = os.path.join(
            android_build_dir,
            "intermediates/core",
        )
        shutil.copyfile(
            os.path.join(android_build_dir_int_core, "core.aar"),
            os.path.join(temp_dir, package_art),
        )
        shutil.copyfile(
            os.path.join(android_build_dir_int_core, "core-sources.jar"),
            os.path.join(temp_dir, sources_art),
        )

        # Unzip source jar in order to generate javadoc
        print("Unziping sources.jar")
        with ZipFile(os.path.join(temp_dir, sources_art), "r") as zip_ref:
            zip_ref.extractall(temp_dir)

        os.mkdir(os.path.join(temp_dir, "javadoc"))

        # Generate javadoc
        print("Generating javadoc")

        # Check if Java version <= 8
        version_infos = subprocess.check_output(
            ["java", "-version"],
            stderr=subprocess.STDOUT
        ).decode("utf-8")
        version = re.search(r'\"(\d+\.\d+).*\"', version_infos).groups()[0]

        if float(version) <= 1.8:
            if os.system("javadoc -Xdoclint:none -quiet -classpath '%s' "
                         "-bootclasspath '%s' go core -d %s" %
                         (
                            os.path.join(temp_dir, sources_art),
                            target_sdk_path,
                            os.path.join(temp_dir, "javadoc")
                         )):
                exit(1)
        else:
            # TODO: add Android javadoc generation using Java version > 1.8
            raise Exception("Can't generate javadoc using Java version > 1.8")

        # Create a jar containing the javadoc
        if os.system("cd %s && jar -cf ../%s ." %
                     (os.path.join(temp_dir, "javadoc"), javadoc_art)):
            exit(1)

        # Generate pom file
        print("Generating pom file")
        with open(os.path.join(current_dir, "pom_template"), "r") as template:
            template_str = template.read()

        context = {
            "group_id": group_id,
            "artifact_id": artifact_id,
            "version": global_version,
            "packaging": packaging,
            "name": core_name,
            "description": description,
            "url": website,
            "licenses": [],
            "has_licenses": False,
            "developers": [],
            "has_developers": False,
        }

        for license in licenses:
            if license["name"] or license["url"] or license["distribution"]:
                context["has_licenses"] = True
                context["licenses"].append({
                    "name": license["name"],
                    "url": license["url"],
                    "distribution": license["distribution"],
                })

        for developer in developers:
            if developer["id"] \
                    or developer["name"] \
                    or developer["email"] \
                    or developer["organization"] \
                    or developer["organization_url"]:
                context["has_developers"] = True
                context["developers"].append({
                    "id": developer["id"],
                    "name": developer["name"],
                    "email": developer["email"],
                    "organization": developer["organization"],
                    "organization_url": developer["organization_url"],
                })

        if scm_conn or scm_dev_conn or scm_url:
            context["scm"] = {
                "connection": scm_conn,
                "developer_connection": scm_dev_conn,
                "url": scm_url,
            }

        rendered = pystache.render(template_str, context)
        rendered_trimed = re.sub(r'\n\s*\n', "\n\n", rendered).replace(
            "\n\n</project>",
            "\n</project>",
        )

        with open(os.path.join(temp_dir, pom_art), "w") as output:
            output.write(rendered_trimed)

        # Create dest directory with maven standard path
        maven_path = "%s/%s/%s/%s" % (
            *group_id.split("."),
            artifact_id,
            global_version,
        )
        dest_dir = os.path.join(android_build_dir_mav, maven_path)
        print("Creating destination directory with maven compliant path: %s" %
              maven_path)
        os.makedirs(dest_dir, exist_ok=True)

        # Copy generated artifacts to dest directory
        print("Copying artifacts to destination directory")
        shutil.copyfile(
            os.path.join(temp_dir, package_art),
            os.path.join(dest_dir, package_art),
        )
        shutil.copyfile(
            os.path.join(temp_dir, sources_art),
            os.path.join(dest_dir, sources_art),
        )
        shutil.copyfile(
            os.path.join(temp_dir, javadoc_art),
            os.path.join(dest_dir, javadoc_art),
        )
        shutil.copyfile(
            os.path.join(temp_dir, pom_art),
            os.path.join(dest_dir, pom_art),
        )

        print("Maven formatting succeeded!")
except Exception as err:
    sys.stderr.write("Error: %s\n" % str(err))
    exit(1)
