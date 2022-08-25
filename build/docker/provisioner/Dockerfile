ARG TAG="v0.0.22"

FROM rancher/local-path-provisioner:${TAG}

FROM kubesphere/distroless-static:nonroot
COPY --from=0 /usr/bin/local-path-provisioner /local-path-provisioner

USER nonroot
ENTRYPOINT ["/local-path-provisioner"]
