# Streaming between server and consumer with gRPC 
This Project aims to define the Routing procedures and rules for streaming messages and requests between the server and consumers. 

## Deafult Client-Server Transmission of messages
A normal Client-Server will transmit messages via a response-request format in which: 
- Request-Reeust architecture
- Synchronous
- Slow and waits (for the previous request, response from the server.)
- Not Scalable (very difficult to scale)

# Solution to Scale our Architecture
- Reject Client-Server Arch. instead adopt **Remote Procedure Calls** Architecture
- It allows the server and client to directly call functions to share messages in the form of streams 
-  Instead of JSON, it uses **protbuf** beacuse the size of its payload is just some fraction of the size of a JSON file, which eventually accelerates communication.

### Steram Archietcture
- Streaming is adopted which will be a flow of data stream which doesn't have to wait for response. 
- Continour flow of Data
- Asynchronous and Faster
- Use case: Microservices, Blockchains, etc

## Different Types of Streaming
### 1. Unary API
Just an asynchronous stream of data been sent in bi directional manner. 
### 2. Server Streaming
Here on a single/simple request of client, an entire stream of response/bunch is sent back to the client. 
### 3. Client Streaming
Here a stream of request is been sent to the server all at once to which the server responds. 
### 4. Bi-directional Streaming
- Client & Server both communicate via streams.
- It doesnt wait for any response! These are totally asynchronous and continous flow of messages & streams.
- Order of the requests, resposnse/ messages are preserved to be presented back as a response from either side.
