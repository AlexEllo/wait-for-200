FROM alpine

RUN apk add --update curl && rm -rf /var/cache/apk/*

COPY ./wait-for-200 /usr/bin/wait-for-200

#RUN chmod +x /wait-for-200.sh

ENTRYPOINT ["/usr/bin/wait-for-200"]
