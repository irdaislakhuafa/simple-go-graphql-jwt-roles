FROM golang:alpine
WORKDIR /apps/simple-go-graphql-jwt-roles
RUN rm -rf app env/*
COPY app app
COPY *.env env/
CMD ["/apps/simple-go-graphql-jwt-roles/app", "-app-mode=PROD"]