# Latest http example

This example show cases how you can run a http webserver with versionify. This example contains a different structure than the http directory.

This structure hqw one latest version. This version contains all the relevant code and operates on a /v.x.x.x sub route.
Any old depercated methods are defined the other directories.

So we take the highest version (latest) and any lower version extends on this version. Also why we call it reverse!