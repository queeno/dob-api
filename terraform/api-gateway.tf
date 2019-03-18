resource "aws_api_gateway_rest_api" "dob_api" {
  name        = "${var.app_name}"
  description = "RESTful API for my DateOfBirth App"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_deployment" "dob_api" {
  rest_api_id = "${aws_api_gateway_rest_api.dob_api.id}"
  stage_name  = "live"

  depends_on = [
    "aws_api_gateway_resource.hello",
    "aws_api_gateway_resource.username",
    "aws_api_gateway_method.get_username",
    "aws_api_gateway_integration.get_username",
    "aws_api_gateway_method.put_username",
    "aws_api_gateway_integration.put_username",
  ]
}

resource "aws_api_gateway_resource" "hello" {
  rest_api_id = "${aws_api_gateway_rest_api.dob_api.id}"
  parent_id   = "${aws_api_gateway_rest_api.dob_api.root_resource_id}"
  path_part   = "hello"
}

resource "aws_api_gateway_resource" "username" {
  rest_api_id = "${aws_api_gateway_rest_api.dob_api.id}"
  parent_id   = "${aws_api_gateway_resource.hello.id}"
  path_part   = "{username}"
}

resource "aws_api_gateway_method" "get_username" {
  rest_api_id   = "${aws_api_gateway_rest_api.dob_api.id}"
  resource_id   = "${aws_api_gateway_resource.username.id}"
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "get_username" {
  rest_api_id             = "${aws_api_gateway_rest_api.dob_api.id}"
  resource_id             = "${aws_api_gateway_resource.username.id}"
  http_method             = "${aws_api_gateway_method.get_username.http_method}"
  type                    = "AWS_PROXY"
  uri                     = "${aws_lambda_function.dob_api.invoke_arn}"
  integration_http_method = "POST"
}

resource "aws_api_gateway_method" "put_username" {
  rest_api_id   = "${aws_api_gateway_rest_api.dob_api.id}"
  resource_id   = "${aws_api_gateway_resource.username.id}"
  http_method   = "PUT"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "put_username" {
  rest_api_id             = "${aws_api_gateway_rest_api.dob_api.id}"
  resource_id             = "${aws_api_gateway_resource.username.id}"
  http_method             = "${aws_api_gateway_method.put_username.http_method}"
  type                    = "AWS_PROXY"
  uri                     = "${aws_lambda_function.dob_api.invoke_arn}"
  integration_http_method = "POST"
}

resource "aws_lambda_permission" "dob_api_get_username" {
  statement_id  = "AllowExecutionFromAPIGateway-${aws_api_gateway_resource.username.id}-${aws_api_gateway_method.get_username.http_method}"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.dob_api.function_name}"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:${var.region}:${data.aws_caller_identity.current.account_id}:${aws_api_gateway_rest_api.dob_api.id}/*/${aws_api_gateway_method.get_username.http_method}${aws_api_gateway_resource.username.path}"
}

resource "aws_lambda_permission" "dob_api_put_username" {
  statement_id  = "AllowExecutionFromAPIGateway-${aws_api_gateway_resource.username.id}-${aws_api_gateway_method.put_username.http_method}"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.dob_api.function_name}"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:${var.region}:${data.aws_caller_identity.current.account_id}:${aws_api_gateway_rest_api.dob_api.id}/*/${aws_api_gateway_method.put_username.http_method}${aws_api_gateway_resource.username.path}"
}
