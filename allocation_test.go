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
			"cluster.routing.allocation.disable_allocation": false
		},
		"transient": {
			"cluster.routing.allocation.disable_allocation": true
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, nil)
	// Output:
	// +--------------+-----------------------------------------------+-------+
	// | SETTING TYPE | SETTING NAME                                  | VALUE |
	// +--------------+-----------------------------------------------+-------+
	// | persistent   | cluster.routing.allocation.disable_allocation | false |
	// | transient    | cluster.routing.allocation.disable_allocation | false |
	// +--------------+-----------------------------------------------+-------+
}
