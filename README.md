Solarium
==

Building Nuget Package

Update the solarium version in Solarium.nuspec
```
$ bazel build //proto:csharp_grpc_lib
$ bazel run @nuget//file:downloaded -- pack 

```
