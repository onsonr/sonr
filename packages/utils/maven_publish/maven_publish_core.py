#!/usr/bin/env python3
from bintray.bintray import Bintray
import yaml
import os
import sys
import datetime

try:
    # Flag used to discard uploaded files on failure
    uploaded = False
    created = False

    # Init bintray API client
    if "BINTRAY_USER" not in os.environ \
            or "BINTRAY_KEY" not in os.environ \
            or "BINTRAY_GPG_PASS" not in os.environ:
        raise Exception("BINTRAY_USER, BINTRAY_KEY and "
                        "BINTRAY_GPG_PASS must be set in environment")

    bintray_user = os.getenv("BINTRAY_USER")
    bintray_key = os.getenv("BINTRAY_KEY")
    bintray_gpg_pass = os.getenv("BINTRAY_GPG_PASS")

    bintray = Bintray(bintray_user, bintray_key)

    # Import Manifest.yml
    current_dir = os.path.dirname(os.path.realpath(__file__))

    with open(os.path.join(current_dir, "../../Manifest.yml"), "r") as stream:
        manifest = yaml.safe_load(stream)

    bintray_orga = manifest["global"]["bintray_orga"]
    bintray_repo = manifest["global"]["android"]["repo"]
    artifact_id = manifest["go_core"]["android"]["artifact_id"]
    description = manifest["go_core"]["android"]["description"]
    licenses = [license["short_name"]
                for license in manifest["global"]["licenses"]]
    vcs_url = manifest["global"]["github"]["git_url"]
    website = manifest["global"]["github"]["url"]
    issue_tracker = manifest["global"]["github"]["issues_url"]
    github_repo = manifest["global"]["github"]["repo"]
    github_notes = manifest["go_core"]["android"]["github_release_notes_file"]
    download_count = manifest["go_core"]["android"]["public_download_numbers"]
    readme_content = manifest["go_core"]["android"]["readme_content"]
    readme_syntax = manifest["go_core"]["android"]["readme_syntax"]
    group_id = manifest["global"]["group_id"]
    publish = manifest["go_core"]["android"]["publish"]
    override = manifest["go_core"]["android"]["override"]

    # Get version from env (CI) or fail
    if "GOMOBILE_IPFS_VERSION" in os.environ:
        global_version = os.getenv("GOMOBILE_IPFS_VERSION")
    else:
        raise Exception("can't publish a dev version")

    # Check if remote package already exists
    package = bintray.search_package(
        subject=bintray_orga,
        repo=bintray_repo,
        package=artifact_id,
    )

    # If remote package doesn't exist, create it
    if len(package) is 1:
        print("Creating new package: %s" % artifact_id)
        bintray.create_package(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
            desc=description,
            licenses=licenses,
            vcs_url=vcs_url,
            website_url=website,
            issue_tracker_url=issue_tracker,
            github_repo=github_repo,
            github_release_notes_file=github_notes,
            public_download_numbers=download_count,
        )
    # If remote package exists, update it
    else:
        print("Updating existing package: %s" % artifact_id)
        bintray.update_package(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
            desc=description,
            licenses=licenses,
            vcs_url=vcs_url,
            website_url=website,
            issue_tracker_url=issue_tracker,
            github_repo=github_repo,
            github_release_notes_file=github_notes,
            public_download_numbers=download_count,
        )

    # Check if remote readme already exists
    try:
        bintray.get_readme(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
        )
        print("Updating existing readme for package: %s" % artifact_id)
    except Exception:
        print("Creating readme for package: %s" % artifact_id)

    # Create or update remote readme
    if readme_content is not None:
        syntax = readme_syntax or "markdown"
        bintray.create_readme(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
            bintray_syntax=syntax,
            bintray_content=readme_content,
        )
    elif github_notes is not None:
        bintray.create_readme(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
            github=github_repo,
        )

    # Check if remote version already exists
    try:
        bintray.get_version(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
            version=global_version,
        )
        version_missing = False
    except Exception:
        version_missing = True

    version_description = "{0}-{1}-{2}".format(
        bintray_repo,
        artifact_id,
        global_version,
    )
    vcs_tag = "v%s" % global_version

    # If remote version doesn't exist or override is enabled
    if version_missing or override:
        if version_missing:
            # Create the remote version
            print("Creating new version: %s for package: %s" %
                  (global_version, artifact_id))
            bintray.create_version(
                subject=bintray_orga,
                repo=bintray_repo,
                package=artifact_id,
                version=global_version,
                description=version_description,
                released=datetime.datetime.now().isoformat(),
                vcs_tag=vcs_tag
            )
        else:
            # Update the existing remote version
            print("Updating existing version: %s for package: %s" %
                  (global_version, artifact_id))
            bintray.update_version(
                subject=bintray_orga,
                repo=bintray_repo,
                package=artifact_id,
                version=global_version,
                description=version_description,
                vcs_tag=vcs_tag
            )
        created = True

        # Upload artifacts
        artifacts = [
            "%s-%s.pom" % (artifact_id, global_version),
            "%s-%s-javadoc.jar" % (artifact_id, global_version),
            "%s-%s-sources.jar" % (artifact_id, global_version),
            "%s-%s.aar" % (artifact_id, global_version),
        ]

        android_build_dir_mav = os.path.join(
            os.path.dirname(os.path.dirname(current_dir)),
            "build/android/maven",
        )
        maven_path = "%s/%s/%s/%s" % (
            *group_id.split("."),
            artifact_id,
            global_version,
        )
        artfacts_local_dir = os.path.join(android_build_dir_mav, maven_path)

        print("Uploading artifacts:")
        for artifact in artifacts:
            sys.stdout.write("    - artifact: %s ..." % artifact)
            sys.stdout.flush()
            bintray.upload_content(
                subject=bintray_orga,
                repo=bintray_repo,
                package=artifact_id,
                version=global_version,
                remote_file_path=os.path.join(maven_path, artifact),
                local_file_path=os.path.join(artfacts_local_dir, artifact),
                override=override,
            )
            sys.stdout.write("\b\b\b- done\n")
            uploaded = True

        print("Signing version: %s for package: %s" %
              (global_version, artifact_id))
        bintray.gpg_sign_version(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
            version=global_version,
            passphrase=bintray_gpg_pass,
        )

        if publish:
            print("Publishing version: %s for package: %s" %
                  (global_version, artifact_id))
            publish = bintray.publish_uploaded_content(
                subject=bintray_orga,
                repo=bintray_repo,
                package=artifact_id,
                version=global_version,
            )

        print("Maven publication succeeded!")

    else:
        print("Maven publication skipped (already exists)")

except Exception as err:
    sys.stderr.write("Error: %s\n" % str(err))
    if created:
        print("Deleting created version")
        bintray.delete_version(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
            version=global_version,
        )
    elif uploaded:
        print("Deleting uploaded content")
        bintray.discard_uploaded_content(
            subject=bintray_orga,
            repo=bintray_repo,
            package=artifact_id,
            version=global_version,
            passphrase=bintray_gpg_pass,
        )
    exit(1)
