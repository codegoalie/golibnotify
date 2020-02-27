// Package golibnotify implements go bindings for libnotify to create, send, and
// update OS level notifications. It does not shell out to `notify-send` so it can
// update existing notifications as well as create new ones.
//
// This package requires CGO and the libnotify (libnotify-dev) shared library
// to be installed.
//
// On Ubuntu or Debian using apt, you can install with:
// sudo apt-get install libnotify-dev
package golibnotify

/*
#cgo pkg-config: libnotify
#include "libnotify/notify.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
)

// SimpleNotifier is an instance of an application sending notifications
type SimpleNotifier struct {
	appName      string
	notification notification
}

// notification is an instance of a particular notification
type notification *C.NotifyNotification

// NewSimpleNotifier initializes a new application to send notifications
func NewSimpleNotifier(applicationName string) *SimpleNotifier {
	C.notify_init(C.CString(applicationName))

	return &SimpleNotifier{appName: applicationName}
}

// ApplicationName returns the current application's initialized name
func (n *SimpleNotifier) ApplicationName() string {
	return n.appName
}

// Show creates a new notification and sends it to the OS
func (n *SimpleNotifier) Show(summary, body, icon string) error {
	cSummary := C.CString(summary)
	cBody := C.CString(body)
	cIcon := C.CString(icon)
	n.notification = C.notify_notification_new(cSummary, cBody, cIcon)

	return show(n.notification)
}

// Update an existing notification with new information
func (n *SimpleNotifier) Update(summary, body, icon string) error {
	if n.notification == nil {
		return n.Show(summary, body, icon)
	}

	cSummary := C.CString(summary)
	cBody := C.CString(body)
	cIcon := C.CString(icon)
	C.notify_notification_update(n.notification, cSummary, cBody, cIcon)

	return show(n.notification)
}

// Close removes the notification from the OS
func (n *SimpleNotifier) Close() error {
	if n.notification == nil {
		return fmt.Errorf("failed to close notification: no notification exists")
	}

	return close(n.notification)
}

func show(notif notification) error {
	var cErr **C.GError
	C.notify_notification_show(notif, cErr)
	if cErr != nil {
		return errors.New(C.GoString((*cErr).message))
	}
	return nil
}

func close(notif notification) error {
	var cErr **C.GError
	C.notify_notification_close(notif, cErr)
	if cErr != nil {
		return errors.New(C.GoString((*cErr).message))
	}
	return nil
}
