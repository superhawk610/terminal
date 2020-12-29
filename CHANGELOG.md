# Changelog

## Version 0.1.0: Export `Terminal`, move setup to New()

This version removed TTY setup from `init()` and into the `New()` constructor
for the newly exported `Terminal` struct. This allows creation of/interaction with
multiple TTYs in parallel, but most importantly prevents the lib from attempting
to initialize a terminal width during test/debug runs.

## Version 0.0.1: Initial Release

This version was released alongside [superhawk610/bar](https://github.com/superhawk610/bar).
