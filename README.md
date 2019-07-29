![solarium](https://raw.githubusercontent.com/Chronojam/Solarium/master/doc/solarium.png)

Solarium
==


Building Nuget Package

Update the solarium version in Solarium.nuspec
```
$ bazel build //proto:csharp_grpc_lib
$ bazel run @nuget//file:downloaded -- pack 

```
