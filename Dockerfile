FROM alpine

COPY ./wait-for-200 /usr/bin/wait-for-200

#RUN chmod +x /wait-for-200.sh

ENTRYPOINT ["/usr/bin/wait-for-200"]
