# Cargo

Blockchain sync for the cult of Cardano

Cargo is a scalable event-based multi-input/multi-output transform system with
a focus on syncing the Cardano blockchain and related ecosystem.

Inspiration for this tool comes from [CDAP](https://github.com/cdapio/cdap)
and its simplified data pipelines as code, and by tools like
[Cardano DB sync](https://github.com/input-output-hk/cardano-db-sync) and
[Oura](https://github.com/txpipe/oura) within the Cardano ecosystem.

## Design goals

Our primary goal is to create an application which can be used for simple use
cases on a local development machine, yet be horizontally scalable at several
key points while remaining interoperable with the overall software ecosystem.

This interoperability comes from creating standard data formats for passing
events between the various layers so necessary data structure transformations
can be applied. For example, an Oura sink could be written for Cargo, which
would allow reusing Oura and its filtered data as an input source.

* Event based architecture
* DAG based workflow definitions
  - Sharable code
* Horizontally scalable
* Configurable inputs (sources)
  - Cardano blockchain (cardano-node)
  - External programs
  - Webhooks
* Configurable transforms (filters)
  - External programs
  - Can be chained together
* Built-in transform primitives
  - JOIN, SPLIT
  - Schema alter (internal)
* Configurable outputs (sinks)
  - AMQP
  - Kafka
  - Database (cardano-db-sync compatible)
  - Webhooks
* Configurable message bus
  - AMQP
  - Kafka
  - In-memory (single process only)
