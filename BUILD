load("@io_bazel_rules_go//go:def.bzl", "go_prefix")

go_prefix("github.com/bolovsky/requester")

load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
  name = "server",
  srcs = ["Server/Server.go"],
)

go_library(
  name = "webworker",
  srcs = glob(["WebWorker/*.go"]),
)

go_binary(
  name = "requester",
  srcs = glob(["main.go"]),
  deps = [":server", ":webworker"],
)
