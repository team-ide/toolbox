FROM centos:7

EXPOSE 18000/tcp

COPY release/go-linux-x64 /data/toolbox
RUN chmod +x /data/toolbox/toolbox
CMD cd /data/toolbox && ./toolbox


# docker build -t teamide/toolbox .


# docker run --name toolbox-18000 -m 256m -p 18000:18000 teamide/toolbox
# docker run -itd --name toolbox-18000 -m 256m -p 18000:18000 --restart=always teamide/toolbox


# docker stop toolbox-18000
# docker rm toolbox-18000