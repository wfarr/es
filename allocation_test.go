package main

func ExampleCommand_Run_runAllocation_one_oh() {
	ts := testServer(`{
		"persistent": {
			"cluster.routing.allocation.enable": "all"
		},
		"transient": {
			"cluster.routing.allocation.enable": "new_primaries"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, nil)
	// Output:
	// +--------------+-----------------------------------+---------------+
	// | SETTING TYPE | SETTING NAME                      | VALUE         |
	// +--------------+-----------------------------------+---------------+
	// | persistent   | cluster.routing.allocation.enable | all           |
	// | transient    | cluster.routing.allocation.enable | new_primaries |
	// +--------------+-----------------------------------+---------------+
}

func ExampleCommand_Run_runAllocation_oh_ninety() {
	ts := testServer(`{
		"persistent": {
			"cluster.routing.allocation.disable_allocation": true
		},
		"transient": {
			"cluster.routing.allocation.disable_replica_allocation": true,
			"indices.recovery.max_bytes_per_sec" : "2gb",
			"indices.recovery.concurrent_streams" : "24",
			"cluster.routing.allocation.node_concurrent_recoveries" : "6"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, nil)
	// Output:
	// +--------------+-------------------------------------------------------+-------+
	// | SETTING TYPE | SETTING NAME                                          | VALUE |
	// +--------------+-------------------------------------------------------+-------+
	// | persistent   | cluster.routing.allocation.disable_allocation         | false |
	// | transient    | cluster.routing.allocation.disable_replica_allocation | true  |
	// +--------------+-------------------------------------------------------+-------+
}
