# Http example

This example show cases how you can run a http webserver with versionify.

## Run the example

```bash
    go run main.go 
```
This will run the example locally on port 8080

try the following:

- [Method foo on version 1](http://localhost:8080/v1.0.0/foo) Working fine. Like normal.
- [Method bar on version 1](http://localhost:8080/v1.0.0/bar) Working fine. Like normal.

- [Method foo using the version 1](http://localhost:8080/v2.0.0/foo) because we haven't specified it. or set a constraint!
- [Method bar overwritten on version 2](http://localhost:8080/v2.0.0/bar) Notice how it doesn't use the version 1 handler!
- [Method only on version 2](http://localhost:8080/v2.0.0/iain)

- [Method NOT on version 3](http://localhost:8080/v3.0.0/iain) Because we've set a constraint to '<= 2.0'