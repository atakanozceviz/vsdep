# Vsdep

Find out which Visual Studio projects needs to build and which tests needs to run by checking differences between a git commit id and HEAD.

Vsdep parses all the solutions (.sln) and projects (.csproj) in a folder and creates dependency graph, then uses that graph for finding dependencies.

## Install

```
go github.com/atakanozceviz/vsdep/vsdep
```

## How to use

```console
$ vsdep -h
Run vsdep with a commit ID in VS project's root.

vsdep [commit]

Some examples:
vsdep HEAD^
vsdep fd32f09
vsdep HEAD^^ -w ../

Commit ID can be path to project folder with leading ID:
vsdep ../otherProject/HEAD^^ -w ../

Usage:
  vsdep [flags]

Flags:
      --config string     config file (default is $HOME/.vsdep.yaml)
  -h, --help              help for vsdep
  -w, --walkpath string   the path to start the search for .sln files (default ".")
```

### Example output:

```console
$ vsdep HEAD^^
```

```json
{
  "solutions": {
    "MyProjectName.sln": "templates/mvc/MyProjectName.sln",
    "Identity.sln": "modules/identity/Identity.sln",
    "Blogging.sln": "modules/blogging/Blogging.sln"
  },
  "tests": {
    "MyProjectName.Application.Tests": "templates/mvc/test/MyProjectName.Application.Tests/MyProjectName.Application.Tests.csproj",
    "MyProjectName.Web.Tests": "templates/mvc/test/MyProjectName.Web.Tests/MyProjectName.Web.Tests.csproj",
    "Identity.Application.Tests": "modules/identity/test/Identity.Application.Tests/Identity.Application.Tests.csproj"
  }
}
```