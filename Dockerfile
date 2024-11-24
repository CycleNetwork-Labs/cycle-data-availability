# CONTAINER FOR BUILDING BINARY
FROM golang:1.21 AS build

# INSTALL DEPENDENCIES
RUN go install github.com/gobuffalo/packr/v2/packr2@v2.8.3

# Set the GOPRIVATE environment variable
ENV GOPRIVATE=github.com/RollNA

# Set up Git to use the personal access token
ARG GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

COPY go.mod go.sum /src/
RUN cd /src && go mod download

# BUILD BINARY
COPY . /src
RUN cd /src/db && packr2
RUN cd /src && make build

# CONTAINER FOR RUNNING BINARY
FROM alpine:3.16.0
COPY --from=build /src/dist/cdk-data-availability /app/cdk-data-availability
EXPOSE 8444
CMD ["/bin/sh", "-c", "/app/cdk-data-availability run"]
