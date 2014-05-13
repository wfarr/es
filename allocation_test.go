package main

func ExampleCommandAllocation_enable_one_oh() {
	ts := testServer(`{
		"version": {
			"number": "1.1.1"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, []string{"enable"})
	// Output:
	// Successfully set cluster.routing.allocation.enable=all
}

func ExampleCommandAllocation_disable_one_oh() {
	ts := testServer(`{
		"version": {
			"number": "1.1.1"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, []string{"disable"})
	// Output:
	// Successfully set cluster.routing.allocation.enable=none
}

func ExampleCommandAllocation_all_one_oh() {
	ts := testServer(`{
		"version": {
			"number": "1.1.1"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, []string{"all"})
	// Output:
	// Successfully set cluster.routing.allocation.enable=all
}

func ExampleCommandAllocation_primaries_one_oh() {
	ts := testServer(`{
		"version": {
			"number": "1.1.1"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, []string{"primaries"})
	// Output:
	// Successfully set cluster.routing.allocation.enable=primaries
}

func ExampleCommandAllocation_new_primaries_one_oh() {
	ts := testServer(`{
		"version": {
			"number": "1.1.1"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, []string{"new_primaries"})
	// Output:
	// Successfully set cluster.routing.allocation.enable=new_primaries
}

func ExampleCommandAllocation_none_one_oh() {
	ts := testServer(`{
		"version": {
			"number": "1.1.1"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, []string{"none"})
	// Output:
	// Successfully set cluster.routing.allocation.enable=none
}

func ExampleCommandAllocation_enable_oh_ninety() {
	ts := testServer(`{
		"version": {
			"number": "0.90.5"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, []string{"enable"})
	// Output:
	// Successfully set cluster.routing.allocation.disable_allocation=false
}

func ExampleCommandAllocation_disable_oh_ninety() {
	ts := testServer(`{
		"version": {
			"number": "0.90.5"
		}
	}`)

	defer ts.Close()
	cluster := makeClusterForTestServer(ts)

	cmdAllocation.Run(cluster, nil, []string{"disable"})
	// Output:
	// Successfully set cluster.routing.allocation.disable_allocation=true
}
