FROM golang:alpine
WORKDIR /apps/simple-go-graphql-jwt-roles
RUN rm -rf app
COPY app app
COPY *.env env/
CMD ["/apps/simple-go-graphql-jwt-roles/app"]