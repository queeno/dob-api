output "api_url" {
  value = "${aws_api_gateway_deployment.dob_api.invoke_url}"
}
