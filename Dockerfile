FROM gcr.io/jenkinsxio/jx-cli-base:0.0.21

ENTRYPOINT ["jx-health"]

COPY ./build/linux/jx-health /usr/bin/jx-health