load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")

def libxkbcommon_repositories():
    maybe(
        http_archive,
        name = "libxkbcommon",
        build_file = Label("//third_party/libxkbcommon:BUILD.libxkbcommon.bazel"),
        sha256 = "560f11c4bbbca10f495f3ef7d3a6aa4ca62b4f8fb0b52e7d459d18a26e46e017",
        strip_prefix = "libxkbcommon-1.5.0",
        urls = [
            "https://xkbcommon.org/download/libxkbcommon-1.5.0.tar.xz",
        ],
    )

    maybe(
        http_archive,
        name = "wayland",
        build_file = Label("//third_party/libxkbcommon:BUILD.wayland.bazel"),
        sha256 = "6dc64d7fc16837a693a51cfdb2e568db538bfdc9f457d4656285bb9594ef11ac",
        strip_prefix = "wayland-1.21.0",
        urls = [
            "https://gitlab.freedesktop.org/wayland/wayland/-/releases/1.21.0/downloads/wayland-1.21.0.tar.xz",
        ],
    )

    maybe(
        http_archive,
        name = "libffi",
        build_file = Label("//third_party/libxkbcommon:BUILD.libffi.bazel"),
        sha256 = "d66c56ad259a82cf2a9dfc408b32bf5da52371500b84745f7fb8b645712df676",
        strip_prefix = "libffi-3.4.4",
        urls = [
            "https://github.com/libffi/libffi/releases/download/v3.4.4/libffi-3.4.4.tar.gz",
        ],
    )

    maybe(
        http_archive,
        name = "libxml2",
        build_file = Label("//third_party/libxkbcommon:BUILD.libxml2.bazel"),
        sha256 = "5d2cc3d78bec3dbe212a9d7fa629ada25a7da928af432c93060ff5c17ee28a9c",
        strip_prefix = "libxml2-2.10.3",
        urls = [
            "https://download.gnome.org/sources/libxml2/2.10/libxml2-2.10.3.tar.xz",
        ],
    )
