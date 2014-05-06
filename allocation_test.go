package main

func ExampleCmdAllocation_one_oh() {
	ts := testServer(`{
		"persistent": {
			"cluster.routing.allocation.enable": "all"
		},
		"transient": {
			"cluster.routing.allocation.enable": "new_primaries"
		}
	}`)

	defer ts.Close()
	cluster := &Cluster{&Client{URL: ts.URL}}

	cmdAllocation.Run(cluster, nil, nil)

	// Output:
	// +--------------+-----------------------------------+---------------+
	// | SETTING TYPE | SETTING NAME                      | VALUE         |
	// +--------------+-----------------------------------+---------------+
	// | persistent   | cluster.routing.allocation.enable | all           |
	// | transient    | cluster.routing.allocation.enable | new_primaries |
	// +--------------+-----------------------------------+---------------+
}


func ExampleCmdAllocation_oh_ninety() {
	ts := testServer(`{
		"persistent": {
			"cluster.routing.allocation.disable_allocation": false
		},
		"transient": {
			"cluster.routing.allocation.disable_allocation": true
		}
	}`)

	defer ts.Close()
	cluster := &Cluster{&Client{URL: ts.URL}}

	cmdAllocation.Run(cluster, nil, nil)

	// Output:
	// +--------------+-----------------------------------------------+-------+
	// | SETTING TYPE | SETTING NAME                                  | VALUE |
	// +--------------+-----------------------------------------------+-------+
	// | persistent   | cluster.routing.allocation.disable_allocation | false |
	// | transient    | cluster.routing.allocation.disable_allocation | false |
	// +--------------+-----------------------------------------------+-------+
}
