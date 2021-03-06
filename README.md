# First project on golang

This is the first time I develop something on GoLang. It is a great language, with a 
strong enviroment and lots of libraries, from which I only saw when developing in Python.

Firstly, it was tough to know and adapt my way of doing encapsulation using Go. But it was not anything 
related with some language construct, but the way the Golang ecosystem allows modules and packages to
coordinate. For this it is essential to work with a great IDE (Integrated Development Enviroment), because
this can solve a lot of problems with the naming of functions, variables and methods and their related
scope inside the packages. After some struggle, it was easier to see where everything has to be. 
I used VSCode and the Go Extension (https://marketplace.visualstudio.com/items?itemName=golang.go)
and it was wonderful.

Tha language itself is strong and well designed. It remembers me the C language - in fact, there are a lot of direct
references like the fmt.Sprintf, strconv.Itoa, etc. Reading the main reference, which is the book "The Go Programming 
Language", written by Kernighan and Donovan (https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-dp-0134190440/dp/0134190440/ref=mt_other?_encoding=UTF8&me=&qid=1635718532), is a great experience. All the source codes are meaningful and representative.

GoLang brought back pointers, and came with JSON structures and slices of arrays and strings. It is one great way to represent JSON data as
a struct, and it is easy to marshall e unmarshall JSON data. But Go has no function named parameters (keywords), just like Python 
and Scala, and it has a different way to declares methods, showing a form of representation that uses an old OOP concept, which 
conceptualizes a class as a method receiver. Not so intuitive, and could bring some readability problems, because the methods 
aren't inside a class scope (like the classes and methods declaration in Java, Scala, C++, etc.).

Yet, Go has a great concurrency model, inspired on CSP (Communicating Sequential Processes), using goroutines and channels.
Used Actors in Scala, that put Go in a similar position, working on a messaging passing model.

After writing some code and get the main ideas, it was easy to get it and put everything working.

# Program Structure

## docker-compose.yml
- It creates 4 services (images/containers): one for an Account API backend, other for a PostgreSQL server, other for
  an Vault server, another for an account_app service (who runs all the tests)

## account / account_rest.go
- REST mappings to the endpoints Fetch, Delete and Create

## account / account.go
- JSON Models for the Account structure

## account / account_test.go
- Unit tests for the Account operations

## cmd / main.go
- Sample usage for this library

## http / rest_client.go
- HTTP REST client with operations GET, POST and DELETE

## http / rest_client_test.go
- HTTP REST client tests

## test_assertions / assertions.go
- Simple library with Unit testing routines 

## Dockerfile
- Creates an image for the Account Testing routine

## run_all_tests.sh
- Bash script that runs prepares and run 3 containers and run all the tests


# Form3 Take Home Exercise

Engineers at Form3 build highly available distributed systems in a microservices environment. Our take home test is designed to evaluate real world activities that are involved with this role. We recognise that this may not be as mentally challenging and may take longer to implement than some algorithmic tests that are often seen in interview exercises. Our approach however helps ensure that you will be working with a team of engineers with the necessary practical skills for the role (as well as a diverse range of technical wizardry). 

## Instructions
The goal of this exercise is to write a client library in Go to access our fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information on how to interact with the API. Please note that the fake account API does not require any authorisation or authentication.

A mapping of account attributes can be found in [models.go](./models.go). Can be used as a starting point, usage of the file is not required.

If you encounter any problems running the fake account API we would encourage you to do some debugging first,
before reaching out for help.

## Submission Guidance

### Shoulds

The finished solution **should:**
- Be written in Go.
- Use the `docker-compose.yaml` of this repository.
- Be a client library suitable for use in another software project.
- Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource.
- Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against the provided fake account API.
- Be simple and concise.
- Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests pass.

### Should Nots

The finished solution **should not:**
- Use a code generator to write the client library.
- Use (copy or otherwise) code from any third party without attribution to complete the exercise, as this will result in the test being rejected.
- Use a library for your client (e.g: go-resty). Libraries to support testing or types like UUID are fine.
- Implement client-side validation.
- Implement an authentication scheme.
- Implement support for the fields `data.attributes.private_identification`, `data.attributes.organisation_identification`
  and `data.relationships`, as they are omitted in the provided fake account API implementation.
- Have advanced features, however discussion of anything extra you'd expect a production client to contain would be useful in the documentation.
- Be a command line client or other type of program - the requirement is to write a client library.
- Implement the `List` operation.
> We give no credit for including any of the above in a submitted test, so please only focus on the "Shoulds" above.

## How to submit your exercise

- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, by copying all files you deem necessary for your submission
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License

Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
