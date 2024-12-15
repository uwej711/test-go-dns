Test DNS behavior
=================

This example program uses github.com/miekg/dns to make DNS requests. On my local machine this only works
with a tiny delay between opening the connection and sending the request due to a network filter of a vpn client
I have to use.

If you have trouble with DNS on lima (1.0.2) based virtual machines, you can try this program to check if you 
have a similar issue than I have.

Play with the timeout in line 21 in main.go and check the output.

Build and run
-------------

You need to have go 1.23 installed. Checkout the code and run 

```shell
go mod tidy
go run main.go
```
