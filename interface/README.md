## This repository showcases GoLang code I developed while actively practicing and honing my GoLang skills.


### Interface Questions.

1. **Basic** Create an interface Greeter with one method Greet() string.
    * Implement Greeter for a Person struct where Greet returns the person’s name.

    *[solution](InterfaceInGo/interfaceBasic)*
2. **Animal Sounds** Define an interface Animal with a method Speak() string.
    * Implement Animal for Dog, Cat, and Cow structs. Test by creating a slice of Animal and calling Speak() on each.

    *[solution](InterfaceInGo/animal)*
3. **Simple Shape Interface** Create a Shape interface with methods Area() float64 and Perimeter() float64.
    * Implement Shape for a Rectangle and Circle struct.

    *[solution](InterfaceInGo/simpleShape)*
4. **Reusable Printer** Write a function PrintDetails that accepts an interface with a Details() string method.
    * Use it with structs like Book and Car that implement the Details method.

    *[solution](InterfaceInGo/reusablePrinter)*
5. **Multiple Structs with One Interface** Define an interface Transport with methods Start() string and Stop() string.
    * Implement it for Car, Bike, and Bus structs.
    * Write a function that takes a slice of Transport and calls Start and Stop for each.

    *[solution](InterfaceInGo/multiStruct)*
6. **Empty Interface** Create a function Describe that accepts an interface{} (empty interface) and prints:
    * The type of the input.
    * The value of the input.
    * Test with integers, strings, and custom structs.

    *[solution](InterfaceInGo/emptyInterface)*
7. **Embedded Interfaces** Define two interfaces Flyer (with Fly() string) and Swimmer (with Swim() string).
    * Create a combined interface Amphibian that embeds both.
    * Implement Amphibian for a Duck struct.

    *[solution](InterfaceInGo/embedded)*

8. **Mock Testing with Interfaces** Define an interface Database with methods Save(data string) and Fetch() string.
    * Implement Database for RealDatabase and MockDatabase.
    * Write a function that uses the interface to interact with the database and test it with MockDatabase.
    *[solution](InterfaceInGo/mockTest)*