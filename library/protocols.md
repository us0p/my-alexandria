---
id: 260207-protocols
tags:
  - architecture
  - networking
created: 2026-02-07
status: draft
---
# Protocols
A protocol is a contract for communication that describes how computer systems and programs must behave.

All parts that want to communicate through that protocol must follow those set of rules.

The protocol doesn't only describes the format of the message but also how and when the messages are sent between the parties.
## Why do protocols exist?
Protocols exist because independent systems cannot assume shared assumptions about timing, format, or order of communication.

Different protocols exist to solve a different communication problems. E.g. TCP solves the unreliability of UDP by adding check-sums and flags to packets to ensure full delivery and correct order.

They can be stacked to achieve more robust forms of communication and can use very different rules and structures of data, like text, bytes and even electricity.
## Examples
### Protocol responsibilities
- Define data's expected format (JSON, text, bytes, etc).
- Define rules of when to send data, like when a request is received.
- Define rules of how to send data, e.g. TCP uses a three-way handshake to establish a connection.
### Counterexample
- Send data without following a standard format.
- Parties sending data without any order or means to organize data transmission.
## References
### Connects with
- [IP](): How computers find each other in a wired network.
- [UDP](): Best-effort datagram transport protocol.
- [TCP](): Reliable datagram transport protocol
- [HTTP](): How Client and Server application can communicate.
- [Protocol Stack](protocol_stack.md): How protocols stack on top of each other and why.
## TL;DR
Protocol is a communication contract with a set of well defined rules that determines how computers in a network must behave in order to communicate successfully.
## Flashcards
- Q: What is a protocol?
- A: A protocol is a contract for communication that describes how computer systems and programs must behave.
- Q: Why protocols exists?
- A: Because independent systems cannot assume shared assumptions about timing, format, or order of communication. 
- Q: What defines a protocol besides message format?
- A: A protocol is also responsible to how and when message must be exchanged.