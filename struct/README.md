## Learning Go structs is a great way to get familiar with how Go handles data types and organizes code. Structs in Go are used to group related data together, and they form the foundation for working with more complex data structures. To help you practice, I'll give you a series of progressively more challenging problems that will allow you to explore Go structs in various contexts.

### Challenge 1: Basic Struct Declaration and Initialization
Problem:

Create a Person struct with the following fields:

    firstName (string)
    lastName (string)
    age (int)

Write a main function where you:

    Declare and initialize an instance of the Person struct.
    Print the details of the Person instance.

Tips:

    Use the fmt.Println() function to print the struct.
    Experiment with both named and unnamed initialization styles.

### Challenge 2: Method on a Struct
Problem:

Create a Rectangle struct with the following fields:

    width (int)
    height (int)

Write a method area() for the Rectangle struct that returns the area of the rectangle. In your main function:

    Create an instance of Rectangle.
    Call the area() method and print the result.

Tips:

    Remember that methods are defined with a receiver type (e.g., func (r Rectangle) area() int { ... }).
    The area of a rectangle is calculated as width * height.

### Challenge 3: Struct with Pointer Receivers
Problem:

Define a BankAccount struct with the following fields:

    accountNumber (string)
    balance (float64)

Add a method deposit(amount float64) that increases the balance by the given amount. Use a pointer receiver to modify the BankAccount instance directly. In the main function:

    Create an instance of BankAccount with an initial balance.
    Call the deposit() method to deposit money into the account.
    Print the updated balance.

Tips:

    Use a pointer receiver for methods that modify the struct's fields.
    You’ll need to pass a pointer to the deposit method (e.g., &account).

### Challenge 4: Struct with Nested Structs
Problem:

Create a Car struct with the following fields:

    make (string)
    model (string)
    year (int)
    engine (a nested Engine struct with fields horsepower (int) and type (string))

Write a method displayInfo() for Car that prints all the information about the car and its engine. In the main function:

    Initialize an instance of Car with a nested Engine.
    Call the displayInfo() method to print the full information.

Tips:

    You can define nested structs inside another struct.
    Use dot notation to access nested fields, e.g., car.engine.horsepower.

### Challenge 5: Struct with Tags and JSON Encoding/Decoding
Problem:

Define a User struct with the following fields:

    id (int)
    firstName (string)
    lastName (string)
    email (string)

Use struct tags to specify JSON field names that differ from the struct field names. Then, in the main function:

    Create a User instance and initialize it with sample data.
    Use encoding/json to convert the User struct to a JSON string and print it.
    Convert the JSON string back into a User struct and print the result.

Tips:

    Use the json package and the json.Marshal() and json.Unmarshal() functions.
    Use struct tags to specify how fields should appear in the JSON (e.g., json:"first_name").

### Challenge 6: Struct with Slice of Structs
Problem:

Create a Team struct with the following fields:

    name (string)
    players (a slice of Player structs)

Define the Player struct with:

    firstName (string)
    lastName (string)
    position (string)

In the main function:

    Create a Team with a list of players.
    Print out the team name and details of each player.

Tips:

    You can create a slice of structs like this: players := []Player{ ... }.
    Use a loop to iterate over the slice of players and print their details.

### Challenge 7: Struct with Multiple Methods
Problem:

Create a Book struct with the following fields:

    title (string)
    author (string)
    price (float64)

Add the following methods to the Book struct:

    applyDiscount(discount float64) – Reduces the price of the book by the given discount percentage.
    getBookInfo() – Returns a string with the book's title, author, and price.

In the main function:

    Create an instance of Book.
    Print the book information using getBookInfo().
    Apply a discount and print the updated price.

Tips:

    For the applyDiscount() method, calculate the discount as: newPrice = price - (price * (discount / 100)).

### Challenge 8: Compare Structs (Equality)
Problem:

Create a Product struct with the following fields:

    name (string)
    category (string)
    price (float64)

Write a function areEqual(product1, product2 Product) bool that compares two Product structs and returns true if they are equal (same name, category, and price). In the main function:

    Create two instances of Product.
    Call the areEqual() function and print the result.

Tips:

    Go does not allow direct comparison of structs with slices or maps, but you can compare basic types directly (e.g., product1.name == product2.name).

Tips:

    The fmt.Stringer interface is implemented by defining the String() method with the signature func (p Product) String() string.

Final Thoughts

These challenges cover a variety of topics and concepts that will help you get comfortable with Go structs. As you work through them, pay attention to how structs and methods interact, how to manage internal state, and how Go’s features (such as method receivers, the sort package, and interfaces) can help you write more efficient and readable code.

Feel free to start with the simpler challenges and gradually increase the complexity as you get more comfortable with Go’s features!