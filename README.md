# Golang DDD Archetype

:sparkles: This is a template that aims to be simple, complete, flexible and configurable for a golang project.

The main objective of this project is to provide a ready-to-use template that respects and follows the principles and patterns of domain-driven design. At the same time, we want to be able to implement good practices in the development and structuring of projects made with Golang.

## About the project

Currently, in my job there is an archetype built for microservices in Golang, which aims to be a template that contains the main technologies that are handled within the company, Rest, gRPC, Kafka, Mongo, GraphQL, SQL, etc.. However, several factors in its creation make it very complex to use and understand, so it does not facilitate the development process.

This project is created with the idea of offering the same functionalities as the company's archetype, serving as a guide with best practice examples and ensuring robustness in development. It is expected that a developer without so much knowledge can learn and develop on the project in a simple way, as well as a developer with more advanced knowledge can adjust it to their needs.

:white_check_mark: I hope it is useful and that you learn something new, feel free to contribute, suggest or make any comment that will help to improve the project. Thank you!

### Future work
- Soporte gRPC
- Eventos
- Conexión SQL

## :hammer: Built with
- Language: Golang >= 1.22.0
- For dependency injection: fx
- Database: MongoDB
- to containerize: Docker
- Live reload: Air

## Getting Started

![golang-archetype-DiagramaGeneral drawio (1)](https://github.com/user-attachments/assets/000690c6-bd54-4180-81d3-cc3c5e895abe)

As can be seen in the image, the internal structure of the project is organized by grouping the code by functionalities, similar to what is proposed in the vertical slice architecture, focused on isolating the business domain from external code and allowing for high cohesion between the different layers. Within each folder containing functionality, the traditional layers used in DDD architecture can be observed, which adhere to the rules for internal communication by using interfaces and the fx library.


### Usage with Docker Compose
````
docker-compose up --build
````
With docker-compose, it is possible to set up and configure a MongoDB database, which connects using the environment variables configured in the .env file. In this way, you simply need to call the endpoints defined in the router.go file to interact with the data.

## Usage

## License

## Contact
