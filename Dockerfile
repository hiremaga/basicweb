FROM busybox:ubuntu-14.04
MAINTAINER Abhi Hiremagalur <abhijit@hiremaga.com>
ADD basicweb basicweb
RUN chmod a+x basicweb
ENV PORT 80
EXPOSE 80
ENTRYPOINT ["/basicweb"]
