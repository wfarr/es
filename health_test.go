package main

func ExampleHealthCommand() {
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
  cluster := &Cluster{URL: ts.URL}

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
