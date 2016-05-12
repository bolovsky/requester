load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "go_binary", "go_library")

go_prefix("github.com/bolovsky/requester")

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
