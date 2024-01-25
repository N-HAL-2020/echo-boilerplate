FROM golang:1.21.4

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

WORKDIR /workspace

RUN \
    # install tools
    go install github.com/cosmtrek/air@latest \
    && go install github.com/go-task/task/v3/cmd/task@latest \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2 \
    # apt
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    lsof \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/* \
    # create non-root user
    && groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && chmod -R a+w /go/pkg

USER $USERNAME

EXPOSE 8080