# Health Probe

This document outlines the implementation of a health probe for multi service architecture in a distributed system.

# How health is measured



**Definition:** _The system is healthy if it is able to serve the requests it receives within an expected time limit._

According to the above definition, the following categories of bad health can be idenitified:

- **Troubled:** The system is uable to fullfill the requests it receives within an expected time window. This would indicate that the system is responding slower than expected. If the system continues to be pressured to serve requests under these conditions its health condition may progress to a more critical level. If the requests are backed off the system may have breathing space to recover.

- **Critical:** The system is unable to fullfill the requests it receives. Backing off requests will not help the system recover.

To underestand this practicaly lets consider a system with two services, Service A and Service B, where Service A depends on Service B. Lets assume that Service B is self sufficient and does not have any other external dependencies.
Clients are able to connect to Service A in order to access and make requests from the system. Based on the above criteria the system can be considered healthy if,

- Clients are able to connect Service A
- Service A is able to connect to Service B
- Service A is able to process requests from clients, send requests to Service B, process responses from Service B and return responses to clients.
- Service B is able to process requests from Service A and send responses to Service A.


Failure to fullfill any of the above criteria would potentially mean the system has started to go into a Troubled state, if the time taken to serve the requests it requests exceeds the acceptable time limit.


