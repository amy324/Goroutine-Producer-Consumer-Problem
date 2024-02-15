

Goroutine Pizzeria: A Go CLI Project on the Consumer/Producer Problem
=====================================================================

Overview
--------

This project is part of a series of short CLI-based programs inspired by classic computer science problems. This series is to demonstrate using Goroutines and Go concurrency features to write solutions to these classic problems.


Problem Statement: The Consumer/Producer Problem
------------------------------------------------

The Consumer/Producer problem is a classic synchronization challenge involving two types of processes: producers that generate data, and consumers that consume the data. These processes share a common, fixed-size buffer or queue. The key challenge is to ensure that producers don't produce data when the buffer is full, and consumers don't attempt to consume data when the buffer is empty.

Project Description
-------------------

In the context of this GoLang CLI project, we simulate a "pizzeria" scenario. Pizzas are produced by a dedicated producer goroutine (`pizzeria` function) and consumed by the main goroutine in the `main` function.

Code Structure
--------------

### Producer (`pizzeria` function):

-   The `Producer` type, implemented in the `pizzeria` function, utilizes a `Producer` struct with channels `data` and `quit`.
-   Continuously produces pizza orders using the `makePizza` function until the number of pizzas reaches a predefined limit (`NumberOfPizzas`).
-   Sends each pizza order to the `data` channel, simulating the attempt to make a pizza.
-   Listens for a quit signal on the `quit` channel to gracefully terminate when the pizzeria is closed.

### Consumer (`main` function):

-   The main function creates a `Producer` instance, initializing channels for data and quit signals.
-   Runs the producer goroutine (`pizzeria`) in the background using the `go` keyword.
-   Consumes pizza orders from the `data` channel in a loop until the number of pizzas made exceeds the limit (`NumberOfPizzas`).
-   Prints status messages based on the success or failure of each pizza order.
-   Closes the pizzeria and prints the final statistics after all pizzas are made.

### Outcome and Performance Review:

-   The code keeps track of the number of successful and failed pizza orders (`pizzasMade`, `pizzasFailed`, `total`).
-   After completing the pizza orders, it provides a performance review based on the number of failed orders, categorizing the performance as Excellent, Satisfactory, Some improvement needed, Improvement needed, or Significant improvement needed.

Getting Started
---------------

1.  Clone this repository to your local machine.
2.  Ensure you have GoLang installed.
3.  Navigate to the project directory in your terminal.
4.  Run the `go run main.go` command to execute the pizzeria simulation.

Conclusion
----------

Explore the code, experiment with variations, and gain insights into how Goroutines and channels can effectively address synchronization challenges, as exemplified by the Consumer/Producer problem in this GoLang CLI project.
