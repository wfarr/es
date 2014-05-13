package main

func ExampleRunHealth() {
	ts := testServer(`{
		"status": "tangerine",
		"cluster_name": "foobar",
		"timed_out" : false,
		"number_of_nodes" : 1,
		"number_of_data_nodes" : 1,
		"active_primary_shards" : 10,
		"active_shards" : 20,
		"relocating_shards" : 2,
		"initializing_shards" : 0,
		"unassigned_shards" : 0
		}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdHealth.Run(cluster, nil, nil)

	// Output:
	// +-----------------------+-----------+
	// | CLUSTER HEALTH        |           |
	// +-----------------------+-----------+
	// | Name                  | foobar    |
	// | Status                | tangerine |
	// | Timed Out             | false     |
	// | Number of Nodes       | 1         |
	// | Number of Data Nodes  | 1         |
	// | Active Primary Shards | 10        |
	// | Active Shards         | 20        |
	// | Relocating Shards     | 2         |
	// | Initializing Shards   | 0         |
	// | Unassigned Shards     | 0         |
	// +-----------------------+-----------+
}

func ExampleRunHealth_indices() {
	ts := testServer(`{
		"status": "tangerine",
		"cluster_name": "foobar",
		"timed_out" : false,
		"number_of_nodes" : 1,
		"number_of_data_nodes" : 1,
		"active_primary_shards" : 10,
		"active_shards" : 20,
		"relocating_shards" : 2,
		"initializing_shards" : 0,
		"unassigned_shards" : 0,
		"indices" : {
			"test1": {
				"status" : "green",
				"number_of_shards" : 1,
				"number_of_replicas" : 1,
				"active_primary_shards" : 1,
				"active_shards" : 2,
				"relocating_shards" : 0,
				"initializing_shards" : 0,
				"unassigned_shards" : 0
			},
			"test2": {
				"status" : "green",
				"number_of_shards" : 1,
				"number_of_replicas" : 1,
				"active_primary_shards" : 1,
				"active_shards" : 2,
				"relocating_shards" : 0,
				"initializing_shards" : 0,
				"unassigned_shards" : 0
			}
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdHealth.Run(cluster, nil, []string{"index"})
	// Output:
	// +-------+--------+--------+----------+-------------------+---------------+------------+--------------+------------+
	// | INDEX | STATUS | SHARDS | REPLICAS | ACT. PRIM. SHARDS | ACTIVE SHARDS | RELOCATING | INITIALIZING | UNASSIGNED |
	// +-------+--------+--------+----------+-------------------+---------------+------------+--------------+------------+
	// | test1 | green  | 1      | 1        | 1                 | 2             | 0          | 0            | 0          |
	// | test2 | green  | 1      | 1        | 1                 | 2             | 0          | 0            | 0          |
	// +-------+--------+--------+----------+-------------------+---------------+------------+--------------+------------+
}
