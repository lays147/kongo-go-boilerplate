# Kong Go Plugin Template

This template can help you to build a Kong Gateway plugin based in the Golang language.

It gives you all the boilerplate for development and production workflows. 

## How to develop

Fork it! :rocket:

```sh
# To build the plugins and docker images
make build-kong
# To start the gateway im dbless mode
make start-kong
```

### Project Structure

| File/Folder     | Purpose                    | 
|-----------------|----------------------------|
|config/kong.conf | Plugin configuration file  |
|config/kong.yaml | Declarative config of Kong |
|kong-plugin-template | source template for the go plugin |
|scripts/go_build.sh | Shell script responsible for building go plugins |
|scripts/go_plugins.conf | Conf file with the names of folders of go plugins |
|docker-compose.yaml | Docker-compose config of kong in dbless mode |
|Dockerfile | Dockerfile with multi-stage build of go plugins and the kong image |