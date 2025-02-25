# Go-Interviews

Go is described as a Statically-typed, compiled language designed by Google around a set of core principles that prioritize Simplicity, Readability, Maintanability and Efficiency.

- **Simplicity**: Go's syntax is designed to be straigthfoward to keep code readable, maintainable and easy to learn. By minimizing syntactic clutter such as requiring no parentheses around conditions in control structures or forcing explicit semicolon usage, the language reduces boilerplate and allows developers to focus on the logic.

- **Readability**: because is clear, straigthfoward code is easier to maintain, debug and collaborate on. This focus on readability is reflected in Go's design choices such as minimalistic syntax, enforced formating with `gofmt` and explicit error handling.

- **Maintainability**: because it recognizes that code is read, modified and extended far more often than it is originally written. This emphasis means that the language and its standar practices are designed to keep the code base simple, consistent and easy to underestand over time.

- **Efficiency**: beacause it was designed from the ground up to meet the demands of high-performance, modern software, especially in networked and concurrent enviroments. Key reasons: compiled to native code, ligthweigth concurrency, efficient memory management (GC), simplicity.

---
### Concurrency vs Parallelism

- **Concurrency**: Refers to the ability of a system to execute multiple task through simultaneous executions or time-sharing, managin interactions and sharing resources.

- **Parallelism**: Is the ability to excecute multiple task at the same time using multiple processors(CPUs).


**In Go** concurrency refers to the ability to handle multiple task independently, while parallelism is about executing those task simultaneously on different CPU cores. Go's ***goroutines*** provide ligthweigth threads to run functions concurrently and the ***channels*** facilitates safe comunication between goroutines.

- **Goroutine**: is a ligthweigth thread of execution managed by the Go runtime.

- **Channels**: are a Data Structure that allows goroutines to communicate and synchronize with each other.

### Race Condition vs. Deadlock

#### Race Condition:
Occurs when two or more threads or processes access shared data concurrently, and the final outcome depends on the unpredictable timing or sequence of execution.
It can lead to inconsistent or unexpected behavior because the operations interleave in unforeseen ways.
Example: Two threads updating the same counter without proper synchronization may result in an incorrect count.

#### Deadlock:
Happens when two or more processes are each waiting for the other to release a resource, leading to a standstill where none of them can proceed.
It is characterized by a circular dependency among processes, where each holds a resource and waits for another.
Example: Process A holds resource X and waits for resource Y, while Process B holds resource Y and waits for resource X.

Key Differences:

**Nature**:
- Race Condition: A timing issue that affects data consistency.
- Deadlock: A resource allocation issue that halts system progress.

**Symptoms**:
- Race Condition: Unpredictable outcomes or corrupted data.
- Deadlock: System freeze or unresponsive behavior as processes wait indefinitely.

**Mitigation**:
- Race Condition: Use proper synchronization mechanisms like locks, mutexes, or atomic operations.
- Deadlock: Employ strategies such as resource ordering, deadlock detection and recovery, or avoiding circular wait conditions.

---

# OOP Concepts in traditional OOP language vs Go

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


---

### REST

**Representational State Transfer** is a software architecture that imposes conditions on how an API should work.
##### Principles of the REST architectural style:
1. **Statelessness**: Each request contains all the information needed to process it, and no client context is stored on the server.
2. **Client-Server Separation**: The clien and server are independent of each other.
3. **Uniform Interface**: Consistent and Standarized communication, typically via HTTP methods.
4. **Cacheability**: Responses should indicate whether they can be cached for future use.
5. **Layered System**: Architecture should allow for layers such as proxies, gateways, and load balancers.
6. **Code on Demand**: servers cans extend client functionallity.

