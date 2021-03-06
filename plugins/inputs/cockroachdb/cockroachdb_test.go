package cockroachdb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/influxdata/telegraf/testutil"
	"github.com/stretchr/testify/require"
)

func TestCockroachdb(t *testing.T) {
	// Create a test server with the const response JSON
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	// Parse the URL of the test server, used to verify the expected host
	u, err := url.Parse(ts.URL)
	require.NoError(t, err)

	// Create a new CockroachDB instance with our given test server
	Cockroachdb := NeCockroachdb()
	Cockroachdb.Servers = []string{ts.URL}

	// Create a test accumulator
	acc := &testutil.Accumulator{}

	// Gather data from the test server
	err = Cockroachdb.Gather(acc)
	require.NoError(t, err)

	// Expect the correct values for all known keys
	expectFields := map[string]interface{}{
		"sys.cpu.user.percent":     float64(0.004999888952466365),
		"sys.cpu.sys.percent":      float64(0.010999755695426005),
		"timeseries.write.bytes":   int(16668854),
		"timeseries.write.samples": int(169916),
		"exec.latency-max":         int(6291455),
	}
	// Expect the correct values for all tags
	expectTags := map[string]string{
		"addressField": "roach1:26257",
		"server":       u.Host,
	}

	acc.AssertContainsTaggedFields(t, "cockroachdb", expectFields, expectTags)
}

