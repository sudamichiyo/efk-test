# efk-test

**efk-stack: verifying efk stack with docker compose**

## Overview

Elasticsearch:

Elasticsearch is a distributed real-time search and analytics engine.
It enables fast and scalable storage of log data and other structured data, allowing for search, analysis, and visualization.
A distinctive feature of Elasticsearch is its high full-text search capability and real-time nature.
Elasticsearch uses a data structure called indexes to store data, enabling fast queries and aggregations.

Fluentd:

Fluentd is a data collection agent that collects log data and event data from various sources, processes and transfers it.
It supports numerous plugins, making it easy to collaborate with different data sources and output destinations.
It plays a role in collecting log data and transferring it to Elasticsearch.
Additionally, Fluentd can perform data processing and filtering, allowing collected data to be shaped and stored.

Kibana:

Kibana is a platform for visualizing and analyzing data stored in Elasticsearch.
Through features such as creating dashboards, executing search queries, graphing, and map displays, it allows gaining insights from data.
Kibana is a highly flexible and powerful visualization tool that supports data exploration and understanding.

## Run Locally

```shell
### Start the docker-compose.
$ docker compose up -d

### Connect to each container.
$ docker compose exec nginx bash
$ docker compose exec fluentd bash
$ docker compose exec elasticsearch bash
$ docker compose exec kibana bash

### Built container list.
$ docker compose ps | awk '{print $1,$2}'
NAME          IMAGE
elasticsearch efk-test-elasticsearch
fluentd       efk-test-fluentd
kibana        efk-test-kibana
nginx         efk-test-nginx

### Monitor real-time logs.
$ docker compose logs -f
```

## References

- https://docs.fluentd.org/configuration/config-file
