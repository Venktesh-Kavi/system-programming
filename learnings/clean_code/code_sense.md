# CODING GUIDELINES 

> Reference: Clean Code (Uncle Bob)

## Why need coding guidelines?

* Aggressive release cycles, quick product to market movement should be accompanied with careful consideration on code. Companies are brought down by poor code
* Impeding caused by bad code is called wading
* We all think of going back and cleaning it later -
    > LeBlanc's Law: "Later equals never"

### Total Cost of Owning a Mess

* As mess builds, productivity of team continues to asymptotically reduce to zero
* As productivity decreases, management does the following:
    * Add more staffs to project
        - New Staffs with the messy code, do not understand the changes which match the design intent and ones which thwarts the design intent.
        - Under pressure to increase producitvity of project, they succumb to creating more messy code.
        - This drives productivity even below zero.

## What next? 

### Grand Redesign

* Team rebels about the messy codebase, management feel the pinch of productivity decrease and agree to redesign
* Tiger team is formed. Everyone wants to be in tiger team -> green field project
* Two teams are in race. The new team should build everything and adds on which the old team does.
* Management does not replace the old system, unless the new system is completely ready.
* In some companies this race as gone for over 10 years quotes "Rober C Martin"

### Attitude

#### What does good code rot to bad one?

Compains from developers
* Requirement changes which thwarts intial design
* Tight schedules
* We harp about bad managers, intolerant customer and poor marketing startegies

#### The fault is on us, bitter pill to swallow!!!

* Tech, project and product managers look up to us for information.
* Don't be shy in telling them what you think
* Most managers want the truth eventhough though they dont act like it. They want good code even when they are obsessing about schedules.
* They may defend the requirements and the schedule, it is our job to defend the code.


### The Art of Clean Code

* Programmers require "code-sense". A programmer with code sense can look at a messy code and see options and variations.
* Code sense helps programmers choose from the best of the variations and transform the code from messy to an elegant one.
* ***Its not about writing clean code once, it is about continuing to maintain clean code.***
    > Leave the campground cleaner than you found it.

## Meaningful Names

1. Use intention revealing names
2. Avoid Disinformation
    - Grouping of accounts as accountList instead of accountGroup
    - hp, aix, sco are unix variable names. These obsure the meanining of the variable
3. Make meaning distinctions - Do not add noise words or numbers just because the compiler demands it.
    - If names must be different, they should also mean something different.
4. Use pronouncable names
    - Eg.., of some poor variable names: genymdhms (generation date, yeat, month, day, hours etc..,)
    - Programming is a social activity ensure that the variable names are pronounable and colloboarated with.
5. Use searchable names
    - MAX_CLASSES_PER_STUDENT is easily greppable than a number 7.
6. Interfaces and Implementation
    - Leave interfaces unadored. eg.., IShapeFactory is unecessary.
    - Distraction and too much information
    - Add encoding to implementation if necessary.
7. Class Names should be Nouns and Method Names should be Verbs
8. Dont be cute - Will be memorable or connect with people who are inline with your humor.
    - Eg.., HolyHandGrenade - May be deleteItems is better.
9. Pick one word per concept
    - eg.., fetch, retrive, get
10. If it depends use solution domain names
    - Use computer algorithm names, pattern names, math names etc..,
    - Its not necessary to choose every name from the problem domain, which might make co-worker to go back and forth with the customer to understand the name.
    - AccountVisitor is an example, a class which embodies Vistor pattern.
11. Add meaningful context as needed.
    - Eg.., firstName, lastName, street, city etc.., implicity mean they belong to an address.
    - What if we saw that11. Add meaningful context as needed.
    -   - Eg.., firstName, lastName, street, city etc.., taken together mean they belong to an class called address.
    -       - If only a variable called state is present it is difficult to understand the context of this variable, adding a meaning context to it makes sense.
    -       - Eg.., addrFirstName, addrState.
12. Dont add gratatious context
    - If you have an application called ULS, it is unecessary to prefix classes with names UlsDataStore etc..,
    - Be helpful to your IDE so that searches with a letter U for eg.., dont result in a large result set.


