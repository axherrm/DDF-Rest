# DDF-Rest

*GET*
**http://localhost:8080/result**
    - returns: 
        - 200 OK 
        - 418 I am a teapot
        - result Json when all votes are done

**http://localhost:8080/hello**
    - returns: 
        - 200 OK


*POST*
**http://localhost:8080/vote?voted=VotedPerson&voter=PersonWhoVotes** 
    - returns:
        - 200 OK  
        - 500 Internal Server Error

**http://localhost:8080/reset**
    -  returns:
        - 200 OK 
        - 500 Internal Server Error