#!/usr/bin/env python3
from bintray.bintray import Bintray
import yaml
import os
import sys
import datetime

try:
    # Flag used to discard uploaded files on failure
    uploaded = False
    published = False

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
    bintray_repo = manifest["global"]["demo_app"]["repo"]
    bintray_package = manifest["ios_demo_app"]["package"]
    description = manifest["ios_demo_app"]["description"]
    licenses = [license["short_name"]
                for license in manifest["global"]["licenses"]]
    vcs_url = manifest["global"]["github"]["git_url"]
    website = manifest["global"]["github"]["url"]
    issue_tracker = manifest["global"]["github"]["issues_url"]
    github_repo = manifest["global"]["github"]["repo"]
    github_notes = manifest["ios_demo_app"]["github_release_notes_file"]
    download_count = manifest["ios_demo_app"]["public_download_numbers"]
    readme_content = manifest["ios_demo_app"]["readme_content"]
    readme_syntax = manifest["ios_demo_app"]["readme_syntax"]
    publish = manifest["ios_demo_app"]["publish"]
    override = manifest["ios_demo_app"]["override"]
    filename = manifest["ios_demo_app"]["filename"]

    # Get version from env (CI) or fail
    if "GOMOBILE_IPFS_VERSION" in os.environ:
        global_version = os.getenv("GOMOBILE_IPFS_VERSION")
    else:
        raise Exception("can't publish a dev version")

    # Check if remote package already exists
    package = bintray.search_package(
        subject=bintray_orga,
        repo=bintray_repo,
        package=bintray_package,
    )

    # If remote package doesn't exist, create it
    if len(package) is 1:
        print("Creating new package: %s" % bintray_package)
        bintray.create_package(
            subject=bintray_orga,
            repo=bintray_repo,
            package=bintray_package,
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
        print("Updating existing package: %s" % bintray_package)
        bintray.update_package(
            subject=bintray_orga,
            repo=bintray_repo,
            package=bintray_package,
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
            package=bintray_package,
        )
        print("Updating existing readme for package: %s" % bintray_package)
    except Exception:
        print("Creating readme for package: %s" % bintray_package)

    # Create or update remote readme
    if readme_content is not None:
        syntax = readme_syntax or "markdown"
        bintray.create_readme(
            subject=bintray_orga,
            repo=bintray_repo,
            package=bintray_package,
            bintray_syntax=syntax,
            bintray_content=readme_content,
        )
    elif github_notes is not None:
        bintray.create_readme(
            subject=bintray_orga,
            repo=bintray_repo,
            package=bintray_package,
            github=github_repo,
        )

    # Check if remote version already exists
    try:
        bintray.get_version(
            subject=bintray_orga,
            repo=bintray_repo,
            package=bintray_package,
            version=global_version,
        )
        version_missing = False
    except Exception:
        version_missing = True

    version_description = "{0}-{1}-{2}".format(
        bintray_repo,
        bintray_package,
        global_version,
    )
    vcs_tag = "v%s" % global_version

    # If remote version doesn't exist or override is enabled
    if version_missing or override:
        if version_missing:
            # Create the remote version
            print("Creating new version: %s for package: %s" %
                  (global_version, bintray_package))
            bintray.create_version(
                subject=bintray_orga,
                repo=bintray_repo,
                package=bintray_package,
                version=global_version,
                description=version_description,
                released=datetime.datetime.now().isoformat(),
                vcs_tag=vcs_tag
            )
        else:
            # Update the existing remote version
            print("Updating existing version: %s for package: %s" %
                  (global_version, bintray_package))
            bintray.update_version(
                subject=bintray_orga,
                repo=bintray_repo,
                package=bintray_package,
                version=global_version,
                description=version_description,
                vcs_tag=vcs_tag
            )

        # Upload artifact
        artifacts = [
            "%s-%s.ipa" % (filename, global_version),
            "%s-%s.plist" % (filename, global_version),
        ]

        ios_build_dir_app = os.path.join(
            os.path.dirname(os.path.dirname(current_dir)),
            "build/ios/app",
        )
        version_path = "%s/%s" % (bintray_package, global_version)
        artifacts_local_dir = os.path.join(ios_build_dir_app, global_version)

        print("Uploading artifacts:")
        for artifact in artifacts:
            sys.stdout.write("    - artifact: %s ..." % artifact)
            sys.stdout.flush()
            bintray.upload_content(
                subject=bintray_orga,
                repo=bintray_repo,
                package=bintray_package,
                version=global_version,
                remote_file_path=os.path.join(version_path, artifact),
                local_file_path=os.path.join(artifacts_local_dir, artifact),
                override=override,
            )
            sys.stdout.write("\b\b\b- done\n")
            uploaded = True

        print("Signing version: %s for package: %s" %
              (global_version, bintray_package))
        bintray.gpg_sign_version(
            subject=bintray_orga,
            repo=bintray_repo,
            package=bintray_package,
            version=global_version,
            passphrase=bintray_gpg_pass,
        )

        if publish:
            print("Publishing version: %s for package: %s" %
                  (global_version, bintray_package))
            publish = bintray.publish_uploaded_content(
                subject=bintray_orga,
                repo=bintray_repo,
                package=bintray_package,
                version=global_version,
            )
            published = True

        print("Bintray publication succeeded!")

    else:
        print("Bintray publication skipped (already exists)")

except Exception as err:
    sys.stderr.write("Error: %s\n" % str(err))
    if published:
        print("Deleting created version")
        bintray.delete_version(
            subject=bintray_orga,
            repo=bintray_repo,
            package=bintray_package,
            version=global_version,
        )
    elif uploaded:
        print("Deleting uploaded content")
        bintray.discard_uploaded_content(
            subject=bintray_orga,
            repo=bintray_repo,
            package=bintray_package,
            version=global_version,
            passphrase=bintray_gpg_pass,
        )
    exit(1)
