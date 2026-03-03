FROM alpine:latest
COPY odd-character-htmx /bin/odd-character-htmx
COPY static/css/style.css /static/css/style.css
ENV PORT=42069
CMD [ "/bin/odd-character-htmx" ]
