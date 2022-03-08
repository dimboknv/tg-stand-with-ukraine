#!/usr/bin/dumb-init /bin/sh

# https://github.com/umputun/baseimage

uid=$(id -u)

if [[ ${uid} -eq 0 ]]; then
    # set UID for user app
    if [[ "${APP_UID}" -ne "1000" ]]; then
        echo "set custom app uid=${APP_UID}"
        sed -i "s/:1000:1000:/:${APP_UID}:${APP_UID}:/g" /etc/passwd
    fi
    chown -R app:app /app
fi

if [[ ${uid} -eq 0 ]]; then
   exec su-exec app "$@"
else
   exec "$@"
fi
