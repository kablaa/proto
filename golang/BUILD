# load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "golang",
    srcs = ["golang.go"], 
    deps = ["//test:test_go_grpc"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "main",
    embed = [":golang"]
)