FROM busybox:ubuntu-14.04
MAINTAINER Abhi Hiremagalur <abhijit@hiremaga.com>
ADD basicweb basicweb
ADD public public
RUN chmod a+x basicweb
EXPOSE 8080
ENTRYPOINT ["/basicweb"]
