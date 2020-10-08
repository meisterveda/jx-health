# Metadata for health

For now this is a file based lookup which could be externalised, internationalised and even replaced all together with a better solution.

Initial metadata files are:
- `Information` links to more details on what a specific health check is

The data is currently in a `./static_data/info.yaml` file

[`go-bindata`](https://github.com/go-bindata/go-bindata) is used to build the static data into an Asset

To regenerate run `make bind` 
