FROM golang as stage0
ENV CGO_ENABLED 0
COPY . /pv

WORKDIR /pv/app/services/portfolio-apis/
RUN go build

FROM alpine:latest
EXPOSE 4567
COPY --from=stage0 /pv/app/services/portfolio-apis/portfolio-apis /pv/portfolio-apis
WORKDIR /pv
CMD [ "./portfolio-apis" ]