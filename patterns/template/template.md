# Design Patterns Notes

## Template Method

- Fixed algorithm structure.
- Common Implementation implemented in the base.
- Abstract steps are defined as abstract methods forcing implementation to be done by subclasses.
- Hook methods: to optionally override in the subclasses.

### Example

```java
class Bevarage  {
    private String name;
    private Integer temparature;
    private List<String> ingredients;
}

class BevarageMaker {

    public Bevarage prepareBevarage(Bevarage b) {
        // fixed algorithm, as list of steps
        Bevarage b = new Bevarage();
        boil();
        brew(b);
        pour(b);
        if condimentsRequired() {
            add(b);
        }
    }

    // default implementation provided, as boiling water is a common step for all bevarages.
    public void boil() {
        System.out.println("Boiling water");
    }

    // abstract method to be provided by subclass, as brewing can vary per bevarage.
    abstract void brew();

    public void pour(Bevarage b) {
        b.setTemplate(85);
    }

    abstract void add(Bevarage b);


    // hook method
    public boolean condimentsRequired() {
        return true;
    }
}

class TeaMaker extends BevarageMaker {

    private void brew(Bevarage b) {
        System.out.println("Brewing tea");
    }

    private void add(Bevarage b) {
        System.out.println("Adding tea condiments");
    }
}

class CoffeeMaker extends BevarageMaker {

    private void brew(Bevarage b) {
        System.out.println("Brewing coffee");
    }

    private void add(Bevarage b) {
        System.out.println("Adding coffee condiments");
    }
}
```

## Strategy Pattern

```java
interface BrewingStrategy {
    brew();
}

class TeaBrewingStrategy implements BrewingStrategy {
    public void brew() {
        // implement custom logic for brewing the tea.
        boil();
        addCondiments();
        filter();
        pour();
    }
}

class CoffeeBrewingStrategy implements BrewingStrategy {
    public void brew() {
        // implement custom logic for brewing the coffee.
        boil();
        pour();
    }
}

class BevarageMaker {
    private BrewingStrategy brewingStrategy;

    public BevarageMaker(BrewingStrategy bs) {
        this.brewingStrategy = bs;
    }
    public void makeBevarage() {
        brewingStrategy.brew();
    }
}

public static void main(String[] args) {
    BevarageMaker maker = new BevarageMaker(new TeaBrewingStrategy());
    maker.makeBevarage();
}
```

## Facade Pattern

* A unified interface wiring up multiple abstract interface which the sub-system has. Provides an abstration for the client and makes it loosely coupled. Law of Delmeter (Principle of least knowledge (an object should know only about its immediate neighbours)).

Refer manual notes.


## Summary

### Algorithm Structure:

- Template Method: The algorithm structure is fixed in the base class and cannot be changed by subclasses.
- Strategy: Each strategy can have its own algorithm structure.

### Inheritance vs Composition:

- Template Method: Uses inheritance - subclasses inherit and override specific steps.
- Strategy: Uses composition - different strategies are injected into the context class

### When to Use:

- Template Method: When you have a fixed series of steps but different implementations of those steps.
- Strategy: When you have completely different algorithms that can be swapped at runtime

### Flexibility:

- Template Method: Less flexible as the algorithm structure is fixed.
- Strategy: More flexible as entire algorithms can be swapped

### Control:

- Template Method: Base class controls the algorithm flow.
- Strategy: Each strategy controls its own flow

## Stratergy Pattern Detailed Notes

- Stratergy is family of algorithms, allowing the clients to be not impacted when one of algorithms change.
- if a collection -> list had a sorting algorithm, we can swap it with another sorting algorithm, without impacting the client.

Refer manual notes

```java


```


---

## SOLID

What is Single Responsibility required?

- A class must have only a single reason to change, which means it should have only one job.

Sometimes in startups or nascent stages, one need not wait for the requirement to find the change. The change can be anticipated by the developer and the the code made to change for single reason or have a single job.

What if we violate this, what can go wrong?.

- Nothing can go wrong, infact all pieces of code can be written in a single file.

