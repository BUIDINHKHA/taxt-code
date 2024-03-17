# Go Version
FROM golang:1.21

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0



# Download the Chrome
RUN apt-get update
RUN apt-get install -y libasound2 fonts-liberation libatk-bridge2.0-0 libatk1.0-0 libatspi2.0-0 libcairo2 libcups2 libdbus-1-3 libdrm2 libglib2.0-0 libgtk-3-0 libgtk-4-1 libnspr4 libnss3 libpango-1.0-0
RUN apt-get install -y libu2f-udev libvulkan1 xdg-utils
RUN wget -O /tmp/chrome.deb https://dl.google.com/linux/chrome/deb/pool/main/g/google-chrome-stable/google-chrome-stable_114.0.5735.90-1_amd64.deb
RUN dpkg -i /tmp/chrome.deb

# Install chromedriver
RUN apt-get install -yqq unzip
RUN wget -O /tmp/chromedriver.zip http://chromedriver.storage.googleapis.com/114.0.5735.90/chromedriver_linux64.zip
RUN unzip /tmp/chromedriver.zip chromedriver -d /usr/local/bin/

RUN mkdir /app
COPY . /app
#COPY ./chromedriver /usr/local/bin/
WORKDIR /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
# Expose port 3000 for the application
EXPOSE 3000

ENTRYPOINT CompileDaemon -build="go build -o main" -command="./main"
