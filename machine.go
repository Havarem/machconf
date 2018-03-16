package main

import (
	"net"
)

// Flag are going to be used by providers, orchestrators and other tooling.
type Flag string

// This is the bare minimum to represent a computer.
type Machine struct {
	ip net.IP
	domain string
	flags []Flag
}

// This is the bare minimum actions a machine can perform.
type Computing interface {
	launch(provision bool) error
	halt() error
	info() (error, Machine)
}