# lib-bpmn-engine

## Build status

[![go test status](https://github.com/rhzs/lib-bpmn-engine/actions/workflows/go-test.yml/badge.svg)](https://github.com/rhzs/lib-bpmn-engine/actions/workflows/go-test.yml)
[![codecov](https://codecov.io/gh/rhzs/lib-bpmn-engine/branch/main/graph/badge.svg?token=J5J6SQ0TPJ)](https://codecov.io/gh/rhzs/lib-bpmn-engine)
[![Documentation Status](https://readthedocs.com/projects/rhzs-lib-bpmn-engine/badge/?version=latest)](https://rhzs-lib-bpmn-engine.readthedocs-hosted.com/en/latest/?badge=latest)

## Project status

* "beta" status, not recommended to use in production
* breaking API changes expected
* contributors welcome

## Documentation

Full documentation with examples: \
https://rhzs.github.io/lib-bpmn-engine/

GoDoc: \
https://pkg.go.dev/github.com/rhzs/lib-bpmn-engine/pkg/bpmn_engine


## Requirements

Go v1.18+

I'm supporting the latest and second-latest version of Go, similar to how Go itself handles releases. 

## BPMN Modelling

All these examples are build with [Camunda Modeler Community Edition](https://camunda.com/de/download/modeler/).
I would like to send a big "thank you", to Camunda for providing such tool.

## Implementation notes

### IDs (process definition, process instance, job, events, etc.)

This engine does use an implementation of [Twitter's Snowflake algorithm](https://en.wikipedia.org/wiki/Snowflake_ID)
which combines some advantages, like it's time based and can be sorted, and it's collision free to a very large extend.
So you can rely on larger IDs were generated later in time, and they will not collide with IDs,
generated on e.g. other nodes of your application in a multi-node installation.

The IDs are structured like this ...
```
+-----------------------------------------------------------+
| 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
+-----------------------------------------------------------+
```

The NodeID is generated out of a hash-function which reads all environment variables.
As a result, this approach allows 4096 unique IDs per node and per millisecond.

## Development of this library and contribution

For development hints and notes, please check [DEVELOPMENT.md](./DEVELOPMENT.md)

For information on contribution, please check [CONTRIBUTING.md](./CONTRIBUTING.md)
