# Contributing

This document outlines some of the conventions on development workflow,
commit message formatting,
contact points and other
resources to make it easier to get your contribution accepted.

## Issues

When filing an issue, make sure to answer these five questions:

1. Version: What version of the project are you using?
2. Environment: What operating system and processor architecture are you using?
3. Step: What did you do?
4. Expect: What did you expect to see?
5. Result: What did you see instead?

## Code

### Contributing flow

This is a rough outline of what a contributor's workflow looks like:

- Fork the repository on GitHub.
- Create a topic branch from where you want to base your work. This is usually develop.
- If necessary, rebase your commits into logical chunks, without errors
- Run tests and make sure all the tests are passed.
- Make sure your commit messages are in the proper format (see below).
- Push your changes to a topic branch in your fork of the repository.
- Submit a pull request to `iotalabs/plug`
- Your PR must receive LGTMs from two maintainers found in the [MAINTAINERS](./MAINTAINERS.md) file.

### Code Style

The coding style suggested by the Golang community is used in TiDB. See the [style doc](https://github.com/golang/go/wiki/CodeReviewComments) for details.

Please follow this style to make TiDB easy to review, maintain and develop.

### Format of the Commit Message

We follow a rough convention for commit messages that is designed to answer two
questions: what changed and why. The subject line should feature the what and
the body of the commit should describe the why.

```
store/localstore: add comment for variable declaration.

Improve documentation.
```

The format can be described more formally as follows:

```
<subsystem>: <what changed>
<BLANK LINE>
<why this change was made>
<BLANK LINE>
<footer>(optional)
```

The first line is the subject and should be no longer than 70 characters, the
second line is always blank, and other lines should be wrapped at 80 characters.
This allows the message to be easier to read on GitHub as well as in various
git tools.

If the change affects more than one subsystem, you can use comma to separate them like `util/codec,util/types:`.

If the change affects many subsystems, you can use ```*``` instead, like ```*:```.

For the why part, if no specific reason for the change,
you can use one of some generic reasons like "Improve documentation.",
"Improve performance.", "Improve robustness.", "Improve test coverage."

### Review

Before Review, keep in mind that:

- Git commit messages should conform to [community standards](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).
- Git commits should represent meaningful milestones or units of work.
- Changed or added code must be well tested. Different kinds of code
- require different testing strategies.
- Changed or added code must pass the project's CI.
- Changes to vendored files must be grouped into a single commit.

Once comments and revisions on the implementation wind down, the reviewers will
add the LGTM comment which marks the PR as merge-able.

Thanks for your contributions!
