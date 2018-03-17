# Machine Configuration Tool

This tool is the main part in a series of other tools in order to manage any
kind of computer, virtual machine, VPS, containers, or any other kind of
processing units.

## Purpose

Has it for now, a machine is define by its Fully Qualified Domain Name, an IP
address and a list of optional tags. These can be enhanced to add other
information needed for Machine Providers.

## Roadmap

This project is the root base of a series of other tools to come. The next
steps may be as follow:

[] Create a Vagrant provider (machconf-vagrant-provider);
[] Create a Network Configuration (netconf);
[] Use Hashicorp Consul/Vault as persistence-data (so no API key in Yaml file);
[] Create a DigitalOcean provider (machconf-do-provider);
[] Create a Google Cloud Platform provider (mach-gcp-provider);
[] Create a Vultr Cloud provider (mach-vultr-provider);
[] Enhance the Network Configuration for multi-provider capabilities.