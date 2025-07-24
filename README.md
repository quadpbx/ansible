# QuadPBX Ansible (Trixie)

This is just a random collection of tools and roles to help configure
QuadPBX machines.

# Prep

Run `make setup` to install all the basic prereqs and set up a
generic .gitconfig - It'll ask you for your name and email. Don't
accept the defaults unless you ACTUALLY ARE xrobau

## Volumes

You probably want to add these as seperate volumes:

* /var/lib/docker
* /usr/local/data
* /usr/local/build

## SSH Keys

Update `group_vars/all/ssh_keys` unless you're actually xrobau. At least add yours!

# Base setup

Run `make base` to set the local machine up as a development
machine.


