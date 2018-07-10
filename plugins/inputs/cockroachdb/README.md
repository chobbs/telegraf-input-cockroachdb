# CockroachDB Plugin

The CockroachDB Telegraf plugin gathers status metrics from one or more nodes in the Cockroach cluster for demo purposes only.

Download and install CockroachDB @https://www.cockroachlabs.com/docs/stable/install-cockroachdb.html

### Configuration:

```toml
# Description
[[inputs.cockroachdb]]
  ## URL's of CockroachDB status endpoint.
  servers = ["http://localhost:8080/_status/nodes/1"]
```

### Measurements & Fields:

CockroachDB provides one measurement named "cockroachdb", with the following fields:

- sys.cpu.user.percent
- sys.cpu.sys.percent
- timeseries.write.bytes
- timeseries.write.samples
- exec.latency-max

### Tags:

All measurements have the following tags:

- server (the host:port of the given server address, ex. `127.0.0.1:8087`)
- addressField (the internal node name received, ex. `roach1:26257`)

### Example Output:

```
$ ./telegraf --config telegraf.conf --input-filter cockroachdb --test
> cockroachdb,addressField=roach1=localhost:8080 sys.cpu.user.prcent=0.004999888952466365,sys.cpu.sys.percent=0.010999755695426005,timeseries.write.bytes=16668854,timeseries.write.samples=169916,exec.latency-max=6291455
```