var response = `
{
  "desc": {
    "nodeId": 1,
    "address": {
      "networkField": "tcp",
      "addressField": "roach1:26257"
    },
    "attrs": {
      "attrs": [
      ]
    },
    "locality": {
      "tiers": [
      ]
    },
    "ServerVersion": {
      "majorVal": 2,
      "minorVal": 0,
      "patch": 0,
      "unstable": 0
    }
  },
  "buildInfo": {
    "goVersion": "go1.10",
    "tag": "v2.0.3",
    "time": "2018/06/18 16:11:33",
    "revision": "91715a9a95edbe716912173204fa4c0fc6724457",
    "cgoCompiler": "gcc 6.3.0",
    "cgoTargetTriple": "x86_64-unknown-linux-gnu",
    "platform": "linux amd64",
    "distribution": "CCL",
    "type": "release",
    "channel": "official-binary",
    "dependencies": null
  },
  "startedAt": "1530575893545315100",
  "updatedAt": "1530594689327860200",
  "metrics": {
    "build.timestamp": 1529338293,
    "clock-offset.meannanos": 0,
    "clock-offset.stddevnanos": 0,
    "distsender.batches": 2895,
    "distsender.batches.partial": 0,
    "distsender.errors.notleaseholder": 1,
    "distsender.rpc.sent": 2896,
    "distsender.rpc.sent.local": 2896,
    "distsender.rpc.sent.nextreplicaerror": 1,
    "exec.error": 1,
    "exec.latency-max": 6291455,
    "exec.latency-p50": 2490367,
    "exec.latency-p75": 3538943,
    "exec.latency-p90": 4718591,
    "exec.latency-p99": 6291455,
    "exec.latency-p99.9": 6291455,
    "exec.latency-p99.99": 6291455,
    "exec.latency-p99.999": 6291455,
    "exec.success": 2895,
    "gossip.bytes.received": 0,
    "gossip.bytes.sent": 0,
    "gossip.connections.incoming": 0,
    "gossip.connections.outgoing": 0,
    "gossip.connections.refused": 0,
    "gossip.infos.received": 0,
    "gossip.infos.sent": 0,
    "liveness.epochincrements": 0,
    "liveness.heartbeatfailures": 2,
    "liveness.heartbeatlatency-max": 5505023,
    "liveness.heartbeatlatency-p50": 3932159,
    "liveness.heartbeatlatency-p75": 4456447,
    "liveness.heartbeatlatency-p90": 5505023,
    "liveness.heartbeatlatency-p99": 5505023,
    "liveness.heartbeatlatency-p99.9": 5505023,
    "liveness.heartbeatlatency-p99.99": 5505023,
    "liveness.heartbeatlatency-p99.999": 5505023,
    "liveness.heartbeatsuccesses": 902,
    "liveness.livenodes": 1,
    "node-id": 1,
    "requests.slow.distsender": 0,
    "round-trip-latency-max": 0,
    "round-trip-latency-p50": 0,
    "round-trip-latency-p75": 0,
    "round-trip-latency-p90": 0,
    "round-trip-latency-p99": 0,
    "round-trip-latency-p99.9": 0,
    "round-trip-latency-p99.99": 0,
    "round-trip-latency-p99.999": 0,
    "sql.bytesin": 0,
    "sql.bytesout": 0,
    "sql.conns": 0,
    "sql.ddl.count": 0,
    "sql.delete.count": 0,
    "sql.distsql.exec.latency-max": 0,
    "sql.distsql.exec.latency-p50": 0,
    "sql.distsql.exec.latency-p75": 0,
    "sql.distsql.exec.latency-p90": 0,
    "sql.distsql.exec.latency-p99": 0,
    "sql.distsql.exec.latency-p99.9": 0,
    "sql.distsql.exec.latency-p99.99": 0,
    "sql.distsql.exec.latency-p99.999": 0,
    "sql.distsql.flows.active": 0,
    "sql.distsql.flows.total": 0,
    "sql.distsql.queries.active": 0,
    "sql.distsql.queries.total": 0,
    "sql.distsql.select.count": 0,
    "sql.distsql.service.latency-max": 0,
    "sql.distsql.service.latency-p50": 0,
    "sql.distsql.service.latency-p75": 0,
    "sql.distsql.service.latency-p90": 0,
    "sql.distsql.service.latency-p99": 0,
    "sql.distsql.service.latency-p99.9": 0,
    "sql.distsql.service.latency-p99.99": 0,
    "sql.distsql.service.latency-p99.999": 0,
    "sql.exec.latency-max": 0,
    "sql.exec.latency-p50": 0,
    "sql.exec.latency-p75": 0,
    "sql.exec.latency-p90": 0,
    "sql.exec.latency-p99": 0,
    "sql.exec.latency-p99.9": 0,
    "sql.exec.latency-p99.99": 0,
    "sql.exec.latency-p99.999": 0,
    "sql.insert.count": 0,
    "sql.mem.admin.current": 0,
    "sql.mem.admin.max-max": 0,
    "sql.mem.admin.max-p50": 0,
    "sql.mem.admin.max-p75": 0,
    "sql.mem.admin.max-p90": 0,
    "sql.mem.admin.max-p99": 0,
    "sql.mem.admin.max-p99.9": 0,
    "sql.mem.admin.max-p99.99": 0,
    "sql.mem.admin.max-p99.999": 0,
    "sql.mem.admin.session.current": 0,
    "sql.mem.admin.session.max-max": 0,
    "sql.mem.admin.session.max-p50": 0,
    "sql.mem.admin.session.max-p75": 0,
    "sql.mem.admin.session.max-p90": 0,
    "sql.mem.admin.session.max-p99": 0,
    "sql.mem.admin.session.max-p99.9": 0,
    "sql.mem.admin.session.max-p99.99": 0,
    "sql.mem.admin.session.max-p99.999": 0,
    "sql.mem.admin.txn.current": 0,
    "sql.mem.admin.txn.max-max": 0,
    "sql.mem.admin.txn.max-p50": 0,
    "sql.mem.admin.txn.max-p75": 0,
    "sql.mem.admin.txn.max-p90": 0,
    "sql.mem.admin.txn.max-p99": 0,
    "sql.mem.admin.txn.max-p99.9": 0,
    "sql.mem.admin.txn.max-p99.99": 0,
    "sql.mem.admin.txn.max-p99.999": 0,
    "sql.mem.client.current": 0,
    "sql.mem.client.max-max": 0,
    "sql.mem.client.max-p50": 0,
    "sql.mem.client.max-p75": 0,
    "sql.mem.client.max-p90": 0,
    "sql.mem.client.max-p99": 0,
    "sql.mem.client.max-p99.9": 0,
    "sql.mem.client.max-p99.99": 0,
    "sql.mem.client.max-p99.999": 0,
    "sql.mem.client.session.current": 0,
    "sql.mem.client.session.max-max": 0,
    "sql.mem.client.session.max-p50": 0,
    "sql.mem.client.session.max-p75": 0,
    "sql.mem.client.session.max-p90": 0,
    "sql.mem.client.session.max-p99": 0,
    "sql.mem.client.session.max-p99.9": 0,
    "sql.mem.client.session.max-p99.99": 0,
    "sql.mem.client.session.max-p99.999": 0,
    "sql.mem.client.txn.current": 0,
    "sql.mem.client.txn.max-max": 0,
    "sql.mem.client.txn.max-p50": 0,
    "sql.mem.client.txn.max-p75": 0,
    "sql.mem.client.txn.max-p90": 0,
    "sql.mem.client.txn.max-p99": 0,
    "sql.mem.client.txn.max-p99.9": 0,
    "sql.mem.client.txn.max-p99.99": 0,
    "sql.mem.client.txn.max-p99.999": 0,
    "sql.mem.conns.current": 0,
    "sql.mem.conns.max-max": 0,
    "sql.mem.conns.max-p50": 0,
    "sql.mem.conns.max-p75": 0,
    "sql.mem.conns.max-p90": 0,
    "sql.mem.conns.max-p99": 0,
    "sql.mem.conns.max-p99.9": 0,
    "sql.mem.conns.max-p99.99": 0,
    "sql.mem.conns.max-p99.999": 0,
    "sql.mem.conns.session.current": 0,
    "sql.mem.conns.session.max-max": 0,
    "sql.mem.conns.session.max-p50": 0,
    "sql.mem.conns.session.max-p75": 0,
    "sql.mem.conns.session.max-p90": 0,
    "sql.mem.conns.session.max-p99": 0,
    "sql.mem.conns.session.max-p99.9": 0,
    "sql.mem.conns.session.max-p99.99": 0,
    "sql.mem.conns.session.max-p99.999": 0,
    "sql.mem.conns.txn.current": 0,
    "sql.mem.conns.txn.max-max": 0,
    "sql.mem.conns.txn.max-p50": 0,
    "sql.mem.conns.txn.max-p75": 0,
    "sql.mem.conns.txn.max-p90": 0,
    "sql.mem.conns.txn.max-p99": 0,
    "sql.mem.conns.txn.max-p99.9": 0,
    "sql.mem.conns.txn.max-p99.99": 0,
    "sql.mem.conns.txn.max-p99.999": 0,
    "sql.mem.distsql.current": 0,
    "sql.mem.distsql.max-max": 0,
    "sql.mem.distsql.max-p50": 0,
    "sql.mem.distsql.max-p75": 0,
    "sql.mem.distsql.max-p90": 0,
    "sql.mem.distsql.max-p99": 0,
    "sql.mem.distsql.max-p99.9": 0,
    "sql.mem.distsql.max-p99.99": 0,
    "sql.mem.distsql.max-p99.999": 0,
    "sql.mem.internal.current": 0,
    "sql.mem.internal.max-max": 0,
    "sql.mem.internal.max-p50": 0,
    "sql.mem.internal.max-p75": 0,
    "sql.mem.internal.max-p90": 0,
    "sql.mem.internal.max-p99": 0,
    "sql.mem.internal.max-p99.9": 0,
    "sql.mem.internal.max-p99.99": 0,
    "sql.mem.internal.max-p99.999": 0,
    "sql.mem.internal.session.current": 0,
    "sql.mem.internal.session.max-max": 0,
    "sql.mem.internal.session.max-p50": 0,
    "sql.mem.internal.session.max-p75": 0,
    "sql.mem.internal.session.max-p90": 0,
    "sql.mem.internal.session.max-p99": 0,
    "sql.mem.internal.session.max-p99.9": 0,
    "sql.mem.internal.session.max-p99.99": 0,
    "sql.mem.internal.session.max-p99.999": 0,
    "sql.mem.internal.txn.current": 0,
    "sql.mem.internal.txn.max-max": 0,
    "sql.mem.internal.txn.max-p50": 0,
    "sql.mem.internal.txn.max-p75": 0,
    "sql.mem.internal.txn.max-p90": 0,
    "sql.mem.internal.txn.max-p99": 0,
    "sql.mem.internal.txn.max-p99.9": 0,
    "sql.mem.internal.txn.max-p99.99": 0,
    "sql.mem.internal.txn.max-p99.999": 0,
    "sql.misc.count": 0,
    "sql.query.count": 0,
    "sql.select.count": 0,
    "sql.service.latency-max": 0,
    "sql.service.latency-p50": 0,
    "sql.service.latency-p75": 0,
    "sql.service.latency-p90": 0,
    "sql.service.latency-p99": 0,
    "sql.service.latency-p99.9": 0,
    "sql.service.latency-p99.99": 0,
    "sql.service.latency-p99.999": 0,
    "sql.txn.abort.count": 0,
    "sql.txn.begin.count": 0,
    "sql.txn.commit.count": 0,
    "sql.txn.rollback.count": 0,
    "sql.update.count": 0,
    "sys.cgo.allocbytes": 72798648,
    "sys.cgo.totalbytes": 82190336,
    "sys.cgocalls": 235651,
    "sys.cpu.sys.ns": 39020000000,
    "sys.cpu.sys.percent": 0.010999755695426005,
    "sys.cpu.user.ns": 30050000000,
    "sys.cpu.user.percent": 0.004999888952466365,
    "sys.fd.open": 19,
    "sys.fd.softlimit": 1048576,
    "sys.gc.count": 43,
    "sys.gc.pause.ns": 8782600,
    "sys.gc.pause.percent": 0,
    "sys.go.allocbytes": 111360584,
    "sys.go.totalbytes": 168122616,
    "sys.goroutines": 116,
    "sys.rss": 212529152,
    "sys.uptime": 18796,
    "timeseries.write.bytes": 16668854,
    "timeseries.write.errors": 0,
    "timeseries.write.samples": 169916,
    "txn.abandons": 0,
    "txn.aborts": 0,
    "txn.autoretries": 0,
    "txn.commits": 899,
    "txn.commits1PC": 898,
    "txn.durations-max": 3670015,
    "txn.durations-p50": 2883583,
    "txn.durations-p75": 3407871,
    "txn.durations-p90": 3670015,
    "txn.durations-p99": 3670015,
    "txn.durations-p99.9": 3670015,
    "txn.durations-p99.99": 3670015,
    "txn.durations-p99.999": 3670015,
    "txn.restarts-max": 0,
    "txn.restarts-p50": 0,
    "txn.restarts-p75": 0,
    "txn.restarts-p90": 0,
    "txn.restarts-p99": 0,
    "txn.restarts-p99.9": 0,
    "txn.restarts-p99.99": 0,
    "txn.restarts-p99.999": 0,
    "txn.restarts.deleterange": 0,
    "txn.restarts.possiblereplay": 0,
    "txn.restarts.serializable": 0,
    "txn.restarts.writetooold": 0
  },
  "storeStatuses": [
    {
      "desc": {
        "storeId": 1,
        "attrs": {
          "attrs": [
          ]
        },
        "node": {
          "nodeId": 1,
          "address": {
            "networkField": "tcp",
            "addressField": "roach1:26257"
          },
          "attrs": {
            "attrs": [
            ]
          },
          "locality": {
            "tiers": [
            ]
          },
          "ServerVersion": {
            "majorVal": 2,
            "minorVal": 0,
            "patch": 0,
            "unstable": 0
          }
        },
        "capacity": {
          "capacity": "511962286915584",
          "available": "455453737746432",
          "used": "55891244",
          "logicalBytes": "37220469",
          "rangeCount": 20,
          "leaseCount": 18,
          "writesPerSecond": 43.92150548843001,
          "bytesPerReplica": {
            "p10": 0,
            "p25": 0,
            "p50": 94,
            "p75": 6278,
            "p90": 31835,
            "pMax": 37145486
          },
          "writesPerReplica": {
            "p10": 0,
            "p25": 0,
            "p50": 0,
            "p75": 0.016911665302134966,
            "p90": 0.6272180892857829,
            "pMax": 43.13451014260106
          }
        }
      },
      "metrics": {
        "addsstable.applications": 0,
        "addsstable.copies": 0,
        "addsstable.proposals": 0,
        "capacity": 511962286915584,
        "capacity.available": 455453737746432,
        "capacity.reserved": 0,
        "capacity.used": 55793626,
        "compactor.compactingnanos": 0,
        "compactor.compactions.failure": 0,
        "compactor.compactions.success": 0,
        "compactor.suggestionbytes.compacted": 0,
        "compactor.suggestionbytes.queued": 0,
        "compactor.suggestionbytes.skipped": 0,
        "gcbytesage": 841984651,
        "intentage": 0,
        "intentbytes": 0,
        "intentcount": 0,
        "keybytes": 33345,
        "keycount": 713,
        "lastupdatenanos": 1530594689321519800,
        "leases.epoch": 17,
        "leases.error": 0,
        "leases.expiration": 1,
        "leases.success": 751,
        "leases.transfers.error": 0,
        "leases.transfers.success": 0,
        "livebytes": 37188944,
        "livecount": 708,
        "queue.consistency.pending": 0,
        "queue.consistency.process.failure": 0,
        "queue.consistency.process.success": 0,
        "queue.consistency.processingnanos": 0,
        "queue.gc.info.abortspanconsidered": 0,
        "queue.gc.info.abortspangcnum": 1,
        "queue.gc.info.abortspanscanned": 1,
        "queue.gc.info.intentsconsidered": 0,
        "queue.gc.info.intenttxns": 0,
        "queue.gc.info.numkeysaffected": 7,
        "queue.gc.info.pushtxn": 0,
        "queue.gc.info.resolvesuccess": 0,
        "queue.gc.info.resolvetotal": 0,
        "queue.gc.info.transactionspangcaborted": 0,
        "queue.gc.info.transactionspangccommitted": 0,
        "queue.gc.info.transactionspangcpending": 0,
        "queue.gc.info.transactionspanscanned": 0,
        "queue.gc.pending": 0,
        "queue.gc.process.failure": 0,
        "queue.gc.process.success": 5,
        "queue.gc.processingnanos": 147793900,
        "queue.raftlog.pending": 0,
        "queue.raftlog.process.failure": 0,
        "queue.raftlog.process.success": 314,
        "queue.raftlog.processingnanos": 933601500,
        "queue.raftsnapshot.pending": 0,
        "queue.raftsnapshot.process.failure": 0,
        "queue.raftsnapshot.process.success": 0,
        "queue.raftsnapshot.processingnanos": 0,
        "queue.replicagc.pending": 0,
        "queue.replicagc.process.failure": 0,
        "queue.replicagc.process.success": 0,
        "queue.replicagc.processingnanos": 0,
        "queue.replicagc.removereplica": 0,
        "queue.replicate.addreplica": 0,
        "queue.replicate.pending": 0,
        "queue.replicate.process.failure": 1776,
        "queue.replicate.process.success": 0,
        "queue.replicate.processingnanos": 130125000,
        "queue.replicate.purgatory": 20,
        "queue.replicate.rebalancereplica": 0,
        "queue.replicate.removedeadreplica": 0,
        "queue.replicate.removereplica": 0,
        "queue.replicate.transferlease": 0,
        "queue.split.pending": 0,
        "queue.split.process.failure": 0,
        "queue.split.process.success": 0,
        "queue.split.processingnanos": 0,
        "queue.tsmaintenance.pending": 0,
        "queue.tsmaintenance.process.failure": 0,
        "queue.tsmaintenance.process.success": 0,
        "queue.tsmaintenance.processingnanos": 0,
        "raft.commandsapplied": 2789,
        "raft.enqueued.pending": 0,
        "raft.heartbeats.pending": 0,
        "raft.process.commandcommit.latency-max": 3014655,
        "raft.process.commandcommit.latency-p50": 557055,
        "raft.process.commandcommit.latency-p75": 1114111,
        "raft.process.commandcommit.latency-p90": 1835007,
        "raft.process.commandcommit.latency-p99": 3014655,
        "raft.process.commandcommit.latency-p99.9": 3014655,
        "raft.process.commandcommit.latency-p99.99": 3014655,
        "raft.process.commandcommit.latency-p99.999": 3014655,
        "raft.process.logcommit.latency-max": 2031615,
        "raft.process.logcommit.latency-p50": 884735,
        "raft.process.logcommit.latency-p75": 1245183,
        "raft.process.logcommit.latency-p90": 1376255,
        "raft.process.logcommit.latency-p99": 2031615,
        "raft.process.logcommit.latency-p99.9": 2031615,
        "raft.process.logcommit.latency-p99.99": 2031615,
        "raft.process.logcommit.latency-p99.999": 2031615,
        "raft.process.tickingnanos": 119916600,
        "raft.process.workingnanos": 6499028100,
        "raft.rcvd.app": 0,
        "raft.rcvd.appresp": 0,
        "raft.rcvd.dropped": 0,
        "raft.rcvd.heartbeat": 0,
        "raft.rcvd.heartbeatresp": 0,
        "raft.rcvd.prevote": 0,
        "raft.rcvd.prevoteresp": 0,
        "raft.rcvd.prop": 0,
        "raft.rcvd.snap": 0,
        "raft.rcvd.timeoutnow": 0,
        "raft.rcvd.transferleader": 0,
        "raft.rcvd.vote": 0,
        "raft.rcvd.voteresp": 0,
        "raft.ticks": 19719,
        "raftlog.behind": 0,
        "raftlog.truncated": 2812,
        "range.adds": 0,
        "range.raftleadertransfers": 0,
        "range.removes": 0,
        "range.snapshots.generated": 0,
        "range.snapshots.normal-applied": 0,
        "range.snapshots.preemptive-applied": 0,
        "range.splits": 0,
        "ranges": 20,
        "ranges.unavailable": 0,
        "ranges.underreplicated": 20,
        "rebalancing.writespersecond": 43.654471389706174,
        "replicas": 20,
        "replicas.commandqueue.combinedqueuesize": 428,
        "replicas.commandqueue.combinedreadcount": 0,
        "replicas.commandqueue.combinedwritecount": 428,
        "replicas.commandqueue.maxoverlaps": 0,
        "replicas.commandqueue.maxreadcount": 0,
        "replicas.commandqueue.maxsize": 428,
        "replicas.commandqueue.maxtreesize": 1,
        "replicas.commandqueue.maxwritecount": 428,
        "replicas.leaders": 20,
        "replicas.leaders_not_leaseholders": 0,
        "replicas.leaseholders": 18,
        "replicas.quiescent": 20,
        "replicas.reserved": 0,
        "requests.backpressure.split": 0,
        "requests.slow.commandqueue": 0,
        "requests.slow.lease": 0,
        "requests.slow.raft": 0,
        "rocksdb.block.cache.hits": 72712,
        "rocksdb.block.cache.misses": 311,
        "rocksdb.block.cache.pinned-usage": 0,
        "rocksdb.block.cache.usage": 552960,
        "rocksdb.bloom.filter.prefix.checked": 6,
        "rocksdb.bloom.filter.prefix.useful": 4,
        "rocksdb.compactions": 1,
        "rocksdb.flushes": 0,
        "rocksdb.memtable.total-size": 59008792,
        "rocksdb.num-sstables": 1,
        "rocksdb.read-amplification": 1,
        "rocksdb.table-readers-mem-estimate": 20480,
        "sysbytes": 11675,
        "syscount": 200,
        "totalbytes": 37220469,
        "tscache.skl.read.pages": 1,
        "tscache.skl.read.rotations": 0,
        "tscache.skl.write.pages": 1,
        "tscache.skl.write.rotations": 0,
        "valbytes": 37187124,
        "valcount": 962
      }
    }
  ],
  "args": [
    "/cockroach/cockroach",
    "start",
    "--insecure"
  ],
  "env": [
    "COCKROACH_CHANNEL=official-docker"
  ],
  "latencies": {
  },
  "activity": {
    "1": {
      "incoming": "0",
      "outgoing": "0",
      "latency": "0"
    }
  }
}
`
