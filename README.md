# golibnotify

`golibnotify` implements go bindings for
[`libnotify`](https://developer.gnome.org/libnotify/unstable/) to create, send,
and update OS level notifications. It does not shell out to `notify-send` so it
can update existing notifications as well as create new ones.

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/codegoalie/golibnotify?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/codegoalie/golibnotify)](https://goreportcard.com/report/github.com/codegoalie/golibnotify)
![GitHub](https://img.shields.io/github/license/codegoalie/golibnotify?style=flat-square)
[![Sourcegraph Badge](https://sourcegraph.com/github.com/codegoalie/golibnotify/-/badge.svg)](https://sourcegraph.com/github.com/codegoalie/golibnotify?badge)

## Installation

This package requires CGO and the libnotify (`libnotify-dev`) shared library to be installed.

On Ubuntu or Debian using `apt`, you can install with `sudo apt-get install libnotify-dev`

## Usage

```go
// Get an instance of a SimpleNotifier
notifier := golibnotify.NewSimpleNotifier("my-cool-app-name")

// Show a new notification
err = notifier.Show("A summary", "Some body text", "a/path/to/an/icon.png")

// Update an existing notification (or send a new one if one hasn't been sent)
err = notifier.Update("A new summary", "Some different body text", "another/path/to/icon.png")

// Remove an existing notification
err = notifier.Close()
```

See also [examples](https://github.com/codegoalie/golibnotify/examples) for a complete example

## Roadmap

- [ ] Add API that returns a Notification instance to manage multiple notifications at once
- [ ] Better support for notification icons (PixBuf?)
- [ ] Support notification timeouts
- [ ] Support notification urgencies

## Contributions

Thanks for wanting to contribute!

1. Open an issue describing your problem, solution, issue, suggestion, feature,
etc. to ensure it's likely to get merged.
1. Fork the project.
1. Make your changes.
1. Open a PR against `master` here.
1. Celebrate being awesome!

## Author

[@codegoalie](https://codegoalie.com)
