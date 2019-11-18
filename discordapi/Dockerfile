FROM golang:1.12.6

RUN apt-get update

# Instal go dep 
# RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

# Add Project Source Code
RUN mkdir -p /go/src/github.com/user/discordapi
COPY . /go/src/github.com/user/discordapi
WORKDIR /go/src/github.com/user/discordapi

# Install go dependencies
# RUN dep ensure 

RUN go install .

CMD if [ ${APP_ENV} = "local" ]; \
	then \
	dep ensure && \
	go get github.com/pilu/fresh && \
	fresh; \
	else \
    discordapi