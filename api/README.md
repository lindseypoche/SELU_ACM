

### Package layout
cmd (for starting up a service)
- acm_be (acm backend)
- acm_bot (acm bot)
internal (contains code that cannot be accessed outside of this project)
  - adapters (ports/externals/plugins/framewowrks for the project)
    - discord 
    - http ()
    - storage (datasource / persistence layer)
      - (db access functions, models, db configuration)
    - web (router, controllers, models)
      - (data must be transformed to model object before being sent to domain)
  - domain (business logic. ie use cases, models)


### DDD Architecture (Hexagonal) Example
internal
- ADAPTERS / INFRASTRUCTURE 
  - large frameworks 
  - sms 
    - sms server
  - web (application backend)
  - http 
  - rest api 
    - controllers / handlers 
  - grpc api 
    - controllers / handlers
  - db / persistence 
    - ORM adaptor
      - mysql 
      - mongo 
  - search 
    - search engine (elasticsearch)
  - logging 
  - email 
    - mailing server
  - files 
  - providers / externals 
  - admin cli/gui
    - views & controllers
  - consumer gui 
    - views & controllers 
- DOMAIN (must be testable without ports or adapters)
  - service / use cases / business logic 
  - repository interface 
  - emailer interface 