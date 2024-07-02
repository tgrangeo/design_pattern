# Golang Design Patterns Practice

Welcome to the Golang Design Patterns Practice repository! This repository contains practical implementations of various design patterns in Golang. The primary goal of this project is to practice, demonstrate, and understand different design patterns in real-world settings. Each pattern is implemented with a focus on clarity and practical use cases.

## Design Patterns

### Creational Patterns
1. **Singleton**: Ensures a class has only one instance and provides a global point of access to it.
2. **Factory Method**: Defines an interface for creating an object, but lets subclasses alter the type of objects that will be created.
3. **Abstract Factory**: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.
4. **Builder**: Separates the construction of a complex object from its representation, allowing the same construction process to create different representations.
5. **Prototype**: Creates new objects by copying an existing object, known as the prototype.

### Structural Patterns
6. **Adapter**: Allows incompatible interfaces to work together by converting the interface of a class into another interface that a client expects.
7. **Bridge**: Decouples an abstraction from its implementation, allowing the two to vary independently.
8. **Composite**: Composes objects into tree structures to represent part-whole hierarchies, allowing clients to treat individual objects and compositions of objects uniformly.
9. **Decorator**: Adds behavior to objects dynamically by wrapping them in an object of a decorator class.
10. **Facade**: Provides a simplified interface to a complex subsystem.
11. **Flyweight**: Reduces the cost of creating and manipulating a large number of similar objects by sharing common parts of the state between them.
12. **Proxy**: Provides a surrogate or placeholder for another object to control access to it.

### Behavioral Patterns
13. **Chain of Responsibility**: Passes a request along a chain of handlers, allowing each handler to either process the request or pass it to the next handler in the chain.
14. **Command**: Encapsulates a request as an object, thereby allowing for parameterization of clients with queues, requests, and operations.
15. **Interpreter**: Implements a specialized language interpreter to interpret expressions or commands in that language.
16. **Iterator**: Provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
17. **Mediator**: Defines an object that encapsulates how a set of objects interact, promoting loose coupling by preventing objects from referring to each other explicitly.
18. **Memento**: Captures and externalizes an object's internal state without violating encapsulation, allowing the object to be restored to this state later.
19. **Observer**: Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
20. **State**: Allows an object to alter its behavior when its internal state changes, appearing as if the object changed its class.
21. **Strategy**: Defines a family of algorithms, encapsulates each one, and makes them interchangeable.
22. **Template Method**: Defines the skeleton of an algorithm in a method, deferring some steps to subclasses without changing the algorithm's structure.
23. **Visitor**: Represents an operation to be performed on the elements of an object structure, allowing you to define a new operation without changing the classes of the elements on which it operates.

## How to Use

Each pattern is implemented in its own directory. To run the examples, ensure you have Golang installed and configured on your machine. Navigate to the desired pattern directory and run the examples using the `go run` command.

```bash
cd singleton
go run main.go
```

---

Happy coding!