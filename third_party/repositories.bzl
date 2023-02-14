load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("//third_party/libxkbcommon:libxkbcommon_repositories.bzl", "libxkbcommon_repositories")

def repositories():
    libxkbcommon_repositories()
