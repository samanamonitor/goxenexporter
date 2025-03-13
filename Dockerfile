FROM golang:latest
RUN <<EOF
apt update
apt upgrade -y
EOF
