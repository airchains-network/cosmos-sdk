

## Introduction

This repository is a cloned version of Cosmos SDK v0.50.3, customized for tracks and stations to streamline the zk rollup proof generation process. The goal is to enhance the efficiency of zk proof generation by sending balance and other necessary details during transactions, reducing the need for additional calls.

## Features

- **Optimized for zk Rollup Proof Generation**: Includes enhancements to send balance and other details during transactions.
- **Reduced Extra Calls**: Streamlines the process, reducing the need for additional balance and detail retrievals.

## Getting Started: Usage as a Go Module
To use this customized Cosmos SDK as a Go module in your project, add the following line to your go.mod file:

```sh
require github.com/cosmos/cosmos-sdk-v0.50.3
go mod tidy
```
OR 
```sh
go get github.com/cosmos/cosmos-sdk-v0.50.3
```
