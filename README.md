This package contains helper functions to encapsulate boiler plate code for interacting with dynamo.
This is specifically for use with lambdas, where configuration variables exist in environment variables.

The region for dynamo to use is assumed to be present in the `REGION` environment variable.

Running the tests requires a `test_conf.json` that specify the `region` and `table name`. See `test_conf_exmaple.json` for an example.
Running `apply` will run `terraform` and generate the `test_conf.json`.