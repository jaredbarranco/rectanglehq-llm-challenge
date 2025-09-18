FROM mcr.microsoft.com/devcontainers/base:ubuntu-22.04

ENV DEBIAN_FRONTEND=noninteractive LANG=C.UTF-8

# Install Node.js, TypeScript, and opencode-ai
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g typescript ts-node opencode-ai

# Install LazyGit (Linux ARM64)
RUN wget https://github.com/jesseduffield/lazygit/releases/download/v0.55.1/lazygit_0.55.1_Linux_arm64.tar.gz -O lazygit.tar.gz \
    && tar -xf lazygit.tar.gz lazygit \
    && install lazygit /usr/local/bin \
    && rm lazygit lazygit.tar.gz

# Install Neovim + Python dependencies + networking tools
RUN apt-get update && apt-get install -y --no-install-recommends \
    neovim \
    python3 \
    python3-pip \
    python3-venv \
    dnsutils \
    iputils-ping \
    net-tools \
    wget \
    curl \
    git \
    bash-completion \
    && rm -rf /var/lib/apt/lists/*

# Non-root user
USER vscode
WORKDIR /workspace

# Clone Neovim config
RUN git clone https://github.com/jaredbarranco/kickstart.nvim /home/vscode/.config/nvim

# Default command
CMD ["bash"]
