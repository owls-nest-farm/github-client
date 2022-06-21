# GitHub Client

## What does this do?

The goal of this project is a simple one: quickly and easily create as many `GitHub` repositories as needed within an **existing** group.

## What does this support?

This project **only** supports the *creation* of the following `GitHub` objects:

- [Repositories]

Its intent is to **only** create them with **only** the required configuration parameters.  The reasoning is to support its only use case, which is just to "bootstrap" one or more projects with all the fixings.  It doesn't really matter what the data **is**, as long as it is **there**.

If you need something more customized, you'll have to do that yourself.

## Requirements

- [`GitHub` API token] with `api` scope
- An extant organization

## Notes

- Both `yaml` and `json` data formats are supported.

## Fields

### Creating [Repositories]

A list of `Repository` objects composed of:

- `name` (string)
- `archived` (bool)
- `private` (bool)
- `tpl_name` (string)
    + The name of a repository which has been designated as a template by having its [`archived` flag] turned on.
- `visibility` (string)
- [`collaborators`] ([]string)
    + Invitations are immediately sent to all listed members.
    + Permissions are hardcoded as `pull`.

## Examples

There are several examples in the `./examples` folder of configs in both `json` and `yaml` formats, both neither are exhaustive.  They should, however, give a good idea of how the configuration should be structured.

### Creating Repositories

```
$ github-client -file examples/github.yaml
$ github-client -file examples/github.json
```

### Deleting Repositories

To tear down what was setup when creating the projects, simply pass the same config file with the `-destroy` flag.

Or, pass another file or make your changes in the same one.  Pick your poison.

```
$ github-client -file examples/github.yaml -destroy
```

> This will **not** ask for confirmation.

## Generating a YAML config file

If you want to quickly create any number of sequentially-named repositories, you can use the following simple program:

```
$ go run tpl/configure.go -out test.yaml -n 100
```

### Parameters

- `out`
    + The name of the generated config file.
    + Defaults to `github.yaml`.

- `n`
    + The number of repositories to create in the given organization.
    + Defaults to 25.

## Developing

If you don't want to compile, you can use `go run`:

```
$ go run *.go -file examples/github.yaml
```

## Acknowledgments

This project uses the [Google's Go library]  for interacting with the [GitHub REST API].

[Organizations]: https://docs.github.com/en/rest/orgs/orgs
[Repositories]: https://docs.github.com/en/rest/repos/repos
[`GitHub` API token]: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token
[`collaborators`]: https://docs.github.com/en/rest/collaborators/collaborators
[`archived` flag]: https://docs.github.com/en/rest/repos/repos#create-a-repository-using-a-template
[Google's Go library]: https://github.com/google/go-github
[GitHub REST API]: https://docs.github.com/en/github-ae@latest/rest

