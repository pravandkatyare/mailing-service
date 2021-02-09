# Mailing-service

[![Go Report Card](https://goreportcard.com/badge/github.com/pravandkatyare/go-grpc-examples)](https://goreportcard.com/report/github.com/itsksaurabh/go-grpc-examples)
[![MIT License](https://img.shields.io/github/license/pravandkatyare/go-grpc-examples?style=social)](./LICENSE)

This repository contains business logic, service endpoints and build configuration for the Mailing Service.

Mailing service acts as a wrapper for multiple mailing services, currently Mailjet's Api has been used
for sending mails and is open for integrating some other mailing services like Amazon SES, SendGrid, etc. 

# Service Endpoints

| Endpoint                     | Action    | Description                                                                |
|:-----------------------------|:----------|:---------------------------------------------------------------------------|
| /send/mail                   | POST      | Send email to recipients                                                   |


## Service Request/Response Models

### Request Object

```
{
    "client": "MJ",
    "mail": [{
        "des": {
            "to": [{ 
                "id": "test@test.com",
                "name": "Test"
            }],
            // optional
            "cc": [{ 
                "id": "test1@test.com",
                "name": "Test1"
            }],
            // optional
            "bcc": [{
                "id": "test2@gmail.com",
                "name": "Test2"
            }]  
        },
        "subject": "Your email subject",
        "textBody": "Your email body",
        "HTMLBody": "Your email's HTMLBody" (optional)
    }]
}
```
* Client code: MJ(MailJet), SES(Amazon SES), SG(SendGrid)

### Response Object

Following response can be expected:

| Scenarios                                  | Status | Response                                          |
|:-------------------------------------------|:-------|:--------------------------------------------------|
| E-mail sent to all the recipients          |  200   | "Mail sent to all recipient/s"                    |
| E-mail sending failed for some recipients  |  200   | Array of recipients to whom mail was not sent.    |
| Unauthorized access to endpoint            |  401   | "Unauthorized"                                    |
| Failing to sent mail due to random reason  |  500   | "Something went wrong"                            |



# Building

### Prerequisites

To build the project, the following tools are required:

* Make - a build tool.
* Go 1.12+ - Go language compiler.
* MailJet Free Developer's account for PUBLIC and PRIVATE API key to access MailJet's mailing service.

### Developer Setup

```bin
cd ~/go/ && \
mkdir -p src/github.com/ && \

git clone https://github.com/pravandkatyare/mailing-service.git ~/go/src/github.com/pravandkatyare/mailing-service

cd ~/go/src/github.com/pravandkatyare/movieManagement
make build 
```
```
make run
```

# Environment Variables

Environment variables need to be set either in the CI/CD system which are passed 
through to the deployment run-time or within the deployment configuration file. 
While running on local system variable are to be exported from terminal. 

Following section describes the variable that are needed to run application, they can 
either be set in `application.yml` file or exported as environment variables(recommended):

* `MAILJET_PUBLIC_API` -  provided by MailJet for accessing their service API
* `MAILJET_PRIVATE_API` -  provided by MailJet for accessing their service API


# Maintainers
[Pravand Katyare](https://www.linkedin.com/in/pravand-katyare/)


## License
[MIT](License) @Pravand Katyare
