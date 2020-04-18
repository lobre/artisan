# Artisan

**In progress**

Really simple declarative configuration management tool. It has been designed to help me reproduce my local development environment easily. I like the simplicity of the Go language and being able to declare my configurations in pure Go feels simple to maintain to me.

As the configuration is in Go, it should be built each time it is modified. It does not seem to be a problem for now. It means I can build the binary each time the code changes. Ideally, the CI could do it. And it could as well upload the binary to GitHub releases with something like [goreleaser](https://github.com/goreleaser/goreleaser).

A little bit like Ansible, this project defines "modules" called "artisans" (hence the name of the repo). Artisans are small modules that implement the `Craft()` method. An artisan has the responsibility to ensure the declaratite configuration is enforced. Each artisan has got a special purpose.

 - Nix: make sure nix packages are installed.
 - Apt: make sure apt packages are installed.
 - File: make sure a file is present at a specific location (file can be in the shape of Go templates that will be processed).
 - Download: retrieve file from a specific location.
 - Archive: extract an archive.
 - Make: execute make tasks.
 - Git: clone a repository.
 - Cmd: launch a specific command (not really declarative but could be useful).

The project should also has a way to tell an artisan to Craft as a specific user. Apt should for instance be execute from a "root" context.

Configurations will be applied on the same machine as the one executing the program. There is for now no plan for building a system that would run on external hosts.
