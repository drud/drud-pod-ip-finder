FROM busybox
ADD main main
ENTRYPOINT ["/main"]
