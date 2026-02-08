# Use AWS App Runner to deploy Go web server

## Context and problem statement

Currently using Jekyll with AWS S3 to host static webpage, but want to move to Go mainly for learning & backend flexibility.

## Considered Options

- AWS App Runner (chosen for ability to run dynamic Go code)
- AWS Lambda (free tier option was major consideration point)
- S3 (limited as Go is dynamic not static)

## Status

Accepted.

## Decision Outcome

Chosen option: "AWS App Runner", because I am able to run a standard Go web server, using a dockerfile. Trade-off from going with Lambda is the $5-$7 monthly charge for 0.25 vCPU instance.
