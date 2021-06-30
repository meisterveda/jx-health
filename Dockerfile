FROM ghcr.io/jenkins-x/jx-boot:latest

ENTRYPOINT ["jx-health"]

COPY ./build/linux/jx-health /usr/bin/jx-health
