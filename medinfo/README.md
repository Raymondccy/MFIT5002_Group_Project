# Bioinfo
cp -r /src/MFIT5002_Group_Project/medinfo/ $GOPATH/src
cp -r /src/medinfo/ /root/go/src


sudo -i #切换当前用户为root用户

Step 1 — Installing Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version


curl -fsSL https://test.docker.com -o test-docker.sh
sudo sh test-docker.sh

sudo apt install make

wget https://go.dev/dl/go1.15.15.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.15.15.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
source $HOME/.profile
go version


mkdir -p $GOPATH/src/
cd $GOPATH/src && git clone https://github.com/Raymondccy/MFIT5002_Group_Project.git
```
在`/etc/hosts`中添加：
```
127.0.0.1  orderer.example.com
127.0.0.1  peer0.org1.example.com
127.0.0.1  peer1.org1.example.com
```
添加依赖：
```
cd medinfo && go mod tidy
```
运行项目：
```
./clean_docker.sh
```
在`127.0.0.1:9000`进行访问