Problems:

    - Maintanence & Complexity: Understanding the code and maintaing will be complex, if multiple logics are combined in a single class.
    - Testing: Testing will be complex, as it would become hard to mock out objects.
    - Tight Coupling: Difficult to swap out implementations (in practical use cases it very rare).
    - Lack of Reusability: If everything is packed in a class, the using client class will not require all the logics packed in this dependent class.

Example of class which follows Single Responsibilty:

```java
class OrderProcessor {

    public void processOrder(Order req) {
        // business logic
        int count = orderRepository.checkInvetory(req.sku);
        if count > 0 {
            orderRepository.updateInventory(req.sku, count - 1);
            orderRepository.createOrder(req);
        }
    }
}
```

Example of class violating Single Resposbility:

```java
class VideoPlayer {
    private List<Video> videos;
    private Database db; // tight coupling, not swappable.

    public void playVideo(int videoid) {
        Video v = findVideoById(videoid);
        if v != null {
            // play
        } else {
            // throw error
        }
    }

    public void addVideo(Video v) {
        videos.add(v);
        saveVideoToDb(v);
    }

    private Video findVideoById(int videoid) {
        return db.findVideoById(videoid);
    }

    private void saveVideoToDb(Video v) {
        db.saveVideo(v);
    }
}

// Responsible for playing and adding video.
// hard to read and maintain, as there are multiple contexts / logics.
// testing is harder: If anything changes on addVideo need to ensure that existing functionality is not broken.
// tight coupling, not swappable.
// not extensible, existing class has to be modified.
```

Q> In the above order processing example, if there are multiple types of processing an order, would a separate class be required to comply with SOLID?
A> Yes

Reason:

- The order processing service, only responsibility is to put a workflow to process an order.
- The different order processing ways can be modelled as a strategy. We can inject the strategy into the order processing service.
- This way the order processing service is not responsible for the different ways of processing an order. The existing order processing service need not be modified if new strategies are introduced or existing strategies are modified.

```java
class OrderProcessor {
    OrderProcessingStrategy os;
    public void processOrder(OrderReq req) {
        os.process(req);
    }
}

interface OrderProcessingStrategy {
    void process(OrderReq req);
}

class StandardOrderStrategy implements OrderProcessingStrategy {
    public void process(OrderReq req) {
        // step A
        // step B
        // step C
    }
}
```

## Open Closed Principle

- Classes should be open for extension like strategy pattern and the order processing example and closed to getting modified.
- Reduces the scope of getting into unwanted errors in already tested classes.
- More maintainable, readable and less complex.
- Swappable and loose coupling.

## Liskov Substitution Principle

- TBD

## Use Case Encountered in work

Any kind of creation/updation + any other operations on an application creation, can have variations for scheme + sector combination.

Example:

1/ Create Application for RLP(Sector) Personal Loan(scheme)
2/ Create Application for RLP Credit Line
3/ Creation Application for SCF, no sector here

Come up with an extensible solution to solve this.

Things to note:

- The request and response payloads for create, update and any other operations is going to be different.
- There can be some commonalities in business logic/activites between a scheme+sector.

Thought Process:

```java
// Thought Process

// why not write like this. This is similar to the diff order processing ways.
// We need to create separate
class ApplicationService {
    private ApplicationRouter router;
    public ApplicationResponse createApplication(ApplicationReq r) {
        // i need to have something to delegate the request to.
        // I need to tell its a create request and pass along the request and get an expected payload.
        // We need a route identifier which can delegate to the appropriate router.
        // To identify a router all routers must comply to a
        switch r.scheme {
            case "rlp":
                createRlpApplication(r);
            case "scf":
                createScfApplication(r);
        }
    }

    private ApplicationResponse createRlpApplication(ApplicationReq r) {
       // step A
       // step B
       // step RLP_A'
    }

    private ApplicationResponse createScfApplication(ApplicationReq r) {
       // step A
       // step B
       // step SCF_A'
    }
}

// violates the S, O in SOLID - order processing example
```

Requirements

- Handle creates for each scheme + sector combination differently. Though there can be commonalities.

