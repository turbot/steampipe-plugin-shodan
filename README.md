![image](https://hub.steampipe.io/images/plugins/turbot/shodan-social-graphic.png)

# Shodan Plugin for Steampipe

* **[Get started →](https://hub.steampipe.io/plugins/turbot/shodan)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/shodan/tables)
* Community: [Slack Channel](https://steampipe.io/community/join)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-shodan/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install shodan
```

Run a query:

```sql
select * from shodan_host_service where ip = '8.8.8.8'
```

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-shodan.git
cd steampipe-plugin-shodan
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/shodan.spc
```

Try it!
```
steampipe query
> .inspect shodan
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [MPL-2.0 open source license](https://github.com/turbot/steampipe-plugin-shodan/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Shodan Plugin](https://github.com/turbot/steampipe-plugin-shodan/labels/help%20wanted)
