FROM golang:1.12.7-buster

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 /usr/local/bin/dep
RUN chmod +x /usr/local/bin/dep

# create application directory in the container
RUN mkdir -p /go/src/app/

# copy the code into application directory in the container
ADD . /go/src/app/

# designate the working directory within the container
WORKDIR /go/src/app/

# copies the Gopkg.toml and Gopkg.lock to WORKDIR
COPY Gopkg.toml Gopkg.lock ./

# Download and install dependencies e.g. github.com/gorilla/mux etc.
RUN dep ensure -vendor-only

RUN go build -o main ./...
CMD ["/go/src/app/main"]