Thought process

- We require a service class for creating an application - (ApplicationService)
- the application service should be responsible only for putting together the workflow for creating an application.
- application service can handle create, update and any other operations of an application.
- application service can have a router to delegate the request to the appropriate router. (Reason is otherwise the application service will have to compose the different strategies for creating/updating an application).

  ```java
      class ApplicationService {
          private CreateUapiRouter createUapiRouter;
          private UpdateUapiRouter updateUapiRouter;

          public ApplicationResponse createApplication(ApplicationReq r) {
              return createUapiRouter.route(r);
          }

          public ApplicationResponse updateApplication(ApplicationReq r) {
              return updateUapiRouter.route(r);
          }
      }
  ```

- `RouterDelegate` acts as a **_facade_** to the different strategies for creating/updating an application.
- `RouterLookUp` acts as a Factory that creates and manages different router implementations.
- Each `Router` (RlpApplicationRouter, ScfApplicationRouter, etc.) represents a different strategy for handling requests. Strategy is used here because the algorithm is different for different scheme combinations. They follow a specific contract from `CreateUapiRouter` and `UpdateUapiRouter`.
- `RouterDelegate` is dependency injected to the service layer.

Changed impl

```java
class ApplicationService {
    RouterDelegate routerDelegate; // dependency injected

    public ApplicationResponse createApplication(ApplicationReq req) {
        routerDelegate.route(ApplicationActivity.class, req, "CREATE")
    }
}

// Facade
class RouterDelegate {

    public <E, T, K> K route(Class<E> eventClass, T request, String operation) {
        RouterLookUp routerLookUp = new RouterLookUp(); // created a new factory.

        K resp = switch operation {
            case "CREATE" -> {
                CreateUapiRouter createUapiRouter = routerLookUp.getRouter(eventClass);
                yield createUapiRouter.route(request);
            }
            case "UPDATE" -> {
                UpdateUapiRouter updateUapiRouter = routerLookUp.getRouter(eventClass);
                yield updateUapiRouter.route(request);
            }
            default -> throw new UnsupportedOperationException("Unsupported operation: " + operation);
        };
    }
}

// Strategy base cont
interface CreateUapiRouter<T, E> {
    E routeCreateRoute(T request);
}

interface UpdateUapiRouter<T, E> {
    E routeUpdateRoute(T request);
}

class RlpCreateApplicationRouter implements CreateUapiRouter<ApplicationReq, ApplicationResponse> {
    public ApplicationResponse routeCreateRequest(ApplicationReq req) {

    }
}


// Factory + Strategy
class RouterLookUp {
    // strategy of composed classes.
    ScfCollectionRouter scfCollectionRouter;
    RlpApplicationRouter rlpApplicationRouter;
    ScfApplicationRouter scfApplicationRouter;

    private static String RLP_PERSONAL_LOAN = "RLP_PERSONAL_LOAN";
    private static String RLP_CREDIT_LINE = "RLP_CREDIT_LINE";

    // clazz is the activity class, used to differentiate the router for the same the routing key.
    // THe slight problem, here is we bring in coopling to
    public <T, E> CreateUapiRouter<T, E> getCreateRouter(Class<T> clazz, String routingKey) {
        switch (routingKey) {
        }
    }
}
```

Why Composition is not a right one for the above use case?

- There is only one behaviour to create an application, having one behaviour doesn't stop us from using composition. Eg.., Loggers.
- The problem is that we will end up creating multiple interfaces to create behaviours, making it less flexible and adds more complexity.
- The use case above is a clear IS-A relationship, with a single cohesive behaviour, fixed algorithm, shallow inheritance hierarchy.

Use inheritance when:

1. There is a strong semantic relationship. Dog extends Animal. (Use it when we can tell the sub-class is truly a specific version of the base class).
2. Code reusability: If most of the methods of the super class are reusable and meaningful, then use inheritance.
3. Base class is stable and unlikely to change.
4. Base class is fixed and doesn't need to change. Inheritnace locks you into static type hierarchy.



## When to use Inheritance? (Okhravi)

