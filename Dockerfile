FROM suchja/wine:latest
MAINTAINER contact@myforexidea.com
USER root
RUN mkdir -p /softwares/mt4/MQL4
COPY files/*.* /softwares/mt4/
COPY files/compiler /softwares/mt4/
COPY files/MQL4 /softwares/mt4/MQL4
run chown -R xclient /softwares/mt4
USER xclient
CMD cd /softwares/mt4 && ./compiler

