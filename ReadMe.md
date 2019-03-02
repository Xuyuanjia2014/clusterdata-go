# Introduction

This project aims to analyse the cluster data in Alibaba (https://github.com/alibaba/clusterdata.git) with the pure golang I/O interfaces.

# Basic Goals

1. The calculation of the total read speed by using bytes Buffer: 30min/279GB
2. The calculation of the total read speed by using goroutine and Channel between threads: 650s/108GB (10% extra time with very low memory).

# Analysis Goals for CSV Batch instances, jobs and tasks

1. A proper format to read all the instances, tasks amd jobs.
2. Leftover resources for each job and  are they larger engouth?
3. How there plan resources compared with machines' usage?

# Analysis Goals for CSV online services

1. Instances' resource requirements and their cotainers usages.
2. Do they have interferences with batch intances?

# Analysis Goals for machine resources

1. The distribution of each machine resources uasage
2. The patterns of various jobs and applications

# Analysis Goals for integrated instances and containers

1. Can instances perform better under heterogeneous environments?
2. The affection of online services.