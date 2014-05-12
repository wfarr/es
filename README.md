# ES

## Usage

```
$ ./es help
Usage: es [command] [options] [arguments]


Commands:

		health    display the health of the cluster
		allocation  control cluster allocation settings

Run 'es help [command]' for details.


$ ./es help health
Usage: es health

	Displays general cluster health information.


$ ./es help allocation
Usage: es allocation [<setting>]

	Manage cluster allocation settings.

	For Elasticsearch clusters running 0.90.x, valid options are:
		* enable
		* disable

	For Elasticsearch clusters running 1.x, valid options are:
		* all
		* primaries
		* new_primaries
		* none

	If no settings is given, display the current cluster allocation settings.
  ```
