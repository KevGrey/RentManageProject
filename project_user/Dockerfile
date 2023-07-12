FROM alpine
WORKDIR /Initial
COPY ./target/project-user .
COPY ./config/config-docker.yaml .
RUN  mkdir config && mv config-docker.yaml config/config.yaml
EXPOSE 8083 8881
ENTRYPOINT ["./project_user"]