FROM golang:1.17 as dependency
WORKDIR /work
ADD ./go.* ./
RUN go mod download

FROM dependency as build
ARG BUILD_SERVICE
WORKDIR /work
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ${BUILD_SERVICE}/cmd/main.go

FROM gcr.io/distroless/static:nonroot-amd64
COPY --from=build /work/app /bin/
ENTRYPOINT ["/bin/app"]