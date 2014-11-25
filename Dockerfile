FROM busybox:ubuntu-14.04
MAINTAINER Abhi Hiremagalur <abhijit@hiremaga.com>
ADD basicweb basicweb
RUN chmod a+x basicweb
EXPOSE 8080
ENTRYPOINT ["/basicweb"]
