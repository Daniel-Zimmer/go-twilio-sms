# go-twilio-sms
(Very) Simple Twilio Service for Sending SMSs

## Overview

**go-twilio-sms** simply creates a service that makes http request to **twilio** to send SMSs.

When the request is unsuccessful, **SendSms()** returns an error containing the message returned by the request.

## Example
```go
package main

import (
	twilio "go-twilio-sms"
	"log"
)

func main() {
  accountSid := "XXXXXX"
  authToken  := "XXXXXX"

  service := twilio.NewService(accountSid, authToken)

  from := "+12021234567"
  to   := "+12027654321"

  err := service.SendSms(from, to, "Your Message Here!")
  if err != nil {	log.Print(err) }
}
```
