# Interfaces Interview Example

This example demonstrates critical concepts about interfaces in Go that are commonly tested in technical interviews.

## Topics Covered

### 1. Basic Interface Definition
- How to define an interface with method signatures
- The `Shape` interface example with `Area()` and `Perimeter()` methods

### 2. Implicit Interface Implementation
- Go's implicit interface satisfaction (no `implements` keyword)
- Multiple types implementing the same interface (`Rectangle`, `Circle`)

### 3. Polymorphism
- Using interfaces to write generic code
- Storing different concrete types in a slice of interfaces
- The `PrintShapeInfo()` function accepting any `Shape`

### 4. Empty Interface (`interface{}` or `any`)
- The empty interface can hold values of any type
- Using `interface{}` for generic parameters

### 5. Type Assertions
- Safe type assertion with "comma, ok" idiom: `value, ok := i.(Type)`
- Unsafe type assertion (can panic)
- Extracting concrete types from interface values

### 6. Type Switch
- Pattern matching on interface types
- Handling different types in a single function
- The `ClassifyType()` example

### 7. Interface Embedding/Composition
- Combining multiple interfaces into one
- The `NamedShape` interface embedding both `Shape` and `Named`
- How to implement composite interfaces

### 8. Nil Interface vs Nil Value (**Important!**)
- A nil interface value: `var p Printer` (nil interface)
- A non-nil interface holding a nil pointer: `var p Printer = (*MyPrinter)(nil)`
- Methods can be called on nil receivers but not nil interfaces
- `p == nil` returns `false` when interface holds nil pointer!

### 9. Interface Conversion
- Converting from empty interface to specific interface
- Type checking before conversion

### 10. Zero Value of Interfaces
- The zero value of an interface is `nil`
- Calling methods on nil interfaces causes panic

## Common Interview Questions This Covers

1. **What is an interface in Go?**
   - A type that specifies a set of method signatures

2. **How does Go's interface implementation differ from other languages?**
   - Implicit satisfaction - no explicit "implements" declaration needed

3. **What is the empty interface?**
   - `interface{}` or `any` can hold any value

4. **How do you safely check if an interface holds a specific type?**
   - Use type assertion with "comma, ok" idiom

5. **What's the difference between a nil interface and an interface holding a nil pointer?**
   - Nil interface: `var i Interface` - the interface itself is nil
   - Interface with nil pointer: `var i Interface = (*Type)(nil)` - interface is not nil, but the value inside is

6. **Can you call a method on a nil interface?**
   - No, this will panic
   - But you CAN call a method on a non-nil interface that holds a nil pointer (if method handles nil receiver)

## Running the Example

```bash
cd interview-examples/interfaces
go run interfaces.go
```

## Key Takeaways for Interviews

- Interfaces are satisfied implicitly in Go
- Empty interface accepts any type
- Always use safe type assertions with "comma, ok" idiom
- Understand the nil interface gotcha: an interface is not nil if it holds a nil pointer
- Interfaces enable polymorphism and loose coupling
- Type switches are useful for handling multiple types
- Interface embedding allows composition
