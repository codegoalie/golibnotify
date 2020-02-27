# golibnotify

`golibnotify` implements go bindings for
[`libnotify`](https://developer.gnome.org/libnotify/unstable/) to create, send,
and update OS level notifications.

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
