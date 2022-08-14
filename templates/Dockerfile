FROM debian:buster

ENV RUNNER_OS="linux"
ENV RUNNER_ARCHITECTURE="X64"
ENV RUNNER_VERSION="2.294.0"
ENV RUNNER_NAME="self-hosted"
ENV RUNNER_TOKEN=""
ENV RUNNER_WORKDIR="/tmp/runner"
ENV GITHUB_OWNER=""
ENV GITHUB_REPOSITORY=""

RUN apt-get update \
    && apt-get install -y \
        curl \
        sudo \
        git \
        tar \
        gnupg2 \
        apt-transport-https \
        ca-certificates  \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -m github && \
    usermod -aG sudo github && \
    echo "%sudo ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

#setup docker runner
RUN curl -sSL https://get.docker.com/ | sh
RUN usermod -aG docker github

USER github
WORKDIR /home/github

RUN curl -O -L https://github.com/actions/runner/releases/download/v$RUNNER_VERSION/actions-runner-$RUNNER_OS-$RUNNER_ARCHITECTURE-$RUNNER_VERSION.tar.gz
RUN tar xzf ./actions-runner-$RUNNER_OS-$RUNNER_ARCHITECTURE-$RUNNER_VERSION.tar.gz
RUN sudo ./bin/installdependencies.sh

COPY --chown=github:github entrypoint.sh ./entrypoint.sh
RUN sudo chmod u+x ./entrypoint.sh

ENTRYPOINT ["/home/github/entrypoint.sh"]