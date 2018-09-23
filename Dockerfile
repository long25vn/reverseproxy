FROM iron/go:dev
WORKDIR /app
ENV DIR=/go/src/dockergo/
ADD . $DIR
RUN cd $DIR; go build -o myapp; cp myapp /app/; cp method.json /app/
ENTRYPOINT ["./myapp"]
EXPOSE 5000