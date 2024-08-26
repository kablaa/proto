# load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc_go//:defs.bzl", "go_proto_library", "go_grpc_library")

proto_library(
    name = "test_proto",
    srcs = ["test.proto"],
)

go_proto_library(
    name = "test_grpc",
    protos = [":test_proto"],
    visibility = ["//visibility:public"],
)

go_grpc_library(
    name = "test_go_grpc",
    protos = [":test_proto"],
    visibility = ["//visibility:public"],
)
