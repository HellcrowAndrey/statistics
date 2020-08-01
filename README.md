STATISTIC SERVICE DEPENDENCY:

POSTGRESQL GOLANG DATABASE DRIVER:
    - go get github.com/lib/pq
GORM GOLANG ORM:    
    - go get -u github.com/jinzhu/gorm
GORILLA MUX GOLANG REST LIB:    
    - go get -u github.com/gorilla/mux
DEPENDENCY INJECTIONS GOLANG LIB:     
    - go get 'go.uber.org/dig@v1'
CONFIG GOLANG LIB:
    - go get github.com/tkanos/gonfig

SWAGGER GOLANG LIBS:
    - go get -u github.com/swaggo/swag/cmd/swag
    - go get -u github.com/swaggo/http-swagger
    - go get github.com/alecthomas/template

STATISTIC SERVICE SWAGGER INSTALL:

Step 1.
    /src/main/go swag init
Step 2.
    add import to main file _ "./docs"  
