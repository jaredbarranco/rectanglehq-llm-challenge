#!/bin/bash

docker run -it --rm \
  -e GOOGLE_API_KEY="$GOOGLE_API_KEY" \
  -v "$(pwd)/scripts:/workspace/scripts" \
  -v "$(pwd)/.config/opencode:/home/vscode/.config/opencode" \
  llm-sandbox \
  bash

