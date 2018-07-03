package cockroachdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Cockroachdb struct {
	Servers []string

	// HTTP client & request
	client *http.Client
}

// NewCockroachdb return a new instance of Cockroachdb with a default http client
func NeCockroachdb() *Cockroachdb {
	tr := &http.Transport{ResponseHeaderTimeout: time.Duration(3 * time.Second)}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(4 * time.Second),
	}
	return &Cockroachdb{client: client}
}

type Cockroach struct {
	Desc struct {
		NodeID  int `json:"nodeId"`
		Address struct {
			NetworkField string `json:"networkField"`
			AddressField string `json:"addressField"`
		} `json:"address"`
		Attrs struct {
			Attrs []interface{} `json:"attrs"`
		} `json:"attrs"`
		Locality struct {
			Tiers []interface{} `json:"tiers"`
		} `json:"locality"`
		ServerVersion struct {
			MajorVal int `json:"majorVal"`
			MinorVal int `json:"minorVal"`
			Patch    int `json:"patch"`
			Unstable int `json:"unstable"`
		} `json:"ServerVersion"`
	} `json:"desc"`
	BuildInfo struct {
		GoVersion       string      `json:"goVersion"`
		Tag             string      `json:"tag"`
		Time            string      `json:"time"`
		Revision        string      `json:"revision"`
		CgoCompiler     string      `json:"cgoCompiler"`
		CgoTargetTriple string      `json:"cgoTargetTriple"`
		Platform        string      `json:"platform"`
		Distribution    string      `json:"distribution"`
		Type            string      `json:"type"`
		Channel         string      `json:"channel"`
		Dependencies    interface{} `json:"dependencies"`
	} `json:"buildInfo"`
	StartedAt string `json:"startedAt"`
	UpdatedAt string `json:"updatedAt"`
	Metrics   struct {
		BuildTimestamp                    int     `json:"build.timestamp"`
		ClockOffsetMeannanos              int     `json:"clock-offset.meannanos"`
		ClockOffsetStddevnanos            int     `json:"clock-offset.stddevnanos"`
		DistsenderBatches                 int     `json:"distsender.batches"`
		DistsenderBatchesPartial          int     `json:"distsender.batches.partial"`
		DistsenderErrorsNotleaseholder    int     `json:"distsender.errors.notleaseholder"`
		DistsenderRPCSent                 int     `json:"distsender.rpc.sent"`
		DistsenderRPCSentLocal            int     `json:"distsender.rpc.sent.local"`
		DistsenderRPCSentNextreplicaerror int     `json:"distsender.rpc.sent.nextreplicaerror"`
		ExecError                         int     `json:"exec.error"`
		ExecLatencyMax                    int     `json:"exec.latency-max"`
		ExecLatencyP50                    int     `json:"exec.latency-p50"`
		ExecLatencyP75                    int     `json:"exec.latency-p75"`
		ExecLatencyP90                    int     `json:"exec.latency-p90"`
		ExecLatencyP99                    int     `json:"exec.latency-p99"`
		ExecLatencyP999                   int     `json:"exec.latency-p99.9"`
		ExecLatencyP9999                  int     `json:"exec.latency-p99.99"`
		ExecLatencyP99999                 int     `json:"exec.latency-p99.999"`
		ExecSuccess                       int     `json:"exec.success"`
		GossipBytesReceived               int     `json:"gossip.bytes.received"`
		GossipBytesSent                   int     `json:"gossip.bytes.sent"`
		GossipConnectionsIncoming         int     `json:"gossip.connections.incoming"`
		GossipConnectionsOutgoing         int     `json:"gossip.connections.outgoing"`
		GossipConnectionsRefused          int     `json:"gossip.connections.refused"`
		GossipInfosReceived               int     `json:"gossip.infos.received"`
		GossipInfosSent                   int     `json:"gossip.infos.sent"`
		LivenessEpochincrements           int     `json:"liveness.epochincrements"`
		LivenessHeartbeatfailures         int     `json:"liveness.heartbeatfailures"`
		LivenessHeartbeatlatencyMax       int     `json:"liveness.heartbeatlatency-max"`
		LivenessHeartbeatlatencyP50       int     `json:"liveness.heartbeatlatency-p50"`
		LivenessHeartbeatlatencyP75       int     `json:"liveness.heartbeatlatency-p75"`
		LivenessHeartbeatlatencyP90       int     `json:"liveness.heartbeatlatency-p90"`
		LivenessHeartbeatlatencyP99       int     `json:"liveness.heartbeatlatency-p99"`
		LivenessHeartbeatlatencyP999      int     `json:"liveness.heartbeatlatency-p99.9"`
		LivenessHeartbeatlatencyP9999     int     `json:"liveness.heartbeatlatency-p99.99"`
		LivenessHeartbeatlatencyP99999    int     `json:"liveness.heartbeatlatency-p99.999"`
		LivenessHeartbeatsuccesses        int     `json:"liveness.heartbeatsuccesses"`
		LivenessLivenodes                 int     `json:"liveness.livenodes"`
		NodeID                            int     `json:"node-id"`
		RequestsSlowDistsender            int     `json:"requests.slow.distsender"`
		RoundTripLatencyMax               int     `json:"round-trip-latency-max"`
		RoundTripLatencyP50               int     `json:"round-trip-latency-p50"`
		RoundTripLatencyP75               int     `json:"round-trip-latency-p75"`
		RoundTripLatencyP90               int     `json:"round-trip-latency-p90"`
		RoundTripLatencyP99               int     `json:"round-trip-latency-p99"`
		RoundTripLatencyP999              int     `json:"round-trip-latency-p99.9"`
		RoundTripLatencyP9999             int     `json:"round-trip-latency-p99.99"`
		RoundTripLatencyP99999            int     `json:"round-trip-latency-p99.999"`
		SqlBytesin                        int     `json:"sql.bytesin"`
		SqlBytesout                       int     `json:"sql.bytesout"`
		SqlConns                          int     `json:"sql.conns"`
		SqlDdlCount                       int     `json:"sql.ddl.count"`
		SqlDeleteCount                    int     `json:"sql.delete.count"`
		SqlDistsqlExecLatencyMax          int     `json:"sql.distsql.exec.latency-max"`
		SqlDistsqlExecLatencyP50          int     `json:"sql.distsql.exec.latency-p50"`
		SqlDistsqlExecLatencyP75          int     `json:"sql.distsql.exec.latency-p75"`
		SqlDistsqlExecLatencyP90          int     `json:"sql.distsql.exec.latency-p90"`
		SqlDistsqlExecLatencyP99          int     `json:"sql.distsql.exec.latency-p99"`
		SqlDistsqlExecLatencyP999         int     `json:"sql.distsql.exec.latency-p99.9"`
		SqlDistsqlExecLatencyP9999        int     `json:"sql.distsql.exec.latency-p99.99"`
		SqlDistsqlExecLatencyP99999       int     `json:"sql.distsql.exec.latency-p99.999"`
		SqlDistsqlFlowsActive             int     `json:"sql.distsql.flows.active"`
		SqlDistsqlFlowsTotal              int     `json:"sql.distsql.flows.total"`
		SqlDistsqlQueriesActive           int     `json:"sql.distsql.queries.active"`
		SqlDistsqlQueriesTotal            int     `json:"sql.distsql.queries.total"`
		SqlDistsqlSelectCount             int     `json:"sql.distsql.select.count"`
		SqlDistsqlServiceLatencyMax       int     `json:"sql.distsql.service.latency-max"`
		SqlDistsqlServiceLatencyP50       int     `json:"sql.distsql.service.latency-p50"`
		SqlDistsqlServiceLatencyP75       int     `json:"sql.distsql.service.latency-p75"`
		SqlDistsqlServiceLatencyP90       int     `json:"sql.distsql.service.latency-p90"`
		SqlDistsqlServiceLatencyP99       int     `json:"sql.distsql.service.latency-p99"`
		SqlDistsqlServiceLatencyP999      int     `json:"sql.distsql.service.latency-p99.9"`
		SqlDistsqlServiceLatencyP9999     int     `json:"sql.distsql.service.latency-p99.99"`
		SqlDistsqlServiceLatencyP99999    int     `json:"sql.distsql.service.latency-p99.999"`
		SqlExecLatencyMax                 int     `json:"sql.exec.latency-max"`
		SqlExecLatencyP50                 int     `json:"sql.exec.latency-p50"`
		SqlExecLatencyP75                 int     `json:"sql.exec.latency-p75"`
		SqlExecLatencyP90                 int     `json:"sql.exec.latency-p90"`
		SqlExecLatencyP99                 int     `json:"sql.exec.latency-p99"`
		SqlExecLatencyP999                int     `json:"sql.exec.latency-p99.9"`
		SqlExecLatencyP9999               int     `json:"sql.exec.latency-p99.99"`
		SqlExecLatencyP99999              int     `json:"sql.exec.latency-p99.999"`
		SqlInsertCount                    int     `json:"sql.insert.count"`
		SqlMemAdminCurrent                int     `json:"sql.mem.admin.current"`
		SqlMemAdminMaxMax                 int     `json:"sql.mem.admin.max-max"`
		SqlMemAdminMaxP50                 int     `json:"sql.mem.admin.max-p50"`
		SqlMemAdminMaxP75                 int     `json:"sql.mem.admin.max-p75"`
		SqlMemAdminMaxP90                 int     `json:"sql.mem.admin.max-p90"`
		SqlMemAdminMaxP99                 int     `json:"sql.mem.admin.max-p99"`
		SqlMemAdminMaxP999                int     `json:"sql.mem.admin.max-p99.9"`
		SqlMemAdminMaxP9999               int     `json:"sql.mem.admin.max-p99.99"`
		SqlMemAdminMaxP99999              int     `json:"sql.mem.admin.max-p99.999"`
		SqlMemAdminSessionCurrent         int     `json:"sql.mem.admin.session.current"`
		SqlMemAdminSessionMaxMax          int     `json:"sql.mem.admin.session.max-max"`
		SqlMemAdminSessionMaxP50          int     `json:"sql.mem.admin.session.max-p50"`
		SqlMemAdminSessionMaxP75          int     `json:"sql.mem.admin.session.max-p75"`
		SqlMemAdminSessionMaxP90          int     `json:"sql.mem.admin.session.max-p90"`
		SqlMemAdminSessionMaxP99          int     `json:"sql.mem.admin.session.max-p99"`
		SqlMemAdminSessionMaxP999         int     `json:"sql.mem.admin.session.max-p99.9"`
		SqlMemAdminSessionMaxP9999        int     `json:"sql.mem.admin.session.max-p99.99"`
		SqlMemAdminSessionMaxP99999       int     `json:"sql.mem.admin.session.max-p99.999"`
		SqlMemAdminTxnCurrent             int     `json:"sql.mem.admin.txn.current"`
		SqlMemAdminTxnMaxMax              int     `json:"sql.mem.admin.txn.max-max"`
		SqlMemAdminTxnMaxP50              int     `json:"sql.mem.admin.txn.max-p50"`
		SqlMemAdminTxnMaxP75              int     `json:"sql.mem.admin.txn.max-p75"`
		SqlMemAdminTxnMaxP90              int     `json:"sql.mem.admin.txn.max-p90"`
		SqlMemAdminTxnMaxP99              int     `json:"sql.mem.admin.txn.max-p99"`
		SqlMemAdminTxnMaxP999             int     `json:"sql.mem.admin.txn.max-p99.9"`
		SqlMemAdminTxnMaxP9999            int     `json:"sql.mem.admin.txn.max-p99.99"`
		SqlMemAdminTxnMaxP99999           int     `json:"sql.mem.admin.txn.max-p99.999"`
		SqlMemClientCurrent               int     `json:"sql.mem.client.current"`
		SqlMemClientMaxMax                int     `json:"sql.mem.client.max-max"`
		SqlMemClientMaxP50                int     `json:"sql.mem.client.max-p50"`
		SqlMemClientMaxP75                int     `json:"sql.mem.client.max-p75"`
		SqlMemClientMaxP90                int     `json:"sql.mem.client.max-p90"`
		SqlMemClientMaxP99                int     `json:"sql.mem.client.max-p99"`
		SqlMemClientMaxP999               int     `json:"sql.mem.client.max-p99.9"`
		SqlMemClientMaxP9999              int     `json:"sql.mem.client.max-p99.99"`
		SqlMemClientMaxP99999             int     `json:"sql.mem.client.max-p99.999"`
		SqlMemClientSessionCurrent        int     `json:"sql.mem.client.session.current"`
		SqlMemClientSessionMaxMax         int     `json:"sql.mem.client.session.max-max"`
		SqlMemClientSessionMaxP50         int     `json:"sql.mem.client.session.max-p50"`
		SqlMemClientSessionMaxP75         int     `json:"sql.mem.client.session.max-p75"`
		SqlMemClientSessionMaxP90         int     `json:"sql.mem.client.session.max-p90"`
		SqlMemClientSessionMaxP99         int     `json:"sql.mem.client.session.max-p99"`
		SqlMemClientSessionMaxP999        int     `json:"sql.mem.client.session.max-p99.9"`
		SqlMemClientSessionMaxP9999       int     `json:"sql.mem.client.session.max-p99.99"`
		SqlMemClientSessionMaxP99999      int     `json:"sql.mem.client.session.max-p99.999"`
		SqlMemClientTxnCurrent            int     `json:"sql.mem.client.txn.current"`
		SqlMemClientTxnMaxMax             int     `json:"sql.mem.client.txn.max-max"`
		SqlMemClientTxnMaxP50             int     `json:"sql.mem.client.txn.max-p50"`
		SqlMemClientTxnMaxP75             int     `json:"sql.mem.client.txn.max-p75"`
		SqlMemClientTxnMaxP90             int     `json:"sql.mem.client.txn.max-p90"`
		SqlMemClientTxnMaxP99             int     `json:"sql.mem.client.txn.max-p99"`
		SqlMemClientTxnMaxP999            int     `json:"sql.mem.client.txn.max-p99.9"`
		SqlMemClientTxnMaxP9999           int     `json:"sql.mem.client.txn.max-p99.99"`
		SqlMemClientTxnMaxP99999          int     `json:"sql.mem.client.txn.max-p99.999"`
		SqlMemConnsCurrent                int     `json:"sql.mem.conns.current"`
		SqlMemConnsMaxMax                 int     `json:"sql.mem.conns.max-max"`
		SqlMemConnsMaxP50                 int     `json:"sql.mem.conns.max-p50"`
		SqlMemConnsMaxP75                 int     `json:"sql.mem.conns.max-p75"`
		SqlMemConnsMaxP90                 int     `json:"sql.mem.conns.max-p90"`
		SqlMemConnsMaxP99                 int     `json:"sql.mem.conns.max-p99"`
		SqlMemConnsMaxP999                int     `json:"sql.mem.conns.max-p99.9"`
		SqlMemConnsMaxP9999               int     `json:"sql.mem.conns.max-p99.99"`
		SqlMemConnsMaxP99999              int     `json:"sql.mem.conns.max-p99.999"`
		SqlMemConnsSessionCurrent         int     `json:"sql.mem.conns.session.current"`
		SqlMemConnsSessionMaxMax          int     `json:"sql.mem.conns.session.max-max"`
		SqlMemConnsSessionMaxP50          int     `json:"sql.mem.conns.session.max-p50"`
		SqlMemConnsSessionMaxP75          int     `json:"sql.mem.conns.session.max-p75"`
		SqlMemConnsSessionMaxP90          int     `json:"sql.mem.conns.session.max-p90"`
		SqlMemConnsSessionMaxP99          int     `json:"sql.mem.conns.session.max-p99"`
		SqlMemConnsSessionMaxP999         int     `json:"sql.mem.conns.session.max-p99.9"`
		SqlMemConnsSessionMaxP9999        int     `json:"sql.mem.conns.session.max-p99.99"`
		SqlMemConnsSessionMaxP99999       int     `json:"sql.mem.conns.session.max-p99.999"`
		SqlMemConnsTxnCurrent             int     `json:"sql.mem.conns.txn.current"`
		SqlMemConnsTxnMaxMax              int     `json:"sql.mem.conns.txn.max-max"`
		SqlMemConnsTxnMaxP50              int     `json:"sql.mem.conns.txn.max-p50"`
		SqlMemConnsTxnMaxP75              int     `json:"sql.mem.conns.txn.max-p75"`
		SqlMemConnsTxnMaxP90              int     `json:"sql.mem.conns.txn.max-p90"`
		SqlMemConnsTxnMaxP99              int     `json:"sql.mem.conns.txn.max-p99"`
		SqlMemConnsTxnMaxP999             int     `json:"sql.mem.conns.txn.max-p99.9"`
		SqlMemConnsTxnMaxP9999            int     `json:"sql.mem.conns.txn.max-p99.99"`
		SqlMemConnsTxnMaxP99999           int     `json:"sql.mem.conns.txn.max-p99.999"`
		SqlMemDistsqlCurrent              int     `json:"sql.mem.distsql.current"`
		SqlMemDistsqlMaxMax               int     `json:"sql.mem.distsql.max-max"`
		SqlMemDistsqlMaxP50               int     `json:"sql.mem.distsql.max-p50"`
		SqlMemDistsqlMaxP75               int     `json:"sql.mem.distsql.max-p75"`
		SqlMemDistsqlMaxP90               int     `json:"sql.mem.distsql.max-p90"`
		SqlMemDistsqlMaxP99               int     `json:"sql.mem.distsql.max-p99"`
		SqlMemDistsqlMaxP999              int     `json:"sql.mem.distsql.max-p99.9"`
		SqlMemDistsqlMaxP9999             int     `json:"sql.mem.distsql.max-p99.99"`
		SqlMemDistsqlMaxP99999            int     `json:"sql.mem.distsql.max-p99.999"`
		SqlMemInternalCurrent             int     `json:"sql.mem.internal.current"`
		SqlMemInternalMaxMax              int     `json:"sql.mem.internal.max-max"`
		SqlMemInternalMaxP50              int     `json:"sql.mem.internal.max-p50"`
		SqlMemInternalMaxP75              int     `json:"sql.mem.internal.max-p75"`
		SqlMemInternalMaxP90              int     `json:"sql.mem.internal.max-p90"`
		SqlMemInternalMaxP99              int     `json:"sql.mem.internal.max-p99"`
		SqlMemInternalMaxP999             int     `json:"sql.mem.internal.max-p99.9"`
		SqlMemInternalMaxP9999            int     `json:"sql.mem.internal.max-p99.99"`
		SqlMemInternalMaxP99999           int     `json:"sql.mem.internal.max-p99.999"`
		SqlMemInternalSessionCurrent      int     `json:"sql.mem.internal.session.current"`
		SqlMemInternalSessionMaxMax       int     `json:"sql.mem.internal.session.max-max"`
		SqlMemInternalSessionMaxP50       int     `json:"sql.mem.internal.session.max-p50"`
		SqlMemInternalSessionMaxP75       int     `json:"sql.mem.internal.session.max-p75"`
		SqlMemInternalSessionMaxP90       int     `json:"sql.mem.internal.session.max-p90"`
		SqlMemInternalSessionMaxP99       int     `json:"sql.mem.internal.session.max-p99"`
		SqlMemInternalSessionMaxP999      int     `json:"sql.mem.internal.session.max-p99.9"`
		SqlMemInternalSessionMaxP9999     int     `json:"sql.mem.internal.session.max-p99.99"`
		SqlMemInternalSessionMaxP99999    int     `json:"sql.mem.internal.session.max-p99.999"`
		SqlMemInternalTxnCurrent          int     `json:"sql.mem.internal.txn.current"`
		SqlMemInternalTxnMaxMax           int     `json:"sql.mem.internal.txn.max-max"`
		SqlMemInternalTxnMaxP50           int     `json:"sql.mem.internal.txn.max-p50"`
		SqlMemInternalTxnMaxP75           int     `json:"sql.mem.internal.txn.max-p75"`
		SqlMemInternalTxnMaxP90           int     `json:"sql.mem.internal.txn.max-p90"`
		SqlMemInternalTxnMaxP99           int     `json:"sql.mem.internal.txn.max-p99"`
		SqlMemInternalTxnMaxP999          int     `json:"sql.mem.internal.txn.max-p99.9"`
		SqlMemInternalTxnMaxP9999         int     `json:"sql.mem.internal.txn.max-p99.99"`
		SqlMemInternalTxnMaxP99999        int     `json:"sql.mem.internal.txn.max-p99.999"`
		SqlMiscCount                      int     `json:"sql.misc.count"`
		SqlQueryCount                     int     `json:"sql.query.count"`
		SqlSelectCount                    int     `json:"sql.select.count"`
		SqlServiceLatencyMax              int     `json:"sql.service.latency-max"`
		SqlServiceLatencyP50              int     `json:"sql.service.latency-p50"`
		SqlServiceLatencyP75              int     `json:"sql.service.latency-p75"`
		SqlServiceLatencyP90              int     `json:"sql.service.latency-p90"`
		SqlServiceLatencyP99              int     `json:"sql.service.latency-p99"`
		SqlServiceLatencyP999             int     `json:"sql.service.latency-p99.9"`
		SqlServiceLatencyP9999            int     `json:"sql.service.latency-p99.99"`
		SqlServiceLatencyP99999           int     `json:"sql.service.latency-p99.999"`
		SqlTxnAbortCount                  int     `json:"sql.txn.abort.count"`
		SqlTxnBeginCount                  int     `json:"sql.txn.begin.count"`
		SqlTxnCommitCount                 int     `json:"sql.txn.commit.count"`
		SqlTxnRollbackCount               int     `json:"sql.txn.rollback.count"`
		SqlUpdateCount                    int     `json:"sql.update.count"`
		SysCgoAllocbytes                  int     `json:"sys.cgo.allocbytes"`
		SysCgoTotalbytes                  int     `json:"sys.cgo.totalbytes"`
		SysCgocalls                       int     `json:"sys.cgocalls"`
		SysCPUSysNs                       int64   `json:"sys.cpu.sys.ns"`
		SysCPUSysPercent                  float64 `json:"sys.cpu.sys.percent"`
		SysCPUUserNs                      int64   `json:"sys.cpu.user.ns"`
		SysCPUUserPercent                 float64 `json:"sys.cpu.user.percent"`
		SysFdOpen                         int     `json:"sys.fd.open"`
		SysFdSoftlimit                    int     `json:"sys.fd.softlimit"`
		SysGcCount                        int     `json:"sys.gc.count"`
		SysGcPauseNs                      int     `json:"sys.gc.pause.ns"`
		SysGcPausePercent                 int     `json:"-"`
		SysGoAllocbytes                   int     `json:"sys.go.allocbytes"`
		SysGoTotalbytes                   int     `json:"sys.go.totalbytes"`
		SysGoroutines                     int     `json:"sys.goroutines"`
		SysRss                            int     `json:"sys.rss"`
		SysUptime                         int     `json:"sys.uptime"`
		TimeseriesWriteBytes              int     `json:"timeseries.write.bytes"`
		TimeseriesWriteErrors             int     `json:"timeseries.write.errors"`
		TimeseriesWriteSamples            int     `json:"timeseries.write.samples"`
		TxnAbandons                       int     `json:"txn.abandons"`
		TxnAborts                         int     `json:"txn.aborts"`
		TxnAutoretries                    int     `json:"txn.autoretries"`
		TxnCommits                        int     `json:"txn.commits"`
		TxnCommits1PC                     int     `json:"txn.commits1PC"`
		TxnDurationsMax                   int     `json:"txn.durations-max"`
		TxnDurationsP50                   int     `json:"txn.durations-p50"`
		TxnDurationsP75                   int     `json:"txn.durations-p75"`
		TxnDurationsP90                   int     `json:"txn.durations-p90"`
		TxnDurationsP99                   int     `json:"txn.durations-p99"`
		TxnDurationsP999                  int     `json:"txn.durations-p99.9"`
		TxnDurationsP9999                 int     `json:"txn.durations-p99.99"`
		TxnDurationsP99999                int     `json:"txn.durations-p99.999"`
		TxnRestartsMax                    int     `json:"txn.restarts-max"`
		TxnRestartsP50                    int     `json:"txn.restarts-p50"`
		TxnRestartsP75                    int     `json:"txn.restarts-p75"`
		TxnRestartsP90                    int     `json:"txn.restarts-p90"`
		TxnRestartsP99                    int     `json:"txn.restarts-p99"`
		TxnRestartsP999                   int     `json:"txn.restarts-p99.9"`
		TxnRestartsP9999                  int     `json:"txn.restarts-p99.99"`
		TxnRestartsP99999                 int     `json:"txn.restarts-p99.999"`
		TxnRestartsDeleterange            int     `json:"txn.restarts.deleterange"`
		TxnRestartsPossiblereplay         int     `json:"txn.restarts.possiblereplay"`
		TxnRestartsSerializable           int     `json:"txn.restarts.serializable"`
		TxnRestartsWritetooold            int     `json:"txn.restarts.writetooold"`
	} `json:"metrics"`
	StoreStatuses []struct {
		Desc struct {
			StoreID int `json:"storeId"`
			Attrs   struct {
				Attrs []interface{} `json:"attrs"`
			} `json:"attrs"`
			Node struct {
				NodeID  int `json:"nodeId"`
				Address struct {
					NetworkField string `json:"networkField"`
					AddressField string `json:"addressField"`
				} `json:"address"`
				Attrs struct {
					Attrs []interface{} `json:"attrs"`
				} `json:"attrs"`
				Locality struct {
					Tiers []interface{} `json:"tiers"`
				} `json:"locality"`
				ServerVersion struct {
					MajorVal int `json:"majorVal"`
					MinorVal int `json:"minorVal"`
					Patch    int `json:"patch"`
					Unstable int `json:"unstable"`
				} `json:"ServerVersion"`
			} `json:"node"`
			Capacity struct {
				Capacity        string  `json:"capacity"`
				Available       string  `json:"available"`
				Used            string  `json:"used"`
				LogicalBytes    string  `json:"logicalBytes"`
				RangeCount      int     `json:"rangeCount"`
				LeaseCount      int     `json:"leaseCount"`
				WritesPerSecond float64 `json:"writesPerSecond"`
				BytesPerReplica struct {
					P10  int `json:"p10"`
					P25  int `json:"p25"`
					P50  int `json:"p50"`
					P75  int `json:"p75"`
					P90  int `json:"p90"`
					PMax int `json:"pMax"`
				} `json:"bytesPerReplica"`
				WritesPerReplica struct {
					P10  int     `json:"p10"`
					P25  int     `json:"p25"`
					P50  int     `json:"p50"`
					P75  float64 `json:"p75"`
					P90  float64 `json:"p90"`
					PMax float64 `json:"pMax"`
				} `json:"writesPerReplica"`
			} `json:"capacity"`
		} `json:"desc"`
		Metrics struct {
			AddsstableApplications                 int     `json:"addsstable.applications"`
			AddsstableCopies                       int     `json:"addsstable.copies"`
			AddsstableProposals                    int     `json:"addsstable.proposals"`
			Capacity                               int64   `json:"capacity"`
			CapacityAvailable                      int64   `json:"capacity.available"`
			CapacityReserved                       int     `json:"capacity.reserved"`
			CapacityUsed                           int     `json:"capacity.used"`
			CompactorCompactingnanos               int     `json:"compactor.compactingnanos"`
			CompactorCompactionsFailure            int     `json:"compactor.compactions.failure"`
			CompactorCompactionsSuccess            int     `json:"compactor.compactions.success"`
			CompactorSuggestionbytesCompacted      int     `json:"compactor.suggestionbytes.compacted"`
			CompactorSuggestionbytesQueued         int     `json:"compactor.suggestionbytes.queued"`
			CompactorSuggestionbytesSkipped        int     `json:"compactor.suggestionbytes.skipped"`
			Gcbytesage                             int     `json:"gcbytesage"`
			Intentage                              int     `json:"intentage"`
			Intentbytes                            int     `json:"intentbytes"`
			Intentcount                            int     `json:"intentcount"`
			Keybytes                               int     `json:"keybytes"`
			Keycount                               int     `json:"keycount"`
			Lastupdatenanos                        int64   `json:"lastupdatenanos"`
			LeasesEpoch                            int     `json:"leases.epoch"`
			LeasesError                            int     `json:"leases.error"`
			LeasesExpiration                       int     `json:"leases.expiration"`
			LeasesSuccess                          int     `json:"leases.success"`
			LeasesTransfersError                   int     `json:"leases.transfers.error"`
			LeasesTransfersSuccess                 int     `json:"leases.transfers.success"`
			Livebytes                              int     `json:"livebytes"`
			Livecount                              int     `json:"livecount"`
			QueueConsistencyPending                int     `json:"queue.consistency.pending"`
			QueueConsistencyProcessFailure         int     `json:"queue.consistency.process.failure"`
			QueueConsistencyProcessSuccess         int     `json:"queue.consistency.process.success"`
			QueueConsistencyProcessingnanos        int     `json:"queue.consistency.processingnanos"`
			QueueGcInfoAbortspanconsidered         int     `json:"queue.gc.info.abortspanconsidered"`
			QueueGcInfoAbortspangcnum              int     `json:"queue.gc.info.abortspangcnum"`
			QueueGcInfoAbortspanscanned            int     `json:"queue.gc.info.abortspanscanned"`
			QueueGcInfoIntentsconsidered           int     `json:"queue.gc.info.intentsconsidered"`
			QueueGcInfoIntenttxns                  int     `json:"queue.gc.info.intenttxns"`
			QueueGcInfoNumkeysaffected             int     `json:"queue.gc.info.numkeysaffected"`
			QueueGcInfoPushtxn                     int     `json:"queue.gc.info.pushtxn"`
			QueueGcInfoResolvesuccess              int     `json:"queue.gc.info.resolvesuccess"`
			QueueGcInfoResolvetotal                int     `json:"queue.gc.info.resolvetotal"`
			QueueGcInfoTransactionspangcaborted    int     `json:"queue.gc.info.transactionspangcaborted"`
			QueueGcInfoTransactionspangccommitted  int     `json:"queue.gc.info.transactionspangccommitted"`
			QueueGcInfoTransactionspangcpending    int     `json:"queue.gc.info.transactionspangcpending"`
			QueueGcInfoTransactionspanscanned      int     `json:"queue.gc.info.transactionspanscanned"`
			QueueGcPending                         int     `json:"queue.gc.pending"`
			QueueGcProcessFailure                  int     `json:"queue.gc.process.failure"`
			QueueGcProcessSuccess                  int     `json:"queue.gc.process.success"`
			QueueGcProcessingnanos                 int     `json:"queue.gc.processingnanos"`
			QueueRaftlogPending                    int     `json:"queue.raftlog.pending"`
			QueueRaftlogProcessFailure             int     `json:"queue.raftlog.process.failure"`
			QueueRaftlogProcessSuccess             int     `json:"queue.raftlog.process.success"`
			QueueRaftlogProcessingnanos            int64   `json:"queue.raftlog.processingnanos"`
			QueueRaftsnapshotPending               int     `json:"queue.raftsnapshot.pending"`
			QueueRaftsnapshotProcessFailure        int     `json:"queue.raftsnapshot.process.failure"`
			QueueRaftsnapshotProcessSuccess        int     `json:"queue.raftsnapshot.process.success"`
			QueueRaftsnapshotProcessingnanos       int     `json:"queue.raftsnapshot.processingnanos"`
			QueueReplicagcPending                  int     `json:"queue.replicagc.pending"`
			QueueReplicagcProcessFailure           int     `json:"queue.replicagc.process.failure"`
			QueueReplicagcProcessSuccess           int     `json:"queue.replicagc.process.success"`
			QueueReplicagcProcessingnanos          int     `json:"queue.replicagc.processingnanos"`
			QueueReplicagcRemovereplica            int     `json:"queue.replicagc.removereplica"`
			QueueReplicateAddreplica               int     `json:"queue.replicate.addreplica"`
			QueueReplicatePending                  int     `json:"queue.replicate.pending"`
			QueueReplicateProcessFailure           int     `json:"queue.replicate.process.failure"`
			QueueReplicateProcessSuccess           int     `json:"queue.replicate.process.success"`
			QueueReplicateProcessingnanos          int     `json:"queue.replicate.processingnanos"`
			QueueReplicatePurgatory                int     `json:"queue.replicate.purgatory"`
			QueueReplicateRebalancereplica         int     `json:"queue.replicate.rebalancereplica"`
			QueueReplicateRemovedeadreplica        int     `json:"queue.replicate.removedeadreplica"`
			QueueReplicateRemovereplica            int     `json:"queue.replicate.removereplica"`
			QueueReplicateTransferlease            int     `json:"queue.replicate.transferlease"`
			QueueSplitPending                      int     `json:"queue.split.pending"`
			QueueSplitProcessFailure               int     `json:"queue.split.process.failure"`
			QueueSplitProcessSuccess               int     `json:"queue.split.process.success"`
			QueueSplitProcessingnanos              int     `json:"queue.split.processingnanos"`
			QueueTsmaintenancePending              int     `json:"queue.tsmaintenance.pending"`
			QueueTsmaintenanceProcessFailure       int     `json:"queue.tsmaintenance.process.failure"`
			QueueTsmaintenanceProcessSuccess       int     `json:"queue.tsmaintenance.process.success"`
			QueueTsmaintenanceProcessingnanos      int     `json:"queue.tsmaintenance.processingnanos"`
			RaftCommandsapplied                    int     `json:"raft.commandsapplied"`
			RaftEnqueuedPending                    int     `json:"raft.enqueued.pending"`
			RaftHeartbeatsPending                  int     `json:"raft.heartbeats.pending"`
			RaftProcessCommandcommitLatencyMax     int     `json:"raft.process.commandcommit.latency-max"`
			RaftProcessCommandcommitLatencyP50     int     `json:"raft.process.commandcommit.latency-p50"`
			RaftProcessCommandcommitLatencyP75     int     `json:"raft.process.commandcommit.latency-p75"`
			RaftProcessCommandcommitLatencyP90     int     `json:"raft.process.commandcommit.latency-p90"`
			RaftProcessCommandcommitLatencyP99     int     `json:"raft.process.commandcommit.latency-p99"`
			RaftProcessCommandcommitLatencyP999    int     `json:"raft.process.commandcommit.latency-p99.9"`
			RaftProcessCommandcommitLatencyP9999   int     `json:"raft.process.commandcommit.latency-p99.99"`
			RaftProcessCommandcommitLatencyP99999  int     `json:"raft.process.commandcommit.latency-p99.999"`
			RaftProcessLogcommitLatencyMax         int     `json:"raft.process.logcommit.latency-max"`
			RaftProcessLogcommitLatencyP50         int     `json:"raft.process.logcommit.latency-p50"`
			RaftProcessLogcommitLatencyP75         int     `json:"raft.process.logcommit.latency-p75"`
			RaftProcessLogcommitLatencyP90         int     `json:"raft.process.logcommit.latency-p90"`
			RaftProcessLogcommitLatencyP99         int     `json:"raft.process.logcommit.latency-p99"`
			RaftProcessLogcommitLatencyP999        int     `json:"raft.process.logcommit.latency-p99.9"`
			RaftProcessLogcommitLatencyP9999       int     `json:"raft.process.logcommit.latency-p99.99"`
			RaftProcessLogcommitLatencyP99999      int     `json:"raft.process.logcommit.latency-p99.999"`
			RaftProcessTickingnanos                int     `json:"raft.process.tickingnanos"`
			RaftProcessWorkingnanos                int64   `json:"raft.process.workingnanos"`
			RaftRcvdApp                            int     `json:"raft.rcvd.app"`
			RaftRcvdAppresp                        int     `json:"raft.rcvd.appresp"`
			RaftRcvdDropped                        int     `json:"raft.rcvd.dropped"`
			RaftRcvdHeartbeat                      int     `json:"raft.rcvd.heartbeat"`
			RaftRcvdHeartbeatresp                  int     `json:"raft.rcvd.heartbeatresp"`
			RaftRcvdPrevote                        int     `json:"raft.rcvd.prevote"`
			RaftRcvdPrevoteresp                    int     `json:"raft.rcvd.prevoteresp"`
			RaftRcvdProp                           int     `json:"raft.rcvd.prop"`
			RaftRcvdSnap                           int     `json:"raft.rcvd.snap"`
			RaftRcvdTimeoutnow                     int     `json:"raft.rcvd.timeoutnow"`
			RaftRcvdTransferleader                 int     `json:"raft.rcvd.transferleader"`
			RaftRcvdVote                           int     `json:"raft.rcvd.vote"`
			RaftRcvdVoteresp                       int     `json:"raft.rcvd.voteresp"`
			RaftTicks                              int     `json:"raft.ticks"`
			RaftlogBehind                          int     `json:"raftlog.behind"`
			RaftlogTruncated                       int     `json:"raftlog.truncated"`
			RangeAdds                              int     `json:"range.adds"`
			RangeRaftleadertransfers               int     `json:"range.raftleadertransfers"`
			RangeRemoves                           int     `json:"range.removes"`
			RangeSnapshotsGenerated                int     `json:"range.snapshots.generated"`
			RangeSnapshotsNormalApplied            int     `json:"range.snapshots.normal-applied"`
			RangeSnapshotsPreemptiveApplied        int     `json:"range.snapshots.preemptive-applied"`
			RangeSplits                            int     `json:"range.splits"`
			Ranges                                 int     `json:"ranges"`
			RangesUnavailable                      int     `json:"ranges.unavailable"`
			RangesUnderreplicated                  int     `json:"ranges.underreplicated"`
			RebalancingWritespersecond             float64 `json:"rebalancing.writespersecond"`
			Replicas                               int     `json:"replicas"`
			ReplicasCommandqueueCombinedqueuesize  int     `json:"replicas.commandqueue.combinedqueuesize"`
			ReplicasCommandqueueCombinedreadcount  int     `json:"replicas.commandqueue.combinedreadcount"`
			ReplicasCommandqueueCombinedwritecount int     `json:"replicas.commandqueue.combinedwritecount"`
			ReplicasCommandqueueMaxoverlaps        int     `json:"replicas.commandqueue.maxoverlaps"`
			ReplicasCommandqueueMaxreadcount       int     `json:"replicas.commandqueue.maxreadcount"`
			ReplicasCommandqueueMaxsize            int     `json:"replicas.commandqueue.maxsize"`
			ReplicasCommandqueueMaxtreesize        int     `json:"replicas.commandqueue.maxtreesize"`
			ReplicasCommandqueueMaxwritecount      int     `json:"replicas.commandqueue.maxwritecount"`
			ReplicasLeaders                        int     `json:"replicas.leaders"`
			ReplicasLeadersNotLeaseholders         int     `json:"replicas.leaders_not_leaseholders"`
			ReplicasLeaseholders                   int     `json:"replicas.leaseholders"`
			ReplicasQuiescent                      int     `json:"replicas.quiescent"`
			ReplicasReserved                       int     `json:"replicas.reserved"`
			RequestsBackpressureSplit              int     `json:"requests.backpressure.split"`
			RequestsSlowCommandqueue               int     `json:"requests.slow.commandqueue"`
			RequestsSlowLease                      int     `json:"requests.slow.lease"`
			RequestsSlowRaft                       int     `json:"requests.slow.raft"`
			RocksdbBlockCacheHits                  int     `json:"rocksdb.block.cache.hits"`
			RocksdbBlockCacheMisses                int     `json:"rocksdb.block.cache.misses"`
			RocksdbBlockCachePinnedUsage           int     `json:"rocksdb.block.cache.pinned-usage"`
			RocksdbBlockCacheUsage                 int     `json:"rocksdb.block.cache.usage"`
			RocksdbBloomFilterPrefixChecked        int     `json:"rocksdb.bloom.filter.prefix.checked"`
			RocksdbBloomFilterPrefixUseful         int     `json:"rocksdb.bloom.filter.prefix.useful"`
			RocksdbCompactions                     int     `json:"rocksdb.compactions"`
			RocksdbFlushes                         int     `json:"rocksdb.flushes"`
			RocksdbMemtableTotalSize               int     `json:"rocksdb.memtable.total-size"`
			RocksdbNumSstables                     int     `json:"rocksdb.num-sstables"`
			RocksdbReadAmplification               int     `json:"rocksdb.read-amplification"`
			RocksdbTableReadersMemEstimate         int     `json:"rocksdb.table-readers-mem-estimate"`
			Sysbytes                               int     `json:"sysbytes"`
			Syscount                               int     `json:"syscount"`
			Totalbytes                             int     `json:"totalbytes"`
			TscacheSklReadPages                    int     `json:"tscache.skl.read.pages"`
			TscacheSklReadRotations                int     `json:"tscache.skl.read.rotations"`
			TscacheSklWritePages                   int     `json:"tscache.skl.write.pages"`
			TscacheSklWriteRotations               int     `json:"tscache.skl.write.rotations"`
			Valbytes                               int     `json:"valbytes"`
			Valcount                               int     `json:"valcount"`
		} `json:"metrics"`
	} `json:"storeStatuses"`
	Args      []string `json:"args"`
	Env       []string `json:"env"`
	Latencies struct {
	} `json:"latencies"`
	Activity struct {
		Num1 struct {
			Incoming string `json:"incoming"`
			Outgoing string `json:"outgoing"`
			Latency  string `json:"latency"`
		} `json:"1"`
	} `json:"activity"`
}

