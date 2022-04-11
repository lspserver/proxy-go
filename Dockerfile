FROM ghcr.io/lspserver/servers:latest

USER craftslab
WORKDIR /home/craftslab
RUN mkdir proxy
COPY . proxy/
RUN apt update && \
    apt install -y upx
RUN make build

USER craftslab
WORKDIR /home/craftslab/proxy
ENTRYPOINT ["./bin/proxy"]
