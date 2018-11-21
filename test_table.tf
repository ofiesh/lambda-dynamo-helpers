variable "region" {}

provider "aws" {
  region = "${var.region}"
}

resource "aws_dynamodb_table" "test_table" {
  name = "${terraform.workspace}_test_lambda_dynamo_helpers"

  "attribute" {
    name = "key"
    type = "S"
  }
  hash_key = "key"
  read_capacity = 1
  write_capacity = 1
}

output "table_name" {
  value = "${aws_dynamodb_table.test_table.name}"
}

output "region" {
  value = "${var.region}"
}