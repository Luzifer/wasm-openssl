#!/bin/bash
set -euo pipefail

(which zip 2>&1 1>/dev/null) || {
	(which apk 2>&1 1>/dev/null) && apk add --update zip
	(which apt-get 2>&1 1>/dev/null) && apt-get update && apt-get install -y zip
}

function step() {
	echo "===> $@..."
}

VERSION=$(git describe --tags --exact-match || echo "ghpublish__notags")
PWD=$(pwd)
godir=${PWD/${GOPATH}\/src\//}
REPO=${REPO:-$(echo ${godir} | cut -d '/' -f 3)}
GHUSER=${GHUSER:-$(echo ${godir} | cut -d '/' -f 2)}
DEPLOYMENT_TAG=${DEPLOYMENT_TAG:-${VERSION}}
PACKAGES=${PACKAGES:-$(echo ${godir} | cut -d '/' -f 1-3)}
BUILD_DIR=${BUILD_DIR:-.build}
DRAFT=${DRAFT:-true}

go version

step "Retrieve dependencies"
go get github.com/aktau/github-release

step "Cleanup build directory if present"
rm -rf ${BUILD_DIR}

step "Compile program"
mkdir ${BUILD_DIR}
GOOS=js GOARCH=wasm go build -ldflags="-X main.version=${VERSION}"  \
	-o "${BUILD_DIR}/openssl.wasm" \
	${PACKAGES}

step "Generate binary SHASUMs"
cd ${BUILD_DIR}
sha256sum * >SHA256SUMS

step "Packing archives"
for file in *; do
	if [ "${file}" = "SHA256SUMS" ]; then
		continue
	fi

  tar -czf "${file%%.*}.tar.gz" "${file}"

	rm "${file}"
done

step "Generate archive SHASUMs"
sha256sum * >>SHA256SUMS
grep -v 'SHA256SUMS' SHA256SUMS >SHA256SUMS.tmp
mv SHA256SUMS.tmp SHA256SUMS

step "Publish builds to Github"

if (test "${VERSION}" == "ghpublish__notags"); then
	echo "No tag present, stopping build now."
	exit 0
fi

if [ -z "${GITHUB_TOKEN}" ]; then
	echo "Please set \$GITHUB_TOKEN environment variable"
	exit 1
fi

if [[ "${DRAFT}" == "true" ]]; then
	step "Create a drafted release"
	github-release release --user ${GHUSER} --repo ${REPO} --tag ${DEPLOYMENT_TAG} --name ${DEPLOYMENT_TAG} --draft || true
else
	step "Create a published release"
	github-release release --user ${GHUSER} --repo ${REPO} --tag ${DEPLOYMENT_TAG} --name ${DEPLOYMENT_TAG} || true
fi

step "Upload build assets"
for file in *; do
	echo "- ${file}"
	github-release upload --user ${GHUSER} --repo ${REPO} --tag ${DEPLOYMENT_TAG} --name ${file} --file ${file}
done

echo -e "\n\n=== Recorded checksums ==="
cat SHA256SUMS

cd -
