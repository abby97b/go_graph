# go_graph

run the server : go run server.go

following are the mutations and query for the graph playground

mutation NewUser{
  createUser(input:{
    FirstName:"faslkfjdsalk;jfalf",
    LastName:"dasfsadf",
    Age:2223,
    Phone:"2222222222",
    Address:"wwdd"
  }){
    FirstName,
    LastName,
    Age,
    Phone,
    Address
  }
}



query getUsers{
  users{
    FirstName,
    LastName,
    Phone,
    id
  }
}



mutation UpdateUser{
  updateUser(input:{
    id:"1443635317331776148"
    FirstName:" changed last",
    LastName:"ban",
    Phone:"9993339991",
    Address:"mumbai pune mumbai"
  }){
    FirstName,
    LastName
    Phone,
    Address
  }
}

mutation{
 deleteUser(input:{
    id:"1443635317331776148"
  }){
  id
  }
}

