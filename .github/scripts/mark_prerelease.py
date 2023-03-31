import os
import re
from github import Github

# Set up a regex pattern to match alpha and beta versions
pattern = re.compile(r".*-(alpha|beta).*", re.IGNORECASE)

# Authenticate with the GitHub API
g = Github(os.environ["GITHUB_TOKEN"])

# Get the current repository
repo = g.get_repo(os.environ["GITHUB_REPOSITORY"])

# Iterate through all releases
for release in repo.get_releases():
    if pattern.match(release.tag_name) and not release.prerelease:
        print(f"Marking {release.tag_name} as pre-release")
        release.update_release(release.title, release.raw_data["body"], prerelease=True)
