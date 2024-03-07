FROM fluentd:v1.14-1
# Specify compatible Fluentd version

USER root

RUN ["gem", "install", "elasticsearch", "--no-document", "--version", "7.17.0"]
# Install fluent-plugin-elasticsearch with specific version

RUN ["gem", "install", "fluent-plugin-elasticsearch", "--no-document", "--version", "5.0.3"]

# Install other plugins
RUN ["gem", "install", "fluent-plugin-detect-exceptions"]
RUN ["gem", "install", "fluent-plugin-concat"]
RUN ["gem", "install", "fluent-plugin-rewrite-tag-filter"]