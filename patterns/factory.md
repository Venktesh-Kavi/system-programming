# Factory Pattern


## Factory Method (Factory Idiom)

- Useful for single responsibility and avoiding code duplication.
- Ideally its good to seggregate the object creation and business implementation choosing logic from the business specific code.

Eg..,

``` java
public class User {

    User indianUser() {}
    User ukUser() {}
    User defaultUser() {}
}
```
- We can also consider create a class to just handle the factory.

``` java

public class User {
}

public class UserFactory {

    public User createIndianUser() {}
    public User createDefaultUser() {}
    public User createUkUser() {}
}
```
- This approach is suitable when the use case is simple were one factory is sufficient and run time polymorphism is not required. 

## Factory Pattern

- Its similar to template pattern, but creates object rather than containing business logic.
- Its different from factory method, as we have polymorphism here.

``` java
// Interface for the Product
interface Product {
    void use();
}

// Concrete implementation of the Product
class ConcreteProduct implements Product {
    public void use() {
        System.out.println("Using the concrete product");
    }
}

// Factory interface to enforce creation method
interface Factory {
    Product createProduct();
}

// Concrete factory implementation
class ConcreteFactory implements Factory {
    public Product createProduct() {
        return new ConcreteProduct();
    }
}

// Client code
public class Client {
    public static void main(String[] args) {
        Factory factory = new ConcreteFactory();
        Product product = factory.createProduct();
        product.use();
    }
}
```

## Abstract Factory

- Creating an entire family of objects that aren't in the same hierarchy.
- Example furniture store: ModernFurniture, ClassicFurniture.

``` java
class ModernFurniture {

    public ModernSofa createSofa(String material) {
        return new ModernSofa();
    }

    public ModernChair createChair(String type) {
        return new ModernChair();
    }
}

interface Sofa {}

class ModernSofa implements Sofa {}
class ClassicSofa implements Sofa {}
```

## Stratergy Pattern

- Strategy Pattern allows us to change the implementation of an algorithm at runtime. Typically we define an interface which is used to apply an algorithm, we implement it multiple times for each possible algorithm.

``` java
interface Discount {
    Integer applyDiscount(Integer val);
}

class ChristmasDiscount implements Discount {

    public Integer applyDiscount(Integer val) {
        return val * 0.2;
    }
}

class EasterDiscount implements Discount {
    public Integer applyDiscount(Integer val) {
        return val * 0.4;
    }
}

Discount d = new EasterDiscount();
Integer out = d.applyDiscount(10)
```

``` java
// java 8 way
interface Discount {
    Double applyDiscount(Integer amount);

    static Discounter applyChristmasDiscount() {
        return amount -> amount * 0.2; // Creates an anonymous implementation of Discount. Its a functional interface
    }

    static Discounter applyEasterDiscount() {
        return amount -> amount * 0.4;
    }
}

Discounter easterDiscountFn = Discount.applyEasterDiscount();
easterDiscountFn.applyDiscount(10);
```
