data "aws_iam_policy_document" "lambda_trust" {
  statement {
    actions = [
      "sts:AssumeRole",
    ]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lambda_to_dynamodb" {
  statement {
    actions = [
      "dynamodb:List*",
      "dynamodb:DescribeReservedCapacity*",
      "dynamodb:DescribeLimits",
      "dynamodb:DescribeTimeToLive",
    ]

    resources = [
      "*",
    ]
  }

  statement {
    actions = [
      "dynamodb:BatchGet*",
      "dynamodb:DescribeStream",
      "dynamodb:DescribeTable",
      "dynamodb:Get*",
      "dynamodb:Query",
      "dynamodb:Scan",
      "dynamodb:BatchWrite*",
      "dynamodb:CreateTable",
      "dynamodb:Delete*",
      "dynamodb:Update*",
      "dynamodb:PutItem",
    ]

    resources = [
      "*",
    ]
  }
}

data "aws_iam_policy_document" "lambda_to_cloudwatch_logs" {
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:DescribeLogStreams",
    ]

    resources = [
      "arn:aws:logs:*:*:*",
    ]
  }
}

resource "aws_iam_role" "dob_api" {
  name               = "${var.app_name}"
  path               = "/"
  assume_role_policy = "${data.aws_iam_policy_document.lambda_trust.json}"
}

resource "aws_iam_role_policy" "lambda_to_cloudwatch_logs" {
  name   = "lambda-to-cloudwatch-logs"
  role   = "${aws_iam_role.dob_api.name}"
  policy = "${data.aws_iam_policy_document.lambda_to_cloudwatch_logs.json}"
}

resource "aws_iam_role_policy" "lambda_to_dynamodb" {
  name   = "lambda-to-dynamodb"
  role   = "${aws_iam_role.dob_api.name}"
  policy = "${data.aws_iam_policy_document.lambda_to_dynamodb.json}"
}
