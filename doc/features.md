## Check UI health
1. check web server connections
2. timeout 5 sec, retry 3
3. ??todo: get response time
4. get status code


## Check SSL cert 
1. check valid ssl cert - establish a TLS connection with the website if succeed it means the website has a valid TLS certificate. To establish a TLS connection we can use the Go crypto/tls package. 
2. 

## Simple testing using Selenium

## prometheus metrics
- push gateway

## Docker file

## GitHub actions


bufio package - to take user input

URLs:
all public facing DNS, public API and looker

UI 
- simple DB
- manage users
- to fetch URL to this go


## Web app
- display web server health
- display web server certificate expiration

## monitoring
- send notification when web server is not responsive + prometheus monitoring
- send notification when web server's SSL expires in 3 months + prometheus monitoring + automatic SSL cert renewal using [ACM](https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html) - check our TF code for [this](https://www.tenable.com/policies/[type]/AC_AWS_0003)