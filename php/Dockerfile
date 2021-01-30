FROM romeoz/docker-apache-php

RUN apt-get update \
    && apt-get -y install php-redis

COPY phpredis.ini /etc/php/7.3/apache2/conf.d/
# CMD ["chmod", "777", "/etc/php/7.3/apache2/conf.d/phpredis.ini"]