- Use inheritance when you require both sub-type polymorphism and code reuse from the base class.

Example 1: Hierarchical code reuse but no sub type polymorphism, then use composition.
```java

class Parent {
    void method1();
}

class ChildA extends Parent {
    public void method2() {
        // do something
        super.method1(); // uses method1 from parent
    }
}

class ChildB extends Parent {
    public void method3() {
        // do something
        super.method1(); // uses method1 from parent
    }
}

// Use object composition
class ChildA {
    private Reusable reusable; // ChildA has a reusable.
    public ChildA(Reusable reusable) {
        this.reusable = reusable;
    }

    public void method2() {
        this.reusable.method1();
    }
}

class ChildB {
    private Reusable reusable; // ChildB has a reusable.
    public ChildB(Reusable reusable) {
        this.reusable = reusable;
    }
    public void method3() {
        this.reusable.method1();
    }
}

class Reusable {
    public void method1() {

    }
}
```
- If inheritance is used just to hierarchically reuse the code from method 1, then it might not be a right fit for inheritance.

Example 2: If the use case is just sub type polymorphism, use interface (which provides light weight polymorphism). As it allows the sub-types to implement multiple interfaces rather a single hierarchy (inheritance)

```java
class Parent {
    public void method1() {

    }
}

class ChildA extends Parent {
    public void method2() {

    }
}

class ChildB extends Parent {
    public void method3() {

    }
}

// Both method 2 and method 3 do not use method 1, inheritnace is introduced just for polymorphsim.

// Solution
interface Parent {
    method();
}

class ChildA implements Parent {
    public void method() {

    }
}

class ChildB implements Parent {
    public void method() {

    }
}

// So if there is code reuse and polymorphism should we always use inheritance, NO using composition with interfaces is much flexible. 
interface Parent {
    method();
} 

class ChildA implements Parent {
    private Reusable reusable; // use composition
    public ChildA(Reusable reusable) {
        this.reusable = reusable;
    }
    public void method() {
        this.reusable.sharedMethod();
    }
}

class ChildB implements Parent {
    private Reusable reusable; // use composition
    public ChildB(Reusable reusable) {
        this.reusable = reusable;
    }
    public void method() {
        this.reusable.sharedMethod();
    }
}

// Reusable code between ChildA and ChildB 
class Reusable {
    public void sharedMethod() {

    }
}

```

## The only time one should sub type polymorphism

- Use sub type polymorphism only when there is variation in behaviour and not data.
- If not used for change in behaviours, conditionals are required.

```java

class IAttack {
    private String name;
    private String damage;
}

class Thunderbolt extends IAttack {
    private String name;
    private String damage;
}

class Fireball extends IAttack {
    private String name;
    private String damage;
}
```
- The above is not a right use of polymorphism, as the different attach types only variations in data (name and damage).
- The fireball and Thuderbolt do not require a class of their own, they rather object/instances of type attack.

```java
class ITarget {

}

interface IAttack {

    void performAttack(ITarget target);
}

class Thunderbolt implements IAttack {
    public void performAttack(ITarget target) {
        target.health -= 10;
    }
}

class Fireball implements IAttack {
    public void performAttack(ITarget target) {
        target.health -= 50;
    }
}
```
- Again this is not a right use of polymorphism, as the method attack just differentiates from what value the target health is reduced, which we can do via a simple class.
- The Thunderbolt and Fireball do not require a class of their own, they rather object/instances of type attack.

```java
// better solution for above
class Attack {
    private String name;
    private String damage;

    public void peformAttack(ITarget target) {
        target.health -= this.damage;
    }
}
```

Behaviour driven polymorphism, in the below example both attack and heal are types of move and they exhibit different behaviour. Attack just decrements the health of target wherease heal resets the health and the armour. This might be a right case of polymorphism.

```java
interface IMove {
    void move(ITarget target);
}

class Attack implements IMove {

    public void move(ITarget target) {
        target.health -= this.damage;
    }
}

class Heal implements IMove {
    public void move(ITarget target) {
        target.health = 100;
        target.armor = 100
    }
}
```




