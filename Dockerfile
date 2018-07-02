FROM 192.168.0.249:20202/hwcse/as-go:1.8.5

COPY ./go113 /home
COPY ./conf /home/conf
RUN chmod +x /home/go113

CMD ["/home/go113"]
