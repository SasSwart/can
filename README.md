# can

The gin templates that are available in the [existing openapi generator](https://openapi-generator.tech/docs/generators/go-gin-server)
require in place modification. This does not suit our workflow. We would like to be able to regenerate our API interface
without having to patch in our own code every time. gRPC generates code that illustrates how to generate code that can
be used without modification.

This repository is intended to apply the pattern used by gRPC to Gin for rest services.

## Usage
Download the latest [release](https://github.com/SasSwart/gin-in-a-can/releases) archive for your platform and extract it.
Copy the [example config file](cmd/petstore/can.yml) to a convenient location and modify it to your needs.

Run the can command as follows:
`/path/to/can --configFile=/path/to/can.yml`

## Design
See [design.md](design.md)