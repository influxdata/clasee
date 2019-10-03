CLAsee - InfluxData CLA Checker
-------------------------------

A github action which checks the influxdata CLA for presence of the PR Author

```
export GITHUB_ACTOR=pauldix
export CLASEE_SECRET="$(cat secrets.json | base64)"

clasee <spreadsheet-id> <spreadsheet-range>
```

Demonstration of a github workflow yaml:

```yaml
on:
  pull_request:
    types: [opened]

jobs:
  test_action:
    runs-on: ubuntu-latest
    name: "Test Action Job"
    steps:
      - uses: actions/checkout@v1
      - name: "Test Action"
        uses: influxdata/clasee@v1
        with:
          spreadsheet: "1jnRZYSw83oa6hcEBb1lxK6nNvXrWnOzPT8Bz9iR4Q8s"
          range: "Form Responses!E:E"
        env:
          CLASEE_SECRET: ${{ secrets.CLASEE_SECRET }}
```
