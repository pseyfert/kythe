load("//tools:build_rules/shims.bzl", "go_library")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "createcmd",
    srcs = [
        "createcmd.go",
        "flag.go",
    ],
    deps = [
        "//kythe/go/platform/kzip",
        "//kythe/go/platform/tools/kzip/flags",
        "//kythe/go/platform/vfs",
        "//kythe/go/util/cmdutil",
        "//kythe/go/util/flagutil",
        "//kythe/go/util/kytheuri",
        "//kythe/go/util/vnameutil",
        "//kythe/proto:analysis_go_proto",
        "//kythe/proto:storage_go_proto",
        "@com_github_golang_protobuf//jsonpb:go_default_library_gen",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library_gen",
        "@com_github_google_subcommands//:go_default_library",
        "@io_bazel_rules_go//proto/wkt:any_go_proto",
    ],
)
