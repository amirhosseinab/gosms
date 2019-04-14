## GO-SMS-IR

[![Build Status](https://travis-ci.org/amirhosseinab/go-sms-ir.svg?branch=master)](https://travis-ci.org/amirhosseinab/go-sms-ir)
[![codecov](https://codecov.io/gh/amirhosseinab/go-sms-ir/branch/master/graph/badge.svg)](https://codecov.io/gh/amirhosseinab/go-sms-ir)
[![Go Report Card](https://goreportcard.com/badge/github.com/amirhosseinab/go-sms-ir)](https://goreportcard.com/report/github.com/amirhosseinab/go-sms-ir)
[![GoDoc](https://godoc.org/github.com/amirhosseinab/go-sms-ir/sms?status.svg)](https://godoc.org/github.com/amirhosseinab/go-sms-ir/sms)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Golang client library over the Restful APIs of [www.sms.ir](https://www.sms.ir).
I changed some of the APIs of this service to be more intuitive. 

In the table below, I indicate which modules and methods of the SMS service implemented in this library.

Modules| APIs
:---:|---
UserInfo| <ul><li>[x] `GetToken`</li><li>[x] `Credit`</li><li>[ ] `SMSLine`</li></ul>
Send-Receive|<ul><li>[ ] `MessageSend`</li><li>[ ] `MessageSend (ReportByDate)`</li><li>[ ] `MessageSend (ReportById)`</li><li>[ ] `MessageSend (ReportByBachkey)`</li><li>[ ] `ReceiveMessage (ByLastID)`</li><li>[ ] `ReceiveMessage (ByDate)`</li></ul>
CustomerClub|<ul><li>[ ] `CustomerClubContact (AddContact)`</li><li>[ ] `CustomerClubContact (UpdateContact)`</li><li>[ ] `CustomerClubContact (GetCategories)`</li><li>[ ] `CustomerClubContact (GetContactsByCategory&ById)`</li><li>[ ] `CustomerClubContact (GetAllContactsByPageID)`</li><li>[ ] `CustomerClub (Send)`</li><li>[ ] `CustomerClub (AddContact&Send)`</li><li>[ ] `CustomerClub (SendToCategories)`</li><li>[ ] `CustomerClub (GetSendMessagesByPagination)`</li><li>[ ] `CustomerClub (GetSendMessagesByPaginationAndLastId)`</li><li>[ ] `CustomerClub (DeleteContactCustomerClub)`</li></ul>
Verification|<ul><li>[ ] `VerificationCode`</li><li>[ ] `UltraFastSend`</li></ul>

---
## Contribution
Any contribution is welcome to this repo. If you find an issue, please report that in the Issues section.
Beyond that, if you like to code in this repo, I appreciate that.
