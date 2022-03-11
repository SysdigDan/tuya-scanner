FROM gcr.io/distroless/static
ENTRYPOINT ["/tuya-scanner"]
COPY tuya-scanner /