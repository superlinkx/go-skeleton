FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install github.com/golang/mock/mockgen@latest
USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
