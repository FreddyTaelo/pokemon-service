# pokemon-service
Get Pokemon details

Microservice to retrieve pokemons
operations
- get 1st 100 pokemon
- get pokemon details by id

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/FreddyTaelo/pokemon-service/tree/main.svg?style=svg)](https://circleci.com/gh/FreddyTaelo/pokemon-service/tree/main)
[![codecov](https://codecov.io/gh/FreddyTaelo/pokemon-service/graph/badge.svg?token=AAVXBDIK7R)](https://codecov.io/gh/FreddyTaelo/pokemon-service)

# Sample API CAll
```
Swagger
curl http://ec2-13-60-74-36.eu-north-1.compute.amazonaws.com:5000/swagger/index.html

REQUEST( Get pokemon details by id)

 curl http://ec2-13-60-74-36.eu-north-1.compute.amazonaws.com:5000/api/pokemon/1  

RESPONSE

{"id":1,"name":"bulbasaur","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/1.png","height":7,"weight":69,"types":["grass","poison"],"stats":[{"name":"hp","value":45},{"name":"attack","value":49},{"name":"defense","value":49},{"name":"special-attack","value":65},{"name":"special-defense","value":65},{"name":"speed","value":45}]}

REQUEST( Get 1st 100 pokemons)

curl http://ec2-13-60-74-36.eu-north-1.compute.amazonaws.com:5000/api/pokemon 

RESPONSE

[{"id":1,"name":"bulbasaur","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png","height":0,"weight":0,"types":null,"stats":null},{"id":2,"name":"ivysaur","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/2.png","height":0,"weight":0,"types":null,"stats":null},{"id":3,"name":"venusaur","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/3.png","height":0,"weight":0,"types":null,"stats":null},{"id":4,"name":"charmander","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/4.png","height":0,"weight":0,"types":null,"stats":null},{"id":5,"name":"charmeleon","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/5.png","height":0,"weight":0,"types":null,"stats":null},{"id":6,"name":"charizard","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/6.png","height":0,"weight":0,"types":null,"stats":null},{"id":7,"name":"squirtle","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/7.png","height":0,"weight":0,"types":null,"stats":null},{"id":8,"name":"wartortle","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/8.png","height":0,"weight":0,"types":null,"stats":null},..]
```

# Instruction to run
```
Using make file

$ make run

```
Todo 

- add more extra test
- optimize memory allocation
- go routine around