##### Status Codes
- **[Informational Responses](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#informational_responses)**(`100 - 199`)
- **[Successful Responses](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#successful_responses)**(`200 - 299`)
- **[Redirecttion Messages](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages)**(`300 - 399`)
- **[Client Error Responses](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses)**(`400 - 499`)
- **[Server Error Responses](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#server_error_responses)**(`500 - 599`)

#### Security
- **Authentication** -> OAuth 2.0, JWT.
- **Athorization** -> Role Base Access Control(RBAC).
- **HTTPS** -> Use TLS/SSL for encryption.
- **Rate Limit** -> Limit the number of request to prevent abuse.
- **CORS** -> Configure Cross Origin Resourse Sharing headers to control access from diferent origins.
- **Security Headers** -> Implement security headers like Content-Security-Policy, X-Frame-Options, etc. to mitigate various Web vulnerabilities.

#### Best Practices
- **Versioning**: Use version numbers in the URL (e.g. /v1/users) for backward compatibility.
- **Filtering**: Use query parameters to filter resources. (e.g. /v1/users?status=active).
- **Sorting**: Use query parameters to sort resources. (e.g. /v1/users?sort=name_asc).
- **Pagination**: Use *limit* and *offset* for large datasets.
- **Error Handling**: Return meaningful Error codes and messages.
- **Documentation**: Use tools like OpenApi (Swagger) for clear, comprenhensive documentation.
- **Caching**: Laverage Caching (server-side or client-side) to improve performance.

##### Benefits:
1. Scalability
2. Flexibility
3. Independence

##### RESTful API request contain:
1. Unique Resourse Identifier (URI)
2. Method (GET, POST, PUT, DELETE)
3. HTTP Headers
4. Data
5. Parameters

---

### Software Principles

#### SOLID

- **Single Responsibility**:la noción de que un objeto solo debería tener una única razón para cambiar.
A class should have only reason to change, meaning that class should have only one job.

  Example: payments system should only implements methods related to payments and for example not for send notifications or emails.
- **Open/Close**: la noción de que las “entidades de software … deben estar abiertas para su extensión, pero cerradas para su modificación”.
Software entities ... should be open for extension, but closed for modification.

  Example: Shape interface with method `Area()`. And two clases `Circle` and `Rectangle` that implements this interface. If we need to add a New shape we can create a new class that implements the `shape interface` without modifying the existing classes.
- **Liskov's Substitution**:la noción de que los “objetos de un programa deberían ser reemplazables por instancias de sus subtipos sin alterar el correcto funcionamiento del programa
states that subclasses should be substitutable for their base classes.

  This means that, given that class B is a subclass of class A, we should be able to pass an object of class B to any method that expects an object of class A and the method should not give any weird output in that case.

  Rectangle and Square example violates Liskov’s Substitution principle when we override setters methods. Creating test for `getArea()`. This is because the call to **`setHeight`** function in the test is setting the width as well and results in an unexpected output.
- **Interface Segregation**: la noción de que “muchas interfaces cliente específicas son mejores que una interfaz de propósito general”
Clients should not be forced to depend upon interfaces that they do not use. Larger interfaces should be split into smaller ones.

  Example: violation of the principle if a interface  ShapeAreaCalculator have methods `calculateArea()` and `calculateVolume()`. If a class square implements the interface then it is forced to implement the `calculateVolume()` method, which it does not need. 

  To overcome this, we can segregate the interface and have two separate interfaces: one for calculating the area and another for calculating the volume. This will allow individual shapes to decide what to implement. Adding Cube class.
- **Dependency Inversion**: la noción de que se debe “depender de abstracciones, no depender de implementaciones”.
La Inyección de Dependencias es uno de los métodos que siguen este principio

  States to classes should depend on abstraction but not on concretion.
  In other words, you must follow abstraction and ensure loose coupling

  Example: 

  ```go
  type NotificationService interface {
	  notify(msg string)
  }

  type EmailNotification struct {}

  func (EmailNotification) notify(msg string) {
	  fmt.Println("Email notification: ", msg)
  }

  type SlackNotification struct {}

  func (SlackNotification) notify(msg string) {
	  fmt.Println("Slack notification: ", msg)
  }

  type Employee struct {
	  notific NotificationService
  }
  ```

  If the `Employee` class depends directly on the `EmailNotification` class, which is a low-level module. This violates the dependency inversion principle.

  To ensured loose coupling. `Employee` is not dependent on any concrete implementation, rather, it depends only on the abstraction `NotificationService`.

#### DRY -> Don't Repeat yourself
#### YAGNI -> You Ain't Gonna Need It
#### KISS -> Keep It Simple, Stupid!