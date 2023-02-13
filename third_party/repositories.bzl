load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("//third_party/libxkbcommon:libxkbcommon_repositories.bzl", "libxkbcommon_repositories")
load("//third_party/tools:wayland_scanner_build.bzl", "wayland_scanner_repositories")

def repositories():
    wayland_scanner_repositories()
    libxkbcommon_repositories()
