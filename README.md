# Go-Interviews

## History

Go is a statically typed, compiled high-level general purpose programming language. It was designed at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and [CSP](https://en.wikipedia.org/wiki/Communicating_sequential_processes)-style concurrency. It is often referred to as Golang because of its former domain name, golang.org, but its proper name is Go.

Statically typed: In computer programming, a type system is a logical system comprising a set of rules that assigns a property called a type (for example, integer, floating point, string) to every term (a word, phrase, or other set of symbols). Usually the terms are various language constructs of a computer program, such as variables, expressions, functions, or modules. A type system dictates the operations that can be performed on a term. For variables, the type system determines the allowed values of that term.

Memory safety is the state of being protected from various software bugs and security vulnerabilities when dealing with memory access, such as buffer overflows and dangling pointers.

Garbage Collection: In computer science, garbage collection (GC) is a form of automatic memory management. The garbage collector attempts to reclaim memory that was allocated by the program, but is no longer referenced; such memory is called garbage.

**Concurrency:** In computer science, concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order, without affecting the outcome. This allows for parallel execution of the concurrent units, which can significantly improve overall speed of the execution in multi-processor and multi-core systems. In more technical terms, concurrency refers to the decomposability of a program, algorithm, or problem into order-independent or partially-ordered components or units of computation.

## Six Main points about Go

- **Static Typing**: Go is statically typed, which means variable types are declared or inferred at initialization and cannot be changed without type conversion. This feature contributes to more predictable code behavior by reducing type-related errors. In Go, every variable has a defined type, providing a higher level of certainty about variable behavior. The Go specification also allows for explicit type declarations or inferred types from the assigned value.
- **Strong Typing**: Go’s strong typing ensures that operations on variables are compatible with their types. For example, adding an integer and a string directly isn’t possible in Go, which avoids bugs that might arise from accidental data type mismatches. This is in contrast to languages like JavaScript, which are more flexible but can result in unexpected runtime behavior.
- **Compiler and Performance**: Go compiles into machine code, resulting in a binary executable that often runs faster than interpreted languages like Python, which translates code at runtime. By pre-compiling, Go reduces runtime overhead and is thus highly suitable for tasks requiring efficient resource use and faster execution. According to benchmarks, Go is known for its speed compared to interpreted languages due to this compilation model.
- **Fast Compilation**: Go’s compilation is optimized for speed, which enables rapid iterations during the development process. This is partly because Go has a single binary output and uses minimal abstractions in its compilation process, so code is ready to run with fewer delays.
- **Concurrency with Goroutines**: Concurrency is a built-in feature of Go, primarily managed through goroutines and channels. Goroutines allow Go to handle multiple tasks simultaneously without the need for complex external libraries. Goroutines are lightweight and highly efficient, making Go a strong choice for applications that require concurrency, such as web servers and data processing.
- **Simplicity and Garbage Collection**: Go’s syntax is designed to be straightforward, aiming for readability and efficiency. Go includes garbage collection, which automatically manages memory allocation and deallocation, helping to reduce the complexity of memory management. This approach aligns with Go’s design philosophy to simplify development


# OOP Concepts in traditional OOP language vs Go

---

### **1. Classes and Objects**
- **Java**: Classes are the blueprint for objects. Objects are instances of classes with state (fields) and behavior (methods).
  ```java
  public class Car {
    private String model;
    public Car(String model) { this.model = model; }
    public void drive() { System.out.println("Driving " + model); }
  }
  Car myCar = new Car("Tesla");
  ```
- **Go**: No classes. Instead, Go uses **structs** (for data) and attaches **methods** to types (including structs).  
  ```go
  type Car struct { Model string }
  func (c Car) Drive() { fmt.Println("Driving", c.Model) }
  myCar := Car{Model: "Tesla"}
  myCar.Drive()
  ```

---

### **2. Inheritance**
- **Java**: Supports **class inheritance** (`extends`) and **interface implementation** (`implements`).  
  ```java
  class ElectricCar extends Car {
    public void charge() { ... }
  }
  ```
- **Go**: **No inheritance**. Uses **composition** (struct embedding) and **interfaces**.  
  ```go
  type ElectricCar struct {
    Car // Embedded struct (composition)
  }
  func (e *ElectricCar) Charge() { ... }
  ```

---

### **3. Polymorphism**
- **Java**: Achieved via inheritance (method overriding) and interfaces.  
  ```java
  interface Drivable { void drive(); }
  class Car implements Drivable { ... }
  ```
- **Go**: Achieved through **interfaces** (implicitly implemented). A type satisfies an interface if it implements its methods.  
  ```go
  type Drivable interface { Drive() }
  // Car implicitly satisfies Drivable if it has a Drive() method.
  ```

---

### **4. Encapsulation**
- **Java**: Uses `private`/`public` modifiers to control access to fields/methods.  
  ```java
  private String model; // Only accessible within the class
  public void drive() { ... } // Accessible everywhere
  ```
- **Go**: Uses **capitalization** for visibility (e.g., `Model` is public, `model` is private). No explicit keywords.  
  ```go
  type Car struct {
    model string // private (lowercase)
    Model string // public (uppercase)
  }
  ```

---

### **5. Abstraction**
- **Java**: Uses `abstract` classes and interfaces to define contracts.  
  ```java
  abstract class Vehicle { abstract void start(); }
  class Car extends Vehicle { void start() { ... } }
  ```
- **Go**: Achieved via **interfaces** and **function types**. No abstract classes.  
  ```go
  type Starter interface { Start() }
  type Car struct{}
  func (c Car) Start() { ... } // Implements Starter
  ```

---

### **6. Constructors**
- **Java**: Explicit constructors with `new` keyword.  
  ```java
  public Car(String model) { this.model = model; }
  Car car = new Car("Tesla");
  ```
- **Go**: No constructors. Uses **factory functions** (e.g., `NewCar()`).  
  ```go
  func NewCar(model string) Car { return Car{Model: model} }
  myCar := NewCar("Tesla")
  ```

---

### **7. Method Overloading**
- **Java**: Supports method overloading (same name, different parameters).  
  ```java
  void drive() { ... }
  void drive(int speed) { ... }
  ```
- **Go**: **No method overloading**. Use unique method names or variadic functions.  
  ```go
  func Drive() { ... }
  func DriveWithSpeed(speed int) { ... }
  ```

---

### **8. Generics**
- **Java**: Uses generics for type-safe code.  
  ```java
  List<String> list = new ArrayList<>();
  ```
- **Go**: Introduced generics in Go 1.18 (simpler syntax, used sparingly).  
  ```go
  func PrintSlice[T any](s []T) { ... }
  ```

---

### **Key Differences Summary**
| **Concept**       | **Java (Traditional OOP)**                | **Go (Idiomatic Approach)**               |
|--------------------|-------------------------------------------|-------------------------------------------|
| **Inheritance**    | Class-based (`extends`/`implements`)      | Composition and struct embedding          |
| **Polymorphism**   | Explicit interface implementation         | Implicit interfaces (duck typing)         |
| **Encapsulation**  | `private`/`public` keywords               | Capitalization for visibility             |
| **Error Handling** | Exceptions (`try`/`catch`)                | Explicit error return values              |
| **Concurrency**    | Threads, `synchronized` blocks            | Goroutines and channels (CSP model)       |
| **Abstraction**    | Abstract classes, interfaces              | Interfaces and function types             |
| **Constructors**   | Constructor methods                       | Factory functions (e.g., `NewX()`)        |

---

### **Why Go Differs**
Go intentionally avoids traditional OOP features like inheritance and overloading to prioritize:
1. **Simplicity**: Less cognitive overhead.
2. **Composition**: Flexible code reuse without rigid hierarchies.
3. **Explicitness**: Clear error handling and minimal hidden behavior.
4. **Concurrency**: Lightweight goroutines over heavy threads.

As Rob Pike noted:  
*"Go is about composition, not inheritance."*  

While Go isn’t strictly OOP, its focus on **interfaces**, **structs**, and **composition** provides a pragmatic way to model real-world problems while maintaining readability and performance.
