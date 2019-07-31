load("//third_party/build_stack_rules_hugo/hugo:rules.bzl", "hugo_site", "hugo_theme")

# Declare a theme 'xmin'.  In this case the `name` and
# `theme_name` are identical, so the `theme_name` could be omitted in this case.
hugo_theme(
    name = "xmin",
    theme_name = "xmin",
    srcs = [
        "@com_github_yihui_hugo_xmin//:files",
    ],
)
# Declare a site. Config file is required.

hugo_site(
    name = "solarium",
    config = "config.toml",
    content = [
        "//site/content:_index.md",
        "//proto:docs",
    ],
    quiet = True,
    theme = ":xmin",
)
