FROM alpine

RUN apk add --update curl && rm -rf /var/cache/apk/*

ADD wait-for-200.sh /wait-for-200.sh

RUN chmod +x /wait-for-200.sh

ENTRYPOINT ["/wait-for-200.sh"]
