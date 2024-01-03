# Go-Reference-Structure
This is an opinionated way to structure your Go web API, that I like to use.

## Project Structure

The project structure is as follows:

* `/cmd/api` - The app's entrypoint. Only startup logic for the API should be written here.

* `/internal` - The private application and library code. This is where the majority of the code should be written.
* `/internal/handlers` - The API handlers for each route. Each handler should be grouped into related handlers, such as auth handlers, health checks etc.
* `/internal/models` - The models used in the application. This is where the database models should be defined.
* `/internal/server` - The server configuration and setup. This is where the server should be configured, and the routes should be registered.
* `/internal/server/routes` - This is where you defined the routes for the API. Each route should be grouped into related routes, such as auth routes, health check routes etc.
* `/tests` - The tests for the application. This is where the tests should be written.

In most of the files, comments are provided to explain why I've structured the code as I have.

## Design choices
The reason I like to structure my projects like this is because it allows me to easily find the code I am looking for. I also like to keep the API handlers separate from the server configuration, as I feel this makes it easier to read and reason about the code.

I find this structure to be very easily scalable, and I can easily navigate to the code I am looking for, regardless of the size of the project. I also - even though it's not the Go standard - like to keep my tests in a separate folder, as I find it confusing to have application code and test code in the same folder.
