FROM apache/kafka:3.8.0

USER root
ARG version=2.6.0

USER appuser

ADD --chown=appuser:appuser https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v$version/opentelemetry-javaagent.jar /tmp/opentelemetry-javaagent.jar
ENV KAFKA_OPTS="-javaagent:/tmp/opentelemetry-javaagent.jar -Dotel.traces.exporter=jaeger -Dotel.metrics.exporter=none -Dotel.logs.exporter=none -Dotel.instrumentation.messaging.experimental.receive-telemetry.enabled=true"