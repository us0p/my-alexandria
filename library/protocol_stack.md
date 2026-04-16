---
id: 20260216-protocol_stack
tags:
  - architecture
  - networking
created: 2026-02-16
status: draft
type: concept
---
## TL;DR
A protocol stack is a suite of protocols grouped together. Each protocol adds capabilities on top of each other. Users interacts with the top most protocol of the stack.

The topmost protocol generates data and the subsequent layers encapsulate that data with their own headers without touching the insides, this allow for fast communication as there's no need to convert formats.

Because of this encapsulation, you can replace protocols in the same layer without having to change other layers. But you can't swap layers among themselves as this breaks their directional dependency.

Spanning layers are layers that can interact with more than two layers in the stack, they add flexibility into the design.
# Protocol Stack
It's an implementation of a computer networking **Protocol Suite**. It's divided into **Layers**, each one having a single specific purpose in mind making the design of the stack simple.

A **Protocol Suite** is the definition of the communication protocols that participate of the protocol stack.

The **Layers** of a protocol stack work together to enable communication between devices.

## How data flows from one layer to another
User applications interacts with the top layer which adds capabilities to lower layers. Each layer adds its own small piece of information (header/trailer) but doesn't modify the actual data from the upper layer.

| Layer           | Responsibility                                                               | Protocol Data Unit (PDU) |
| --------------- | ---------------------------------------------------------------------------- | ------------------------ |
| **Application** | Create original data (e.g. HTTP request)                                     | **Data**                 |
| **Transport**   | Add transport headers (TCP/UDP) on top of **Data**                           | **Segment**              |
| **Network**     | Add ip header on top of **Segment**                                          | **Packet**               |
| **Data Link**   | Add frame header + trailer (MAC addresses, error check) on top of **Packet** | **Frame**                |
| **Physical**    | Converts **Frame** into bits/signals on wire, fiver or radio                 | **Bits**                 |

This wrapping process is **encapsulation** and it's what makes this design scalable. On the receiving end, each layer only reads its own header and ignores the rest.

> Lower layers treat upper-layer data as opaque payload, so they don't care about format.

Since each layer solves a specific problem, they don't need to understand each other's data.

This allows you to swap protocols on the same layer and the suite would still work without any changes in other layers.

Example:
```plaintext
-- Original Suite
HTTP
TCP
IP
Ethernet
Fiber

-- Replaces Ethernet by Wi-Fi
HTTP
TCP
IP
Wi-Fi
Fiber
```

Everything from IP above still working unchanged.
## Layer Directional Dependency
Layers abstract different scope of communications:
- **Physical**: raw signal transmission
- **Data Link**: local network delivery
- **Network**: routing between networks
- **Transport**: end-to-end host communication
- **Application**: user protocols

 Layers are only independent in implementation, this directional dependency exists because topmost layers depends on services provided by the layer below. 
```plaintext
Application
    ↓ needs
Transport
    ↓ needs
Network
    ↓ needs
Data Link
    ↓ needs
Physical

-- Original Suite
HTTP
TCP
IP
Ethernet
Fiber

-- Changing layer order breaks directional dependency
Wi-Fi
TCP
IP
HTTP
Fiber
```

## Spanning Layers
A protocol stack makes use of a **Spanning Layer** which is a layer that can interact with multiple layers **simultaneously** in the stack to provide services relevant to the whole stack. An example of a spanning layer is security mechanisms like encryption which can appear at transport layers ([TLS]()) and application layers ([HTTPS]()). It provides flexibility at different levels of the stack.

Spanning layers exists because clear layering isn't realistic, a system needs features that affect many layers simultaneously:
- Security
- Monitoring/Telemetry
- Network management
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
## Flashcards
- Q: What is a protocol stack?
- A: A protocol stack is a suite of protocols grouped together. Each protocol adds capabilities on top of each other.
- Q: What is a protocol suite?
- A: Is the definition of the communication protocols that participate of the protocol stack.
- Q: What is a spanning layer and what is their purpose?
- A: Is a layer that can interact with more than two layers in the stack, they add functionality required by more than one layer.
- Q: Why do protocol stacks exist?
- A: This modular design exists to make it easier to create more robust forms of communication by connecting one protocol into the other, without having to create translation layers for each combination of protocols.
- Q: How data flows from one layer to another?
- A: Each layer encapsulates the information received from the layer above with their own header without touching the insides.
- Q: What is the Protocol Data Unit of each layer?
- A: Data, Segment, Packet, Frame, Bits.
- Q: How can we manage specific protocols inside the stack? Can we replace them with other? Can we change their ordering? Why?
- A: You can change protocol in the same layer without affecting the overall system, but you can not swap layers as this would break their directional dependency.