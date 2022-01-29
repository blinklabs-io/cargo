# Cargo

Blockchain sync for the cult of Cardano

Cargo is a scalable event-based filter/transform system with a focus on syncing the Cardano blockchain.
It was inspired by tools like [Cardano DB sync](https://github.com/input-output-hk/cardano-db-sync) and
[Oura](https://github.com/txpipe/oura), but our design goals differ.

## Design goals

* scalable event-based architecture
    - able to run as a single process or scale horizontally across multiple hosts
* configurable inputs
    - Cardano blockchain
    - external programs
* configurable filter/transform steps
    - external programs
    - filter/transform step can generate events as input for other steps
* configurable outputs
    - AMQP
    - Kafka
    - webhooks
    - database
* pluggable message bus beween tiers
    - AMQP
    - Kafka
    - in-memory (for running as a single process only)
