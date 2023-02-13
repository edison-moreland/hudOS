"""Rules for building hud apps"""

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")
load("@rules_pkg//pkg:mappings.bzl", "pkg_attributes", "pkg_filegroup", "pkg_files")

def hudapp(app, systemd_units, binaries = [], post_deploy = [], configs = []):
    pkg_files(
        name = "hudapp_bin",
        prefix = "bin",
        srcs = binaries,
        attributes = pkg_attributes(
            mode = "700",
            user = "hud",
            group = "hud",
        ),
    )

    pkg_files(
        name = "hudapp_services",
        prefix = "services",
        srcs = systemd_units,
    )

    pkg_files(
        name = "hudapp_post_deploy",
        prefix = "post_deploy",
        srcs = post_deploy,
    )

    pkg_files(
        name = "hudapp_configs",
        prefix = "configs",
        srcs = configs,
    )

    pkg_filegroup(
        name = "hudapp",
        prefix = "",
        srcs = [
            ":hudapp_bin",
            ":hudapp_services",
            ":hudapp_post_deploy",
            ":hudapp_configs",
        ],
        visibility = ["//visibility:public"],
    )

    pkg_tar(
        name = "hudapp_{}".format(app),
        srcs = [
            ":hudapp",
        ],
        visibility = ["//visibility:public"],
    )
