cp ./myweb_backend /data/lpf/prod

cd /data/lpf/prod

mv ./myweb_backend main

chmod 777 main

docker stop docker_myweb1

docker rm docker_myweb1

docker run -d -p 8001:3000 -v /data/lpf/prod:/usr/src/myapp --name docker_myweb1 lvpf/docker-main:2.0.0
