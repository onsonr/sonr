<p align="center">
  <a href="" rel="noopener">
 <img width=600px src="./docs/static/cover.jpg" alt="Project logo"></a>
</p>

<h1 align="center"><bold>Sonr</bold></h1>

<div align="center">

[![CodeFactor](https://www.codefactor.io/repository/github/sonr-io/sonr/badge)](https://www.codefactor.io/repository/github/sonr-io/sonr)
  [![Status](https://img.shields.io/badge/status-active-success.svg)](https://sonr.io)
  [![Go Reference](https://pkg.go.dev/badge/github.com/sonr-io/sonr.svg)](https://pkg.go.dev/github.com/sonr-io/sonr)
  [![Go Report Card](https://goreportcard.com/badge/github.com/sonr-io/sonr)](https://goreportcard.com/report/github.com/sonr-io/sonr)
  [![GitHub Issues](https://img.shields.io/github/issues/sonr-io/sonr.svg)](https://github.com/sonr-io/sonr/issues)
  [![GitHub Pull Requests](https://img.shields.io/github/issues-pr/sonr-io/sonr.svg)](https://github.com/sonr-io/sonr/pulls)
  [![License](https://img.shields.io/badge/license-GPLv3-blue.svg)](/LICENSE)

</div>

---

<p align="center"> Build <strong>privacy-preserving</strong>, <strong>user-centric applications</strong>, on a robust, rapid-scaling platform designed for intereoperability, and total digital autonomy.
    <br>
</p>

## 📝 Table of Contents
- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [Contributing](./docs/guides/CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>
Sonr aims to be the most immersive and powerful DWeb experience for both Users and Developers alike. We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that’s composable and interoperable with all existing protocols.

For a more in-depth technical look into the Sonr ecosystem please refer to the [Architecture Decision Records](./docs/architecture/GUIDE.md).


## 🏁 Getting Started <a name = "getting_started"></a>
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#-deployment-) for notes on how to deploy the project on a live system.

### Prerequisites
What things you need to install the software and how to install them.
- [Go](https://golang.org/doc/install)
- [Ignite CLI](https://github.com/ignite/cli)
- [Protocol Buffers](https://grpc.io)
- [GRPCLI](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md)

### Installing
A step by step series of examples that tell you how to get a development env running.

Get the chain running
```bash
ignite chain serve
```

Test it with `grpc_cli`
```bash
grpc_cli ls 137.184.190.146:9090 sonrio.sonr.bucket.Query
```
Outputs: 
```
Params
WhereIs
WhereIsByCreator
WhereIsAll
```

## 🔧 Running the tests <a name = "tests"></a>
Explain how to run the automated tests for this system.

### Break down into end to end tests

Run a motor test
```bash
cd pkg/motor
go test -run CreateAccount
```

Run <b>ALL</b> motor tests
```bash
cd pkg/motor
go test
```

Run the simulation tests
```bash
ignite chain simulate
```

### And coding style tests
We run `goplz` on our source code. Generally, you can get that as part of the [Go Extension in VS Code](https://marketplace.visualstudio.com/items?itemName=golang.Go)

## 🎈 Usage <a name="usage"></a>
TODO: Give quick overview of how to use the system, with link to [Structure](./docs/guides/STRUCTURE.md)
- [Running the Makefile](./docs/guides/USAGE.md#running-the-makefile)
- [Interacting with the Motor](./docs/guides/USAGE.md#interacting-with-the-motor)
- [Submitting a Proposal](./docs/guides//USAGE.md#submitting-a-proposal)

## 🚀 Deployment <a name = "deployment"></a>
TODO: Insert quick motor v. highway explanation with link to [docs](https://docs.sonr.io).
- [Running the Highway](./docs/guides/DEPLOYMENT.md#running-the-highway-node)
- [Building the Motor](./docs/guides/DEPLOYMENT.md#binding-the-motor-library)

## ⛏️ Built Using <a name = "built_using"></a>
- [Libp2p](https://github.com/libp2p/libp2p) - Networking layer
- [Cosmos](https://github.com/cosmos-sdk/cosmos) - Blockchain Framework
- [IPFS](https://github.com/ipfs/ipfs) - Storage Module
- [HNS](https://handshake.org/) - Decentralized DNS

## ✍️ Authors <a name = "authors"></a>

- [Prad Nukala](https://github.com/prnk28)
- [Nick Tindle](https://github.com/ntindle)
- [Josh Long](https://github.com/joshLong145)
- [Brayden Cloud](https://github.com/mcjcloud)
- [Ian Perez](https://github.com/brokecollegekidwithaclothingobsession)

See also the list of [contributors](https://github.com/sonr-io/sonr/contributors) who participated in this project.

## 🎉 Acknowledgements <a name = "acknowledgement"></a>
- [W3C](https://www.w3.org/)
- [Protocol Labs](https://protocol.ai/)
- [Ignite](https://ignite.com/)
- [AE Studio](https://ae.studio/)


Partners, collaborators, compliance, or just pure appreciation!
