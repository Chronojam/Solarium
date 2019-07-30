![solarium](https://raw.githubusercontent.com/Chronojam/Solarium/master/docs/solarium.png)

Solarium
==


Quick Start
== 
**Download**  
Linux:
```
$ curl -LO solarium https://storage.googleapis.com/solarium/bin/$(curl -s https://storage.googleapis.com/solarium/bin/latest)/linux_amd64/solarium
```
OSX:
```
$ curl -LO solarium https://storage.googleapis.com/solarium/bin/$(curl -s https://storage.googleapis.com/solarium/bin/latest)/darwin_amd64/solarium
```
Windows:
Download from this link: [This Link](https://storage.googleapis.com/solarium/bin/2.0.2/windows_amd64/solarium.exe)

Or if you have curl installed
```
$ curl -LO https://storage.googleapis.com/solarium/bin/2.0.2/windows_amd64/solarium.exe
```

Run with:

```
./solarium
```

Developing
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
