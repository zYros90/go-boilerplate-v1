# Clean architecture by Robert Martin alias Uncle Bob:
Layers: Core - Adapters | User interfaces - Main

## Core
Holds pure business logic. All components have no external dependencies.

Core is devided into core components.

#### Core Components
These components separate the app into independent elements:

* Use cases: Only high-level policy and they orchestrate the other elements. They can be triggered by users or jobs

* Entities: The purest objects and business rules. They can be simple or complex data structures. They don't know anything about the persistence or delivery mechanism. 

* Ports: Interfaces used by the `Use cases`.

## Adapter
Implements a port interface used by the core. Adapter (low-level) is used by the use-case (high-level). Multiple adapters for the same port are possible. Extending of the adapter behavior by using design patterns like Composite or Decorator are possible.

## User Interface
This layer can have multiple delivery mechanism (REST API, gRPC, CLI, ...)

## Main
This layer wires everything up