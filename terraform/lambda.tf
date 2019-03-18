locals {
  lambda_file_name = "dob-api.zip"
}

data "archive_file" "dob_api_zip" {
  type        = "zip"
  source_file = "../dob-api"
  output_path = "${local.lambda_file_name}"
}

resource "aws_lambda_function" "dob_api" {
  filename         = "${local.lambda_file_name}"
  function_name    = "${var.app_name}"
  role             = "${aws_iam_role.dob_api.arn}"
  handler          = "${var.app_name}"
  source_code_hash = "${data.archive_file.dob_api_zip.output_base64sha256}"
  runtime          = "go1.x"
  publish          = true
  timeout          = "30"
}
