---
id: 20260216-protocol_stack
tags:
  - architecture
  - networking
created: 2026-02-16
status: draft
---
# Protocol Stack
It's an implementation of a computer networking **Protocol Suite**. It's divided into **Layers**, each one having a single specific purpose in mind making the design of the stack simple.

A **Protocol Suite** is the definition of the communication protocols that participate of the protocol stack.

The **Layers** of a protocol stack work together to enable communication between devices.

User applications interacts with the top layer which adds capabilities to lower layers.

A protocol stack is divided into three major sections:
- **Media**: How devices communicate with each other.
- **Transport**: How data is moved from one place to another.
- **Applications**: How applications communicate with each other.

An OS will have two well defined interfaces:
- **Media to Transport Layers**: Defines how transport protocols make use of particular media and hardware. It's associated with a device driver, e.g.: How **TCP/IP (Transport Layer)** would talk to the network interface controller.
- **Transport to Application layers**: Defines how application programs make use of the transport layers, e.g.: Defines how a web browser talk to **TCP/IP**.

A protocol stack makes use of a **Spanning Layer** which is a layer that can cut across multiple layers of the stack to provide services used by multiple layers. An example of a spanning layer is security mechanisms like encryption which can appear at transport layers ([TLS]()) and application layers ([HTTPS]()). It provides flexibility at different levels of the stack.
## Why do Protocol Stack exists?
Each protocol tries to solve a single problem. A protocol stack stacks protocols on top of each other so we can have more robust forms of communication. The simple form of each protocol give us the flexibility we need to create that stack.

Each protocol of a stack complements an important characteristic that's missing in a lower one. For example the TCP adds error correction and reliability on top of IP.
## Examples
### Example of protocol stack layers
- Application layer (HTTP)
- Transport Layer (TCP)
- Internet Layer (IP)
- Link Layer (Ethernet/Wi-fi)
- Physical Layer (IEEE 802.3ab)
### Counterexample
- Creating a conversion layer that converts from one protocol to another. In this example we would to create a conversion mechanism for each protocol combination.
## References
### Connects with
- [Protocols](protocols.md): What is a protocol and what is its purpose?
- [TCP/IP Model]: Original protocol stack organization
- [OSI Model]: Modern protocol stack organization
## Questions
- How does data flows from one protocol to another? And how is this made simple so conversions aren't needed?
## TL;DR
A protocol stack is a suite of protocols grouped together. Each protocol adds capabilities on top of each other. Users interacts with the top most protocol of the stack.

Spanning layers are layers that can interact with more than two layers in the stack, they add flexibility into the design.
## Flashcards
- Q: What is a protocol stack?
- A: A protocol stack is a suite of protocols grouped together. Each protocol adds capabilities on top of each other.
- Q: What is a protocol suite?
- A: Is the definition of the communication protocols that participate of the protocol stack.
- Q: What is a spanning layer and what is their purpose?
- A: Is a layer that can interact with more than two layers in the stack, they add flexibility into the design. 
- Q: Why do protocol stacks exist?
- A: This modular design exists to make it easier to create more robust forms of communication by connecting one protocol into the other, without having to create translation layers for each combination of protocols.
- Q: What are the two interfaces an OS will usually have to interact with protocols?
- A:  Media to Transport Layers and Transport to Application layers.
- Q: What is the purpose of the media to transport layer interface of an OS?
- A: Define how transport protocols make use of particular media and hardware. 
- Q: What is the purpose of the transport to application layer interface of an OS?
- A: Define how application programs make use of the transport layers.