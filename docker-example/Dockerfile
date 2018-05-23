FROM ubuntu
MAINTAINER XieWei
RUN mkdir -p /app
COPY main /app
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/bin/bash", "-c", "/app/main"]