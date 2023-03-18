import os
import requests
import json
from github import Github
import subprocess

def generate_changelog(diff):
    # Use the OpenAI API to generate a changelog
    api_key = os.environ["OPENAI_API_KEY"]
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json",
    }
    data = {
        "prompt": f"Generate a changelog for a new software release based on the following git diff:\n\n{diff}\n",
        "max_tokens": 200,
    }
    response = requests.post(
        "https://api.openai.com/v1/engines/davinci-codex/completions",
        headers=headers,
        data=json.dumps(data),
    )
    changelog = response.json()["choices"][0]["text"].strip()
    return changelog

def get_latest_and_previous_tags():
    g = Github(os.environ["GITHUB_TOKEN"])
    repo = g.get_repo(f"{os.environ['GITHUB_REPOSITORY']}")
    tags = repo.get_tags()
    return tags[0], tags[1]

def get_git_diff(latest_tag, previous_tag):
    diff = subprocess.check_output(["git", "diff", previous_tag.name, latest_tag.name], text=True)
    return diff

if __name__ == "__main__":
    latest_tag, previous_tag = get_latest_and_previous_tags()
    diff = get_git_diff(latest_tag, previous_tag)
    changelog = generate_changelog(diff)
    latest_release = get_latest_release()
    new_body = f"{latest_release.body}\n\n{changelog}"
    latest_release.update_release(latest_release.title, new_body)
    print(f"::set-output name=changelog::{changelog}")
