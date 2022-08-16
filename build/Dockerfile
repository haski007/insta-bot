# Create a base layer with linkerd-await froma recent release.
FROM docker.io/curlimages/curl:latest as linkerd
ARG LINKERD_AWAIT_VERSION=v0.2.0
RUN curl -sSLo /tmp/linkerd-await https://github.com/olix0r/linkerd-await/releases/download/release%2F${LINKERD_AWAIT_VERSION}/linkerd-await
RUN chmod 755 /tmp/linkerd-await

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && mkdir /app

COPY --from=linkerd /tmp/linkerd-await /linkerd-await
ADD app /app

WORKDIR /app

ENTRYPOINT ["/linkerd-await", "--shutdown", "--"]
CMD  ["/app/app"]
