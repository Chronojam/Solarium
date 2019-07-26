workspace(name = "com_github_chronojam_solarium")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

# Protobuf

#http_archive(
#    name = "com_google_protobuf",
#    sha256 = "8eb5ca331ab8ca0da2baea7fc0607d86c46c80845deca57109a5d637ccb93bb4",
#    strip_prefix = "protobuf-3.9.0",
#    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.9.0.zip"],
#)
#load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
http_archive(
    name = "com_google_protobuf",
    sha256 = "5eb85831c3fcdacfe18f00f9b258cba0b81ca89ad63b80b835ca9f00693fdd5c",
    strip_prefix = "protobuf-3.9.0-rc1",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.9.0-rc1.zip"],
)
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
protobuf_deps()
http_archive(
    name = "com_github_grpc_grpc",
    urls = ["https://github.com/grpc/grpc/archive/v1.22.0.zip"],
    strip_prefix = "grpc-1.22.0",
)

http_archive(
    name = "build_stack_rules_proto",
    urls = ["https://github.com/stackb/rules_proto/archive/b93b544f851fdcd3fc5c3d47aee3b7ca158a8841.tar.gz"],
    sha256 = "c62f0b442e82a6152fcd5b1c0b7c4028233a9e314078952b6b04253421d56d61",
    strip_prefix = "rules_proto-b93b544f851fdcd3fc5c3d47aee3b7ca158a8841",
)

load("@build_stack_rules_proto//csharp:deps.bzl", "csharp_proto_compile")

csharp_proto_compile()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

# Register C# toolchains.

http_archive(
    name = "io_bazel_rules_dotnet",
    urls = [
        "https://github.com/bazelbuild/rules_dotnet/archive/0.0.3.tar.gz",
    ],
    strip_prefix = "rules_dotnet-0.0.3",
)

load(
    "@io_bazel_rules_dotnet//dotnet:defs.bzl",
    "core_register_sdk",
    "net_register_sdk",
    "mono_register_sdk",
    "dotnet_register_toolchains",
    "dotnet_repositories",
    "nuget_package",
)

core_version = "v2.1.503"

dotnet_register_toolchains(
    core_version  = core_version,
)

# For .NET Core:
core_register_sdk(
    core_version = core_version,
    name = "core_sdk",
)
# For .NET Framework:
#net_register_sdk(
#    "net471",
#    name = "net_sdk",
#)

# For Mono:
#mono_register_sdk()
dotnet_repositories()
load("@build_stack_rules_proto//csharp/nuget:packages.bzl", nuget_packages = "packages")
nuget_packages()

# Adding these manually down below.
#load("@build_stack_rules_proto//csharp/nuget:nuget.bzl", "nuget_protobuf_packages")
#nuget_protobuf_packages()
#load("@build_stack_rules_proto//csharp/nuget:nuget.bzl", "nuget_grpc_packages")
#nuget_grpc_packages()

# Register Golang toolchains.

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
    ],
    sha256 = "8df59f11fb697743cbb3f26cfb8750395f30471e9eabde0d174c3aebc7a1cd39",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
    ],
    sha256 = "be9296bfd64882e3c08e3283c58fcb461fa6dd3c171764fcc4cf322f60615a9b",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# Golang Dependencies

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:rTIdg5QFRR7XCaK4LCjBiPbx8j4DQRpdYMnGn/bJUEU=",
    version = "v0.0.0-20190628185345-da137c7871d7",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:Ygq9/SRJX9+dU0WCIICM8RkWvDw03lvB77hrhJnpxfU=",
    version = "v0.0.0-20190716160619-c506a9f90610",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:J0UbZOIrCAl+fpTOf8YLs4dJo8L/owV4LYVtAXQoPkw=",
    version = "v1.22.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:LepdCS8Gf/MVejFIt8lsiexZATdoGVyp5bcyS+rYoUI=",
    version = "v0.0.0-20190712062909-fae7ac547cb7",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_bwmarrin_discordgo",
    importpath = "github.com/bwmarrin/discordgo",
    sum = "h1:kMED/DB0NR1QhRcalb85w0Cu3Ep2OrGAqZH1R5awQiY=",
    version = "v0.19.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:HuIa8hRrWRSrqYzx1qI49NNxhdi2PrY7gxVSq1JjLDc=",
    version = "v0.0.0-20190701094942-4def268fd1a4",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:WDFjx/TMzVgy9VdMMQi2K2Emtwi2QcUQsztZ/zLaH/Q=",
    version = "v1.4.0",
)


