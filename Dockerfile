FROM scratch
COPY mqtt-homekit-light /
ENTRYPOINT ["/mqtt-homekit-light"]