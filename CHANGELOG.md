## v0.6.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#36](https://github.com/turbot/steampipe-plugin-shodan/pull/36))

## v0.6.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#34](https://github.com/turbot/steampipe-plugin-shodan/pull/34))
- Recompiled plugin with Go version `1.21`. ([#34](https://github.com/turbot/steampipe-plugin-shodan/pull/34))

## v0.5.0 [2023-04-07]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#26](https://github.com/turbot/steampipe-plugin-shodan/pull/26))

## v0.4.0 [2022-09-26]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#23](https://github.com/turbot/steampipe-plugin-shodan/pull/23))
- Recompiled plugin with Go version `1.19`. ([#23](https://github.com/turbot/steampipe-plugin-shodan/pull/23))

## v0.3.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#19](https://github.com/turbot/steampipe-plugin-shodan/pull/19))

## v0.3.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#16](https://github.com/turbot/steampipe-plugin-shodan/pull/16))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#18](https://github.com/turbot/steampipe-plugin-shodan/pull/18))

## v0.2.0 [2021-12-08]

_Enhancements_

- Recompiled plugin with Go version 1.17 ([#12](https://github.com/turbot/steampipe-plugin-shodan/pull/12))
- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#11](https://github.com/turbot/steampipe-plugin-shodan/pull/11))

## v0.1.0 [2021-08-06]

_Enhancements_

- Updated: README.md and docs/index.md have been updated with the latest formats ([#8](https://github.com/turbot/steampipe-plugin-shodan/pull/8))

_Bug fixes_

- Fixed: All Shodan website links now point to the correct URLs ([#7](https://github.com/turbot/steampipe-plugin-shodan/pull/7))

## v0.0.1 [2021-03-18]

_What's new?_

- New tables added
  - [shodan_account_profile](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_account_profile)
  - [shodan_api_info](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_api_info)
  - [shodan_dns_reverse](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_dns_reverse)
  - [shodan_domain](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_domain)
  - [shodan_exploit](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_exploit)
  - [shodan_host](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_host)
  - [shodan_host_service](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_host_service)
  - [shodan_port](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_port)
  - [shodan_protocol](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_protocol)
  - [shodan_scan](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_scan)
  - [shodan_service](https://hub.steampipe.io/plugins/turbot/shodan/tables/shodan_service)
