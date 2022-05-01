<h1 align="center">Go-DDos 💥</h1>
<img align="right" src="./zorro.svg" height="169"> 
A DDoS attack is a cybercrime in which the attacker floods a server with internet traffic to prevent users from accessing connected online services and sites.

In this repository, I show how a simple DDoS works, but remember: __It is for educational use only__, don't execute this script in real world applications!


## How To Run

Clone the project

```bash
  git clone https://github.com/franciscofeo/go-ddos.git
```

Go to the project directory

```bash
  cd go-ddos
```


Start the application

```bash
  go run main.go <url> <workers quantity>
```



## Example

After started a server locally in http://localhost:8080, we just have to write this code below to execute the attack with 300 workers:

```bash
  go run main http://localhost:8080 300
```