func (c *Cockroachdb) Description() string {
	return "Inserts health status data from CockroachDB Cluster for demonstration purposes"
}

var sampleConfig = `
  ## URL of CockroachDB Health endpoint.
  # servers = ["http://localhost:8080"]
`

func (c *Cockroachdb) SampleConfig() string {
	return sampleConfig
}

// Reads light stats from all configured servers.
func (c *Cockroachdb) Gather(acc telegraf.Accumulator) error {
	// Default to a single node at localhost (default adminport)
	if len(c.Servers) == 0 {
		c.Servers = []string{"http://localhost:8080/"}
	}

	// Range over all servers, gathering stats. Returns early in case of any error.
	for _, u := range c.Servers {
		acc.AddError(c.gatherLights(u, acc))
	}

	return nil
}

// Gathers _status from a single node, adding them to the accumulator
func (c *Cockroachdb) gatherLights(s string, acc telegraf.Accumulator) error {
	// Parse the given URL to extract the server tag
	u, err := url.Parse(s)
	if err != nil {
		return fmt.Errorf("Unable to parse given server url %s: %s", s, err)
	}

	// Perform the GET request to the cockroachdb /_status/nodes/1 endpoint
	resp, err := c.client.Get(s + "/_status/nodes/1")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Successful responses will always return status code 200
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Cockroachdb responded with unexepcted status code %d", resp.StatusCode)
	}

	// Decode the response JSON into a new lights struct
	stats := &Cockroach{}
	if err := json.NewDecoder(resp.Body).Decode(stats); err != nil {
		return fmt.Errorf("unable to decode Cockroachdb response: %s", err)
	}

	// Build a map of tags
	tags := map[string]string{
		"addressField": stats.Desc.Address.AddressField,
		"server":       u.Host,
	}

	// Build a map of field values
	fields := map[string]interface{}{
		"sys.cpu.user.percent":     stats.Metrics.SysCPUUserPercent,
		"sys.cpu.sys.percent":      stats.Metrics.SysCPUSysPercent,
		"timeseries.write.bytes":   stats.Metrics.TimeseriesWriteBytes,
		"timeseries.write.samples": stats.Metrics.TimeseriesWriteSamples,
		"exec.latency-max":         stats.Metrics.ExecLatencyMax,
	}

	// Accumulate the tags and values
	acc.AddFields("cockroachdb", fields, tags)

	return nil
}

func init() {
	inputs.Add("cockroachdb", func() telegraf.Input {
		return NeCockroachdb()
	})
}
