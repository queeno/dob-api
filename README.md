# dob-api

[![codecov.io](http://codecov.io/github/queeno/dob-api/coverage.svg?branch=master)](http://codecov.io/github/queeno/dob-api?branch=master)

[![travis-ci.com](https://travis-ci.com/queeno/dob-api.svg?branch=master)](https://travis-ci.com/queeno/dob-api.svg?branch=master)

A simple, scalable, highly available RESTful API written in Go, never to forget again your friends' birthdays.

The application supports two modes of operation:
- **Local mode:** An HTTP server is span up is bound to the following local address: 0.0.0.0:8000
- **Serverless mode:** This mode introduces support for AWS Go 1.x Lambda functions.

The application is written to detect the appropriate mode of operation automatically.

## Build and run locally

Only follow this step if you wish to run the application in your local environment.

Please install [Go 1.12](https://golang.org/doc/install) on your system.

```shell
go get github.com/queeno/dob-api
```
will place the project in `$GOPATH/src/github.com/queeno/dob-api` and downloads
the associated dependencies. A statically-linked binary will be also automatically produced in `$GOPATH/bin`.

In order to run `dob-api`, just run the binary:

```shell
./dob-api
2019/03/19 20:02:34 Starting server on 0.0.0.0:8000
```

Then just curl:

```shell
curl -XPUT localhost:8000/hello/simon -d '{ "dateOfBirth": "2011-09-02"}'
```

And because today is the 19th March, you'll get the following message:

```shell
curl localhost:8000/hello/simon
{"message":"Hello, simon! Your birthday is in 167 day(s)"}
```

In order to shutdown the server, please just send a SIGINT to the process:

```shell
kill -INT 76364
```

or just press CTRL+C.

## Deploy in AWS

Please make sure that the following packages are installed:
- curl
- jq
- openssl
- terraform

The terraform directory contains the infrastructure configuration code.
Make sure the AWS credentials are injected in your environment before running it.

The `run-terraform.sh` script makes sure the latest dob-api release is always deployed with terraform.

A new release is published automatically within GitHub releases every time the
master branch is updated. A `.travis.yml` file is included in the repo
to achieve this task.

In order to create infrastructure and deploy the application simply run:

`./run-terraform.sh apply`

Terraform will output the API Gateway endpoint to query (ie. https://kkxnq4br7k.execute-api.eu-west-2.amazonaws.com/live)

```shell
curl -XPUT https://kkxnq4br7k.execute-api.eu-west-2.amazonaws.com/live/hello/simon -d '{ "dateOfBirth": "2010-02-01" }'
```

sets Simon's date of birth to 1st February 2010.

Today's the 19th March, hence the following request returns:

```shell
curl https://kkxnq4br7k.execute-api.eu-west-2.amazonaws.com/live/hello/simon
{"message":"Hello, simon! Your birthday is in 319 day(s)"}
```

### Cleaning up
Please don't forget to run

```shell
./run-terraform.sh destroy
```

in order to clean up!

## Infrastructure architectural overview

<p align="center">
  <img src="img/infrastructure_diagram.png?raw=true" alt="App architectural overview"/>
</p>

The above diagram shows how *dob-api* can be deployed within AWS to
leverage the public cloud power in order to maximise
resource elasticity and scalability as well as optimising the running costs.

An API gateway is provisioned which exposes a frontend API on an arbitrary Amazon URL.
This serves the two endpoints:
- PUT `/hello/<username>`
- GET `/hello/<username>`

The API gateway is configured to automatically trigger the associated lambda function
running the app. This interacts with DynamoDB, the AWS no-SQL database service, where
the data resides.

The application includes infrastructure deployment scripts which implement this
architecture in AWS.

Scalability limits depend entirely on the AWS configuration and commercials.
Technically speaking, the application can be scaled infinitely.

## Application architectural overview

<p align="center">
  <img src="img/app_diagram.png?raw=true" alt="App architectural overview"/>
</p>

The application has been written in Go 1.12 following a bottom-up approach.
The above diagram shows the logical building blocks, or in Go terms, *packages*.
Each package contains one or more classes, which implement its functionality
and least one public interface, which allows these classes to be plugged
interchangeably to supporting objects.

The *db* package provides DB integration functionality. This is described by
the following interface:

```go
type Database interface {
  Open() error
  Close()
  Put(string, string) error
  Get(string) (string, error)
}
```

The *db* package includes two implementation of this interface:
- **DynamoDB**, which provides AWS DynamoDB support
- **Bolt**, which provides [Bolt](https://github.com/boltdb/bolt) support.
Bolt is a lightweight, minimal and simple database written in go. Data is stored
in a local file.

The *app* package contains the main functionality of the project. In other words,
it contains the functions that elaborate response for a given query.
There are two classes inside this package:

- **App**, a public class where the core logic resides.
- **Validator**, a private class, used by App to validate the input data.

The app class is described by the following interface:

```go
type MyApp interface {
  UpdateUsername(string, string) (error)
  GetDateOfBirth(string) (string, error)
}
```

The *api* package contains the **Api** class, a public API implementation
for local runs.

The *lambda* package contains the **Lambda** class, which includes an AWS
lambda handler to serve API Gateway requests.

The *main* package contains the logic to automatically detect a lambda run
vs a local run and trigger the relevant logic.

### Testing considerations

Each class is unit tested using the [Go testify](https://godoc.org/github.com/stretchr/testify/suite) suite. Integration tests are also provided within the package.

Where appropriate, mock objects are created to isolate the testing scope to the single
class.

#### Limitations

- The integration tests only cover local running components (AWS infrastructure is not in scope for now).
- There is missing logic to replicate the count days functionality for testing
reasons.
