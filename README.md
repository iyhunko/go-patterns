# Go Patterns

A collection of Go design patterns, concurrency patterns, and practical examples demonstrating best practices in Go programming.

## Overview

This repository contains working implementations of various design patterns and concurrency patterns in Go. Each pattern is contained in its own directory with a complete, runnable example.

## 📚 Contents

### Design Patterns

- **[builder-pattern](./builder-pattern/)** - Implements the Builder pattern for constructing complex objects step by step
- **[decorator-pattern](./decorator-pattern/)** - Demonstrates the Decorator pattern for dynamically adding functionality to objects

### Concurrency Patterns

Go's concurrency primitives (goroutines and channels) enable powerful patterns for concurrent programming:

- **[channels-fan-in-pattern](./channels-fan-in-pattern/)** - Merge multiple input channels into a single output channel
- **[channels-fan-out](./channels-fan-out/)** - Distribute work from a single input channel to multiple workers
- **[channels-generator-micro-pattern](./channels-generator-micro-pattern/)** - Generate values and send them to a channel
- **[channels-pipeline-pattern](./channels-pipeline-pattern/)** - Chain multiple stages of processing using channels
- **[channels-select](./channels-select/)** - Use the select statement for multiplexing channel operations
- **[channels-tee-pattern](./channels-tee-pattern/)** - Split a single input channel into two identical output channels
- **[channels-worker-pool](./channels-worker-pool/)** - Implement a worker pool pattern for concurrent task processing
- **[done-channel-pattern](./done-channel-pattern/)** - Use a done channel to signal cancellation and cleanup
- **[filter-pattern](./filter-pattern/)** - Filter values from a channel based on a predicate function

### Go Concepts & Fundamentals

- **[closed-channels](./closed-channels/)** - Demonstrates behavior of closed channels and how to handle them
- **[nil-channels-block](./nil-channels-block/)** - Shows how nil channels block indefinitely
- **[interface-type-assersion](./interface-type-assersion/)** - Examples of type assertions with interfaces
- **[goroutines-fair-execution-time](./goroutines-fair-execution-time/)** - Demonstrates fair scheduling of goroutines

### Interview Examples

- **[interview-examples](./interview-examples/)** - Common Go interview topics including:
  - **[interfaces](./interview-examples/interfaces/)** - Comprehensive interface concepts (implicit implementation, type assertions, nil interface gotchas)
  - **[arrays](./interview-examples/arrays/)** - Array declaration, initialization, value semantics, and comparison
  - **[maps-simple](./interview-examples/maps-simple/)** - Basic map operations and iteration
  - **[slice-simple](./interview-examples/slice-simple/)** - Slice fundamentals, length, capacity, and append behavior
  - **[slice-and-colons](./interview-examples/slice-and-colons/)** - Slice expressions and backing array pitfalls
  - **[append-slice-with-colon](./interview-examples/append-slice-with-colon/)** - Advanced slice manipulation
  - **[strings](./interview-examples/strings/)** - String immutability, UTF-8 encoding, runes vs bytes, and common operations

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher (some examples use generics)

### Running the Examples

Each pattern is self-contained. To run any example:

```bash
cd <pattern-directory>
go run main.go
```

For example:

```bash
cd builder-pattern
go run main.go
```

## 📖 Pattern Descriptions

### Builder Pattern
The Builder pattern provides a flexible solution for constructing complex objects. It separates the construction of a complex object from its representation, allowing the same construction process to create different representations.

### Decorator Pattern  
The Decorator pattern allows behavior to be added to individual objects dynamically, without affecting other objects from the same class. This example uses a coffee shop scenario where you can add different ingredients to a base coffee.

### Fan-In Pattern
The Fan-In pattern merges multiple input channels into a single output channel. This is useful when you have multiple producers that you want to consolidate into a single stream of data.

### Fan-Out Pattern
The Fan-Out pattern distributes work from a single channel to multiple workers, allowing concurrent processing of tasks.

### Worker Pool Pattern
The Worker Pool pattern manages a fixed number of workers that process tasks from a queue. This is efficient for controlling concurrency and resource usage.

### Pipeline Pattern
The Pipeline pattern chains multiple stages of processing together, where the output of one stage becomes the input to the next stage.

### Done Channel Pattern
The Done Channel pattern uses a channel to signal when work should stop, enabling graceful shutdown and cleanup.

## 🤝 Contributing

Contributions are welcome! Feel free to submit pull requests with new patterns or improvements to existing ones.

## 📝 License

This project is open source and available for educational purposes.

## 🔗 Resources

- [The Go Programming Language](https://golang.org/)
- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) by Katherine Cox-Buday
- [Go Design Patterns](https://github.com/tmrts/go-patterns)
