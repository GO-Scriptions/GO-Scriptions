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


func main(){
        // build image and container from Dockerfile

        // stop container
        dockerstop := exec.Command("sudo", "docker", "stop", "pgcontainer")
        dockerstop.Run()

        // remove container
        dockerdelete := exec.Command("sudo", "docker", "container", "rm", "pgcontainer")
        dockerdelete.Run()

        // remove image
        dockerrm := exec.Command("sudo","docker", "rmi", "pgimage")
        dockerrm.Run()

        // build image
        dockerbuild := exec.Command("sudo", "docker", "build", "-t", "pgimage", ".")
        dockerbuild.Run()

        // build container
        dockerrun := exec.Command("sudo", "docker", "run", "--name=pgcontainer", "-p=5432:5432", "-d", "pgimage")
        dockerrun.Run()

        // to open psql shell- sudo docker exec -it pgcontainer psql -U postgres
        // to list tables in database in psql: \dt
        // to view rows in doctors table: select * from doctors;
        // to exit psql: \q
}


