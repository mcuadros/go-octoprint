load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")
load("@bazel_gazelle//:def.bzl", "gazelle")

buildifier(
    name = "buildifier",
)

# gazelle:prefix github.com/conrey-engineering/go-octoprint
gazelle(name = "gazelle")

go_library(
    name = "go-octoprint",
    srcs = [
        "client.go",
        "common.go",
        "connection.go",
        "files.go",
        "job.go",
        "printer.go",
        "settings.go",
        "system.go",
        "version.go",
    ],
    importpath = "github.com/conrey-engineering/go-octoprint",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go-octoprint_test",
    srcs = [
        "common_test.go",
        "files_test.go",
        "printer_test.go",
        "settings_test.go",
        "system_test.go",
    ],
    embed = [":go-octoprint"],
    deps = ["@com_github_stretchr_testify//assert"],
)
