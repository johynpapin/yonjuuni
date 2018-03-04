# Yonjuuni (四十二)

Yonjuuni is a web application to learn vocabulary for readers of the KKLC book.

## Status

Yonjuuni is currently under development. The website is not usable.

## Structure

The project is organized into micro-services.
The WEB application uses the Vue framework, and communicates with microservices.

## Installation

If you want to test the actual project, you can follow this little guide:

First, you have to install [Docker](https://docs.docker.com/install/) and [Docker Compose](https://docs.docker.com/compose/install/).

Then, you can clone the repository:
```zsh
git clone https://github.com/johynpapin/yonjuuni.git
```

Finaly, you can build and run the project:
```zsh
docker-compose build && docker-compose run
```
