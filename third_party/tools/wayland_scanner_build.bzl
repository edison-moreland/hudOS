load("@rules_foreign_cc//foreign_cc:defs.bzl", "meson")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def wayland_scanner_repositories():
    _ALL_CONTENT = """\
filegroup(
    name = "all_srcs",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"],
)
"""

    maybe(
        http_archive,
        name = "wayland_tool",
        build_file_content = _ALL_CONTENT,
        sha256 = "6dc64d7fc16837a693a51cfdb2e568db538bfdc9f457d4656285bb9594ef11ac",
        strip_prefix = "wayland-1.21.0",
        urls = [
            "https://gitlab.freedesktop.org/wayland/wayland/-/releases/1.21.0/downloads/wayland-1.21.0.tar.xz",
        ],
    )

def wayland_scanner_tool(name, srcs, **kwargs):
    tags = ["manual"] + kwargs.pop("tags", [])

    meson(
        name = "{}.build".format(name),
        lib_source = srcs,
        options = {
            "libraries": "false",
            "scanner": "true",
            "tests": "false",
            "documentation": "false",
            "dtd_validation": "false",
            "pkgconfig.relocatable": "true",
        },
        visibility = ["//visibility:public"],
        out_binaries = [
            "wayland-scanner",
        ],
        out_data_dirs = [
            "lib/",
        ],
        tags = tags,
        **kwargs
    )

    native.filegroup(
        name = name,
        srcs = ["{}.build".format(name)],
        output_group = "gen_dir",
        tags = tags,
    )

def _current_toolchain_impl(ctx):
    toolchain = ctx.toolchains[ctx.attr._toolchain]

    if toolchain.data.target:
        return [
            toolchain,
            platform_common.TemplateVariableInfo(toolchain.data.env),
            DefaultInfo(
                files = toolchain.data.target.files,
                runfiles = toolchain.data.target.default_runfiles,
            ),
        ]
    return [
        toolchain,
        platform_common.TemplateVariableInfo(toolchain.data.env),
        DefaultInfo(),
    ]

current_wayland_scanner_toolchain = rule(
    implementation = _current_toolchain_impl,
    attrs = {
        "_toolchain": attr.string(default = str(Label("//third_party/tools:wayland_scanner_toolchain"))),
    },
    incompatible_use_toolchain_transition = True,
    toolchains = [
        str(Label("//third_party/tools:wayland_scanner_toolchain")),
    ],
)