## Functions 

1. Functions should be small, but how small?
    - Depending on the context functions can range from 5 - 10 lines as per ULS guidelines
    - Each function should be transparently obvious
    - Each should tell a story
    - Each should lead to a compelling order
2. Do one thing
    - Functions should do one thing, they should do it well.They should do it only.
3. One level of abstraction per function
    - mixing very high level, intemediate and low level abstractions
    - x.getHtml() high level of abstraction, PageParse.parsePage() is intermediate, StringBuffer.append() is low level
    - All the above in a single method, just reduces the cognitive sense of the reader.
4. Step down rule
    - Code should be top-down narrative
    - Every function to be followed by those at the next level of abstraction, descend one level of abstraction at a time
    - Reading a program should be like a set of `TO paragraphs
    - To do x which requires Y, To perform Y etc..,
5. Switch Statements 
    - switch statements always tend to do N things
    - Bury switch statements at a level so that they are not repeated anywhere
    - Violates SRP and OCP, more than one reason for change and changes whenever new types are added respectively.
    - Employee eg:

    ```
        public Money calculatePay(Employee e) {
            switch(e.type()) {
                case COMMISIONED:
                    return calculateCommisionpay(e);
                case SALAIRED:
                    return calculateSalariedPay(e);
                default:
                    throw new IllegalArgumentException("unknown");
            }
        }
    
    This function violates SRP because it can change for multiple, independent reasons:

    Changes to Employee Types

        1/ Adding a new employee type (e.g., Contractor, Intern)
        2/ Removing an existing employee type
        3/ Modifying the employee type enumeration


    Changes to Pay Calculation Rules

        1/ If the structure of how any type calculates pay changes
        2/ If error handling for pay calculation needs to change
        3/ If validation rules for pay calculation change


    Changes to Error Handling

        1/ Modifying how invalid employee types are handled
        2/ Adding new exception types or error conditions
        3/ Changing error reporting requirements


    Changes to Business Process Flow

        1/ If the order of operations needs to change
        2/ If pre or post-processing steps need to be added
        3/ If auditing or logging requirements change


    Changes to the Method's Interface

        1/ If the parameter type needs to change
        2/ If the return type needs to change
        3/ If additional context information needs to be passed


    The core issue is that this method is doing too much:

        1/ Type checking
        2/ Routing logic
        3/ Error handling
        4/ Orchestrating different pay calculations

    Violation of OCP
        - Adding a new employee type requires modifying existing code
        - The switch statement must be updated
        - Risk of breaking existing functionality
        - Hard to maintain as the number of employee types grows
        -
    [Ref Better Swtich Case](https://github.com/Venktesh-Kavi/programming-fundamentals/blob/main/clean_code/better_switch_cases.java)
    ```
`
### Arguments 

- Prefer no arguments to a method. Arguments adds interpretation complexity and makes testing also a lot tougher as we need to mock each parameter
- Each argument comes with a cognitive compleixty for the reader to understand.
- Prefer order no args > monadic > diadic > triadic (think twice before having a triadic argument)
- Examples of monadic
    - boolean fileExsits(File file) -> OK (asks a question about an argument)
    - InputStream fileOpen(File file) -> OK (transforms an input and returns an output)
    - void triggerFileOpenFailedEvent(File file) -> OK (method without return are acceptable for events like this)
    - void transform(StringBuffer sb) -> transforming sb inplace in the method is less preferred than performing StringBuffer transform(StringBuffer sb). (Clearer to the reader)

- Avoid Flag arguments (flags are clear violation of SRP, method does more than one thing)
    - Rather break down into two functions
- Less preference to Dyadic functions does not mean to avoid something like this Point(int x, int y) -> This is absolutely fine.
    - If arguments are ordered components of a single value.
    - Avoid something like asset(expected, actual) -> these two args have no correlation (user has to remember that arg ordering, this adds to reader complexity).
- For triads think twice (consider making objects which seem to group together)
