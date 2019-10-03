CLAsee - InfluxData CLA Checker
-------------------------------

A github action which checks the influxdata CLA for presence of the PR Author

```
GITHUB_ACTOR=pauldix
CLASEE_SECRET="$(cat secrets.json | base64)"

clasee <spreadsheet-id> <spreadsheet-range>
```
