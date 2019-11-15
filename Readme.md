Notes:

Download docker:
sudo apt install docker.io

Build a custom image from the Dockerfile:
sudo docker build -t postgres .

Build and run a container:
sudo docker run -d --name pgcontainer -p 5432:5432 postgres

Open psql shell:
sudo docker exec -it pgcontainer psql -U postgres

List tables in database in psql:
\dt

View rows in doctors table:
select * from doctors;

Exit psql:
\q


package main

import "os/exec"



