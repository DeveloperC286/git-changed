#!/usr/bin/env sh

set -o errexit

# Get current version.
current_version=$(cat VERSION)

# If the tag already exist then exit.
new_tag=${current_version}
git tag -l | grep -q "^${new_tag}$" && exit 0

# Get latest tag.
latest_tag=$(git tag --sort=-committerdate | head -1)

# Generate the release description.
release_description=$("${CARGO_HOME}/bin/git-cliff" "${latest_tag}.." --tag "${new_tag}" --strip all)

# Create the new release.
/usr/local/bin/release-cli create \
	--name "${new_tag}" \
	--description "${release_description}" \
	--tag-name "${new_tag}" \
	--ref "${CI_COMMIT_SHA}" \
	--assets-link '{"name":"linux-amd64-binary.zip","url":"https://gitlab.com/DeveloperC/git-changed/-/jobs/artifacts/'${new_tag}'/download?job=release-binary-compiling-linux-amd64"}' \
	--assets-link '{"name":"darwin-amd64-binary.zip","url":"https://gitlab.com/DeveloperC/git-changed/-/jobs/artifacts/'${new_tag}'/download?job=release-binary-compiling-darwin-amd64"}'
