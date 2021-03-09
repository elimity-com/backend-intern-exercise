# Elimity backend internship exercise

Thank you for applying for the **Elimity Backend Developer Intern** position! This exercise is designed to give you an
opportunity to show off your programming skills that would be relevant to work at Elimity, and to give you an idea of
what type of work you would be doing during your internship. The code in this repository provides a Go module as a
starting point for the assignments. It implements `gothub`, a simple CLI for tracking public GitHub repositories. The
actual assignments are listed below.

## Installation

The `gothub` CLI can be installed using `go get` with Go 1.16+:

```sh
$ go get github.com/elimity-com/backend-intern-exercise/cmd/gothub
```

## Usage

```sh
$ gothub help
Simple CLI for tracking public GitHub repositories.

Usage:
  gothub help
  gothub track [-interval=<interval>]

Commands:
  help  Show usage information
  track Track public GitHub repositories

Options:
  -interval=<interval> Repository update interval, greater than zero [default: 10s]
```

## Assignments

Currently `gothub` only supports regularly printing the most recently updated public GitHub repositories. The
assignments consist of extending its functionality. More specifically, we ask you to implement these improvements:

1. GitHub applies very strict rate limiting for anonymous requests, causing `gothub` to fail when it contacts the API
   too frequently. As stated [here](https://docs.github.com/en/rest/reference/search#rate-limit), the rate limit could
   be increased by using authenticated requests. Therefore we ask you to add a `token-file` option to the `gothub` CLI,
   which should support authenticating requests by reading a
   [GitHub personal access token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token)
   read from the given file path. If not provided, `gothub` should stick to anonymous requests.

2. Currently `gothub` only prints repository names. We ask you to extend this to a table-formatted output with columns
   for the repository's owner, name, last update timestamp and star count. If the repository belongs to an organization,
   list the organization's name in the owner column. The result should look like this (including headers):

```
Owner   | Name      | Updated at (UTC)    | Star count
golang  | go        | 2021-03-08T11:29:52 | 12345
google  | go-github | 2021-03-07T12:34:56 | 5432
angular | angular   | 2021-02-29T05:43:21 | 43210
```

3. The repositories outputted by `gothub` are only filtered by their accessibility, as they are all public. We ask you
   to add a `min-stars` option, which causes `gothub` to filter out repositories with a star count below the given
   value. If not provided, `gothub` should use a minimum star count of zero.

### Grading

Your solutions for the assignments will be graded according to the following points:

- compliance with the listed requirements,
- consistency with the existing code,
- understanding of Go fundamentals,
- adherence to Elimity's style guidelines:
  - strong types over implicit assumptions,
  - immutability over mutability, and
  - simple over clever.

### Practicalities

- You can use any library you want, but we prefer the standard library.
- These practices are important on the job, but won't affect your grades:
  - proper version control,
  - testing (which is complex for a CLI without business logic).
