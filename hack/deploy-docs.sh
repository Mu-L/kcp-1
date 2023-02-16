#!/usr/bin/env bash

# Copyright 2023 The KCP Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

REPO_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)

if [[ -n "${GITHUB_REF_NAME:-}" ]]; then
  VERSION="${VERSION:-$GITHUB_REF_NAME}"
else
  VERSION=${VERSION:-$(git rev-parse --abbrev-ref HEAD)}
fi

if echo "$VERSION" | grep '^release-[0-9]'; then
  VERSION=v$(echo "$VERSION" | cut -d - -f 2)
elif echo "$VERSION" | grep '^v[0-9]\+\.[0-9]\+'; then
  VERSION=$(echo "$VERSION" | grep -o '^v[0-9]\+\.[0-9]\+')
fi

MIKE_OPTIONS=()

if [[ -n "${CI:-}" ]]; then
  MIKE_OPTIONS+=(--push)
  git config user.name kcp-docs-bot
  git config user.email no-reply@kcp.io
fi

if [[ -n "${LOCAL:-}" ]]; then
  mike deploy --config-file "${REPO_ROOT}/docs/mkdocs.yml" "${MIKE_OPTIONS[@]}" "$VERSION"
else
  docker run --rm -it \
    -v "$REPO_ROOT/.git":/.git \
    -v "$REPO_ROOT/docs":/docs \
    -w /.git \
    kcp-docs \
    mike deploy --config-file /docs/mkdocs.yml "${MIKE_OPTIONS[@]}" "$VERSION"
fi
