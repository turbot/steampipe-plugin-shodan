## Note for users new to Steampipe
The Shodan Plugin for Steampipe can be managed automatically with the Steampipe CLI.
For more information on how to get started view the [documentation](https://hub.steampipe.io/plugins/turbot/shodan).

# The Shodan Plugin for Steampipe

Use SQL to query Shodan APIs including host, DNS and exploit information. For example:

```sql
select * from shodan_host_service where ip = '8.8.8.8'
```

- [Documentation](https://hub.steampipe.io/plugins/turbot/shodan)
- [Tables & schemas](https://hub.steampipe.io/plugins/turbot/shodan/tables)
- [Shodan plugin issues](https://github.com/turbot/steampipe-plugin-shodan/issues)
- [Steampipe issues](https://github.com/turbot/steampipe/issues)
- [Discussion forums](https://github.com/turbot/steampipe/discussions)

## Requirements

- [Steampipe](https://steampipe.io/downloads) v0.3.3 or greater
- [Go](https://golang.org/doc/install) 1.12.3 (recommended)

## Building the Plugin

Install Steampipe: [Instructions](https://steampipe.io/downloads)

Clone repository:

```sh
$ mkdir -p $GOPATH/src/github.com/turbot; cd $GOPATH/src/github.com/turbot
$ git clone git@github.com:turbot/steampipe-plugin-shodan
```

Enter the plugin directory and build the plugin:

```sh
$ cd $GOPATH/src/github.com/turbot/steampipe-plugin-shodan
$ make
```

Copy configuration files from repo to local steampipe config:
```sh
$ cp $GOPATH/src/github.com/turbot/steampipe-plugin-shodan/config/* ~/.steampipe/config
```

View and update the configuration file settings for your environment [see the plugin docs](https://hub.steampipe.io/plugins/turbot/shodan) for additional info:
```sh
$ pico ~/.steampipe/config/shodan.spc
```

## Using the Plugin

During the `make` process, the script will output the plugin to `~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/shodan@latest/` which is the default location for steampipe plugins. Restart Steampipe if already running. Then try a test query:

```sql
select * from shodan_host_service where ip = '8.8.8.8'
```

## Developing the Plugin

To add features to the Plugin, install [Go](http://www.golang.org) and configure your your [GOPATH](http://golang.org/doc/code.html#GOPATH)

Compile the Plugin by running `make`. The Plugin binary will output to your Steampipe plugin directory.

```sh
$ make
```

## Community

The Steampipe community can be found on [GitHub Discussions](https://github.com/turbot/steampipe/discussions), where you can ask questions, voice ideas, and share your projects. Our [Code of Conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md) applies to all Steampipe community channels.

## Contributing

Please see [CONTRIBUTING.md](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md).

`Help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Shodan Plugin](https://github.com/turbot/steampipe-plugin-shodan/labels/help%20wanted)

## License

By contributing to Steampipe and Steampipe plugins you agree that your contributions will be licensed as defined on the [LICENSE](LICENSE) file.