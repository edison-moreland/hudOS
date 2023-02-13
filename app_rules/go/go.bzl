load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("//app_rules:app.bzl", "hudapp")

def hudapp_go(app, embed, setcap = ""):
    go_binary(
        name = "{}_arm".format(app),
        embed = embed,
        visibility = ["//visibility:public"],
        goos = "linux",
        goarch = "arm64",
    )

    hudapp_go_service(
        name = "systemd_unit",
        app = app,
        prefix = "/opt/hud",
    )

    hudapp_go_post_deploy(
        name = "post_deploy",
        app = app,
        setcap = setcap,
        prefix = "/opt/hud",
    )

    hudapp(
        app = app,
        binaries = [
            ":{}_arm".format(app),
        ],
        systemd_units = [
            ":systemd_unit",
        ],
        post_deploy = [
            ":post_deploy",
        ],
    )

def _hudapp_go_service_impl(ctx):
    out_service = ctx.actions.declare_file("hud_{}.service".format(ctx.attr.app))
    ctx.actions.expand_template(
        template = ctx.file._service_template,
        output = out_service,
        substitutions = {
            "{prefix}": ctx.attr.prefix,
            "{app}": ctx.attr.app,
        },
    )

    return [
        DefaultInfo(files = depset([out_service])),
    ]

hudapp_go_service = rule(
    implementation = _hudapp_go_service_impl,
    attrs = {
        "prefix": attr.string(mandatory = True),
        "app": attr.string(mandatory = True),
        "_service_template": attr.label(
            allow_single_file = [".service"],
            default = Label("//app_rules/go:template.service"),
        ),
    },
)

def _hudapp_go_post_deploy_impl(ctx):
    out_script = ctx.actions.declare_file("{}-{}.sh".format(ctx.attr.priority, ctx.attr.app))
    ctx.actions.expand_template(
        template = ctx.file._post_deploy_template,
        output = out_script,
        substitutions = {
            "{prefix}": ctx.attr.prefix,
            "{app}": ctx.attr.app,
            "{setcap}": ctx.attr.setcap,
        },
    )

    return [
        DefaultInfo(files = depset([out_script])),
    ]

hudapp_go_post_deploy = rule(
    implementation = _hudapp_go_post_deploy_impl,
    attrs = {
        "prefix": attr.string(mandatory = True),
        "app": attr.string(mandatory = True),
        "setcap": attr.string(mandatory = False),
        "priority": attr.string(default = "100"),
        "_post_deploy_template": attr.label(
            allow_single_file = [".sh"],
            default = Label("//app_rules/go:post_deploy.sh"),
        ),
    },
)
