package main

func ExampleCommandSettings() {
	ts := testServer(`{
    "persistent": {
      "cluster.routing.allocation.disable_allocation": "true"
    },
    "transient": {
      "cluster.routing.allocation.disable_replica_allocation": "true",
      "indices.recovery.max_bytes_per_sec" : "2gb",
      "indices.recovery.concurrent_streams" : "24",
      "cluster.routing.allocation.node_concurrent_recoveries" : "6"
    }
  }`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdSettings.Run(cluster, nil, nil)
	// Output:
	// +-------------------------------------------------------+-------+
	// | SETTING NAME                                          | VALUE |
	// +-------------------------------------------------------+-------+
	// | PERSISTENT SETTINGS                                   |       |
	// | cluster.routing.allocation.disable_allocation         | true  |
	// |                                                       |       |
	// | TRANSIENT SETTINGS                                    |       |
	// | cluster.routing.allocation.disable_replica_allocation | true  |
	// | indices.recovery.max_bytes_per_sec                    | 2gb   |
	// | indices.recovery.concurrent_streams                   | 24    |
	// | cluster.routing.allocation.node_concurrent_recoveries | 6     |
	// +-------------------------------------------------------+-------+
}
