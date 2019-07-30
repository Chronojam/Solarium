exports_files(
    [
        "_index.md",
    ]
)

# This is obviously dumb, but hugo_site needs it to exist in a content/ dir
# so...


filegroup(
    name = "proto_documentation.md",
    srcs = [
        "//proto:docs",
    ],
    visibility = ["//visibility:public"],
)