### Generated by the tool
nuget_package(
    name = "system.interactive.async",
    package = "system.interactive.async",
    version = "3.2.0",
    sha256 = "8d8c3296247b4e86c7d0bdaae6bdf6447939e2cf59e613debb9231da8e2fb978",
    core_lib = {
        "netcoreapp2.0": "lib/netstandard2.0/System.Interactive.Async.dll",
        "netcoreapp2.1": "lib/netstandard2.0/System.Interactive.Async.dll",
    },
    net_lib = {
        "net45": "lib/net45/System.Interactive.Async.dll",
        "net451": "lib/net45/System.Interactive.Async.dll",
        "net452": "lib/net45/System.Interactive.Async.dll",
        "net46": "lib/net46/System.Interactive.Async.dll",
        "net461": "lib/net46/System.Interactive.Async.dll",
        "net462": "lib/net46/System.Interactive.Async.dll",
        "net47": "lib/net46/System.Interactive.Async.dll",
        "net471": "lib/net46/System.Interactive.Async.dll",
        "net472": "lib/net46/System.Interactive.Async.dll",
        "netstandard1.0": "lib/netstandard1.0/System.Interactive.Async.dll",
        "netstandard1.1": "lib/netstandard1.0/System.Interactive.Async.dll",
        "netstandard1.2": "lib/netstandard1.0/System.Interactive.Async.dll",
        "netstandard1.3": "lib/netstandard1.3/System.Interactive.Async.dll",
        "netstandard1.4": "lib/netstandard1.3/System.Interactive.Async.dll",
        "netstandard1.5": "lib/netstandard1.3/System.Interactive.Async.dll",
        "netstandard1.6": "lib/netstandard1.3/System.Interactive.Async.dll",
        "netstandard2.0": "lib/netstandard2.0/System.Interactive.Async.dll",
    },
    mono_lib = "lib/net46/System.Interactive.Async.dll",
    core_files = {
        "netcoreapp2.0": [
           "lib/netstandard2.0/System.Interactive.Async.dll",
           "lib/netstandard2.0/System.Interactive.Async.xml",
        ],
        "netcoreapp2.1": [
           "lib/netstandard2.0/System.Interactive.Async.dll",
           "lib/netstandard2.0/System.Interactive.Async.xml",
        ],
    },
    net_files = {
        "net45": [
           "lib/net45/System.Interactive.Async.dll",
           "lib/net45/System.Interactive.Async.xml",
        ],
        "net451": [
           "lib/net45/System.Interactive.Async.dll",
           "lib/net45/System.Interactive.Async.xml",
        ],
        "net452": [
           "lib/net45/System.Interactive.Async.dll",
           "lib/net45/System.Interactive.Async.xml",
        ],
        "net46": [
           "lib/net46/System.Interactive.Async.dll",
           "lib/net46/System.Interactive.Async.xml",
        ],
        "net461": [
           "lib/net46/System.Interactive.Async.dll",
           "lib/net46/System.Interactive.Async.xml",
        ],
        "net462": [
           "lib/net46/System.Interactive.Async.dll",
           "lib/net46/System.Interactive.Async.xml",
        ],
        "net47": [
           "lib/net46/System.Interactive.Async.dll",
           "lib/net46/System.Interactive.Async.xml",
        ],
        "net471": [
           "lib/net46/System.Interactive.Async.dll",
           "lib/net46/System.Interactive.Async.xml",
        ],
        "net472": [
           "lib/net46/System.Interactive.Async.dll",
           "lib/net46/System.Interactive.Async.xml",
        ],
        "netstandard1.0": [
           "lib/netstandard1.0/System.Interactive.Async.dll",
           "lib/netstandard1.0/System.Interactive.Async.xml",
        ],
        "netstandard1.1": [
           "lib/netstandard1.0/System.Interactive.Async.dll",
           "lib/netstandard1.0/System.Interactive.Async.xml",
        ],
        "netstandard1.2": [
           "lib/netstandard1.0/System.Interactive.Async.dll",
           "lib/netstandard1.0/System.Interactive.Async.xml",
        ],
        "netstandard1.3": [
           "lib/netstandard1.3/System.Interactive.Async.dll",
           "lib/netstandard1.3/System.Interactive.Async.xml",
        ],
        "netstandard1.4": [
           "lib/netstandard1.3/System.Interactive.Async.dll",
           "lib/netstandard1.3/System.Interactive.Async.xml",
        ],
        "netstandard1.5": [
           "lib/netstandard1.3/System.Interactive.Async.dll",
           "lib/netstandard1.3/System.Interactive.Async.xml",
        ],
        "netstandard1.6": [
           "lib/netstandard1.3/System.Interactive.Async.dll",
           "lib/netstandard1.3/System.Interactive.Async.xml",
        ],
        "netstandard2.0": [
           "lib/netstandard2.0/System.Interactive.Async.dll",
           "lib/netstandard2.0/System.Interactive.Async.xml",
        ],
    },
    mono_files = [
        "lib/net46/System.Interactive.Async.dll",
        "lib/net46/System.Interactive.Async.xml",
    ],
)
nuget_package(
    name = "grpc.core.api",
    package = "grpc.core.api",
    version = "1.22.0",
    sha256 = "75517461b8c601ea85e444bfc00edbea930b789d11f09c59d667c0dbd6d5d5cd",
    core_lib = {
        "netcoreapp2.0": "lib/netstandard2.0/Grpc.Core.Api.dll",
        "netcoreapp2.1": "lib/netstandard2.0/Grpc.Core.Api.dll",
    },
    net_lib = {
        "net45": "lib/net45/Grpc.Core.Api.dll",
        "net451": "lib/net45/Grpc.Core.Api.dll",
        "net452": "lib/net45/Grpc.Core.Api.dll",
        "net46": "lib/net45/Grpc.Core.Api.dll",
        "net461": "lib/net45/Grpc.Core.Api.dll",
        "net462": "lib/net45/Grpc.Core.Api.dll",
        "net47": "lib/net45/Grpc.Core.Api.dll",
        "net471": "lib/net45/Grpc.Core.Api.dll",
        "net472": "lib/net45/Grpc.Core.Api.dll",
        "netstandard1.5": "lib/netstandard1.5/Grpc.Core.Api.dll",
        "netstandard1.6": "lib/netstandard1.5/Grpc.Core.Api.dll",
        "netstandard2.0": "lib/netstandard2.0/Grpc.Core.Api.dll",
    },
    mono_lib = "lib/net45/Grpc.Core.Api.dll",
    core_deps = {
        "net45": [
           "@system.interactive.async//:net45_net",
        ],
        "net451": [
           "@system.interactive.async//:net451_net",
        ],
        "net452": [
           "@system.interactive.async//:net452_net",
        ],
        "net46": [
           "@system.interactive.async//:net46_net",
        ],
        "net461": [
           "@system.interactive.async//:net461_net",
        ],
        "net462": [
           "@system.interactive.async//:net462_net",
        ],
        "net47": [
           "@system.interactive.async//:net47_net",
        ],
        "net471": [
           "@system.interactive.async//:net471_net",
        ],
        "net472": [
           "@system.interactive.async//:net472_net",
        ],
        "netstandard1.5": [
           "@system.interactive.async//:netstandard1.5_net",
        ],
        "netstandard1.6": [
           "@system.interactive.async//:netstandard1.6_net",
        ],
        "netstandard2.0": [
           "@system.interactive.async//:netstandard2.0_net",
           "@system.memory//:netstandard2.0_net",
        ],
    },
    net_deps = {
        "net45": [
           "@system.interactive.async//:net45_net",
        ],
        "net451": [
           "@system.interactive.async//:net451_net",
        ],
        "net452": [
           "@system.interactive.async//:net452_net",
        ],
        "net46": [
           "@system.interactive.async//:net46_net",
        ],
        "net461": [
           "@system.interactive.async//:net461_net",
        ],
        "net462": [
           "@system.interactive.async//:net462_net",
        ],
        "net47": [
           "@system.interactive.async//:net47_net",
        ],
        "net471": [
           "@system.interactive.async//:net471_net",
        ],
        "net472": [
           "@system.interactive.async//:net472_net",
        ],
        "netstandard1.5": [
           "@system.interactive.async//:netstandard1.5_net",
        ],
        "netstandard1.6": [
           "@system.interactive.async//:netstandard1.6_net",
        ],
        "netstandard2.0": [
           "@system.interactive.async//:netstandard2.0_net",
           "@system.memory//:netstandard2.0_net",
        ],
    },
    mono_deps = [
        "@system.interactive.async//:mono",
    ],
    core_files = {
        "netcoreapp2.0": [
           "lib/netstandard2.0/Grpc.Core.Api.dll",
           "lib/netstandard2.0/Grpc.Core.Api.pdb",
           "lib/netstandard2.0/Grpc.Core.Api.xml",
        ],
        "netcoreapp2.1": [
           "lib/netstandard2.0/Grpc.Core.Api.dll",
           "lib/netstandard2.0/Grpc.Core.Api.pdb",
           "lib/netstandard2.0/Grpc.Core.Api.xml",
        ],
    },
    net_files = {
        "net45": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "net451": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "net452": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "net46": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "net461": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "net462": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "net47": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "net471": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "net472": [
           "lib/net45/Grpc.Core.Api.dll",
           "lib/net45/Grpc.Core.Api.pdb",
           "lib/net45/Grpc.Core.Api.xml",
        ],
        "netstandard1.0": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.1": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.2": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.3": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.4": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.5": [
           "lib/netstandard1.5/Grpc.Core.Api.dll",
           "lib/netstandard1.5/Grpc.Core.Api.pdb",
           "lib/netstandard1.5/Grpc.Core.Api.xml",
        ],
        "netstandard1.6": [
           "lib/netstandard1.5/Grpc.Core.Api.dll",
           "lib/netstandard1.5/Grpc.Core.Api.pdb",
           "lib/netstandard1.5/Grpc.Core.Api.xml",
        ],
        "netstandard2.0": [
           "lib/netstandard2.0/Grpc.Core.Api.dll",
           "lib/netstandard2.0/Grpc.Core.Api.pdb",
           "lib/netstandard2.0/Grpc.Core.Api.xml",
        ],
    },
    mono_files = [
        "lib/net45/Grpc.Core.Api.dll",
        "lib/net45/Grpc.Core.Api.pdb",
        "lib/net45/Grpc.Core.Api.xml",
    ],
)
nuget_package(
    name = "grpc.core",
    package = "grpc.core",
    version = "1.22.0",
    sha256 = "4b5039d9446f907b208591584d89ea4cd164ce00c2928ecbdb73ff9cf713a453",
    core_lib = {
        "netcoreapp2.0": "lib/netstandard2.0/Grpc.Core.dll",
        "netcoreapp2.1": "lib/netstandard2.0/Grpc.Core.dll",
    },
    net_lib = {
        "net45": "lib/net45/Grpc.Core.dll",
        "net451": "lib/net45/Grpc.Core.dll",
        "net452": "lib/net45/Grpc.Core.dll",
        "net46": "lib/net45/Grpc.Core.dll",
        "net461": "lib/net45/Grpc.Core.dll",
        "net462": "lib/net45/Grpc.Core.dll",
        "net47": "lib/net45/Grpc.Core.dll",
        "net471": "lib/net45/Grpc.Core.dll",
        "net472": "lib/net45/Grpc.Core.dll",
        "netstandard1.5": "lib/netstandard1.5/Grpc.Core.dll",
        "netstandard1.6": "lib/netstandard1.5/Grpc.Core.dll",
        "netstandard2.0": "lib/netstandard2.0/Grpc.Core.dll",
    },
    mono_lib = "lib/net45/Grpc.Core.dll",
    core_deps = {
        "net45": [
           "@grpc.core.api//:net45_net",
           "@system.interactive.async//:net45_net",
        ],
        "net451": [
           "@grpc.core.api//:net451_net",
           "@system.interactive.async//:net451_net",
        ],
        "net452": [
           "@grpc.core.api//:net452_net",
           "@system.interactive.async//:net452_net",
        ],
        "net46": [
           "@grpc.core.api//:net46_net",
           "@system.interactive.async//:net46_net",
        ],
        "net461": [
           "@grpc.core.api//:net461_net",
           "@system.interactive.async//:net461_net",
        ],
        "net462": [
           "@grpc.core.api//:net462_net",
           "@system.interactive.async//:net462_net",
        ],
        "net47": [
           "@grpc.core.api//:net47_net",
           "@system.interactive.async//:net47_net",
        ],
        "net471": [
           "@grpc.core.api//:net471_net",
           "@system.interactive.async//:net471_net",
        ],
        "net472": [
           "@grpc.core.api//:net472_net",
           "@system.interactive.async//:net472_net",
        ],
        "netstandard1.5": [
           "@grpc.core.api//:netstandard1.5_net",
           "@system.interactive.async//:netstandard1.5_net",
        ],
        "netstandard1.6": [
           "@grpc.core.api//:netstandard1.6_net",
           "@system.interactive.async//:netstandard1.6_net",
        ],
        "netstandard2.0": [
           "@grpc.core.api//:netstandard2.0_net",
           "@system.interactive.async//:netstandard2.0_net",
        ],
    },
    net_deps = {
        "net45": [
           "@grpc.core.api//:net45_net",
           "@system.interactive.async//:net45_net",
        ],
        "net451": [
           "@grpc.core.api//:net451_net",
           "@system.interactive.async//:net451_net",
        ],
        "net452": [
           "@grpc.core.api//:net452_net",
           "@system.interactive.async//:net452_net",
        ],
        "net46": [
           "@grpc.core.api//:net46_net",
           "@system.interactive.async//:net46_net",
        ],
        "net461": [
           "@grpc.core.api//:net461_net",
           "@system.interactive.async//:net461_net",
        ],
        "net462": [
           "@grpc.core.api//:net462_net",
           "@system.interactive.async//:net462_net",
        ],
        "net47": [
           "@grpc.core.api//:net47_net",
           "@system.interactive.async//:net47_net",
        ],
        "net471": [
           "@grpc.core.api//:net471_net",
           "@system.interactive.async//:net471_net",
        ],
        "net472": [
           "@grpc.core.api//:net472_net",
           "@system.interactive.async//:net472_net",
        ],
        "netstandard1.5": [
           "@grpc.core.api//:netstandard1.5_net",
           "@system.interactive.async//:netstandard1.5_net",
        ],
        "netstandard1.6": [
           "@grpc.core.api//:netstandard1.6_net",
           "@system.interactive.async//:netstandard1.6_net",
        ],
        "netstandard2.0": [
           "@grpc.core.api//:netstandard2.0_net",
           "@system.interactive.async//:netstandard2.0_net",
        ],
    },
    mono_deps = [
        "@grpc.core.api//:mono",
        "@system.interactive.async//:mono",
    ],
    core_files = {
        "netcoreapp2.0": [
           "lib/netstandard2.0/Grpc.Core.dll",
           "lib/netstandard2.0/Grpc.Core.pdb",
           "lib/netstandard2.0/Grpc.Core.xml",
        ],
        "netcoreapp2.1": [
           "lib/netstandard2.0/Grpc.Core.dll",
           "lib/netstandard2.0/Grpc.Core.pdb",
           "lib/netstandard2.0/Grpc.Core.xml",
        ],
    },
    net_files = {
        "net45": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "net451": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "net452": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "net46": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "net461": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "net462": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "net47": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "net471": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "net472": [
           "lib/net45/Grpc.Core.dll",
           "lib/net45/Grpc.Core.pdb",
           "lib/net45/Grpc.Core.xml",
        ],
        "netstandard1.0": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.1": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.2": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.3": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.4": [
           "lib/",
           "lib/net45/",
           "lib/netstandard1.5/",
           "lib/netstandard2.0/",
        ],
        "netstandard1.5": [
           "lib/netstandard1.5/Grpc.Core.dll",
           "lib/netstandard1.5/Grpc.Core.pdb",
           "lib/netstandard1.5/Grpc.Core.xml",
        ],
        "netstandard1.6": [
           "lib/netstandard1.5/Grpc.Core.dll",
           "lib/netstandard1.5/Grpc.Core.pdb",
           "lib/netstandard1.5/Grpc.Core.xml",
        ],
        "netstandard2.0": [
           "lib/netstandard2.0/Grpc.Core.dll",
           "lib/netstandard2.0/Grpc.Core.pdb",
           "lib/netstandard2.0/Grpc.Core.xml",
        ],
    },
    mono_files = [
        "lib/net45/Grpc.Core.dll",
        "lib/net45/Grpc.Core.pdb",
        "lib/net45/Grpc.Core.xml",
    ],
)
nuget_package(
    name = "google.protobuf",
    package = "google.protobuf",
    version = "3.9.0-rc1",
    sha256 = "ade0c45790ebdb675a75df1939048e6f2a7bc17e73d1ec3b5855c7a15a16685e",
    core_lib = {
        "netcoreapp2.0": "lib/netstandard2.0/Google.Protobuf.dll",
        "netcoreapp2.1": "lib/netstandard2.0/Google.Protobuf.dll",
    },
    net_lib = {
        "net45": "lib/net45/Google.Protobuf.dll",
        "net451": "lib/net45/Google.Protobuf.dll",
        "net452": "lib/net45/Google.Protobuf.dll",
        "net46": "lib/net45/Google.Protobuf.dll",
        "net461": "lib/net45/Google.Protobuf.dll",
        "net462": "lib/net45/Google.Protobuf.dll",
        "net47": "lib/net45/Google.Protobuf.dll",
        "net471": "lib/net45/Google.Protobuf.dll",
        "net472": "lib/net45/Google.Protobuf.dll",
        "netstandard1.0": "lib/netstandard1.0/Google.Protobuf.dll",
        "netstandard1.1": "lib/netstandard1.0/Google.Protobuf.dll",
        "netstandard1.2": "lib/netstandard1.0/Google.Protobuf.dll",
        "netstandard1.3": "lib/netstandard1.0/Google.Protobuf.dll",
        "netstandard1.4": "lib/netstandard1.0/Google.Protobuf.dll",
        "netstandard1.5": "lib/netstandard1.0/Google.Protobuf.dll",
        "netstandard1.6": "lib/netstandard1.0/Google.Protobuf.dll",
        "netstandard2.0": "lib/netstandard2.0/Google.Protobuf.dll",
    },
    mono_lib = "lib/net45/Google.Protobuf.dll",
    core_deps = {
        "netstandard2.0": [
           "@system.memory//:netstandard2.0_net",
        ],
    },
    net_deps = {
        "netstandard2.0": [
           "@system.memory//:netstandard2.0_net",
        ],
    },
    core_files = {
        "netcoreapp2.0": [
           "lib/netstandard2.0/Google.Protobuf.dll",
           "lib/netstandard2.0/Google.Protobuf.pdb",
           "lib/netstandard2.0/Google.Protobuf.xml",
        ],
        "netcoreapp2.1": [
           "lib/netstandard2.0/Google.Protobuf.dll",
           "lib/netstandard2.0/Google.Protobuf.pdb",
           "lib/netstandard2.0/Google.Protobuf.xml",
        ],
    },
    net_files = {
        "net45": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "net451": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "net452": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "net46": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "net461": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "net462": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "net47": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "net471": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "net472": [
           "lib/net45/Google.Protobuf.dll",
           "lib/net45/Google.Protobuf.pdb",
           "lib/net45/Google.Protobuf.xml",
        ],
        "netstandard1.0": [
           "lib/netstandard1.0/Google.Protobuf.dll",
           "lib/netstandard1.0/Google.Protobuf.pdb",
           "lib/netstandard1.0/Google.Protobuf.xml",
        ],
        "netstandard1.1": [
           "lib/netstandard1.0/Google.Protobuf.dll",
           "lib/netstandard1.0/Google.Protobuf.pdb",
           "lib/netstandard1.0/Google.Protobuf.xml",
        ],
        "netstandard1.2": [
           "lib/netstandard1.0/Google.Protobuf.dll",
           "lib/netstandard1.0/Google.Protobuf.pdb",
           "lib/netstandard1.0/Google.Protobuf.xml",
        ],
        "netstandard1.3": [
           "lib/netstandard1.0/Google.Protobuf.dll",
           "lib/netstandard1.0/Google.Protobuf.pdb",
           "lib/netstandard1.0/Google.Protobuf.xml",
        ],
        "netstandard1.4": [
           "lib/netstandard1.0/Google.Protobuf.dll",
           "lib/netstandard1.0/Google.Protobuf.pdb",
           "lib/netstandard1.0/Google.Protobuf.xml",
        ],
        "netstandard1.5": [
           "lib/netstandard1.0/Google.Protobuf.dll",
           "lib/netstandard1.0/Google.Protobuf.pdb",
           "lib/netstandard1.0/Google.Protobuf.xml",
        ],
        "netstandard1.6": [
           "lib/netstandard1.0/Google.Protobuf.dll",
           "lib/netstandard1.0/Google.Protobuf.pdb",
           "lib/netstandard1.0/Google.Protobuf.xml",
        ],
        "netstandard2.0": [
           "lib/netstandard2.0/Google.Protobuf.dll",
           "lib/netstandard2.0/Google.Protobuf.pdb",
           "lib/netstandard2.0/Google.Protobuf.xml",
        ],
    },
    mono_files = [
        "lib/net45/Google.Protobuf.dll",
        "lib/net45/Google.Protobuf.pdb",
        "lib/net45/Google.Protobuf.xml",
    ],
)
nuget_package(
    name = "system.runtime.compilerservices.unsafe",
    package = "system.runtime.compilerservices.unsafe",
    version = "4.5.2",
    sha256 = "f1e5175c658ed8b2fbb804cc6727b6882a503844e7da309c8d4846e9ca11e4ef",
    core_lib = {
        "netcoreapp2.0": "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netcoreapp2.1": "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
    },
    net_lib = {
        "net45": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "net451": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "net452": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "net46": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "net461": "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
        "net462": "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
        "net47": "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
        "net471": "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
        "net472": "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netstandard1.0": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netstandard1.1": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netstandard1.2": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netstandard1.3": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netstandard1.4": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netstandard1.5": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netstandard1.6": "ref/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
        "netstandard2.0": "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
    },
    mono_lib = "ref/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
    core_files = {
        "netcoreapp2.0": [
           "lib/netcoreapp2.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netcoreapp2.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netcoreapp2.1": [
           "lib/netcoreapp2.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netcoreapp2.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
    },
    net_files = {
        "net45": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "net451": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "net452": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "net46": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "net461": [
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "net462": [
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "net47": [
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "net471": [
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "net472": [
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netstandard1.0": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netstandard1.1": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netstandard1.2": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netstandard1.3": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netstandard1.4": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netstandard1.5": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netstandard1.6": [
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard1.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
        "netstandard2.0": [
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
           "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.xml",
        ],
    },
    mono_files = [
        "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.dll",
        "lib/netstandard2.0/System.Runtime.CompilerServices.Unsafe.xml",
    ],
)
nuget_package(
    name = "system.memory",
    package = "system.memory",
    version = "4.5.3",
    sha256 = "0af97b45b45b46ef6a2b37910568dabd492c793da3859054595d523e2a545859",
    core_lib = {
        "netcoreapp2.0": "lib/netstandard2.0/System.Memory.dll",
    },
    net_lib = {
        "net45": "lib/netstandard1.1/System.Memory.dll",
        "net451": "lib/netstandard1.1/System.Memory.dll",
        "net452": "lib/netstandard1.1/System.Memory.dll",
        "net46": "lib/netstandard1.1/System.Memory.dll",
        "net461": "lib/netstandard2.0/System.Memory.dll",
        "net462": "lib/netstandard2.0/System.Memory.dll",
        "net47": "lib/netstandard2.0/System.Memory.dll",
        "net471": "lib/netstandard2.0/System.Memory.dll",
        "net472": "lib/netstandard2.0/System.Memory.dll",
        "netstandard1.1": "lib/netstandard1.1/System.Memory.dll",
        "netstandard1.2": "lib/netstandard1.1/System.Memory.dll",
        "netstandard1.3": "lib/netstandard1.1/System.Memory.dll",
        "netstandard1.4": "lib/netstandard1.1/System.Memory.dll",
        "netstandard1.5": "lib/netstandard1.1/System.Memory.dll",
        "netstandard1.6": "lib/netstandard1.1/System.Memory.dll",
        "netstandard2.0": "lib/netstandard2.0/System.Memory.dll",
    },
    mono_lib = "lib/netstandard2.0/System.Memory.dll",
    core_deps = {
        "net45": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net45_system.buffers.dll",
           "@system.runtime.compilerservices.unsafe//:net45_net",
        ],
        "net451": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net451_system.buffers.dll",
           "@system.runtime.compilerservices.unsafe//:net451_net",
        ],
        "net452": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net452_system.buffers.dll",
           "@system.runtime.compilerservices.unsafe//:net452_net",
        ],
        "net46": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net46_system.buffers.dll",
           "@system.runtime.compilerservices.unsafe//:net46_net",
        ],
        "net461": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net461_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net461_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net461_net",
        ],
        "net462": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net462_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net462_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net462_net",
        ],
        "net47": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net47_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net47_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net47_net",
        ],
        "net471": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net471_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net471_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net471_net",
        ],
        "net472": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net472_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net472_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net472_net",
        ],
        "netstandard1.1": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.1_net",
        ],
        "netstandard1.2": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.2_net",
        ],
        "netstandard1.3": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.3_net",
        ],
        "netstandard1.4": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.4_net",
        ],
        "netstandard1.5": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.5_net",
        ],
        "netstandard1.6": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.6_net",
        ],
        "netstandard2.0": [
           "@system.runtime.compilerservices.unsafe//:netstandard2.0_net",
        ],
    },
    net_deps = {
        "net45": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net45_system.buffers.dll",
           "@system.runtime.compilerservices.unsafe//:net45_net",
        ],
        "net451": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net451_system.buffers.dll",
           "@system.runtime.compilerservices.unsafe//:net451_net",
        ],
        "net452": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net452_system.buffers.dll",
           "@system.runtime.compilerservices.unsafe//:net452_net",
        ],
        "net46": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net46_system.buffers.dll",
           "@system.runtime.compilerservices.unsafe//:net46_net",
        ],
        "net461": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net461_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net461_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net461_net",
        ],
        "net462": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net462_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net462_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net462_net",
        ],
        "net47": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net47_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net47_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net47_net",
        ],
        "net471": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net471_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net471_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net471_net",
        ],
        "net472": [
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net472_system.buffers.dll",
           "@io_bazel_rules_dotnet//dotnet/stdlib.net:net472_system.numerics.vectors.dll",
           "@system.runtime.compilerservices.unsafe//:net472_net",
        ],
        "netstandard1.1": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.1_net",
        ],
        "netstandard1.2": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.2_net",
        ],
        "netstandard1.3": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.3_net",
        ],
        "netstandard1.4": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.4_net",
        ],
        "netstandard1.5": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.5_net",
        ],
        "netstandard1.6": [
           "@system.runtime.compilerservices.unsafe//:netstandard1.6_net",
        ],
        "netstandard2.0": [
           "@system.runtime.compilerservices.unsafe//:netstandard2.0_net",
        ],
    },
    mono_deps = [
        "@io_bazel_rules_dotnet//dotnet/stdlib:system.buffers.dll",
        "@io_bazel_rules_dotnet//dotnet/stdlib:system.numerics.vectors.dll",
        "@system.runtime.compilerservices.unsafe//:mono",
    ],
    core_files = {
        "netcoreapp2.0": [
           "lib/netstandard2.0/System.Memory.dll",
           "lib/netstandard2.0/System.Memory.xml",
        ],
    },
    net_files = {
        "net45": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "net451": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "net452": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "net46": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "net461": [
           "lib/netstandard2.0/System.Memory.dll",
           "lib/netstandard2.0/System.Memory.xml",
        ],
        "net462": [
           "lib/netstandard2.0/System.Memory.dll",
           "lib/netstandard2.0/System.Memory.xml",
        ],
        "net47": [
           "lib/netstandard2.0/System.Memory.dll",
           "lib/netstandard2.0/System.Memory.xml",
        ],
        "net471": [
           "lib/netstandard2.0/System.Memory.dll",
           "lib/netstandard2.0/System.Memory.xml",
        ],
        "net472": [
           "lib/netstandard2.0/System.Memory.dll",
           "lib/netstandard2.0/System.Memory.xml",
        ],
        "netstandard1.1": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "netstandard1.2": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "netstandard1.3": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "netstandard1.4": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "netstandard1.5": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "netstandard1.6": [
           "lib/netstandard1.1/System.Memory.dll",
           "lib/netstandard1.1/System.Memory.xml",
        ],
        "netstandard2.0": [
           "lib/netstandard2.0/System.Memory.dll",
           "lib/netstandard2.0/System.Memory.xml",
        ],
    },
    mono_files = [
        "lib/netstandard2.0/System.Memory.dll",
        "lib/netstandard2.0/System.Memory.xml",
    ],
)
nuget_package(
    name = "microsoft.netcore.platforms",
    package = "microsoft.netcore.platforms",
    version = "1.1.2",
    sha256 = "adaf848ef23c5f302a56f8b4cde237b77ac88e0c6e721026bcf7a86e8a3af8b1",
)
### End of generated by the tool
