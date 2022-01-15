# Satori/go.uuid replacement package 

This package is simply a shimming replacement package for the old satori/go.uuid package that suffers from an RCE.

The package uses type aliases and functions to forcibly replace this package with github.com/google/uuid

In the most part this should be a drop in unless you're using a lot of functionality that was specific to the old satori package.

