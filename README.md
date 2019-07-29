![solarium](https://raw.githubusercontent.com/Chronojam/Solarium/master/docs/solarium.png)

Solarium
==

**Building Nuget Package**

Update the solarium version in Solarium.nuspec
```
$ bazel build //proto:csharp_grpc_lib
$ bazel run @nuget//file:downloaded -- pack 
```

**Building the Documentation**
```
$ bazel build //proto:docs
$ cp bazel-bin/proto/docs.md docs/docs.md
```

GRPC Documentation 
==
[GRPCDocs](https://github.com/Chronojam/Solarium/blob/master/docs/docs.md)
