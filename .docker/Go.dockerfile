FROM golang:1.26 AS golang

ARG GO_USER_ID=1001
ARG GO_USER_NAME=krait

ENV GO_USER_ID=${GO_USER_ID}
ENV GO_USER_NAME=${GO_USER_NAME}

ENV GOCACHE=/tmp/.cache/go/build
ENV GOMODCACHE=/tmp/.cache/go/pkg/mod
#ENV GOBIN=/go/bin
# ENV GOPATH=/usr/local
# ENV GOROOT=/opt/go
#ENV GOPROXY=https://proxy.golang.org,direct
#ENV GOSUMDB='sum.golang.org'

# ENV PATH="/go/bin:${PATH}"

COPY --chmod=0755 .docker/scripts/go-* /usr/local/bin/

RUN apt-get update && apt-get install -y net-tools sudo && \
    \
    groupadd -g "${GO_USER_ID}" "${GO_USER_NAME}" && \
    useradd -m -u "${GO_USER_ID}" -g "${GO_USER_ID}" "${GO_USER_NAME}" && \
    echo "${GO_USER_NAME} ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers && \
    echo "User ${GO_USER_NAME} created with ID: ${GO_USER_ID}"

# RUN mkdir -p /opt/go
RUN go env
RUN bash /usr/local/bin/go-install-vscode-tools

COPY .docker/zshrc /home/${GO_USER_NAME}/.zshrc


# ----

FROM golang AS devcontainer

