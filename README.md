Solarium


![solarium](https://raw.githubusercontent.com/chronojam/solarium/master/doc/solarium.png)

==

Building Nuget Package

Update the solarium version in Solarium.nuspec
```
$ bazel build //proto:csharp_grpc_lib
$ bazel run @nuget//file:downloaded -- pack 

```
