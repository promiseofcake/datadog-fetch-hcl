# datadog-fetch-hcl

Tool to fetch [Datadog](https://www.datadoghq.com/) Dashboards and output an HCL
representation of them to stdout.

Useful for managing Dashboard state in: [TerraForm](https://github.com/hashicorp/terraform)

## installation

```bash
go get github.com/promiseofcake/datadog-fetch-hcl
```

## usage

Ensure your Datadog API / APP keys are exported

```bash
export DATADOG_API_KEY=foo
export DATADOG_APP_KEY=bar
```

Run via the following, you can redirect to a `.tf` file as you wish.

```bash
datadog-fetch-hcl -id <dashboard id> -title <resource title>
```

## limitations

Hackery:
- Resource title is passed in via CLI, not pulled from the remote dashboard (due to HCL encoder limitations)

Missing features:
- Datadog events overlays
- Precision / Aggreation metrics
- Probably lots else
