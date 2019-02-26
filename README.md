# hbxpack


# running app 

docker build -t app .
docker run -p 8000:8000 app

# running web
docker build -t web
docker run -p 80:80 web

#docker-compose
docker-compose build
docker-compose up

docker-compose scale app=3
docker-compose scale web=3
