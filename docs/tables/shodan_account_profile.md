---
title: "Steampipe Table: shodan_account_profile - Query Shodan Account Profiles using SQL"
description: "Allows users to query Shodan Account Profiles, specifically the information related to the Shodan account. This includes data such as credit balance, display name, and the member status of the user."
---

# Table: shodan_account_profile - Query Shodan Account Profiles using SQL

A Shodan Account Profile is a user account on Shodan, a search engine for internet-connected devices. The profile includes information about the user, such as their display name, email, and credit balance. The profile also indicates whether the user is a member of the Shodan platform.

## Table Usage Guide

The `shodan_account_profile` table provides insights into user profiles within Shodan. As a cybersecurity analyst, you can explore profile-specific details through this table, including display name, email, and credit balance. Use this table to uncover information about the users, such as their membership status and the number of query credits they have available.

## Examples

### List the scans
Explore your Shodan account profile to gain insights into the details of your scans. This could be useful in identifying potential security vulnerabilities and understanding areas for improvement.

```sql
select
  *
from
  shodan_account_profile
```