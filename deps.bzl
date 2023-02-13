load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        sum = "h1:U+s90UTSYgptZMwQh2aRr3LuazLJIa+Pg3Kc1ylSYVY=",
        version = "v2.0.0-20190314233015-f79a8a8ca69d",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_fogleman_gg",
        importpath = "github.com/fogleman/gg",
        sum = "h1:/7zJX8F6AaYQc57WQCyN9cAIz+4bCJGO9B+dyW29am8=",
        version = "v1.3.0",
    )

    go_repository(
        name = "com_github_go_ble_ble",
        importpath = "github.com/go-ble/ble",
        sum = "h1:wCW6nm32DzgPEmKK8GPJj0D1ZRGrnUgfiGsXaJoClNc=",
        version = "v0.0.0-20230130210458-dd4b07d15402",
    )
    go_repository(
        name = "com_github_golang_freetype",
        importpath = "github.com/golang/freetype",
        sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
        version = "v0.0.0-20170609003504-e2365dfdc4a0",
    )

    go_repository(
        name = "com_github_juullabs_oss_cbgo",
        importpath = "github.com/JuulLabs-OSS/cbgo",
        sum = "h1:A5JdglvFot1J9qYR0POZ4qInttpsVPN9lqatjaPp2ro=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:mweAR1A6xJ3oS2pRaGiHgQ4OO8tzTaLawm8vnODuwDk=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:6Su7aK7lXmJ/U79bYtBjLNaha4Fs1Rg9plHpcH+vvnE=",
        version = "v0.1.6",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:wuysRhFDzyxgEmMf5xjvJ2M9dZoWAXNNr5LSBS7uHXY=",
        version = "v0.0.12",
    )
    go_repository(
        name = "com_github_mgutz_ansi",
        importpath = "github.com/mgutz/ansi",
        sum = "h1:j7+1HpAFS1zy5+Q4qx1fWh90gTKwiN4QCGoY9TWyyO4=",
        version = "v0.0.0-20170206155736-9520e82c474b",
    )
    go_repository(
        name = "com_github_mgutz_logxi",
        importpath = "github.com/mgutz/logxi",
        sum = "h1:n8cgpHzJ5+EDyDri2s/GC7a9+qK3/YEGnBsd0uS/8PY=",
        version = "v0.0.0-20161027140823-aebf8a7d67ab",
    )
    go_repository(
        name = "com_github_neurlang_wayland",
        importpath = "github.com/neurlang/wayland",
        sum = "h1:qaKR1d2/CMdNRIGaE8z5lnFwFjno/ppJjcd4X6r0Wvw=",
        version = "v0.0.0-20220924004225-434478269863",
    )
    go_repository(
        name = "com_github_nfnt_resize",
        importpath = "github.com/nfnt/resize",
        sum = "h1:zYyBkD/k9seD2A7fsi6Oo2LfFZAehjjQMERAvZLEDnQ=",
        version = "v0.0.0-20180221191011-83c6a9932646",
    )

    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_raff_goble",
        importpath = "github.com/raff/goble",
        sum = "h1:JtoVdxWJ3tgyqtnPq3r4hJ9aULcIDDnPXBWxZsdmqWU=",
        version = "v0.0.0-20190909174656-72afc67d6a99",
    )
    go_repository(
        name = "com_github_rkusa_gm",
        importpath = "github.com/rkusa/gm",
        sum = "h1:isLzfNyCDEFaua2vQtlM7cr5PplqEBHqzDf8AdX6te8=",
        version = "v0.0.0-20160203133140-acc6b6d0b64a",
    )

    go_repository(
        name = "com_github_russross_blackfriday_v2",
        importpath = "github.com/russross/blackfriday/v2",
        sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:1N5EYkVAPEywqZRJd7cwnRtCb6xJx7NH3T3WUTF980Q=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_urfave_cli",
        importpath = "github.com/urfave/cli",
        sum = "h1:gsqYFH8bb9ekPA12kRo0hfjngWQjkJPlN9R0N78BoUo=",
        version = "v1.22.2",
    )
    go_repository(
        name = "com_github_vulkan_go_vulkan",
        importpath = "github.com/vulkan-go/vulkan",
        sum = "h1:PS3d1+TTLuWkae7njZu+RAs9j8s9phKJ2ElSxq2r1t0=",
        version = "v0.0.0-20220916223707-3ea69233b9d2",
    )
    go_repository(
        name = "com_github_yalue_native_endian",
        importpath = "github.com/yalue/native_endian",
        sum = "h1:MZkkXsOTwRURxf/dM/AhbLfRUBQY+CjMdJRXp9bqd2U=",
        version = "v1.0.1",
    )

    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
        version = "v0.0.0-20161208181325-20d25e280405",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
        version = "v2.2.2",
    )
    go_repository(
        name = "org_golang_x_image",
        importpath = "golang.org/x/image",
        sum = "h1:fqpd0EBDzlHRCjiphRR5Zo/RSWWQlWv34418dnEixWk=",
        version = "v0.0.0-20210220032944-ac19c3e999fb",
    )

    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:QAkhGVjOxMa+n4mlsAWeAU+BMZmimQAaNiMu+iUi94E=",
        version = "v0.0.0-20211204120058-94396e421777",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
        version = "v0.3.0",
    )
