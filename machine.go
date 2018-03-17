package main

import (
	"net"
)

// Flag are going to be used by providers, orchestrators and other tooling.
type Flag string

// This is the bare minimum to represent a computer.
type Machine struct {
	Ip net.IP
	Domain string
	Flags []Flag
}

// This is the bare minimum actions a machine can perform.
type Computing interface {
	Launch(provision bool) error
	Halt() error
	Info() (error, Machine)
}