# VersionTool

Perform automated arithmetic on semantic versions.

This repo contains the code for a CLI called `versiontool`
as well as a github action to employ it in workflows.

## Usage

`versiontool [options] <cmd> [<args>...]`

Options:
  --prefix=<string>   A prefix to use with versions. (e.g "v" for v1.2.3)
  --strict, -S        Absolutely require the presence of the prefix

### Subcommands

<dl>
<dt>increment</dt> <dd>increments a version</dd>
<dt>decrement</dt> <dd>decrements a version</dd>
<dt>cut</dt> <dd>trims a version</dd>
<dt>sort</dt> <dd>sorts a list of version from stdin</dd>
<dt>highest</dt> <dd>returns the highest (i.e. most recent) version from a list on stdin</dd>
</dl>

#### versiontool increment

Usage: `versiontool increment [options] <version>`

```
Options:
	--no-reset, -R             Do not reset smaller increments
	--major=<int>, -M <int>    Increment major version by <int>
	--minor=<int>, -m <int>    Increment minor version by <int>
	--patch=<int>, -p <int>    Increment patch version by <int>
```

Resetting will clear or set to zero parts of the version less significant
than the parts being incremented. This is almost always what.

With no arguments provided, the default is -p1 - that is, to bump the patch
version by one.

Example:
```
> versiontool increment -M 1 1.2.3-rc1
2.0.0

> versiontool increment -M 1 -m 1 1.2.3-rc1
2.1.0

> versiontool increment --no-reset -M 1 1.2.3-rc1
2.2.3-rc1

> versiontool increment --no-reset -M 1 -m 1 1.2.3-rc1
2.3.3-rc1
```

#### versiontool cut
Usage: `versiontool cut [options] <version>`

```
Options:
	--major, -M  Cut after major version
	--minor, -m  Cut after minor version
	--patch, -p  Cut after patch version
```

Example:
```
> versiontool cut -M 1.2.3-rc1
1

> versiontool cut -m 1.2.3-rc1
1.2

> versiontool cut -p 1.2.3
1.2.3
```

## Github Action

Packaged as a Github Action, that takes in a version (or a file of versions for `sort`/`highest`)
and outputs the result.

Example:
```yaml

# ...
    - id: highest-tag
      uses: ymeadows/versiontool@v1.1.0
      with:
        version: ./tag-list.txt
        operation: highest
    - id: inc-tag
      uses: ymeadows/versiontool@v1.1.0
      with:
        version: ${{ steps.highest-tag.outputs.result }}
        operation: increment # other options: decrement, cut, sort
        prefix: ${{ inputs.prefix }}
        strict: true
        # flags: --minor=1
# ...
```

Also see: `ymeadows/github-actions-public/rolling-versions` which has outputs `major` and `major-minor`
so that you can quickly generate GHA or Docker tags
