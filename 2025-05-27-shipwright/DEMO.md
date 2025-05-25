# Live demo

## Overview

This live demo showcases the 3 main components of **shipwright**, `Build`, `BuildStrategy` and `BuildRun`.

It takes a source code written in go, leverage `ko` as builder and runs shipwright on a local `KinD` cluster.

## Pre-requisite

Ensure that a Kubernetes is up-and-running with `Shipwright` installed.

If not, please refer to [kind]() and [shipwright]() official documentation.

## Demo

*The live demo is related with the presentation slide available [here]()*

### Context

Following the personas and team topology introduced during the talk, we have the following responsibilities:

| Persona | Role | Topology team | Responsibilities |
|---------|------|---------------|------------------|
| Taro | Platform engineer | Platform team | Manage the organisation strategies and standardise it |
| Blackberry | DevOps engineer | Enabling team | Facilitate the connection between the Build-as-a-Service from the platform team and the usage of it by the stream-aligned team |
| Coco | Lead developer | Stream-align team | Indirect end-user of the Build-as-a-Service, focus on the business and features development |

The organisation the personas are working for only develop microservices in Go and only Go. All of them are cloud-native apps and are design to be Kubernates friendly.


### Step-by-step

```bash
k shp build create cnb-build --source-git-url="git@github.com:rxinui/talks.git" \
   --source-context-dir="2025-05-27-shipwright/cmd" \
   --source-git-clone-secret="ssh-auth" \
   --output-image="ghcr.io/rxinui/cnb:latest" \
   --strategy-kind="ClusterBuildStrategy" \
   --strategy-name="ko" \
   --output-image-push-secret="ghcr"

```











### Platform engineer put in place the organisation strategy

*Note*

1. Create a simple `ko` strategy that takes a main package
2. Should use at least one external lib
3. Should take a LDFLAGS to inject a `BuildStrategy` parameter such as `BuildAuthor`
4. Should generate and display a unique build id to display along build author

### Lead developer just focus on the code

Refers to `BuildStrategy`

### DevOps engineer link the code with the strategy

1. Take the code from github
2. Write a trigger on github push
3. Enable vuln scan
4. Push image output to ghcr



