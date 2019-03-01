# Introduction

This project aims to analyse the cluster data in Alibaba (https://github.com/alibaba/clusterdata.git) with the pure golang I/O interfaces.

# Basic Goals

1. The calculation of the total read speed by using bytes Buffer: 30min/279GB
2. The calculation of the total read speed by using goroutine and Channel between threads: 650s/108GB (10% extra time with very low memory).

# Analysis Goals
