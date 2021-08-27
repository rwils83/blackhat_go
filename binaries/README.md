# Binaries

## What? Why? 
Because, I was bored and wanted to mess around with go build.

## Lets say I want to use them?
Well, go ahead. They are built for windows because I was using windows at the time. Understand,
the echoserver will open port 20080, the tcpclient will proxy requests to port 80 to what ever
you specify in -send=. Normally I use that with -send=127.0.0.1:20080 along with the echoserver for playing around.
The getrequest binary makes a get request to https://www.google.com/robots.txt and writes the results to a file named
Results.txt in the directory it was run from. 

## Again, why?
Again, I was bored.

## Seriously, why?
Seriously, I get bored too. 