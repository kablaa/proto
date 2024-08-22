load("@rules_proto_grpc_go//:defs.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "test_proto",
    srcs = ["test.proto"],
)

go_proto_library(
    name = "test_go_proto",
    protos = [":test_proto"]
)