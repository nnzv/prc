FROM golang:1.18-buster

WORKDIR /prc
COPY . .
CMD ["sh", "mk.sh", "t"]
