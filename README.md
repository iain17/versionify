# Versionify

This project takes care of code versioning. This is done by registering the versions and adding methods to these versions.
Versions automatically inherit the methods registered in previous versions. This is done by building on top of [Hashicorp's Versioning Library for Go](https://github.com/hashicorp/go-version).

A common use case is versioning of a API. The routes and handlers of each version is not something we'd like to duplicate in our code base. Nor do we want to specify which older version it should also call take the routes from. This should all be done automatically and that's where this package comes in.

## Supports

The following use cases are supported out of the box:
- [Mux routes](http://www.gorillatoolkit.org/pkg/mux)

## Examples

See the examples directory for examples on how to use this package.