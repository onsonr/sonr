If you want to push a tag, run the following command first. It removes all local tags and pulls the remote tags so you don't clobber them.

`git tag -l | xargs git tag -d && git fetch --tags`
