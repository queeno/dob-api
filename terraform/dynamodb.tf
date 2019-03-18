resource "aws_dynamodb_table" "dob_api" {
  name           = "DateOfBirths"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "username"

  attribute {
    name = "username"
    type = "S"
  }
}
