# Green Go

Green Go is a tiny utility to check the health status of web endpoints you'd love to keep
track of.

## Usage

The binary which you can compile, (or soon to be available as a downloadable executable) should
be executed with one additional command line argument to provide the list of endpoints that you
wish to monitor. The structure of the yaml is pretty straightforward.

```yaml
- endpoint: blog.svivekkrishna.cc
  status: 200
  protocol: https   
  port: 443 
```

The command line can be invoked with

```bash
green-go -f config.yaml
```

## Development

The project is a straightforward GoLang project. In the pursuit of supporting both a CLI and a 
library which could be then used to build a web interface, the main logic is abstracted into `lib`
and the CLI aspects in `cli`. A `makefile` has been made available to support with the commonly
used commands during development.

## Disclaimer

This project is my Dojo to learn GoLang. There will be lots of experimentation to help me on that
journey.