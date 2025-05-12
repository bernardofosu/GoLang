# 💡 OOP Concepts in Python, JavaScript, and Go

We’ll go step by step and explain the main Object-Oriented Programming (OOP) concepts — and then show how they look in:

- 🐍 Python  
- 🛠️ JavaScript  
- 🐹 Go (Golang)  

---

## 🎓 OOP Concepts Explained

There are 4 key pillars of OOP:

1. Encapsulation  
2. Abstraction  
3. Inheritance  
4. Polymorphism  

Let’s explore each with simple examples.

---

## 1️⃣ Encapsulation – "Protect the data"

**Definition:**  
Encapsulation means bundling data (variables) and behavior (methods) into one unit and restricting direct access to some of the object's components.

### 🐍 Python
```python
class Person:
    def __init__(self, name):
        self.__name = name  # private variable

    def get_name(self):
        return self.__name
```

### 🛠️ JavaScript
```javascript
class Person {
    #name; // private field

    constructor(name) {
        this.#name = name;
    }

    getName() {
        return this.#name;
    }
}
```

### 🐹 Go
```go
type person struct {      // lowercase → private to package
    name string           // private field
}

func (p person) GetName() string {
    return p.name         // public method
}
```

---

## 2️⃣ Abstraction – "Show only the necessary"

**Definition:**  
Abstraction hides unnecessary details and shows only the relevant parts of an object.

### 🐍 Python
```python
from abc import ABC, abstractmethod

class Animal(ABC):
    @abstractmethod
    def make_sound(self):
        pass
```

### 🛠️ JavaScript
```javascript
class Animal {
    makeSound() {
        throw "Must implement makeSound()";
    }
}
```

### 🐹 Go
```go
type Animal interface {
    MakeSound()
}
```
> Any type that has a `MakeSound()` method automatically satisfies the interface.

---

## 3️⃣ Inheritance – "Reuse code from parent"

**Definition:**  
Inheritance allows a class (child) to inherit fields and methods from another class (parent).

### 🐍 Python
```python
class Animal:
    def speak(self):
        print("Animal speaks")

class Dog(Animal):
    def speak(self):
        print("Dog barks")
```

### 🛠️ JavaScript
```javascript
class Animal {
    speak() {
        console.log("Animal speaks");
    }
}

class Dog extends Animal {
    speak() {
        console.log("Dog barks");
    }
}
```

### 🐹 Go
```go
type Animal struct {
    Sound string
}

type Dog struct {
    Animal  // embedded struct
}

func (a Animal) Speak() {
    fmt.Println(a.Sound)
}
```

---

## 4️⃣ Polymorphism – "Same method, different behavior"

**Definition:**  
Polymorphism allows the same function or method name to behave differently depending on the object type.

### 🐍 Python
```python
class Animal:
    def speak(self):
        print("Animal speaks")

class Dog(Animal):
    def speak(self):
        print("Dog barks")

def make_it_speak(animal):
    animal.speak()
```

### 🛠️ JavaScript
```javascript
class Animal {
    speak() {
        console.log("Animal speaks");
    }
}
class Dog extends Animal {
    speak() {
        console.log("Dog barks");
    }
}

function makeItSpeak(animal) {
    animal.speak();
}
```

### 🐹 Go
```go
type Animal interface {
    Speak()
}

type Dog struct{}
type Cat struct{}

func (Dog) Speak() {
    fmt.Println("Dog barks")
}

func (Cat) Speak() {
    fmt.Println("Cat meows")
}

func MakeItSpeak(a Animal) {
    a.Speak()
}
```

---

## 🧠 Recap Table

| Concept       | Python                  | JavaScript               | Go                              |
|---------------|--------------------------|---------------------------|----------------------------------|
| Encapsulation | `__var`, getters/setters | `#var`, getters/setters   | Capital/lowercase naming        |
| Abstraction   | ABC + `@abstractmethod`  | Base class with empty fn  | Interfaces                      |
| Inheritance   | `class B(A)`             | `class B extends A`       | Embedding (composition)         |
| Polymorphism  | Function accepts base class | Function accepts base class | Interfaces + method satisfaction |

---

## 💬 Real-World: Functions vs OOP — Which is used more?

🔥 Great question! It depends on the context and the language.

### 🧱 1. OOP — Where It's Common
**Used in:**
- Enterprise apps (banking, CRMs)
- Mobile apps (Swift, Kotlin)
- Game dev (Unity/C#)
- GUI apps
- Large codebases

**Why:**
- Models real-world things
- Modular, reusable
- Extensible (inheritance)
- Interfaces for flexibility

**Languages:**
- Java, C#, Python, C++, TypeScript, Swift

---

### 🧪 2. Functional / Procedural — Where It's Common

**Used in:**
- Scripting & automation (Python, Go, Bash)
- Data science (Python, R)
- Web APIs (Node.js, Go)
- DevOps tools (Terraform)
- CLI tools, microservices

**Why:**
- Simpler and faster to write
- Stateless — good for the cloud
- Avoids over-engineering
- Encourages immutability

**Languages:**
- Go, JavaScript, Python, Rust, Haskell

---

## 🧠 So, Which One?

| Use Case                | Preferred Style       |
|-------------------------|------------------------|
| Simple scripts/tasks    | ✅ Functions            |
| Microservices/APIs      | ✅ Functions + Interfaces |
| Business apps/teams     | ✅ OOP                  |
| Games / UI apps         | ✅ OOP                  |
| Concurrent systems      | ✅ Functional           |

---

In practice, many apps combine both:  
**Start small with functions → grow with OOP if needed.**

