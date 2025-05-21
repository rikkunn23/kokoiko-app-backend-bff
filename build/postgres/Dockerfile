FROM postgres:14
RUN apt-get update && apt-get install -y \
    postgresql-14-postgis-3 \
    locales
RUN sed -i -E 's/# (ja_JP.UTF-8)/\1/' /etc/locale.gen
RUN locale-gen
RUN update-locale LANG=ja_JP.UTF-8
ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
EXPOSE 5432
ENTRYPOINT [ "bash", "/postgres/set_up_db.sh" ]
