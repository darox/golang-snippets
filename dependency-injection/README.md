Very simple example of dependency injection. 

1. There's a `coffeeMachine`
2. The `coffeeMachine`depends on a `cleaner` 
3. The `cleaner`imerplements the `cleanable` interface
4. The `cleaner` is injected into the `coffeeMachine` via the constructor
5. There's a function `setTools` that allows to change the `cleaner` tools at runtime
6. Through this approach the `cleaner` implementation can be changed at runtime by running `coffeeMachine.cleaner = dryCleaner`

Have fun 👨‍💻👩‍💻