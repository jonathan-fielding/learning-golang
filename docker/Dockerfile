# The base go-image
FROM golang:1.15.7-alpine3.13

# Add Git support to enable us to install libraries
RUN apk add git
 
# Create a directory for the application
RUN mkdir /app
 
# Copy files from the current directory to the image's application directory
COPY . /app
 
# Set docker images working directory
WORKDIR /app

# Install the go-cache library we are using in our application
RUN go get github.com/patrickmn/go-cache
 
# go build will build an executable file named main in the current directory
RUN go build -o main . 
 
# Run the generated "main" executable
CMD [ "/app/main" ]