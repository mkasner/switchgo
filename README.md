# Switchgo

Switch go versions with one command.

# Why use multiple go versions
For example goracle driver doesn't work with go1.6, so for some projects go1.5.x must be used.

# Instructions
1. Unpack go files to /usr/local (for example /usr/local/go-1.5.4, /usr/local/go-1.6.2)
2. Run sudo switchgo -new go-1.5.4
3. Symbolic link /usr/local/go will be created
