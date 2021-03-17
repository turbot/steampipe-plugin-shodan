module github.com/turbot/steampipe-plugin-shodan

go 1.15

replace github.com/shadowscatcher/shodan => github.com/e-gineer/shodan v1.0.6-0.20210316235051-ba6827fed7ce

require (
	github.com/shadowscatcher/shodan v1.0.5
	github.com/turbot/steampipe-plugin-sdk v0.2.3
)
