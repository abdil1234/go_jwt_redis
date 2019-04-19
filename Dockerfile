FROM golang:alpine

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/rest_api/
# Add the source code:
ADD . $SRC_DIR
WORKDIR $SRC_DIR

# Build it:
RUN cd $SRC_DIR; go build -o rest_go; 
ENTRYPOINT ["./rest_go"]
